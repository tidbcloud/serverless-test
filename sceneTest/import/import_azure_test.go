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

func TestAzureImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eAzureImport"))
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
			Type: imp.IMPORTSOURCETYPEENUM_AZURE_BLOB,
			AzureBlob: &imp.AzureBlobSource{
				Uri:      config.ImportAzureURI,
				AuthType: imp.IMPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
				SasToken: &config.AzureSASToken,
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

func TestAzureNoPrivilegeImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eAzureNoPrivilegeImport"))
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
			Type: imp.IMPORTSOURCETYPEENUM_AZURE_BLOB,
			AzureBlob: &imp.AzureBlobSource{
				Uri:      config.ImportAzureURI,
				AuthType: imp.IMPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
				SasToken: &config.ImportAzureSASTokenNoPrivilege,
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
	err = expectFail(err, "error")
	if err != nil {
		t.Fatal("test failed -> ", zap.Error(err), zap.String("importId", *i.ImportId))
	} else {
		logger.Info("import failed as expected")
	}
}
