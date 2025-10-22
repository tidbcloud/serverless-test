package auditlg

import (
	"context"
	"log"
	"testing"

	"github.com/tidbcloud/serverless-test/config"
)

const (
	clusterID = "cluster-1nv3z3t2"
	region    = ""
)

func TestAuditLogGeneration(t *testing.T) {
	ctx := context.Background()
	log.Println("test audit log generation")

	cfg := config.LoadConfig().AuditLog
	t.Logf("start to test audit log generation, cluster_id:%s, region:%s",
		cfg.ClusterID,
		cfg.Region)

}
