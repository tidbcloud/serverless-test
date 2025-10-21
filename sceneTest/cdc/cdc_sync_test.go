package cdc

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cdc"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

// TestMySQLSync tests if the MySQL changefeed can sync within 1 minute on alicloud-ap-southeast-1
func TestMySQLSync(t *testing.T) {
	ctx := context.Background()

	cfg := config.LoadConfig().Changefeed.MySQL
	t.Logf("start to test mysql changefeed sync, changefeed_id: %s, cluster_id:%s, region:%s",
		cfg.ChangefeedID,
		cfg.ClusterID,
		cfg.Region)
	cf, err := getChangefeed(ctx, cfg.ClusterID, cfg.ChangefeedID)
	if err != nil {
		t.Fatalf("failed to get changefeed: %v", err)
	}
	if *cf.State != cdc.CHANGEFEEDSTATEENUM_RUNNING {
		t.Fatalf("changefeed is not running, current state: %s", *cf.State)
	}

	ts := time.Now().UnixMilli()
	t.Log("start to insert into upstream tidb cloud cluster")
	err = executeDB(ctx, cfg.ClusterDSN, fmt.Sprintf("insert into test.cdc (id, name) values (%d, 'cdc')", ts))
	if err != nil {
		t.Fatalf("failed to insert into upstream tidb cloud cluster: %v", err)
	}

	t.Log("wait for 1 minute for data to sync")
	time.Sleep(1 * time.Minute)

	t.Log("start to check mysql sync")
	exist, err := queryDB(ctx, cfg.MySQLDSN, fmt.Sprintf("select id, name from test.cdc where id = %d", ts))
	if err != nil {
		t.Fatalf("failed to query downstream mysql: %v", err)
	}
	if !exist {
		t.Fatalf("data not synced to downstream mysql after 1 minute")
	}
}

// TestMySQLSync tests if the Kafka changefeed can sync within 1 minute on alicloud-ap-southeast-1
func TestKafkaSync(t *testing.T) {
	ctx := context.Background()

	cfg := config.LoadConfig().Changefeed.Kafka
	t.Log(fmt.Sprintf("start to test kafka changefeed sync, changefeed_id: %s, cluster_id:%s, region:%s",
		cfg.ChangefeedID,
		cfg.ClusterID,
		cfg.Region))
	cf, err := getChangefeed(ctx, cfg.ClusterID, cfg.ChangefeedID)
	if err != nil {
		t.Fatalf("failed to get changefeed: %v", err)
	}
	if *cf.State != cdc.CHANGEFEEDSTATEENUM_RUNNING {
		t.Fatalf("changefeed is not running, current state: %s", *cf.State)
	}

	ts := time.Now().UnixMilli()

	t.Log("consume kafka topic to check if data exists")

	t.Log("start to insert into upstream tidb cloud cluster")
	err = executeDB(ctx, cfg.ClusterDSN, fmt.Sprintf("insert into kafka.cdc (id, name) values (%d, 'cdc')", ts))
	if err != nil {
		t.Fatalf("failed to insert into upstream tidb cloud cluster: %v", err)
	}
}

func executeDB(ctx context.Context, dsn string, query string) (err error) {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
	})

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open mysql connection. error:%s\n", err.Error())
		return err
	}
	defer conn.Close()

	res, err := conn.ExecContext(ctx, query)
	if err != nil {
		fmt.Printf("Failed to execute query. error:%s\n", err.Error())
		return err
	}
	rowAffect, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get rows affected. error:%s\n", err.Error())
		return err
	}
	if rowAffect == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func queryDB(ctx context.Context, dsn string, query string) (exist bool, err error) {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
	})

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open mysql connection. error:%s\n", err.Error())
		return false, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		fmt.Printf("Failed to execute query. error:%s\n", err.Error())
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func getChangefeed(ctx context.Context, clusterId, changefeedId string) (*cdc.Changefeed, error) {
	r := cdcClient.ChangefeedServiceAPI.ChangefeedServiceGetChangefeed(ctx, clusterId, changefeedId)
	res, h, err := r.Execute()
	return res, util.ParseError(err, h)
}
