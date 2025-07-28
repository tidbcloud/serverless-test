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
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/imp"
)

const (
	localFilePath = "../../data/test.a.csv"
)

// TestLocalImport tests successful local file import via upload
func TestLocalImport(t *testing.T) {
	ctx := context.Background()

	// Clean up existing table
	if err := cleanupTestTable(ctx); err != nil {
		t.Fatalf("Failed to cleanup test table: %v", err)
	}

	t.Log("Starting local file import test")

	// Upload file and create import
	importID, err := createLocalImport(ctx)
	if err != nil {
		t.Fatalf("Failed to create local import: %v", err)
	}

	// Wait for import completion
	if err := waitImport(ctx, importID); err != nil {
		t.Fatalf("Import failed, importId: %s, error: %v", importID, err)
	}

	t.Log("Local import completed successfully")
}

// createLocalImport handles the complete local import process: upload + import creation
func createLocalImport(ctx context.Context) (string, error) {
	const (
		targetDatabase = "test"
		targetTable    = "a"
	)

	// Step 1: Start upload process
	uploadID, uploadURL, err := startFileUpload(ctx, targetDatabase, targetTable)
	if err != nil {
		return "", err
	}

	// Step 2: Upload the file
	if err := uploadFile(localFilePath, uploadURL); err != nil {
		return "", err
	}

	// Step 3: Complete upload
	if err := completeFileUpload(ctx, uploadID); err != nil {
		return "", err
	}

	// Step 4: Create import
	return createImportFromUpload(ctx, uploadID, targetDatabase, targetTable)
}

// startFileUpload initiates the file upload process
func startFileUpload(ctx context.Context, targetDatabase, targetTable string) (string, string, error) {
	startUploadContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	request := importClient.ImportServiceAPI.ImportServiceStartUpload(startUploadContext, clusterId)
	request = request.FileName(localFilePath)
	request = request.PartNumber(1)
	request = request.TargetDatabase(targetDatabase)
	request = request.TargetTable(targetTable)

	uploadResult, resp, err := request.Execute()
	if err := util.ParseError(err, resp); err != nil {
		return "", "", err
	}

	return *uploadResult.UploadId, uploadResult.UploadUrl[0], nil
}

// completeFileUpload finalizes the upload process
func completeFileUpload(ctx context.Context, uploadID string) error {
	_, resp, err := importClient.ImportServiceAPI.ImportServiceCompleteUpload(ctx, clusterId).
		UploadId(uploadID).
		Parts([]imp.CompletePart{}).
		Execute()

	return util.ParseError(err, resp)
}

// createImportFromUpload creates an import from the uploaded file
func createImportFromUpload(ctx context.Context, uploadID, targetDatabase, targetTable string) (string, error) {
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	body := imp.ImportServiceCreateImportBody{
		ImportOptions: imp.ImportOptions{
			FileType: imp.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &imp.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: imp.ImportSource{
			Type: imp.IMPORTSOURCETYPEENUM_LOCAL,
			Local: &imp.LocalSource{
				UploadId:       uploadID,
				TargetDatabase: targetDatabase,
				TargetTable:    targetTable,
			},
		},
	}

	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, clusterId)
	r = r.Body(body)

	importTask, resp, err := r.Execute()
	if err := util.ParseError(err, resp); err != nil {
		return "", err
	}

	return *importTask.ImportId, nil
}

// uploadFile uploads a file to the specified URL
func uploadFile(filePath string, url string) error {
	// Open and read file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read file content into buffer
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		return err
	}

	// Create HTTP request
	request, err := http.NewRequest(http.MethodPut, url, buffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "multipart/form-data")

	// Execute upload
	client := &http.Client{}
	_, err = client.Do(request)
	return err
}
