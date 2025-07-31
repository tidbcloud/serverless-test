package cluster

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/lithammer/shortuuid/v4"
	"github.com/stretchr/testify/require"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cluster"
)

var testClusterId string
var region = "regions/aws-us-east-1"

// go test -v sceneTest/cluster/* -project-id {project-id}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestCreateCluster(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)

	// Cleanup: Delete any existing cluster with the ClusterTest- prefix
	clusterName := "ClusterTest-" + shortuuid.New()
	t.Logf("Cleaning up any existing cluster with ClusterTest- prefix")

	// Clean up existing cluster using similar pattern to branch test
	if err := cleanupExistingCluster(ctx, projectId, "ClusterTest-"); err != nil {
		t.Logf("Failed to cleanup existing cluster: %v", err)
	}

	t.Logf("Creating cluster: %s", clusterName)

	// Set spending limit
	spendLimit := int32(100)

	// Create cluster configuration
	clusterBody := cluster.TidbCloudOpenApiserverlessv1beta1Cluster{
		DisplayName: clusterName,
		Region: cluster.Commonv1beta1Region{
			Name: &region,
		},
		SpendingLimit: &cluster.ClusterSpendingLimit{
			Monthly: &spendLimit,
		},
		//EncryptionConfig: &cluster.V1beta1ClusterEncryptionConfig{
		//	EnhancedEncryptionEnabled: pointer.ToBool(true),
		//},
		//HighAvailabilityType: cluster.CLUSTERHIGHAVAILABILITYTYPE_REGIONAL.Ptr(),
	}

	// Add project labels if project ID is provided
	if projectId != "" {
		clusterBody.Labels = &map[string]string{"tidb.cloud/project": projectId}
	}

	resp, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceCreateCluster(ctx).
		Cluster(clusterBody).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to create cluster: %v", err)
	}

	assert.NotNil(resp)
	assert.NotEmpty(*resp.ClusterId)
	assert.Equal(clusterName, resp.DisplayName)
	assert.Equal(region, *resp.Region.Name)

	testClusterId = *resp.ClusterId

	// Wait for cluster to become active
	activeCluster, err := waitForClusterActive(ctx, testClusterId)
	if err != nil {
		t.Fatalf("Cluster failed to become active: %v", err)
	}

	t.Logf("Cluster created successfully - ID: %s, Name: %s, Region: %s",
		testClusterId, activeCluster.DisplayName, *activeCluster.Region.Name)
}

// cleanupExistingCluster removes any existing test cluster with the given cluster name prefix
func cleanupExistingCluster(ctx context.Context, projectId, clusterNamePrefix string) error {
	req := clusterClient.ServerlessServiceAPI.ServerlessServiceListClusters(ctx).PageSize(100)

	if projectId != "" {
		projectFilter := fmt.Sprintf("projectId=%s", projectId)
		req = req.Filter(projectFilter)
	}

	clusters, h, err := req.Execute()
	if err := util.ParseError(err, h); err != nil {
		return fmt.Errorf("failed to list clusters: %w", err)
	}

	// Find and delete existing test clusters with the prefix
	for _, clu := range clusters.Clusters {
		if strings.HasPrefix(clu.DisplayName, clusterNamePrefix) {
			fmt.Printf("Found existing cluster with prefix %s: %s, deleting it\n", clusterNamePrefix, clu.DisplayName)
			deleteCluster(*clu.ClusterId)
		}
	}

	return nil
}

// deleteCluster deletes a cluster by its ID
func deleteCluster(clusterID string) {
	ctx := context.Background()

	_, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceDeleteCluster(ctx, clusterID).Execute()
	if err := util.ParseError(err, h); err != nil {
		fmt.Printf("Failed to delete cluster %s: %v\n", clusterID, err)
	} else {
		fmt.Printf("Successfully deleted cluster %s\n", clusterID)
	}
}

func TestGetCluster(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)

	resp, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, testClusterId).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to get cluster: %v", err)
	}

	assert.NotNil(resp)
	assert.Equal(testClusterId, *resp.ClusterId)
	assert.NotEmpty(resp.DisplayName)
	assert.Equal(region, *resp.Region.Name)
	assert.NotNil(resp.SpendingLimit)
	assert.Equal(int32(100), *resp.SpendingLimit.Monthly)

	t.Logf("Cluster retrieved successfully - ID: %s, Name: %s", *resp.ClusterId, resp.DisplayName)
}

func TestUpdateCluster(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)

	// 1. Update displayName
	newDisplayName := "ClusterTest-" + shortuuid.New()
	updateDisplayNameBody := cluster.V1beta1ServerlessServicePartialUpdateClusterBody{
		Cluster: &cluster.RequiredTheClusterToBeUpdated{
			DisplayName: &newDisplayName,
		},
		UpdateMask: "displayName",
	}

	resp, h, err := clusterClient.ServerlessServiceAPI.ServerlessServicePartialUpdateCluster(ctx, testClusterId).
		Body(updateDisplayNameBody).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to update cluster display name: %v", err)
	}
	assert.NotNil(resp)
	assert.Equal(newDisplayName, resp.DisplayName)

	// Verify the update
	getResp, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, testClusterId).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to get cluster after display name update: %v", err)
	}
	assert.Equal(newDisplayName, getResp.DisplayName)
	t.Logf("Updated display name to: %s", newDisplayName)

	// 2. Update spendingLimit
	newSpendLimit := int32(200)
	updateSpendingLimitBody := cluster.V1beta1ServerlessServicePartialUpdateClusterBody{
		Cluster: &cluster.RequiredTheClusterToBeUpdated{
			SpendingLimit: &cluster.ClusterSpendingLimit{
				Monthly: &newSpendLimit,
			},
		},
		UpdateMask: "spendingLimit",
	}
	resp, h, err = clusterClient.ServerlessServiceAPI.ServerlessServicePartialUpdateCluster(ctx, testClusterId).
		Body(updateSpendingLimitBody).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to update cluster spending limit: %v", err)
	}
	assert.NotNil(resp)
	assert.NotNil(resp.SpendingLimit)
	assert.Equal(newSpendLimit, *resp.SpendingLimit.Monthly)

	// Verify the update
	getResp, h, err = clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, testClusterId).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to get cluster after spending limit update: %v", err)
	}
	assert.Equal(newSpendLimit, *getResp.SpendingLimit.Monthly)
	t.Logf("Updated spending limit to: %d", newSpendLimit)

	// 3. Update automatedBackupPolicy
	backupRetentionDays := int32(7)
	backupTime := "02:00"
	updateBackupPolicyBody := cluster.V1beta1ServerlessServicePartialUpdateClusterBody{
		Cluster: &cluster.RequiredTheClusterToBeUpdated{
			AutomatedBackupPolicy: &cluster.V1beta1ClusterAutomatedBackupPolicy{
				RetentionDays: &backupRetentionDays,
				StartTime:     &backupTime,
			},
		},
		UpdateMask: "automatedBackupPolicy",
	}
	resp, h, err = clusterClient.ServerlessServiceAPI.ServerlessServicePartialUpdateCluster(ctx, testClusterId).
		Body(updateBackupPolicyBody).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to update cluster automated backup policy: %v", err)
	}
	assert.NotNil(resp)
	assert.NotNil(resp.AutomatedBackupPolicy)
	assert.Equal(backupRetentionDays, *resp.AutomatedBackupPolicy.RetentionDays)

	// Verify the update
	getResp, h, err = clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, testClusterId).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to get cluster after automated backup policy update: %v", err)
	}
	assert.Equal(backupRetentionDays, *getResp.AutomatedBackupPolicy.RetentionDays)
	assert.Equal(backupTime, *getResp.AutomatedBackupPolicy.StartTime)
	t.Logf("Updated automated backup policy retention days to: %d", backupRetentionDays)

	// 4. Update endpoints (labels as a substitute since endpoints might not be available)
	newLabels := map[string]string{
		"test-label":  "test-value",
		"environment": "testing",
	}
	updateLabelsBody := cluster.V1beta1ServerlessServicePartialUpdateClusterBody{
		Cluster: &cluster.RequiredTheClusterToBeUpdated{
			Labels: &newLabels,
		},
		UpdateMask: "labels",
	}
	resp, h, err = clusterClient.ServerlessServiceAPI.ServerlessServicePartialUpdateCluster(ctx, testClusterId).
		Body(updateLabelsBody).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to update cluster labels: %v", err)
	}
	assert.NotNil(resp)
	assert.NotNil(resp.Labels)
	assert.Equal("test-value", (*resp.Labels)["test-label"])
	assert.Equal("testing", (*resp.Labels)["environment"])

	// Verify the update
	getResp, h, err = clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, testClusterId).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to get cluster after labels update: %v", err)
	}
	assert.Equal("test-value", (*getResp.Labels)["test-label"])
	assert.Equal("testing", (*getResp.Labels)["environment"])
	t.Logf("Updated labels - test-label: %s, environment: %s",
		(*getResp.Labels)["test-label"], (*getResp.Labels)["environment"])
}

func TestDeleteCluster(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)

	_, h, err := clusterClient.ServerlessServiceAPI.ServerlessServiceDeleteCluster(ctx, testClusterId).Execute()
	if err := util.ParseError(err, h); err != nil {
		t.Fatalf("Failed to delete cluster: %v", err)
	}
	t.Logf("Cluster deleted successfully: %s", testClusterId)

	// Verify the cluster is deleted by trying to get it
	_, h, err = clusterClient.ServerlessServiceAPI.ServerlessServiceGetCluster(ctx, testClusterId).Execute()
	assert.Error(err, "Cluster should not exist after deletion")
}
