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

// TestS3ArnImport tests successful S3 import with valid role ARN
func TestS3ArnImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting S3 import test with role ARN")

	// Build S3 source with valid role ARN
	cfg := config.LoadConfig()
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.URI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
		RoleArn:  &cfg.Import.S3.RoleARN,
	}

	// Create and execute import
	importID, err := createS3Import(ctx, s3Source)
	if err != nil {
		t.Fatalf("Failed to create S3 import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("S3 import completed successfully")
}

// TestS3AccessKeyImport tests successful S3 import with valid access key
func TestS3AccessKeyImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting S3 import test with access key")

	// Build S3 source with valid access key
	cfg := config.LoadConfig()
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.URI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
		AccessKey: &imp.S3SourceAccessKey{
			Id:     cfg.S3.AccessKeyID,
			Secret: cfg.S3.SecretAccessKey,
		},
	}

	// Create and execute import
	importID, err := createS3Import(ctx, s3Source)
	if err != nil {
		t.Fatalf("Failed to create S3 import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("S3 import completed successfully")
}

// TestS3ArnNoPrivilegeImport tests S3 import with insufficient role privileges
func TestS3ArnNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting S3 import test with no privilege role ARN")

	// Build S3 source with no privilege role ARN
	cfg := config.LoadConfig()
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.URI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
		RoleArn:  &cfg.Import.S3.RoleARNNoPrivilege,
	}

	// Create import and expect failure
	importID, err := createS3Import(ctx, s3Source)
	if err != nil {
		t.Fatalf("Failed to create S3 import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, importID); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "is not authorized to perform: s3:ListBucket"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", importID, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}

// TestS3ArnDiffExternalIDImport tests S3 import with different external ID
func TestS3ArnDiffExternalIDImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting S3 import test with different external ID")

	// Build S3 source with different external ID role ARN
	cfg := config.LoadConfig()
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.URI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
		RoleArn:  &cfg.Import.S3.RoleARNDiffExternalID,
	}

	// Create import and expect failure
	importID, err := createS3Import(ctx, s3Source)
	if err != nil {
		t.Fatalf("Failed to create S3 import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, importID); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "is not authorized to perform: sts:AssumeRole on resource"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", importID, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}

// TestS3AccessKeyNoPrivilegeImport tests S3 import with insufficient access key privileges
func TestS3AccessKeyNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting S3 import test with no privilege access key")

	// Build S3 source with no privilege access key
	cfg := config.LoadConfig()
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.URI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
		AccessKey: &imp.S3SourceAccessKey{
			Id:     cfg.Import.S3.AccessKeyIDNoPrivilege,
			Secret: cfg.Import.S3.SecretAccessKeyNoPrivilege,
		},
	}

	// Create import and expect failure
	importID, err := createS3Import(ctx, s3Source)
	if err != nil {
		t.Fatalf("Failed to create S3 import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, importID); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "AccessDenied"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", importID, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}

// createS3Import creates an S3 import with the provided S3 source configuration
func createS3Import(ctx context.Context, s3Source *imp.S3Source) (string, error) {
	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	// Build import request body
	body := imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &imp.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_S3,
			S3:   s3Source,
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
