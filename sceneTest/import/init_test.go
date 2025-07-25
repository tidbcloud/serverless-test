package imp

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pingcap/log"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
	"go.uber.org/zap"
)

var (
	clusterId    string
	importClient *imp.APIClient
	db           *sql.DB
)

func init() {
	flag.StringVar(&clusterId, "cid", "", "")
}

func setup() {
	cfg := config.LoadConfig()
	var err error
	importClient, err = NewImportClient(cfg)
	if err != nil {
		panic(err)
	}
	db, err = NewDB(cfg)
	if err != nil {
		panic(err)
	}
}

func NewDB(cfg *config.Config) (*sql.DB, error) {
	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: cfg.Import.Cluster.Host,
	})
	if err != nil {
		log.Fatal("failed to register tls config -> ", zap.Error(err))
	}
	db, err = sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:4000)/test?tls=tidb",
		cfg.Import.Cluster.User, cfg.Import.Cluster.Password, cfg.Import.Cluster.Host),
	)
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)
	if err != nil {
		log.Fatal("failed to connect database -> ", zap.Error(err))
	}
	return db, nil
}

func NewImportClient(cfg *config.Config) (*imp.APIClient, error) {
	httpclient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(cfg.Endpoint.Serverless)
	if err != nil {
		return nil, err
	}
	importCfg := imp.NewConfiguration()
	importCfg.HTTPClient = httpclient
	importCfg.Host = serverlessURL.Host
	importCfg.UserAgent = util.UserAgent
	return imp.NewAPIClient(importCfg), nil
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func waitImport(ctx context.Context, importID string) error {
	// Wait for the import task to finish
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Minute)

	for {
		select {
		case <-ticker.C:
			getContext, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			r := importClient.ImportServiceAPI.ImportServiceGetImport(getContext, clusterId, importID)
			i, resp, err := r.Execute()
			err = util.ParseError(err, resp)
			if err != nil {
				return err
			}
			if *i.State == imp.IMPORTSTATEENUM_COMPLETED {
				if i.HasTotalSize() && strings.EqualFold(*i.TotalSize, "0") {
					return errors.New("import succeeded but no data imported")
				}
				return nil
			} else if *i.State == imp.IMPORTSTATEENUM_FAILED {
				if i.Message == nil {
					return errors.New("import failed")
				}
				return errors.New(*i.Message)
			} else if *i.State == imp.IMPORTSTATEENUM_CANCELING || *i.State == imp.IMPORTSTATEENUM_CANCELED {
				return errors.New("import task cancelled")
			}
		case <-timeout:
			return errors.New("timed out to wait import task complete")
		}
	}
}

func expectFail(err error, errorMsg string) error {
	if err != nil {
		if strings.Contains(err.Error(), errorMsg) {
			return nil
		}
		return fmt.Errorf("import failed, but not as expected. expected: %s, actual: %s", errorMsg, err.Error())
	}
	return errors.New("import should fail, but succeeded")
}
