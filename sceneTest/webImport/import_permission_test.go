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

func TestS3Arn(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
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
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestS3ArnNoPrivilege(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArnNoPrivilege,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestS3ArnDiffExternalID(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      config.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportS3RoleArnDiffExternalID,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestOSSArn(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
			Oss: &consoleimportapi.OSSSource{
				Uri:      config.ImportOSSURI,
				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportOSSRoleArn,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestOSSArnNoPrivilege(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
			Oss: &consoleimportapi.OSSSource{
				Uri:      config.ImportOSSURI,
				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportOSSRoleArnNoPrivilege,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestOSSArnDiffExternalID(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
			Oss: &consoleimportapi.OSSSource{
				Uri:      config.ImportOSSURI,
				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &config.ImportOSSRoleArnDiffExternalID,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}
