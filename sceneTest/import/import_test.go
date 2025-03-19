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

func TestParquetImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`ppp`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	body := &imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_PARQUET,
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_S3,
			S3: &imp.S3Source{
				Uri:      config.ImportS3ParquetURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     config.S3AccessKeyId,
					Secret: config.S3SecretAccessKey,
				},
			},
		},
	}
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	if body != nil {
		r = r.Body(*body)
	}
	i, resp, err := r.Execute()
	err = util.ParseError(err, resp)
	if err != nil {
		t.Fatal(err)
	}
	err = waitImport(ctx, *i.ImportId)
	if err != nil {
		t.Fatalf("import failed, importId: %s, error: %s", *i.ImportId, err.Error())
	}
	t.Log("import finished")
}

func TestSchemaCompressImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	body := &imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &imp.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_S3,
			S3: &imp.S3Source{
				Uri:      config.ImportS3SchemaCompressURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     config.S3AccessKeyId,
					Secret: config.S3SecretAccessKey,
				},
			},
		},
	}
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	if body != nil {
		r = r.Body(*body)
	}
	i, resp, err := r.Execute()
	err = util.ParseError(err, resp)
	if err != nil {
		t.Fatal(err)
	}

	err = waitImport(ctx, *i.ImportId)
	if err != nil {
		t.Fatalf("import failed, importId: %s, error: %s", *i.ImportId, err.Error())
	}
	t.Log("import finished")
}

func TestSchemaTypeMisMatchedImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	body := &imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &imp.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_S3,
			S3: &imp.S3Source{
				Uri:      config.ImportS3SchemaTypeMisMatchedURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     config.S3AccessKeyId,
					Secret: config.S3SecretAccessKey,
				},
			},
		},
	}
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	if body != nil {
		r = r.Body(*body)
	}
	i, resp, err := r.Execute()
	err = util.ParseError(err, resp)
	if err != nil {
		t.Fatal(err)
	}

	err = waitImport(ctx, *i.ImportId)
	err = expectFail(err, "failed to cast value as int(11) for column `name`")
	if err != nil {
		t.Fatalf("test failed, importId: %s, err: %s", *i.ImportId, err.Error())
	} else {
		t.Log("import failed as expected")
	}
}

func TestSchemaColumnNumberMismatchedImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	body := &imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &imp.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_S3,
			S3: &imp.S3Source{
				Uri:      config.ImportS3SchemaColumnNumberMismatchedURI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     config.S3AccessKeyId,
					Secret: config.S3SecretAccessKey,
				},
			},
		},
	}
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	if body != nil {
		r = r.Body(*body)
	}
	i, resp, err := r.Execute()
	err = util.ParseError(err, resp)
	if err != nil {
		t.Fatal(err)
	}

	err = waitImport(ctx, *i.ImportId)
	err = expectFail(err, "TiDB schema `test`.`a` doesn't have the default value for number")
	if err != nil {
		t.Fatalf("test failed, importId: %s, err: %s", *i.ImportId, err.Error())
	} else {
		t.Log("import failed as expected")
	}
}
