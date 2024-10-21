package export

import (
	"context"
	"flag"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/shiyuhang0/serverless-scene-test/config"
	"github.com/shiyuhang0/serverless-scene-test/util"
	"github.com/stretchr/testify/assert"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/export"
)

var (
	clusterId    string
	exportClient *export.APIClient
)

func init() {
	flag.StringVar(&clusterId, "cid", "", "")
}

func setup() {
	var err error
	config.InitializeConfig()
	exportClient, err = NewExportClient()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func TestExportToLocalAndDownload(t *testing.T) {
	ctx := context.Background()
	t.Log("start to create export")
	res, err := CreateExport(ctx, exportClient, clusterId, nil)
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
	for _, exportFile := range exportFilesRes.Files {
		// download file
		downloadRes, err := util.GetResponse(*exportFile.Url)
		if err != nil {
			t.Fatal(err)
		}
		if downloadRes.ContentLength <= 0 {
			t.Fatal("file is empty")
		}
		file, err := os.Create(*exportFile.Name)
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(file, downloadRes.Body)
		// check file
		targetFileInfo, err := os.Stat(*exportFile.Name)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, targetFileInfo.Name(), *exportFile.Name)
	}
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
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.FileType, fileType)
	assert.Equal(t, *exp.ExportOptions.ParquetFormat.Compression, parquetCompressionType)
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
	res, err := CreateExport(ctx, exportClient, clusterId, body)
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
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.FileType, fileType)
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
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.Compression, compression)
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
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.Filter.Sql, sql)
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
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}
	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)

	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.ExportOptions.Filter.Table.Where, where)
	assert.Equal(t, exp.ExportOptions.Filter.Table.Patterns[0], table0)
	assert.Equal(t, exp.ExportOptions.Filter.Table.Patterns[1], table1)
}

func TestExportToS3(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_S3
	s3AccessKeyId := config.S3AccessKeyId
	s3SecretKeyId := config.S3SecretAccessKey
	exportS3Uri := config.S3URI
	if s3AccessKeyId == "" || s3SecretKeyId == "" || exportS3Uri == "" {
		t.Fatal("s3 access key or secret key or uri is empty")
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

	t.Log("start to create export")
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, *exp.Target.S3.Uri, exportS3Uri)
}

func TestExportToAzure(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_AZURE_BLOB
	azureUri := config.AzureURI
	azureSASToken := config.AzureSASToken
	if azureUri == "" || azureSASToken == "" {
		t.Fatal("azure uri or sas token is empty")
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

	t.Log("start to create export")
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, exp.Target.AzureBlob.Uri, azureUri)
}

func TestExportToGCS(t *testing.T) {
	ctx := context.Background()

	exportType := export.EXPORTTARGETTYPEENUM_GCS
	gcsUri := config.GCSURI
	gcsServiceAccountKey := config.GCSServiceAccountKey
	if gcsUri == "" || gcsServiceAccountKey == "" {
		t.Fatal("gcs uri or service account key is empty")
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

	t.Log("start to create export")
	res, err := CreateExport(ctx, exportClient, clusterId, body)
	if err != nil {
		t.Fatal(err)
	}

	exp := checkServerlessExportState(ctx, t, clusterId, *res.ExportId)
	assert.Equal(t, *exp.State, export.EXPORTSTATEENUM_SUCCEEDED)
	assert.Equal(t, exp.Target.Gcs.Uri, gcsUri)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func checkServerlessExportState(ctx context.Context, t *testing.T, clusterId, exportId string) *export.Export {
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

func NewExportClient() (*export.APIClient, error) {
	httpclient := &http.Client{
		Transport: util.NewDigestTransport(config.PublicKey, config.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(config.ServerlessEndpoint)
	if err != nil {
		return nil, err
	}
	exportCfg := export.NewConfiguration()
	exportCfg.HTTPClient = httpclient
	exportCfg.Host = serverlessURL.Host
	exportCfg.UserAgent = util.UserAgent
	return export.NewAPIClient(exportCfg), nil
}

func CreateExport(ctx context.Context, c *export.APIClient, clusterId string, body *export.ExportServiceCreateExportBody) (*export.Export, error) {
	r := c.ExportServiceAPI.ExportServiceCreateExport(ctx, clusterId)
	if body != nil {
		r = r.Body(*body)
	} else {
		r = r.Body(*export.NewExportServiceCreateExportBody())
	}
	res, h, err := r.Execute()
	return res, util.ParseError(err, h)
}
