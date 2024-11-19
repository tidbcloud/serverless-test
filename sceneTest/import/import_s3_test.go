package imp

import (
	"context"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/pingcap/log"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
	"go.uber.org/zap"
)

func TestS3ArnNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eS3ArnNoPrivilegeImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatal("failed to drop table -> ", zap.Error(err))
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
				Uri:      config.ImportS3URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArnNoPrivilege,
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
	err = expectFail(err, "is not authorized to perform: s3:ListBucket")
	if err != nil {
		t.Fatal("test failed -> ", zap.Error(err), zap.String("importId", *i.ImportId))
	} else {
		logger.Info("import failed as expected")
	}
}

func TestS3ArnDiffExternalIDImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eS3ArnDiffExternalIDImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatal("failed to drop table -> ", zap.Error(err))
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
				Uri:      config.ImportS3URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArnDiffExternalID,
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
	err = expectFail(err, "is not authorized to perform: sts:AssumeRole on resource")
	if err != nil {
		t.Fatal("test failed -> ", zap.Error(err), zap.String("importId", *i.ImportId))
	} else {
		logger.Info("import failed as expected")
	}
}

func TestS3AccessKeyNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eS3AccessKeyNoPrivilegeImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatal("failed to drop table -> ", zap.Error(err))
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
				Uri:      config.ImportS3URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     config.ImportS3AccessKeyIdNoPrivilege,
					Secret: config.ImportS3SecretAccessKeyNoPrivilege,
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
	err = expectFail(err, "AccessDenied")
	if err != nil {
		t.Fatal("test failed -> ", zap.Error(err), zap.String("importId", *i.ImportId))
	} else {
		logger.Info("import failed as expected")
	}
}

func TestS3AccessKeyImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eS3AccessKeyImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatal("failed to drop table -> ", zap.Error(err))
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
				Uri:      config.ImportS3URI,
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
		t.Fatal("import failed -> ", zap.Error(err), zap.String("importId", *i.ImportId))
	}
	logger.Info("import finished")
}

func TestS3ArnImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eS3ArnImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatal("failed to drop table -> ", zap.Error(err))
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
				Uri:      config.ImportS3URI,
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
		t.Fatal("import failed -> ", zap.Error(err), zap.String("importId", *i.ImportId))
	}
	logger.Info("import finished")
}
