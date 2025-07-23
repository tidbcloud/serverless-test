package branch

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/branch"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cluster"
)

func setup() {
	cfg := config.LoadConfig()
	var err error
	branchClient, err = NewBranchClient(cfg)
	if err != nil {
		panic(err)
	}
	clusterClient, err = NewClusterClient(cfg)
	if err != nil {
		panic(err)
	}
	clu, err := CreateCluster(cfg.ProjectID)
	if err != nil {
		panic(err)
	}
	clusterId = *clu.ClusterId
	println(fmt.Sprintf("cluster created successfully, clusterId is %s, region is %s", *clu.ClusterId, *clu.Region.Name))
}

func NewBranchClient(cfg *config.Config) (*branch.APIClient, error) {
	httpclient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(cfg.ServerlessEndpoint)
	if err != nil {
		return nil, err
	}
	bcfg := branch.NewConfiguration()
	bcfg.HTTPClient = httpclient
	bcfg.Host = serverlessURL.Host
	bcfg.UserAgent = util.UserAgent
	return branch.NewAPIClient(bcfg), nil
}

func NewClusterClient(cfg *config.Config) (*cluster.APIClient, error) {
	httpclient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(cfg.ServerlessEndpoint)
	if err != nil {
		return nil, err
	}
	ccfg := cluster.NewConfiguration()
	ccfg.HTTPClient = httpclient
	ccfg.Host = serverlessURL.Host
	ccfg.UserAgent = util.UserAgent
	return cluster.NewAPIClient(ccfg), nil
}

func CreateCluster(projectId string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ctx := context.Background()
	clusterName := "branch-test"

	req := clusterClient.ServerlessServiceAPI.ServerlessServiceListClusters(ctx)
	req = req.PageSize(100)
	if projectId != "" {
		projectFilter := fmt.Sprintf("projectId=%s", projectId)
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
		Region: cluster.Commonv1beta1Region{
			Name: &region,
		},
	}
	if strings.Contains(region, "aws") {
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
	if projectId != "" {
		clusterBody.Labels = &map[string]string{"tidb.cloud/project": projectId}
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
