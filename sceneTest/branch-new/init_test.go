package branchnew

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/branch"
)

var (
	branchClient *branch.APIClient
	clusterId    string
)

func init() {
	pflag.StringVar(&clusterId, "cluster-id", "", "")
}

func setup() {
	cfg := config.LoadConfig()
	if strings.TrimSpace(clusterId) == "" {
		log.Panic("clusterId is required, use -cid")
	}

	var err error
	branchClient, err = NewBranchClient(cfg)
	if err != nil {
		log.Panicf("failed to create branch client: %v", err)
	}

	log.Printf("Use existing cluster for branch-new tests - ClusterID: %s", clusterId)
}

func NewBranchClient(cfg *config.Config) (*branch.APIClient, error) {
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}

	serverlessURL, err := util.ValidateApiUrl(cfg.Endpoint.Serverless)
	if err != nil {
		return nil, fmt.Errorf("invalid serverless endpoint: %w", err)
	}

	branchCfg := branch.NewConfiguration()
	branchCfg.HTTPClient = httpClient
	branchCfg.Host = serverlessURL.Host
	branchCfg.UserAgent = util.UserAgent

	return branch.NewAPIClient(branchCfg), nil
}

func cleanupAllBranches(ctx context.Context, clusterId string) error {
	branches, err := listAllBranches(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("list branches for cleanup: %w", err)
	}
	if len(branches) == 0 {
		log.Printf("No branches need cleanup for cluster %s", clusterId)
		return nil
	}

	sortBranchesForCleanup(branches)

	for _, bran := range branches {
		if bran.BranchId == nil || strings.TrimSpace(*bran.BranchId) == "" {
			continue
		}

		_, h, err := branchClient.BranchServiceAPI.BranchServiceDeleteBranch(ctx, clusterId, *bran.BranchId).Execute()
		if err := util.ParseError(err, h); err != nil {
			log.Printf("Cleanup failed to delete branch %s: %v", branchLabel(bran), err)
			continue
		}

		log.Printf("Cleanup deleted branch %s", branchLabel(bran))
	}

	remaining, err := listAllBranches(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("list remaining branches: %w", err)
	}
	if len(remaining) == 0 {
		return nil
	}

	labels := make([]string, 0, len(remaining))
	for _, bran := range remaining {
		labels = append(labels, branchLabel(bran))
	}

	return fmt.Errorf("remaining branches after cleanup: %s", strings.Join(labels, ", "))
}

func listAllBranches(ctx context.Context, clusterId string) ([]branch.Branch, error) {
	const pageSize int32 = 100

	branches := make([]branch.Branch, 0)
	pageToken := ""

	for {
		req := branchClient.BranchServiceAPI.BranchServiceListBranches(ctx, clusterId).PageSize(pageSize)
		if pageToken != "" {
			req = req.PageToken(pageToken)
		}

		resp, h, err := req.Execute()
		if err := util.ParseError(err, h); err != nil {
			return nil, err
		}

		branches = append(branches, resp.GetBranches()...)
		pageToken = resp.GetNextPageToken()
		if pageToken == "" {
			return branches, nil
		}
	}
}

func sortBranchesForCleanup(branches []branch.Branch) {
	parentByID := make(map[string]string, len(branches))
	depthCache := make(map[string]int, len(branches))

	for _, bran := range branches {
		if bran.BranchId == nil || bran.ParentId == nil {
			continue
		}
		parentByID[*bran.BranchId] = *bran.ParentId
	}

	var depth func(string, map[string]bool) int
	depth = func(branchID string, visiting map[string]bool) int {
		if cached, ok := depthCache[branchID]; ok {
			return cached
		}
		parentID, ok := parentByID[branchID]
		if !ok || parentID == "" || visiting[branchID] {
			return 0
		}

		visiting[branchID] = true
		branchDepth := 1 + depth(parentID, visiting)
		delete(visiting, branchID)
		depthCache[branchID] = branchDepth
		return branchDepth
	}

	sort.SliceStable(branches, func(i, j int) bool {
		leftDepth := 0
		if branches[i].BranchId != nil {
			leftDepth = depth(*branches[i].BranchId, map[string]bool{})
		}

		rightDepth := 0
		if branches[j].BranchId != nil {
			rightDepth = depth(*branches[j].BranchId, map[string]bool{})
		}

		return leftDepth > rightDepth
	})
}

func branchLabel(bran branch.Branch) string {
	branchID := "<unknown-id>"
	if bran.BranchId != nil && strings.TrimSpace(*bran.BranchId) != "" {
		branchID = *bran.BranchId
	}

	displayName := bran.DisplayName
	if displayName == "" {
		displayName = "<unknown-name>"
	}

	return fmt.Sprintf("%s(%s)", displayName, branchID)
}
