package branch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/branch"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cluster"
)

// setup initializes the test environment by creating API clients and a test cluster
func setup() {
	cfg := config.LoadConfig()

	// Initialize branch client
	var err error
	branchClient, err = NewBranchClient(cfg)
	if err != nil {
		log.Panicf("failed to create branch client: %v", err)
	}

	// Initialize cluster client
	clusterClient, err = NewClusterClient(cfg)
	if err != nil {
		log.Panicf("failed to create cluster client: %v", err)
	}

	// Create test cluster
	cluster, err := CreateCluster(cfg.ProjectID, "branch-test")
	if err != nil {
		log.Panicf("failed to create test cluster: %v", err)
	}

	clusterId = *cluster.ClusterId
	log.Printf("Test cluster created successfully - ClusterID: %s, Region: %s", *cluster.ClusterId, *cluster.Region.Name)
}

// NewBranchClient creates a new branch API client with the given configuration
func NewBranchClient(cfg *config.Config) (*branch.APIClient, error) {
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}

	serverlessURL, err := util.ValidateApiUrl(cfg.ServerlessEndpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid serverless endpoint: %w", err)
	}

	branchCfg := branch.NewConfiguration()
	branchCfg.HTTPClient = httpClient
	branchCfg.Host = serverlessURL.Host
	branchCfg.UserAgent = util.UserAgent

	return branch.NewAPIClient(branchCfg), nil
}

// NewClusterClient creates a new cluster API client with the given configuration
func NewClusterClient(cfg *config.Config) (*cluster.APIClient, error) {
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}

	serverlessURL, err := util.ValidateApiUrl(cfg.ServerlessEndpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid serverless endpoint: %w", err)
	}

	clusterCfg := cluster.NewConfiguration()
	clusterCfg.HTTPClient = httpClient
	clusterCfg.Host = serverlessURL.Host
	clusterCfg.UserAgent = util.UserAgent

	return cluster.NewAPIClient(clusterCfg), nil
}

// CreateCluster creates a new test cluster with the specified project ID and cluster name
func CreateCluster(projectId, clusterName string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ctx := context.Background()

	// Clean up any existing test cluster
	if err := cleanupExistingCluster(ctx, projectId, clusterName); err != nil {
		return nil, fmt.Errorf("failed to cleanup existing cluster: %w", err)
	}

	// Create new cluster configuration
	clusterBody, err := buildClusterConfig(projectId, clusterName)
	if err != nil {
		return nil, fmt.Errorf("failed to build cluster config: %w", err)
	}

	// Create the cluster
	resp, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceCreateCluster(ctx).
		Cluster(clusterBody).Execute()
	if err := util.ParseError(err, h); err != nil {
		return nil, fmt.Errorf("failed to create cluster: %w", err)
	}

	// Wait for cluster to become active
	activeCluster, err := waitForClusterActive(ctx, *resp.ClusterId)
	if err != nil {
		return nil, fmt.Errorf("cluster failed to become active: %w", err)
	}

	return activeCluster, nil
}

// cleanupExistingCluster removes any existing test cluster with the given project ID and cluster name
func cleanupExistingCluster(ctx context.Context, projectId, clusterName string) error {
	req := clusterClient.ServerlessServiceAPI.ServerlessServiceListClusters(ctx).PageSize(100)

	if projectId != "" {
		projectFilter := fmt.Sprintf("projectId=%s", projectId)
		req = req.Filter(projectFilter)
	}

	clusters, h, err := req.Execute()
	if err := util.ParseError(err, h); err != nil {
		return fmt.Errorf("failed to list clusters: %w", err)
	}

	// Find and delete existing test cluster
	for _, clu := range clusters.Clusters {
		if clu.DisplayName == clusterName {
			DeleteCluster(*clu.ClusterId)
			break
		}
	}

	return nil
}

// buildClusterConfig creates a cluster configuration based on the project ID and cluster name
func buildClusterConfig(projectId, clusterName string) (cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	region := config.GetRandomRegion()

	clusterBody := cluster.TidbCloudOpenApiserverlessv1beta1Cluster{
		DisplayName: clusterName,
		Region: cluster.Commonv1beta1Region{
			Name: &region,
		},
	}

	// Configure based on cloud provider
	if strings.Contains(region, "aws") {
		// AWS configuration with spending limit
		spendLimit := int32(100)
		clusterBody.SpendingLimit = &cluster.ClusterSpendingLimit{
			Monthly: &spendLimit,
		}
	} else {
		minRcu := int64(2000)
		maxRcu := int64(4000)
		clusterBody.AutoScaling = &cluster.V1beta1ClusterAutoScaling{
			MinRcu: &minRcu,
			MaxRcu: &maxRcu,
		}
	}

	// Add project labels if project ID is provided
	if projectId != "" {
		clusterBody.Labels = &map[string]string{"tidb.cloud/project": projectId}
	}

	return clusterBody, nil
}

// waitForClusterActive waits for the cluster to reach active state with timeout
func waitForClusterActive(ctx context.Context, clusterID string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	const (
		checkInterval = 10 * time.Second
		timeout       = 3 * time.Minute
	)

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	timeoutTimer := time.After(timeout)

	for {
		select {
		case <-ticker.C:
			c, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, clusterID).Execute()
			if err := util.ParseError(err, h); err != nil {
				log.Printf("Failed to get cluster status: %v", err)
				continue
			}

			if *c.State == cluster.COMMONV1BETA1CLUSTERSTATE_ACTIVE {
				return c, nil
			}

		case <-timeoutTimer:
			return nil, errors.New("cluster creation timed out after 3 minutes")
		}
	}
}

// DeleteCluster deletes a cluster by its ID
func DeleteCluster(clusterID string) {
	ctx := context.Background()

	_, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceDeleteCluster(ctx, clusterID).Execute()
	if err := util.ParseError(err, h); err != nil {
		log.Printf("Failed to delete cluster %s: %v", clusterID, err)
	} else {
		log.Printf("Successfully deleted cluster %s", clusterID)
	}
}
