package imp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/pingcap/log"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
	"go.uber.org/zap"
)

const (
	localFilePath = "../../data/test.a.csv"
)

func TestLocalImport(t *testing.T) {
	ctx := context.Background()
	logger := log.L().With(zap.String("test", "e2eLocalImport"))
	_, err := db.Exec("DROP TABLE IF EXISTS a")
	if err != nil {
		logger.Fatal("failed to drop table -> ", zap.Error(err))
	}

	logger.Info("start upload")
	targetDatabase := "test"
	targetTable := "a"
	startUploadContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	request := importClient.ImportServiceAPI.ImportServiceStartUpload(startUploadContext, clusterId)
	request = request.FileName(localFilePath)
	request = request.PartNumber(1)
	request = request.TargetDatabase(targetDatabase)
	request = request.TargetTable(targetTable)
	u, resp, err := request.Execute()
	err = util.ParseError(err, resp)
	if err != nil {
		t.Fatal(err)
	}

	err = uploadFile(localFilePath, u.UploadUrl[0])
	if err != nil {
		t.Fatal("upload file failed", err)
	}

	_, resp, err = importClient.ImportServiceAPI.ImportServiceCompleteUpload(ctx, clusterId).UploadId(*u.UploadId).Parts([]imp.CompletePart{}).Execute()
	err = util.ParseError(err, resp)
	if err != nil {
		t.Fatal(err)
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
			Type: imp.IMPORTSOURCETYPEENUM_LOCAL,
			Local: &imp.LocalSource{
				UploadId:       *u.UploadId,
				TargetDatabase: targetDatabase,
				TargetTable:    targetTable,
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
		logger.Fatal("import failed -> ", zap.Error(err))
	}
	logger.Info("import finished")
}

func uploadFile(filePath string, url string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, url, buffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "multipart/form-data")
	client := &http.Client{}
	_, err = client.Do(request)
	return err
}
