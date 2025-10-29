package auditlg

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/auditlog"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestAuditLogGeneration(t *testing.T) {
	ctx := context.Background()
	log.Println("test audit log generation")

	cfg := config.LoadConfig().AuditLog
	t.Logf("start to test audit log generation, cluster_id:%s, region:%s",
		cfg.ClusterID,
		cfg.Region)

	file, err := getLatestAuditLogFile(ctx, cfg.ClusterID)
	if err != nil {
		t.Fatalf("failed to get audit log files: %v", err)
	}
	if file == nil {
		t.Fatalf("no audit log files found for cluster in last 24 hour %s", cfg.ClusterID)
	}
	if time.Since(*file.CreateTime) > 30*time.Minute {
		t.Fatalf("no audit log files found in last 30 minutes, the recent file generate at: %s", file.CreateTime.Format(time.DateTime))
	}
	log.Println(fmt.Sprintf("audit log file found at %s", file.CreateTime.Format(time.DateTime)))
}

func getLatestAuditLogFile(ctx context.Context, clusterID string) (*auditlog.AuditLogFile, error) {
	now := time.Now().UTC()

	today := now.Format("2006-01-02")
	resp, err := getAuditLogFiles(ctx, clusterID, today)
	if err != nil {
		return nil, err
	}
	if len(resp.AuditLogFiles) > 0 {
		return &resp.AuditLogFiles[len(resp.AuditLogFiles)-1], nil
	}

	yesterday := now.AddDate(0, 0, -1).UTC().Format("2006-01-02")
	resp, err = getAuditLogFiles(ctx, clusterID, yesterday)
	if err != nil {
		return nil, err
	}
	if len(resp.AuditLogFiles) > 0 {
		return &resp.AuditLogFiles[len(resp.AuditLogFiles)-1], nil
	}
	return nil, nil
}

func getAuditLogFiles(ctx context.Context, clusterID, date string) (*auditlog.ListAuditLogFilesResponse, error) {
	r := auditLogClient.DatabaseAuditLogServiceAPI.DatabaseAuditLogServiceListAuditLogFiles(ctx, clusterID).Date(date).PageSize(1000)
	res, h, err := r.Execute()
	return res, util.ParseError(err, h)
}
