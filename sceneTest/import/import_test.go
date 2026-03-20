package imp

import (
	"context"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
)

// TestParquetImport tests successful Parquet file import
func TestParquetImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	_, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS `test`.`ppp`")
	if err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting Parquet import test")

	cfg := config.LoadConfig()

	importOptions := imp.ImportOptions{FileType: imp.IMPORTFILETYPEENUM_PARQUET}
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.ParquetURI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
		AccessKey: &imp.S3SourceAccessKey{
			Id:     cfg.S3.AccessKeyID,
			Secret: cfg.S3.SecretAccessKey,
		},
	}

	importID, err := createS3Import(ctx, importOptions, s3Source)
	if err != nil {
		t.Fatalf("Failed to create Parquet import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("Parquet import completed successfully")
}

// TestSchemaCompressImport tests successful CSV import with schema compression
func TestSchemaCompressImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting schema compress import test")

	cfg := config.LoadConfig()

	importOptions := imp.ImportOptions{
		FileType: imp.IMPORTFILETYPEENUM_CSV,
		CsvFormat: &imp.CSVFormat{
			Separator: pointer.ToString(";"),
		},
	}
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.SchemaCompressURI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
		AccessKey: &imp.S3SourceAccessKey{
			Id:     cfg.S3.AccessKeyID,
			Secret: cfg.S3.SecretAccessKey,
		},
	}

	importID, err := createS3Import(ctx, importOptions, s3Source)
	if err != nil {
		t.Fatalf("Failed to create schema compress import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("Schema compress import completed successfully")
}

// TestSchemaTypeMismatchedImport tests import with type mismatch error
func TestSchemaTypeMismatchedImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting schema type mismatch import test")

	cfg := config.LoadConfig()

	importOptions := imp.ImportOptions{
		FileType: imp.IMPORTFILETYPEENUM_CSV,
		CsvFormat: &imp.CSVFormat{
			Separator: pointer.ToString(";"),
		},
	}
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.SchemaTypeMismatchedURI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
		AccessKey: &imp.S3SourceAccessKey{
			Id:     cfg.S3.AccessKeyID,
			Secret: cfg.S3.SecretAccessKey,
		},
	}

	importID, err := createS3Import(ctx, importOptions, s3Source)
	if err != nil {
		t.Fatalf("Failed to create type mismatch import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, importID); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "CastValueError: when encoding 1-th data row in file test.a.csv:0"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", importID, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatalf("Import %s should have failed but succeeded", importID)
}

// TestSchemaColumnNumberMismatchedImport tests import with column number mismatch error
func TestSchemaColumnNumberMismatchedImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting schema column number mismatch import test")

	cfg := config.LoadConfig()

	importOptions := imp.ImportOptions{
		FileType: imp.IMPORTFILETYPEENUM_CSV,
		CsvFormat: &imp.CSVFormat{
			Separator: pointer.ToString(";"),
		},
	}
	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.SchemaColumnNumberMismatchedURI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
		AccessKey: &imp.S3SourceAccessKey{
			Id:     cfg.S3.AccessKeyID,
			Secret: cfg.S3.SecretAccessKey,
		},
	}

	importID, err := createS3Import(ctx, importOptions, s3Source)
	if err != nil {
		t.Fatalf("Failed to create column mismatch import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, importID); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "TiDB schema `test`.`a` doesn't have the default value for number"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", importID, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatalf("Import %s should have failed but succeeded", importID)
}

// TestZeroDateImport verifies zero date strings can be imported without errors
func TestZeroDateImport(t *testing.T) {
	ctx := context.Background()

	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS `test`.`zero_date`"); err != nil {
		t.Fatalf("Failed to drop zero date import table: %v", err)
	}

	cfg := config.LoadConfig()

	importOptions := imp.ImportOptions{
		FileType: imp.IMPORTFILETYPEENUM_CSV,
		CsvFormat: &imp.CSVFormat{
			Separator: pointer.ToString(";"),
		},
	}

	s3Source := &imp.S3Source{
		Uri:      cfg.Import.S3.ZeroDateURI,
		AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
		AccessKey: &imp.S3SourceAccessKey{
			Id:     cfg.S3.AccessKeyID,
			Secret: cfg.S3.SecretAccessKey,
		},
	}

	importID, err := createS3Import(ctx, importOptions, s3Source)
	if err != nil {
		t.Fatalf("Failed to create zero date import: %v", err)
	}

	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Zero date import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("Zero date import completed successfully")
}
