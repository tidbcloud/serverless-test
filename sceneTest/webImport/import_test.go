package imp

import (
	"context"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/require"
	"github.com/tidbcloud/serverless-test/config"
	consoleimportapi "github.com/tidbcloud/serverless-test/pkg/console/import"
	"github.com/tidbcloud/serverless-test/util"
)

func TestImportWithoutTargetTables(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)

	// Clean up existing table
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	assert.NoError(err)

	t.Log("Starting import without target tables")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ConsoleImportServiceCreateImportBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &consoleimportapi.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      cfg.Import.S3.URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &consoleimportapi.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
				},
			},
		},
		Creator: nil,
	})
	i, resp, err := r.Execute()
	err = util.ParseError(err, resp)
	assert.NoError(err)
	err = waitImport(ctx, *i.Id)
	assert.NoError(err)
}

func TestImportWithTargetTables(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)

	// Clean up and prepare test table
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`b`")
	assert.NoError(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `test`.`b` (name VARCHAR(20) NOT NULL, age INT NOT NULL)")
	assert.NoError(err)

	t.Log("Starting import with target tables")
	startImportContext, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceCreateImport(startImportContext, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ConsoleImportServiceCreateImportBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
			CsvFormat: &consoleimportapi.CSVFormat{
				Separator: pointer.ToString(";"),
			},
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      cfg.Import.S3.URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &consoleimportapi.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
				},
				TargetTableInfos: []consoleimportapi.ImportTargetTableInfo{
					{
						TargetTable: &consoleimportapi.ConsoleTable{
							Database: pointer.ToString("test"),
							Table:    pointer.ToString("b"),
						},
						CustomFile: pointer.ToString("test.a.csv"),
					},
				},
			},
		},
	})
	i, resp, err := r.Execute()
	err = util.ParseError(err, resp)
	assert.NoError(err)
	err = waitImport(ctx, *i.Id)
	assert.NoError(err)

	// Verify data was imported
	query, err := db.Query("SELECT COUNT(*) FROM `test`.`b`")
	assert.NoError(err)
	defer query.Close()

	var count int
	if query.Next() {
		_ = query.Scan(&count)
	}
	assert.Greater(count, 0, "table test.b should not be empty")
}
