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

func TestOSSAccessKeyNoPrivilegeImport(t *testing.T) {
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
			Type: imp.IMPORTSOURCETYPEENUM_OSS,
			Oss: &imp.OSSSource{
				Uri:      config.ImportOSSURI,
				AuthType: imp.IMPORTOSSAUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.OSSSourceAccessKey{
					Id:     config.ImportOSSAccessKeyIdNoPrivilege,
					Secret: config.ImportOSSSecretAccessKeyNoPrivilege,
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

func TestOSSAccessKeyImport(t *testing.T) {
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
			Type: imp.IMPORTSOURCETYPEENUM_OSS,
			Oss: &imp.OSSSource{
				Uri:      config.ImportOSSURI,
				AuthType: imp.IMPORTOSSAUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &imp.OSSSourceAccessKey{
					Id:     config.ImportOSSAccessKeyId,
					Secret: config.ImportOSSSecretAccessKey,
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
