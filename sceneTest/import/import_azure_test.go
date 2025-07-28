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

// TestAzureImport tests successful Azure Blob import with valid SAS token
func TestAzureImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting Azure import test")

	// Create and execute import
	importID, err := createAzureImport(ctx, false)
	if err != nil {
		t.Fatalf("Failed to create Azure import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("Azure import completed successfully")
}

// TestAzureNoPrivilegeImport tests Azure Blob import with insufficient privileges
func TestAzureNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting Azure import test with no privilege")

	// Create and execute import with no privilege token
	importID, err := createAzureImport(ctx, true)
	if err != nil {
		t.Fatalf("Failed to create Azure import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, importID); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "AzBlobAccessDenied"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", importID, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}

// createAzureImport creates an Azure Blob import with the specified configuration
func createAzureImport(ctx context.Context, useNoPrivilegeToken bool) (string, error) {
	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()

	// Determine which SAS token to use
	var sasToken string
	if useNoPrivilegeToken {
		sasToken = cfg.Import.Azure.SASTokenNoPrivilege
	} else {
		sasToken = cfg.Import.Azure.SASToken
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
			Type: imp.IMPORTSOURCETYPEENUM_AZURE_BLOB,
			AzureBlob: &imp.AzureBlobSource{
				Uri:      cfg.Import.Azure.URI,
				AuthType: imp.IMPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
				SasToken: &sasToken,
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
