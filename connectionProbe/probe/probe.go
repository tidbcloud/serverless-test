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

const probeTimeoutSec = 10

func ProbeDB(ctx context.Context, db *DBConfig, notifyCh chan<- *NotifyInfo, jobIdex int) (err error) {
	start := time.Now()
	fmt.Printf("Job %d start probing %s(%d) at %s\n", jobIdex, db.ClusterID, db.Port, start.Format("2006-01-02 15:04:05"))
	defer func() {
		latencyMS := time.Since(start).Milliseconds()
		now := time.Now().Format("2006-01-02 15:04:05")
		if err != nil {
			fmt.Printf("[%s] Probe failed: %s(%d) error:%v\n", now, db.ClusterID, db.Port, err)
			notifyCh <- &NotifyInfo{db, false, latencyMS, err.Error(), true, start.Format("2006-01-02 15:04:05"), now}
		} else {
			if latencyMS > probeTimeoutSec*1000 {
				fmt.Printf("[%s] Probe too much time: %s(%d)\n", now, db.ClusterID, db.Port)
				notifyCh <- &NotifyInfo{db, true, latencyMS, fmt.Sprintf("probe more than %ds", latencyMS/1000), true, start.Format("2006-01-02 15:04:05"), now}
			} else {
				fmt.Printf("[%s] Probe success: %s(%d)\n", now, db.ClusterID, db.Port)
				notifyCh <- &NotifyInfo{db, true, latencyMS, "", false, start.Format("2006-01-02 15:04:05"), now}
			}
		}
	}()

	mysql.RegisterTLSConfig(db.ClusterID, &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: db.Host,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/test?tls=%s&timeout=10s", db.User, db.Password, db.Host, db.Port, db.ClusterID)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open mysql connection: %s(%d) error:%s\n", db.ClusterID, db.Port, err.Error())
		return err
	}
	defer conn.Close()

	conn.SetMaxIdleConns(0)

	_, err = conn.Query("SHOW DATABASES;")
	return err
	// // probe the connection with a timeout context
	// cancelCtx, cancel := context.WithTimeout(ctx, probeTimeoutSec*time.Second)
	// defer cancel()
	// if err := conn.PingContext(cancelCtx); err != nil {
	// 	if cancelCtx.Err() == context.DeadlineExceeded {
	// 		return fmt.Errorf("connection timeout(%ds)", probeTimeoutSec)
	// 	}
	// 	return err
	// }
}
