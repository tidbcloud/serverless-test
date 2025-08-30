package probe

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
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

func ProbeDB(ctx context.Context, db *DBConfig, notifyCh chan<- *NotifyInfo) (err error) {
	defer func() {
		if err != nil {
			fmt.Printf("Probe failed: %s(%d) error:%s\n", db.ClusterID, db.Port, err.Error())
			notifyCh <- &NotifyInfo{db, false, err.Error()}
		} else {
			fmt.Printf("Probe success: %s(%d)\n", db.ClusterID, db.Port)
			notifyCh <- &NotifyInfo{db, true, ""}
		}
	}()

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: db.Host,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/test?tls=tidb&timeout=5s", db.User, db.Password, db.Host, db.Port)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer conn.Close()

	// probe the connection with a timeout context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return errors.New("connection timeout(5s)")
		}
		return err
	}
	return nil
}
