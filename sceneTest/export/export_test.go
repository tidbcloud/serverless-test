package export

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/lithammer/shortuuid/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/export"
)

var (
	clusterId    string
	exportClient *export.APIClient
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestExportToLocalAndDownload(t *testing.T) {
	ctx := context.Background()

	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, nil)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)

	t.Log("start to list download files")
	exportFilesReq := exportClient.ExportServiceAPI.ExportServiceListExportFiles(ctx, clusterId, *res.ExportId)
	exportFilesReq = exportFilesReq.GenerateUrl(true)
	exportFilesRes, h, err := exportFilesReq.Execute()
	err = util.ParseError(err, h)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("start to download files")
	path := shortuuid.New()
	for _, exportFile := range exportFilesRes.Files {
		// download file
		downloadRes, err := util.GetResponse(*exportFile.Url)
		if err != nil {
			t.Fatal(err)
		}
		if downloadRes.ContentLength <= 0 {
			t.Fatal("file is empty")
		}
		fileName := *exportFile.Name + path
		file, err := os.Create(fileName)
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(file, downloadRes.Body)
		// check file
		targetFileInfo, err := os.Stat(fileName)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, targetFileInfo.Name(), fileName)
	}

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportWithParquetFile(t *testing.T) {
	ctx := context.Background()

	fileType := export.EXPORTFILETYPEENUM_PARQUET
	parquetCompressionType := export.EXPORTPARQUETCOMPRESSIONTYPEENUM_SNAPPY
	body := &export.ExportServiceCreateExportBody{
		ExportOptions: &export.ExportOptions{
			FileType: &fileType,
			ParquetFormat: &export.ExportOptionsParquetFormat{
				Compression: &parquetCompressionType,
			},
		},
	}

	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.FileType, fileType)
	assert.Equal(t, *exp.ExportOptions.ParquetFormat.Compression, parquetCompressionType)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportWithCSVFile(t *testing.T) {
	ctx := context.Background()

	fileType := export.EXPORTFILETYPEENUM_CSV
	separator := ",,"
	delimiter := "\"\""
	skipHeader := true
	nullValue := "NULL"
	body := &export.ExportServiceCreateExportBody{
		ExportOptions: &export.ExportOptions{
			FileType: &fileType,
			CsvFormat: &export.ExportOptionsCSVFormat{
				Separator:  &separator,
				Delimiter:  *export.NewNullableString(&delimiter),
				SkipHeader: &skipHeader,
				NullValue:  *export.NewNullableString(&nullValue),
			},
		},
	}
	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.FileType, fileType)
	assert.Equal(t, *exp.ExportOptions.CsvFormat.SkipHeader, skipHeader)
	assert.Equal(t, *exp.ExportOptions.CsvFormat.Separator, separator)
	assert.Equal(t, *exp.ExportOptions.CsvFormat.Delimiter.Get(), delimiter)
	assert.Equal(t, *exp.ExportOptions.CsvFormat.NullValue.Get(), nullValue)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportWithSQLFile(t *testing.T) {
	ctx := context.Background()

	fileType := export.EXPORTFILETYPEENUM_SQL
	body := &export.ExportServiceCreateExportBody{
		ExportOptions: &export.ExportOptions{
			FileType: &fileType,
		},
	}
	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.FileType, fileType)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportWithCompression(t *testing.T) {
	ctx := context.Background()

	compression := export.EXPORTCOMPRESSIONTYPEENUM_ZSTD
	body := &export.ExportServiceCreateExportBody{
		ExportOptions: &export.ExportOptions{
			Compression: &compression,
		},
	}
	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.Compression, compression)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportWithSQLFilter(t *testing.T) {
	ctx := context.Background()

	sql := "SELECT * FROM test.test where id = 1"
	body := &export.ExportServiceCreateExportBody{
		ExportOptions: &export.ExportOptions{
			Filter: &export.ExportOptionsFilter{
				Sql: &sql,
			},
		},
	}
	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.Filter.Sql, sql)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportWithTableFilter(t *testing.T) {
	ctx := context.Background()

	table0 := "test.test"
	table1 := "test.test1"
	tables := []string{table0, table1}
	where := "id = 1"
	body := &export.ExportServiceCreateExportBody{
		ExportOptions: &export.ExportOptions{
			Filter: &export.ExportOptionsFilter{
				Table: &export.ExportOptionsFilterTable{
					Patterns: tables,
					Where:    &where,
				},
			},
		},
	}
	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.Filter.Table.Where, where)
	assert.Equal(t, exp.ExportOptions.Filter.Table.Patterns[0], table0)
	assert.Equal(t, exp.ExportOptions.Filter.Table.Patterns[1], table1)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportToS3AccessKey(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_S3
	s3AccessKeyId := config.S3AccessKeyId
	s3SecretKeyId := config.S3SecretAccessKey
	exportS3Uri := config.S3URI
	if s3AccessKeyId == "" || s3SecretKeyId == "" || exportS3Uri == "" {
		t.Fatal("s3 access key or secret key or uri is empty")
	}
	if strings.HasSuffix(exportS3Uri, "/") {
		exportS3Uri = exportS3Uri + shortuuid.New()
	} else {
		exportS3Uri = exportS3Uri + "/" + shortuuid.New()
	}

	body := export.NewExportServiceCreateExportBody()
	body.Target = &export.ExportTarget{
		Type: &exportType,
		S3: &export.S3Target{
			Uri:       &exportS3Uri,
			AuthType:  export.EXPORTS3AUTHTYPEENUM_ACCESS_KEY,
			AccessKey: export.NewS3TargetAccessKey(s3AccessKeyId, s3SecretKeyId),
		},
	}

	t.Logf("start to create export to s3: %s", exportS3Uri)
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.Target.S3.Uri, exportS3Uri)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportToS3RoleArn(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_S3
	roleArn := config.S3RoleArn
	exportS3Uri := config.S3URI
	if roleArn == "" || exportS3Uri == "" {
		t.Fatal("s3 role arn or uri is empty")
	}
	if strings.HasSuffix(exportS3Uri, "/") {
		exportS3Uri = exportS3Uri + shortuuid.New()
	} else {
		exportS3Uri = exportS3Uri + "/" + shortuuid.New()
	}

	body := export.NewExportServiceCreateExportBody()
	body.Target = &export.ExportTarget{
		Type: &exportType,
		S3: &export.S3Target{
			Uri:      &exportS3Uri,
			AuthType: export.EXPORTS3AUTHTYPEENUM_ROLE_ARN,
			RoleArn:  &roleArn,
		},
	}

	t.Logf("start to create export to s3: %s", exportS3Uri)
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.Target.S3.Uri, exportS3Uri)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportToAzure(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_AZURE_BLOB
	azureUri := config.AzureURI
	azureSASToken := config.AzureSASToken
	if azureUri == "" || azureSASToken == "" {
		t.Fatal("azure uri or sas token is empty")
	}
	if strings.HasSuffix(azureUri, "/") {
		azureUri = azureUri + shortuuid.New()
	} else {
		azureUri = azureUri + "/" + shortuuid.New()
	}

	body := export.NewExportServiceCreateExportBody()
	body.Target = &export.ExportTarget{
		Type: &exportType,
		AzureBlob: &export.AzureBlobTarget{
			Uri:      azureUri,
			AuthType: export.EXPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
			SasToken: &azureSASToken,
		},
	}

	t.Logf("start to create export to %s", azureUri)
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, exp.Target.AzureBlob.Uri, azureUri)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestExportToGCS(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_GCS
	gcsUri := config.GCSURI
	gcsServiceAccountKey := config.GCSServiceAccountKey
	if gcsUri == "" || gcsServiceAccountKey == "" {
		t.Fatal("gcs uri or service account key is empty")
	}
	if strings.HasSuffix(gcsUri, "/") {
		gcsUri = gcsUri + shortuuid.New()
	} else {
		gcsUri = gcsUri + "/" + shortuuid.New()
	}

	body := export.NewExportServiceCreateExportBody()
	body.Target = &export.ExportTarget{
		Type: &exportType,
		Gcs: &export.GCSTarget{
			Uri:               gcsUri,
			AuthType:          export.EXPORTGCSAUTHTYPEENUM_SERVICE_ACCOUNT_KEY,
			ServiceAccountKey: &gcsServiceAccountKey,
		},
	}

	t.Logf("start to create export to %s", gcsUri)
	res, err := CreateExport(ctx, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, exp.Target.Gcs.Uri, gcsUri)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func TestCancelExport(t *testing.T) {
	ctx := context.Background()

	t.Log("start to create export")
	res, err := CreateExport(ctx, clusterId, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("start to cancel export")
	var cancelRes *export.Export
	var h *http.Response
	for i := 0; i < 3; i++ {
		cancelRes, h, err = exportClient.ExportServiceAPI.ExportServiceCancelExport(ctx, clusterId, *res.ExportId).Execute()
		if err == nil || !strings.Contains(err.Error(), "Internal Server Error") {
			break
		}
		time.Sleep(time.Millisecond * 10)
	}
	if err != nil {
		t.Fatal(util.ParseError(err, h))
	}
	assert.Equal(t, *cancelRes.State, export.EXPORTSTATEENUM_CANCELED)

	DeleteExport(ctx, clusterId, *res.ExportId)
}

func checkServerlessExportState(ctx context.Context, t *testing.T, clusterId, exportId string) *export.Export {
	t.Logf("start to check the state of %s", exportId)
	ticker := time.NewTicker(time.Second * 10)
	timeout := time.After(time.Minute * 3)
	for {
		select {
		case <-ticker.C:
			res, h, err := exportClient.ExportServiceAPI.ExportServiceGetExport(ctx, clusterId, exportId).Execute()
			if util.ParseError(err, h) != nil {
				t.Logf("get export failed: %s", util.ParseError(err, h).Error())
				continue
			}
			t.Logf("get export with state %s", *res.State)
			if *res.State != export.EXPORTSTATEENUM_RUNNING {
				return res
			}
		case <-timeout:
			t.Fatal("export timeout")
		}
	}
}

func CreateExport(ctx context.Context, clusterId string, body *export.ExportServiceCreateExportBody) (*export.Export, error) {
	r := exportClient.ExportServiceAPI.ExportServiceCreateExport(ctx, clusterId)
	if body != nil {
		r = r.Body(*body)
	} else {
		r = r.Body(*export.NewExportServiceCreateExportBody())
	}
	res, h, err := r.Execute()
	return res, util.ParseError(err, h)
}

func DeleteExport(ctx context.Context, clusterId, exportId string) error {
	_, h, err := exportClient.ExportServiceAPI.ExportServiceDeleteExport(ctx, clusterId, exportId).Execute()
	return util.ParseError(err, h)
}
