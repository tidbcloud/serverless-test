package imp

import (
	"context"
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pingcap/log"
	"github.com/tidbcloud/serverless-test/config"
	consoleimportapi "github.com/tidbcloud/serverless-test/pkg/console/import"
	tidbcloudlogin "github.com/tidbcloud/serverless-test/pkg/login"
	"github.com/tidbcloud/serverless-test/util"
	"go.uber.org/zap"
)

var (
	clusterId    string
	projectId    string
	orgId        string
	importClient *consoleimportapi.APIClient
	db           *sql.DB
)

func init() {
	flag.StringVar(&clusterId, "cid", "", "")
	flag.StringVar(&projectId, "pid", "", "")
	flag.StringVar(&orgId, "oid", "", "")
}

func setup() {
	var err error
	config.InitializeConfig()
	importClient, err = NewImportConsoleClient()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	db, err = NewDB()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func NewDB() (*sql.DB, error) {
	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: config.ImportClusterHost,
	})
	if err != nil {
		log.Fatal("failed to register tls config -> ", zap.Error(err))
	}
	db, err = sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:4000)/test?tls=tidb",
		config.ImportClusterUser, config.ImportClusterPassWord, config.ImportClusterHost),
	)
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)
	if err != nil {
		log.Fatal("failed to connect database -> ", zap.Error(err))
	}
	return db, nil
}

func NewImportConsoleClient() (*consoleimportapi.APIClient, error) {
	c := tidbcloudlogin.WebApiLoginContext{
		Host:              config.ConsoleApiHost,
		Auth0Domain:       config.Auth0Domain,
		Auth0ClientID:     config.Auth0ClientID,
		Auth0ClientSecret: config.Auth0ClientSecret,
		UserEmail:         config.UserEmail,
	}
	token, err := c.LoginAndGetToken(context.Background())
	if err != nil {
		return nil, err
	}
	httpclient := &http.Client{
		Transport: util.NewBearerTransport(token),
	}
	importCfg := consoleimportapi.NewConfiguration()
	importCfg.HTTPClient = httpclient
	importCfg.Host = config.ConsoleApiHost
	importCfg.UserAgent = util.UserAgent
	importCfg.Scheme = "https"
	return consoleimportapi.NewAPIClient(importCfg), nil
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

//func waitImport(ctx context.Context, importID string) error {
//	// Wait for the import task to finish
//	ticker := time.NewTicker(10 * time.Second)
//	defer ticker.Stop()
//
//	timeout := time.After(5 * time.Minute)
//
//	for {
//		select {
//		case <-ticker.C:
//			getContext, cancel := context.WithTimeout(ctx, 30*time.Second)
//			defer cancel()
//			r := importClient.ImportServiceAPI.ImportServiceGetImport(getContext, clusterId, importID)
//			i, resp, err := r.Execute()
//			err = util.ParseError(err, resp)
//			if err != nil {
//				return err
//			}
//			if *i.State == imp.IMPORTSTATEENUM_COMPLETED {
//				if i.HasTotalSize() && strings.EqualFold(*i.TotalSize, "0") {
//					return errors.New("import succeeded but no data imported")
//				}
//				return nil
//			} else if *i.State == imp.IMPORTSTATEENUM_FAILED {
//				if i.Message == nil {
//					return errors.New("import failed")
//				}
//				return errors.New(*i.Message)
//			} else if *i.State == imp.IMPORTSTATEENUM_CANCELING || *i.State == imp.IMPORTSTATEENUM_CANCELED {
//				return errors.New("import task cancelled")
//			}
//		case <-timeout:
//			return errors.New("timed out to wait import task complete")
//		}
//	}
//}
//
//func expectFail(err error, errorMsg string) error {
//	if err != nil {
//		if strings.Contains(err.Error(), errorMsg) {
//			return nil
//		}
//		return fmt.Errorf("import failed, but not as expected. expected: %s, actual: %s", errorMsg, err.Error())
//	}
//	return errors.New("import should fail, but succeeded")
//}
