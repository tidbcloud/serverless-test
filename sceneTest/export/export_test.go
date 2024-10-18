package export

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/shiyuhang0/serverless-scene-test/client"
	"github.com/shiyuhang0/serverless-scene-test/config"
	"github.com/shiyuhang0/serverless-scene-test/pkg/tidbcloud/v1beta1/export"
	"github.com/stretchr/testify/assert"
)

var (
	clusterId   string
	exportId    string
	cloudClient *client.ClientDelegate
)

func init() {
	flag.StringVar(&clusterId, "cid", "", "")
	flag.StringVar(&exportId, "eid", "", "")
}

func setup() {
	var err error
	flag.Parse()
	cloudClient, err = config.GetClient()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func TestCreateExport(t *testing.T) {
	ctx := context.Background()
	t.Log("start create export")
	exp, err := cloudClient.CreateExport(ctx, clusterId, export.NewExportServiceCreateExportBody())
	if err != nil {
		t.Fatal(err)
	}
	exp = checkServerlessExportState(ctx, t, clusterId, *exp.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
}

func TestCreateExportToS3(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_S3
	s3AccessKeyId, s3SecretKeyId := config.GetS3AccessKey()
	exportS3Uri := config.GetS3URI()
	if s3AccessKeyId == "" || s3SecretKeyId == "" || exportS3Uri == "" {
		t.Fatal("s3 access key or secret key or uri is empty")
	}

	body := export.NewExportServiceCreateExportBody()
	body.Target = &export.ExportTarget{
		Type: &exportType,
		S3: &export.S3Target{
			Uri:       &exportS3Uri,
			AuthType:  export.EXPORTS3AUTHTYPEENUM_ACCESS_KEY,
			AccessKey: export.NewS3TargetAccessKey(s3AccessKeyId, s3SecretKeyId),
		},
	}

	t.Log("start create export")
	exp, err := cloudClient.CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp = checkServerlessExportState(ctx, t, clusterId, *exp.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.Target.S3.Uri, exportS3Uri)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func checkServerlessExportState(ctx context.Context, t *testing.T, clusterId, exportId string) *export.Export {
	ticker := time.NewTicker(time.Second * 10)
	timeout := time.After(time.Minute * 3)
	t.Log("start check export state")
	for {
		select {
		case <-ticker.C:
			exp, err := cloudClient.GetExport(ctx, clusterId, exportId)
			if err != nil {
				t.Logf("get export failed: %s", err.Error())
				continue
			}
			if *exp.State != export.EXPORTSTATEENUM_RUNNING {
				t.Logf("export finished with state %s", *exp.State)
				return exp
			}
		case <-timeout:
			t.Fatal("export timeout")
		}
	}
}
