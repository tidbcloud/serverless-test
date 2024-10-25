package branch

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/branch"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cluster"
)

func setup() {
	var err error
	config.InitializeConfig()
	branchClient, err = NewBranchClient()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	clusterClient, err = NewClusterClient()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	clu, err := CreateCluster()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	clusterId = *clu.ClusterId
	println(fmt.Sprintf("cluster created successfully, clusterId is %s, region is %s", *clu.ClusterId, *clu.Region.Name))
}

func NewBranchClient() (*branch.APIClient, error) {
	httpclient := &http.Client{
		Transport: util.NewDigestTransport(config.PublicKey, config.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(config.ServerlessEndpoint)
	if err != nil {
		return nil, err
	}
	cfg := branch.NewConfiguration()
	cfg.HTTPClient = httpclient
	cfg.Host = serverlessURL.Host
	cfg.UserAgent = util.UserAgent
	return branch.NewAPIClient(cfg), nil
}

func NewClusterClient() (*cluster.APIClient, error) {
	httpclient := &http.Client{
		Transport: util.NewDigestTransport(config.PublicKey, config.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(config.ServerlessEndpoint)
	if err != nil {
		return nil, err
	}
	cfg := cluster.NewConfiguration()
	cfg.HTTPClient = httpclient
	cfg.Host = serverlessURL.Host
	cfg.UserAgent = util.UserAgent
	return cluster.NewAPIClient(cfg), nil
}

func CreateCluster() (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ctx := context.Background()
	clusterName := "branch-test"

	req := clusterClient.ServerlessServiceAPI.ServerlessServiceListClusters(ctx)
	req = req.PageSize(100)
	if config.ProjectId != "" {
		projectFilter := fmt.Sprintf("projectId=%s", config.ProjectId)
		req = req.Filter(projectFilter)
	}
	clusters, h, err := req.Execute()
	err = util.ParseError(err, h)
	if err != nil {
		return nil, err
	}

	for _, clu := range clusters.Clusters {
		if clu.DisplayName == clusterName {
			DeleteCluster(*clu.ClusterId)
			break
		}
	}

	var spendLimit int32 = 100
	region := config.GetRandomRegion()
	clusterBody := cluster.TidbCloudOpenApiserverlessv1beta1Cluster{
		DisplayName: clusterName,
		SpendingLimit: &cluster.ClusterSpendingLimit{
			Monthly: &spendLimit,
		},
		Region: cluster.Commonv1beta1Region{
			Name: &region,
		},
	}
	if config.ProjectId != "" {
		clusterBody.Labels = &map[string]string{"tidb.cloud/project": config.ProjectId}
	}

	resp, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceCreateCluster(ctx).Cluster(clusterBody).Execute()
	err = util.ParseError(err, h)
	if err != nil {
		return nil, err
	}

	resp, err = checkServerlessState(ctx, *resp.ClusterId)
	if err != nil {
		return nil, err
	}
	if *resp.State != cluster.COMMONV1BETA1CLUSTERSTATE_ACTIVE {
		return nil, errors.New("create cluster failed, state is" + string(*resp.State))
	}
	return resp, nil
}

func checkServerlessState(ctx context.Context, clusterId string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ticker := time.NewTicker(time.Second * 10)
	timeout := time.After(time.Minute * 3)
	for {
		select {
		case <-ticker.C:
			res, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, clusterId).Execute()
			if util.ParseError(err, h) != nil {
				println("get cluster failed: " + util.ParseError(err, h).Error())
				continue
			}
			if *res.State != cluster.COMMONV1BETA1CLUSTERSTATE_CREATING {
				return res, nil
			}
		case <-timeout:
			return nil, errors.New("create cluster timeout")
		}
	}
}

func DeleteCluster(clusterId string) {
	ctx := context.Background()
	_, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceDeleteCluster(ctx, clusterId).Execute()
	err = util.ParseError(err, h)
	if err != nil {
		println("delete cluster failed: " + err.Error())
	} else {
		println("delete cluster success")
	}
}
