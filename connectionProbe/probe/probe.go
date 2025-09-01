package probe

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Name      string `yaml:"name"`
	ClusterID string `yaml:"cluster_id"`
	Host      string `yaml:"host"`
	User      string `yaml:"user"`
	Port      int    `yaml:"port"`
	Region    string `yaml:"region"`
	Plan      string `yaml:"plan"`
	Password  string `yaml:"password"`
	TiDBPool  string `yaml:"tidb_pool"`
}

const probeTimeoutSec = 8

func ProbeDB(ctx context.Context, db *DBConfig, notifyCh chan<- *NotifyInfo) (err error) {
	start := time.Now()
	defer func() {
		latencyMS := time.Since(start).Milliseconds()
		now := time.Now().Format("2006-01-02 15:04:05")
		if err != nil {
			fmt.Printf("[%s] Probe failed: %s(%d) start time: %s error:%s\n", now, db.ClusterID, db.Port, start.Format("2006-01-02 15:04:05"), err.Error())
			notifyCh <- &NotifyInfo{db, false, latencyMS, err.Error()}
		} else {
			fmt.Printf("[%s] Probe success: %s(%d)\n", now, db.ClusterID, db.Port)
			notifyCh <- &NotifyInfo{db, true, latencyMS, ""}
		}
	}()

	mysql.RegisterTLSConfig(db.ClusterID, &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: db.Host,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/test?tls=%s&timeout=%ds", db.User, db.Password, db.Host, db.Port, db.ClusterID, probeTimeoutSec)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open mysql connection: %s(%d) error:%s\n", db.ClusterID, db.Port, err.Error())
		return err
	}
	defer conn.Close()

	conn.SetMaxIdleConns(0)

	// probe the connection with a timeout context
	cancelCtx, cancel := context.WithTimeout(ctx, probeTimeoutSec*time.Second)
	defer cancel()
	if err := conn.PingContext(cancelCtx); err != nil {
		if cancelCtx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("connection timeout(%ds)", probeTimeoutSec)
		}
		return err
	}
	return nil
}
