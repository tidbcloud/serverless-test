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

func TestS3ArnNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()
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
				Uri:      cfg.Import.S3.URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.Import.S3.RoleARNNoPrivilege,
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
		t.Fatalf("test failed, importId: %s, err: %s", *i.ImportId, err.Error())
	} else {
		t.Log("import failed as expected")
	}
}

func TestS3ArnDiffExternalIDImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()
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
				Uri:      cfg.Import.S3.URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.Import.S3.RoleARNDiffExternalID,
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
		t.Fatalf("test failed, importId: %s, err: %s", *i.ImportId, err.Error())
	} else {
		t.Log("import failed as expected")
	}
}

func TestS3AccessKeyNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()
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
				Uri:      cfg.Import.S3.URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     cfg.Import.S3.AccessKeyIDNoPrivilege,
					Secret: cfg.Import.S3.SecretAccessKeyNoPrivilege,
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
		t.Fatalf("test failed, importId: %s, err: %s", *i.ImportId, err.Error())
	} else {
		t.Log("import failed as expected")
	}
}

func TestS3AccessKeyImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()
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
				Uri:      cfg.Import.S3.URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
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

func TestS3ArnImport(t *testing.T) {
	ctx := context.Background()
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	if err != nil {
		t.Fatalf("failed to drop table, err: %s", err.Error())
	}

	t.Log("start import")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()
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
				Uri:      cfg.Import.S3.URI,
				AuthType: imp.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.Import.S3.RoleARN,
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
