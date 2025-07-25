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

// TestParquetImport tests successful Parquet file import
func TestParquetImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	_, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS `test`.`ppp`")
	if err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting Parquet import test")

	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()

	// Build import request body
	body := imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_PARQUET,
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_S3,
			S3: &imp.S3Source{
				Uri:      cfg.Import.S3.ParquetURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
				},
			},
		},
	}

	// Execute import request
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	r = r.Body(body)

	importTask, resp, err := r.Execute()
	if err := util.ParseError(err, resp); err != nil {
		t.Fatalf("Failed to create Parquet import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, *importTask.ImportId); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", *importTask.ImportId, err)
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

	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()

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
			S3: &imp.S3Source{
				Uri:      cfg.Import.S3.SchemaCompressURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
				},
			},
		},
	}

	// Execute import request
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	r = r.Body(body)

	importTask, resp, err := r.Execute()
	if err := util.ParseError(err, resp); err != nil {
		t.Fatalf("Failed to create schema compress import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, *importTask.ImportId); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", *importTask.ImportId, err)
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

	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()

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
			S3: &imp.S3Source{
				Uri:      cfg.Import.S3.SchemaTypeMismatchedURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
				},
			},
		},
	}

	// Execute import request
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	r = r.Body(body)

	importTask, resp, err := r.Execute()
	if err := util.ParseError(err, resp); err != nil {
		t.Fatalf("Failed to create type mismatch import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, *importTask.ImportId); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "failed to cast value as int(11) for column `name`"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", *importTask.ImportId, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}

// TestSchemaColumnNumberMismatchedImport tests import with column number mismatch error
func TestSchemaColumnNumberMismatchedImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting schema column number mismatch import test")

	// Set timeout for import creation
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()

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
			S3: &imp.S3Source{
				Uri:      cfg.Import.S3.SchemaColumnNumberMismatchedURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
				},
			},
		},
	}

	// Execute import request
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	r = r.Body(body)

	importTask, resp, err := r.Execute()
	if err := util.ParseError(err, resp); err != nil {
		t.Fatalf("Failed to create column mismatch import: %v", err)
	}

	// Wait for import and expect failure
	if err := waitImport(ctx, *importTask.ImportId); err != nil {
		// Check if failure is expected
		if expectErr := expectFail(err, "TiDB schema `test`.`a` doesn't have the default value for number"); expectErr != nil {
			t.Fatalf("Test failed, importId: %s, err: %v", *importTask.ImportId, expectErr)
		}
		t.Log("Import failed as expected")
		return
	}

	t.Fatal("Import should have failed but succeeded")
}
