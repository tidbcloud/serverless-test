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
	cfg := config.LoadConfig()

	var err error
	importClient, err = NewImportConsoleClient(cfg)
	if err != nil {
		log.Fatal("failed to create import client", zap.Error(err))
	}

	db, err = NewDB(cfg)
	if err != nil {
		log.Fatal("failed to create database connection", zap.Error(err))
	}
}

func NewDB(cfg *config.Config) (*sql.DB, error) {
	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: cfg.Import.ClusterHost,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register tls config: %w", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:4000)/test?tls=tidb",
		cfg.Import.ClusterUser, cfg.Import.ClusterPassword, cfg.Import.ClusterHost,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	return db, nil
}

func NewImportConsoleClient(cfg *config.Config) (*consoleimportapi.APIClient, error) {
	loginCtx := tidbcloudlogin.WebApiLoginContext{
		Host:              cfg.ConsoleAPIHost,
		Auth0Domain:       cfg.Auth0Domain,
		Auth0ClientID:     cfg.Auth0ClientID,
		Auth0ClientSecret: cfg.Auth0ClientSecret,
		UserEmail:         cfg.UserEmail,
	}

	token, err := loginCtx.LoginAndGetToken(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to login and get token: %w", err)
	}

	httpClient := &http.Client{
		Transport: util.NewBearerTransport(token),
	}

	importCfg := consoleimportapi.NewConfiguration()
	importCfg.HTTPClient = httpClient
	importCfg.Host = cfg.ConsoleAPIHost
	importCfg.UserAgent = util.UserAgent
	importCfg.Scheme = "https"

	return consoleimportapi.NewAPIClient(importCfg), nil
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func waitImport(ctx context.Context, importID string) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Minute)

	for {
		select {
		case <-ticker.C:
			getContext, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			r := importClient.ImportServiceAPI.ImportServiceGetImport(getContext, orgId, projectId, clusterId, importID)
			importTask, resp, err := r.Execute()
			if err := util.ParseError(err, resp); err != nil {
				return fmt.Errorf("failed to get import status: %w", err)
			}

			switch *importTask.State {
			case consoleimportapi.IMPORTSTATEENUM_COMPLETED:
				if importTask.HasTotalSize() && strings.EqualFold(*importTask.TotalSize, "0") {
					return errors.New("import succeeded but no data was imported")
				}
				return nil
			case consoleimportapi.IMPORTSTATEENUM_FAILED:
				if importTask.Message == nil {
					return errors.New("import failed with no error message")
				}
				return errors.New(*importTask.Message)
			case consoleimportapi.IMPORTSTATEENUM_CANCELING, consoleimportapi.IMPORTSTATEENUM_CANCELED:
				return errors.New("import task was cancelled")
			default:
				// Import is still in progress, continue waiting
				continue
			}
		case <-timeout:
			return errors.New("import task timed out after 5 minutes")
		}
	}
}
