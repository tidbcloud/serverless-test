package branch

import (
	"context"
	"errors"
	"github.com/lithammer/shortuuid/v4"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/branch"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cluster"
)

var (
	clusterClient *cluster.APIClient
	branchClient  *branch.APIClient
	clusterId     string
)

// go test -v sceneTest/branch/* -project-id {project-id} -config {config}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestResetAndCreateFromBranch(t *testing.T) {
	ctx := context.Background()

	name := "test-" + shortuuid.New()
	t.Logf("create branch: %s", name)
	body := &branch.Branch{DisplayName: name}
	bran, err := createBranch(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	bran, err = checkBranchState(ctx, clusterId, *bran.BranchId, t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *bran.State, branch.BRANCHSTATE_ACTIVE)

	t.Log("reset branch")
	bran, h, err := branchClient.BranchServiceAPI.BranchServiceResetBranch(ctx, clusterId, *bran.BranchId).Execute()
	err = util.ParseError(err, h)
	if err != nil {
		t.Fatal(err)
	}
	bran, err = checkBranchState(ctx, clusterId, *bran.BranchId, t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *bran.State, branch.BRANCHSTATE_ACTIVE)

	parentId := *bran.BranchId
	name2 := "test-" + shortuuid.New()
	t.Logf("create branch %s from branch", name2)
	bran, err = createBranch(ctx, clusterId, &branch.Branch{
		DisplayName: name2,
		ParentId:    &parentId,
	})
	if err != nil {
		t.Fatal(err)
	}
	bran, err = checkBranchState(ctx, clusterId, *bran.BranchId, t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *bran.State, branch.BRANCHSTATE_ACTIVE)
	assert.Equal(t, *bran.ParentId, parentId)
}

func TestCreateFromBranch(t *testing.T) {
	ctx := context.Background()

	name := "test-" + shortuuid.New()
	body := &branch.Branch{DisplayName: name}
	t.Logf("create branch: %s", name)
	bran, err := createBranch(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	bran, err = checkBranchState(ctx, clusterId, *bran.BranchId, t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *bran.State, branch.BRANCHSTATE_ACTIVE)

	t.Log("create branch from branch")
	parentId := *bran.BranchId
	name2 := "test-" + shortuuid.New()
	body2 := &branch.Branch{
		DisplayName: name2,
		ParentId:    &parentId,
	}
	bran, err = createBranch(ctx, clusterId, body2)
	if err != nil {
		t.Fatal(err)
	}
	bran, err = checkBranchState(ctx, clusterId, *bran.BranchId, t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *bran.State, branch.BRANCHSTATE_ACTIVE)
	assert.Equal(t, *bran.ParentId, parentId)
}

func TestSpecifyTimestamp(t *testing.T) {
	ctx := context.Background()

	parentTimeStamp := time.Now()
	time.Sleep(time.Second * 1)
	name := "test-" + shortuuid.New()
	body := &branch.Branch{DisplayName: name}
	body.SetParentTimestamp(parentTimeStamp)
	t.Logf("create branch: %s", name)
	bran, err := createBranch(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	bran, err = checkBranchState(ctx, clusterId, *bran.BranchId, t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *bran.State, branch.BRANCHSTATE_ACTIVE)
}

func createBranch(ctx context.Context, clusterId string, body *branch.Branch) (*branch.Branch, error) {
	req := branchClient.BranchServiceAPI.BranchServiceCreateBranch(ctx, clusterId)
	if body != nil {
		req = req.Branch(*body)
	} else {
		req = req.Branch(branch.Branch{})
	}
	resp, h, err := req.Execute()
	err = util.ParseError(err, h)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func checkBranchState(ctx context.Context, clusterId, branchId string, t *testing.T) (*branch.Branch, error) {
	ticker := time.NewTicker(time.Second * 10)
	timeout := time.After(time.Minute * 5)
	for {
		select {
		case <-ticker.C:
			res, h, err := branchClient.BranchServiceAPI.BranchServiceGetBranch(ctx, clusterId, branchId).Execute()
			if util.ParseError(err, h) != nil {
				t.Logf("get branch failed: %s", util.ParseError(err, h).Error())
				continue
			}
			t.Logf("get branch with state %s", *res.State)
			if *res.State == branch.BRANCHSTATE_ACTIVE {
				return res, nil
			}
		case <-timeout:
			return nil, errors.New("create branch timeout")
		}
	}
}
