package imp

import (
	"context"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
)

// TestOSSAccessKeyImport tests successful OSS import with valid access key
func TestOSSAccessKeyImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting OSS import test with access key")

	// Create and execute import
	importID, err := createOSSImport(ctx, false)
	if err != nil {
		t.Fatalf("Failed to create OSS import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("OSS import completed successfully")
}

// TestOSSAccessKeyNoPrivilegeImport tests OSS import with insufficient privileges
func TestOSSAccessKeyNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting OSS import test with no privilege access key")

	// Create import and expect failure
	_, err := createOSSImport(ctx, true)
	if err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "OSS access deny"); expectErr != nil {
			t.Fatalf("Create import failed, err: %v", expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}

// createOSSImport creates an OSS import with the specified configuration
func createOSSImport(ctx context.Context, useNoPrivilegeKey bool) (string, error) {
	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()

	// Determine which access key to use
	var accessKeyID, secretAccessKey string
	if useNoPrivilegeKey {
		accessKeyID = cfg.Import.OSS.AccessKeyIDNoPrivilege
		secretAccessKey = cfg.Import.OSS.SecretAccessKeyNoPrivilege
	} else {
		accessKeyID = cfg.Import.OSS.AccessKeyID
		secretAccessKey = cfg.Import.OSS.SecretAccessKey
	}

	// Build import request body
	body := imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &imp.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_OSS,
			Oss: &imp.OSSSource{
				Uri:      cfg.Import.OSS.URI,
				AuthType: imp.IMPORTOSSAUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.OSSSourceAccessKey{
					Id:     accessKeyID,
					Secret: secretAccessKey,
				},
			},
		},
	}

	// Execute import request
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	r = r.Body(body)

	importTask, resp, err := r.Execute()
	if err := util.ParseError(err, resp); err != nil {
		return "", err
	}

	return *importTask.ImportId, nil
}
