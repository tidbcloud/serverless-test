package cluster

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/pflag"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cluster"
)

var (
	clusterClient *cluster.APIClient
	projectId     string
)

func init() {
	pflag.StringVar(&projectId, "project-id", "", "")
}

// setup initializes the test environment by creating API clients
func setup() {
	cfg := config.LoadConfig()

	// Initialize cluster client
	var err error
	clusterClient, err = NewClusterClient(cfg)
	if err != nil {
		log.Panicf("failed to create cluster client: %v", err)
	}
}

// NewClusterClient creates a new cluster API client with the given configuration
func NewClusterClient(cfg *config.Config) (*cluster.APIClient, error) {
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}

	serverlessURL, err := util.ValidateApiUrl(cfg.Endpoint.Serverless)
	if err != nil {
		return nil, fmt.Errorf("invalid serverless endpoint: %w", err)
	}

	clusterCfg := cluster.NewConfiguration()
	clusterCfg.HTTPClient = httpClient
	clusterCfg.Host = serverlessURL.Host
	clusterCfg.UserAgent = util.UserAgent

	return cluster.NewAPIClient(clusterCfg), nil
}

// waitForClusterActive waits for the cluster to reach active state with timeout
func waitForClusterActive(ctx context.Context, clusterID string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	const (
		checkInterval = 10 * time.Second
		timeout       = 5 * time.Minute
	)

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	timeoutTimer := time.After(timeout)

	for {
		select {
		case <-ticker.C:
			c, h, err := clusterClient.ClusterServiceAPI.ClusterServiceGetCluster(ctx, clusterID).Execute()
			if err := util.ParseError(err, h); err != nil {
				log.Printf("Failed to get cluster status: %v", err)
				continue
			}

			if *c.State == cluster.COMMONV1BETA1CLUSTERSTATE_ACTIVE {
				return c, nil
			}

		case <-timeoutTimer:
			return nil, errors.New("cluster creation timed out after 5 minutes")
		}
	}
}
