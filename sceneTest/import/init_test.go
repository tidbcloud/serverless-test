package imp

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pingcap/log"
	"github.com/spf13/pflag"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
	"go.uber.org/zap"
)

// Test state variables
var (
	clusterId    string
	importClient *imp.APIClient
	db           *sql.DB
)

func init() {
	pflag.StringVar(&clusterId, "cid", "", "Cluster ID for import tests")
}

// setup initializes the test environment by creating API clients and database connection
func setup() {
	cfg := config.LoadConfig()

	// Initialize import client
	var err error
	importClient, err = NewImportClient(cfg)
	if err != nil {
		log.Fatal("Failed to create import client", zap.Error(err))
	}

	// Initialize database connection
	db, err = NewDB(cfg)
	if err != nil {
		log.Fatal("Failed to create database connection", zap.Error(err))
	}
}

// NewDB creates a new database connection with TLS configuration
func NewDB(cfg *config.Config) (*sql.DB, error) {
	// Register TLS configuration for TiDB
	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: cfg.Import.ClusterHost,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register TLS config: %w", err)
	}

	// Create database connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:4000)/test?tls=tidb",
		cfg.Import.ClusterUser,
		cfg.Import.ClusterPassword,
		cfg.Import.ClusterHost,
	)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	return db, nil
}

// NewImportClient creates a new import API client with the given configuration
func NewImportClient(cfg *config.Config) (*imp.APIClient, error) {
	// Create HTTP client with digest authentication
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}

	// Validate and parse serverless endpoint
	serverlessURL, err := util.ValidateApiUrl(cfg.Endpoint.Serverless)
	if err != nil {
		return nil, fmt.Errorf("invalid serverless endpoint: %w", err)
	}

	// Configure import API client
	importCfg := imp.NewConfiguration()
	importCfg.HTTPClient = httpClient
	importCfg.Host = serverlessURL.Host
	importCfg.UserAgent = util.UserAgent

	return imp.NewAPIClient(importCfg), nil
}

// TestMain runs before all tests and sets up the test environment
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

// cleanupTestTable removes the test table if it exists
func cleanupTestTable(ctx context.Context) error {
	_, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS `test`.`a`")
	return err
}

// waitImport waits for an import task to complete and returns an error if it fails
func waitImport(ctx context.Context, importID string) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Minute)

	for {
		select {
		case <-ticker.C:
			// Get import status
			getContext, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			r := importClient.ImportServiceAPI.ImportServiceGetImport(getContext, clusterId, importID)
			importTask, resp, err := r.Execute()
			if err := util.ParseError(err, resp); err != nil {
				return fmt.Errorf("failed to get import status: %w", err)
			}

			// Check import state
			switch *importTask.State {
			case imp.IMPORTSTATEENUM_COMPLETED:
				if importTask.HasTotalSize() && strings.EqualFold(*importTask.TotalSize, "0") {
					return errors.New("import succeeded but no data was imported")
				}
				return nil
			case imp.IMPORTSTATEENUM_FAILED:
				if importTask.Message == nil {
					return errors.New("import failed with no error message")
				}
				return errors.New(*importTask.Message)
			case imp.IMPORTSTATEENUM_CANCELING, imp.IMPORTSTATEENUM_CANCELED:
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

// expectFail validates that an error contains the expected error message
func expectFail(err error, expectedErrorMsg string) error {
	if err == nil {
		return errors.New("import should have failed, but succeeded")
	}

	if strings.Contains(err.Error(), expectedErrorMsg) {
		return nil
	}

	return fmt.Errorf("import failed, but not as expected. expected: %s, actual: %s",
		expectedErrorMsg, err.Error())
}
