package imp

import (
	"context"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/tidbcloud/pkgv2/log"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
	"go.uber.org/zap"
)

func TestParquetImport(t *testing.T) {
	ctx := context.Background()
	logger := log.WithContextL(ctx).With(zap.String("test", "e2eParquetImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS ppp")
	if err != nil {
		t.Fatal("failed to drop table -> ", err)
	}

	logger.Info("start import")
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
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
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
		t.Fatal("import failed -> ", err)
	}
	logger.Info("import finished")
}

func TestSchemaCompressImport(t *testing.T) {
	ctx := context.Background()
	logger := log.WithContextL(ctx).With(zap.String("test", "e2eSchemaCompressImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS a")
	if err != nil {
		t.Fatal("failed to drop table -> ", err)
	}

	logger.Info("start import")
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
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
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
		t.Fatal("import failed -> ", err)
	}
	logger.Info("import finished")
}

func TestSchemaTypeMisMatchedImport(t *testing.T) {
	ctx := context.Background()
	logger := log.WithContextL(ctx).With(zap.String("test", "e2eSchemaTypeMisMatchedImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS a")
	if err != nil {
		t.Fatal("failed to drop table -> ", err)
	}

	logger.Info("start import")
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
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
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
	expectFail(err, "failed to cast value as int(11) for column `name`", logger, t)
}

func TestSchemaColumnNumberMismatchedImport(t *testing.T) {
	ctx := context.Background()
	logger := log.WithContextL(ctx).With(zap.String("test", "e2eSchemaColumnNumberMismatchedImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS a")
	if err != nil {
		t.Fatal("failed to drop table -> ", err)
	}

	logger.Info("start import")
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
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
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
	expectFail(err, "TiDB schema `test`.`a` doesn't have the default value for number", logger, t)
}
