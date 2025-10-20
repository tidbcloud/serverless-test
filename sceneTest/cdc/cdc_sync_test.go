package cdc

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tidbcloud/serverless-test/config"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

// TestMySQLSync tests if the MySQL changefeed can sync within 1 minute on alicloud-ap-southeast-1
func TestMySQLSync(t *testing.T) {
	ctx := context.Background()

	cfg := config.LoadConfig().Changefeed.MySQL
	t.Log(fmt.Sprintf("start to test mysql changefeed sync, changefeed_id: %s, cluster_id:%s, region:%s",
		cfg.ChangefeedID,
		cfg.ClusterID,
		cfg.Region))

	ts := time.Now().UnixMilli()
	t.Log("start to insert into upstream tidb cloud cluster")
	err := executeDB(ctx, cfg.ClusterDSN, fmt.Sprintf("insert into test.cdc (id, name) values (%d, 'cdc')", ts))
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
	// ctx := context.Background()

	// cfg := config.LoadConfig().Changefeed.Kafka
	// t.Log(fmt.Sprintf("start to test kafka changefeed sync, changefeed_id: %s, cluster_id:%s, region:%s",
	// 	cfg.ChangefeedID,
	// 	cfg.ClusterID,
	// 	cfg.Region))

	// ts := time.Now().UnixMilli()
	// t.Log("start to insert into upstream tidb cloud cluster")
	// err := executeDB(ctx, cfg.ClusterDSN, fmt.Sprintf("insert into test.cdc (id, name) values (%d, 'cdc')", ts))
	// if err != nil {
	// 	t.Fatalf("failed to insert into upstream tidb cloud cluster: %v", err)
	// }

	// t.Log("wait for 1 minute for data to sync")
	// time.Sleep(1 * time.Minute)

	// t.Log("start to check kafka sync")
}

func executeDB(ctx context.Context, dsn string, query string) (err error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		fmt.Printf("Failed to parse DSN %s error:%s\n", dsn, err.Error())
		return err
	}
	host := strings.Split(cfg.Addr, ":")[0]

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: host,
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
		fmt.Printf("No rows affected\n")
		return errors.New("no rows affected")
	}
	return nil
}

func queryDB(ctx context.Context, dsn string, query string) (exist bool, err error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		fmt.Printf("Failed to parse DSN %s error:%s\n", dsn, err.Error())
		return false, err
	}
	host := strings.Split(cfg.Addr, ":")[0]

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: host,
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
