package imp

import (
	"context"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/require"
	"github.com/tidbcloud/serverless-test/config"
	consoleimportapi "github.com/tidbcloud/serverless-test/pkg/console/import"
	"github.com/tidbcloud/serverless-test/util"
)

func TestPrecheckWithoutTable(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	assert.NoError(err)
	r := importClient.ImportServiceAPI.ImportServicePrecheck(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServicePrecheckBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToBool(false), result.IsTruncated)
	util.EqualPointerValues(assert, pointer.ToString(""), result.ErrorMessage)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalTablesCount)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalDataFilesCount)
	// Verify table result
	assert.ElementsMatch([]consoleimportapi.TableResult{
		{
			TargetTable: &consoleimportapi.ConsoleTable{
				Database: pointer.ToString("test"),
				Table:    pointer.ToString("a"),
			},
			ErrorMessage:      pointer.ToString(""),
			MatchedDataFiles:  []string{"test.a.csv"},
			MatchedSchemaFile: pointer.ToString("test.a-schema.sql.gz"),
			UseSchemaFile:     pointer.ToBool(true),
		},
	}, result.TargetTables)
}

func TestPrecheckWithEmptyTable(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	assert.NoError(err)
	_, err = db.Exec("CREATE TABLE `test`.`a` (name VARCHAR(20) NOT NULL, age INT NOT NULL)")
	assert.NoError(err)
	r := importClient.ImportServiceAPI.ImportServicePrecheck(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServicePrecheckBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToBool(false), result.IsTruncated)
	util.EqualPointerValues(assert, pointer.ToString(""), result.ErrorMessage)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalTablesCount)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalDataFilesCount)
	// Verify table result
	assert.ElementsMatch([]consoleimportapi.TableResult{
		{
			TargetTable: &consoleimportapi.ConsoleTable{
				Database: pointer.ToString("test"),
				Table:    pointer.ToString("a"),
			},
			ErrorMessage:      pointer.ToString(""),
			MatchedDataFiles:  []string{"test.a.csv"},
			MatchedSchemaFile: pointer.ToString("test.a-schema.sql.gz"),
			UseSchemaFile:     pointer.ToBool(false),
		},
	}, result.TargetTables)
}

func TestPrecheckWithNonEmptyTable(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	assert.NoError(err)
	_, err = db.Exec("CREATE TABLE `test`.`a` (name VARCHAR(20) NOT NULL, age INT NOT NULL)")
	assert.NoError(err)
	_, err = db.Exec("INSERT INTO `test`.`a` VALUES ('Alice', 30), ('Bob', 25)")
	assert.NoError(err)
	r := importClient.ImportServiceAPI.ImportServicePrecheck(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServicePrecheckBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToBool(false), result.IsTruncated)
	util.EqualPointerValues(assert, pointer.ToString("Found 1 table(s) with error: test.a"), result.ErrorMessage)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalTablesCount)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalDataFilesCount)
	// Verify table result
	assert.ElementsMatch([]consoleimportapi.TableResult{
		{
			TargetTable: &consoleimportapi.ConsoleTable{
				Database: pointer.ToString("test"),
				Table:    pointer.ToString("a"),
			},
			ErrorMessage:      pointer.ToString("table test.a is not empty"),
			MatchedDataFiles:  []string{"test.a.csv"},
			MatchedSchemaFile: pointer.ToString("test.a-schema.sql.gz"),
			UseSchemaFile:     pointer.ToBool(false),
		},
	}, result.TargetTables)
}

func TestPrecheckTruncatedResult(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	r := importClient.ImportServiceAPI.ImportServicePrecheck(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServicePrecheckBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_SQL,
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      "s3://tidbcloud-samples-us-east-1/import-data/gharchive_dev/",
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToBool(true), result.IsTruncated)
	util.EqualPointerValues(assert, pointer.ToString(""), result.ErrorMessage)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalTablesCount)
	util.EqualPointerValues(assert, pointer.ToString("997"), result.TotalDataFilesCount)
}

func TestPrecheckCustomMappingWithoutTable(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	assert.NoError(err)
	r := importClient.ImportServiceAPI.ImportServicePrecheck(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServicePrecheckBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
				TargetTableInfos: []consoleimportapi.ImportTargetTableInfo{
					{
						TargetTable: &consoleimportapi.ConsoleTable{
							Database: pointer.ToString("test"),
							Table:    pointer.ToString("a"),
						},
						CustomFile: pointer.ToString("test.a.csv"),
					},
				},
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToBool(false), result.IsTruncated)
	util.EqualPointerValues(assert, pointer.ToString("Found 1 table(s) with error: test.a"), result.ErrorMessage)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalTablesCount)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalDataFilesCount)
	// Verify table result
	assert.ElementsMatch([]consoleimportapi.TableResult{
		{
			TargetTable: &consoleimportapi.ConsoleTable{
				Database: pointer.ToString("test"),
				Table:    pointer.ToString("a"),
			},
			ErrorMessage:      pointer.ToString("table test.a does not exist"),
			MatchedDataFiles:  []string{"test.a.csv"},
			MatchedSchemaFile: pointer.ToString(""),
			UseSchemaFile:     pointer.ToBool(false),
		},
	}, result.TargetTables)
}

func TestPrecheckCustomMappingWithNonEmptyTable(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	assert.NoError(err)
	_, err = db.Exec("CREATE TABLE `test`.`a` (name VARCHAR(20) NOT NULL, age INT NOT NULL)")
	assert.NoError(err)
	_, err = db.Exec("INSERT INTO `test`.`a` VALUES ('Alice', 30), ('Bob', 25)")
	assert.NoError(err)
	r := importClient.ImportServiceAPI.ImportServicePrecheck(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServicePrecheckBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
				TargetTableInfos: []consoleimportapi.ImportTargetTableInfo{
					{
						TargetTable: &consoleimportapi.ConsoleTable{
							Database: pointer.ToString("test"),
							Table:    pointer.ToString("a"),
						},
						CustomFile: pointer.ToString("test.a.csv"),
					},
				},
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToBool(false), result.IsTruncated)
	util.EqualPointerValues(assert, pointer.ToString("Found 1 table(s) with error: test.a"), result.ErrorMessage)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalTablesCount)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalDataFilesCount)
	// Verify table result
	assert.ElementsMatch([]consoleimportapi.TableResult{
		{
			TargetTable: &consoleimportapi.ConsoleTable{
				Database: pointer.ToString("test"),
				Table:    pointer.ToString("a"),
			},
			ErrorMessage:      pointer.ToString("table test.a is not empty"),
			MatchedDataFiles:  []string{"test.a.csv"},
			MatchedSchemaFile: pointer.ToString(""),
			UseSchemaFile:     pointer.ToBool(false),
		},
	}, result.TargetTables)
}

func TestPrecheckCustomMappingWithEmptyTable(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	_, err := db.Exec("DROP TABLE IF EXISTS `test`.`a`")
	assert.NoError(err)
	_, err = db.Exec("CREATE TABLE `test`.`a` (name VARCHAR(20) NOT NULL, age INT NOT NULL)")
	assert.NoError(err)
	r := importClient.ImportServiceAPI.ImportServicePrecheck(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServicePrecheckBody{
		ImportOptions: consoleimportapi.ImportOptions{
			FileType: consoleimportapi.IMPORTFILETYPEENUM_CSV,
		},
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArn,
				TargetTableInfos: []consoleimportapi.ImportTargetTableInfo{
					{
						TargetTable: &consoleimportapi.ConsoleTable{
							Database: pointer.ToString("test"),
							Table:    pointer.ToString("a"),
						},
						CustomFile: pointer.ToString("test.a.csv"),
					},
				},
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToBool(false), result.IsTruncated)
	util.EqualPointerValues(assert, pointer.ToString(""), result.ErrorMessage)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalTablesCount)
	util.EqualPointerValues(assert, pointer.ToString("1"), result.TotalDataFilesCount)
	// Verify table result
	assert.ElementsMatch([]consoleimportapi.TableResult{
		{
			TargetTable: &consoleimportapi.ConsoleTable{
				Database: pointer.ToString("test"),
				Table:    pointer.ToString("a"),
			},
			ErrorMessage:      pointer.ToString(""),
			MatchedDataFiles:  []string{"test.a.csv"},
			MatchedSchemaFile: pointer.ToString(""),
			UseSchemaFile:     pointer.ToBool(false),
		},
	}, result.TargetTables)
}
