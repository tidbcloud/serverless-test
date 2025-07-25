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

// TestGcsImport tests successful Google Cloud Storage import with valid service account key
func TestGcsImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting GCS import test")

	// Create and execute import
	importID, err := createGcsImport(ctx, false)
	if err != nil {
		t.Fatalf("Failed to create GCS import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("GCS import completed successfully")
}

// TestGcsNoPrivilegeImport tests GCS import with insufficient privileges
func TestGcsNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting GCS import test with no privilege")

	// Create and execute import with no privilege service account key
	importID, err := createGcsImport(ctx, true)
	if err != nil {
		t.Fatalf("Failed to create GCS import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, importID); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "Permission 'storage.objects.list' denied on resource"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", importID, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}

// createGcsImport creates a Google Cloud Storage import with the specified configuration
func createGcsImport(ctx context.Context, useNoPrivilegeKey bool) (string, error) {
	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()

	// Determine which service account key to use
	var serviceAccountKey string
	if useNoPrivilegeKey {
		serviceAccountKey = cfg.Import.GCS.ServiceAccountKeyNoPrivilege
	} else {
		serviceAccountKey = cfg.Import.GCS.ServiceAccountKey
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
			Type: imp.IMPORTSOURCETYPEENUM_GCS,
			Gcs: &imp.GCSSource{
				Uri:               cfg.Import.GCS.URI,
				AuthType:          imp.IMPORTGCSAUTHTYPEENUM_SERVICE_ACCOUNT_KEY,
				ServiceAccountKey: &serviceAccountKey,
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
