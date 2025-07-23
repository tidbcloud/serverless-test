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
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      cfg.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.ImportS3RoleARN,
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
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      cfg.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.ImportS3RoleARNNoPrivilege,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString("S3 access deny, please check your S3 access key or role arn and uri."), result.GetBaseResp().ErrMsg)
}

func TestS3ArnDiffExternalID(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      cfg.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.ImportS3RoleARNDiffExternalID,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString("S3 access deny, please check your S3 access key or role arn and uri."), result.GetBaseResp().ErrMsg)
}

func TestS3AccessKey(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      cfg.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &consoleimportapi.S3SourceAccessKey{
					Id:     cfg.S3AccessKeyID,
					Secret: cfg.S3SecretAccessKey,
				},
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestS3AccessKeyNoPrivilege(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_S3,
			S3: &consoleimportapi.S3Source{
				Uri:      cfg.ImportS3URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &consoleimportapi.S3SourceAccessKey{
					Id:     cfg.ImportS3AccessKeyIDNoPrivilege,
					Secret: cfg.ImportS3SecretAccessKeyNoPrivilege,
				},
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString("S3 access deny, please check your S3 access key or role arn and uri."), result.GetBaseResp().ErrMsg)
}

func TestGCSServiceAccountKey(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_GCS,
			Gcs: &consoleimportapi.GCSSource{
				Uri:               cfg.ImportGCSURI,
				AuthType:          consoleimportapi.IMPORTGCSAUTHTYPEENUM_SERVICE_ACCOUNT_KEY,
				ServiceAccountKey: &cfg.ImportGCSServiceAccountKey,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestGCSServiceAccountKeyNoPrivilege(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_GCS,
			Gcs: &consoleimportapi.GCSSource{
				Uri:               cfg.ImportGCSURI,
				AuthType:          consoleimportapi.IMPORTGCSAUTHTYPEENUM_SERVICE_ACCOUNT_KEY,
				ServiceAccountKey: &cfg.ImportGCSServiceAccountKeyNoPrivilege,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString("GCS access deny, please check your GCS Service Account Key and uri."), result.GetBaseResp().ErrMsg)
}

func TestAzureSASToken(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_AZURE_BLOB,
			AzureBlob: &consoleimportapi.AzureBlobSource{
				Uri:      cfg.ImportAzureURI,
				AuthType: consoleimportapi.IMPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
				SasToken: &cfg.ImportAzureSASToken,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestAzureSASTokenNoPrivilege(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_AZURE_BLOB,
			AzureBlob: &consoleimportapi.AzureBlobSource{
				Uri:      cfg.ImportAzureURI,
				AuthType: consoleimportapi.IMPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
				SasToken: &cfg.ImportAzureSASTokenNoPrivilege,
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString("Azure Blob access deny, please check your Azure Blob SAS token and uri."), result.GetBaseResp().ErrMsg)
}

func TestOSSAccessKey(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
			Oss: &consoleimportapi.OSSSource{
				Uri:      cfg.ImportOSSURI,
				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ACCESS_KEY,
				AccessKey: consoleimportapi.NewOSSSourceAccessKey(
					cfg.ImportOSSAccessKeyID, cfg.ImportOSSSecretAccessKey),
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
}

func TestOSSAccessKeyNoPrivilege(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	cfg := config.LoadConfig()
	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
		Source: consoleimportapi.ImportSource{
			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
			Oss: &consoleimportapi.OSSSource{
				Uri:      cfg.ImportOSSURI,
				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ACCESS_KEY,
				AccessKey: consoleimportapi.NewOSSSourceAccessKey(
					cfg.ImportOSSAccessKeyIDNoPrivilege, cfg.ImportOSSSecretAccessKeyNoPrivilege),
			},
		},
	})
	result, resp, err := r.Execute()
	err = util.ParseError(err, resp)

	assert.NoError(err)
	assert.NotNil(result)
	util.EqualPointerValues(assert, pointer.ToString("OSS access deny, please check your OSS access key and uri."), result.GetBaseResp().ErrMsg)
}

// TODO open this test when OSS role arn is supported
//func TestOSSArn(t *testing.T) {
//	ctx := context.Background()
//	assert := require.New(t)
//	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
//	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
//		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
//		Source: consoleimportapi.ImportSource{
//			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
//			Oss: &consoleimportapi.OSSSource{
//				Uri:      config.ImportOSSURI,
//				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ROLE_ARN,
//				RoleArn:  &config.ImportOSSRoleArn,
//			},
//		},
//	})
//	result, resp, err := r.Execute()
//	err = util.ParseError(err, resp)
//
//	assert.NoError(err)
//	assert.NotNil(result)
//	util.EqualPointerValues(assert, pointer.ToString(""), result.GetBaseResp().ErrMsg)
//}
//
//func TestOSSArnNoPrivilege(t *testing.T) {
//	ctx := context.Background()
//	assert := require.New(t)
//	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
//	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
//		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
//		Source: consoleimportapi.ImportSource{
//			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
//			Oss: &consoleimportapi.OSSSource{
//				Uri:      config.ImportOSSURI,
//				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ROLE_ARN,
//				RoleArn:  &config.ImportOSSRoleArnNoPrivilege,
//			},
//		},
//	})
//	result, resp, err := r.Execute()
//	err = util.ParseError(err, resp)
//
//	assert.NoError(err)
//	assert.NotNil(result)
//	util.EqualPointerValues(assert, pointer.ToString("OSS access deny, please check your OSS access key and uri."), result.GetBaseResp().ErrMsg)
//}
//
//func TestOSSArnDiffExternalID(t *testing.T) {
//	ctx := context.Background()
//	assert := require.New(t)
//	r := importClient.ImportServiceAPI.ImportServiceValidateImport(ctx, orgId, projectId, clusterId)
//	r = r.Body(consoleimportapi.ImportServiceValidateImportBody{
//		ValidationType: consoleimportapi.IMPORTVALIDATIONTYPEENUM_SOURCE_ACCESS_CHECK,
//		Source: consoleimportapi.ImportSource{
//			Type: consoleimportapi.IMPORTSOURCETYPEENUM_OSS,
//			Oss: &consoleimportapi.OSSSource{
//				Uri:      config.ImportOSSURI,
//				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ROLE_ARN,
//				RoleArn:  &config.ImportOSSRoleArnDiffExternalID,
//			},
//		},
//	})
//	result, resp, err := r.Execute()
//	err = util.ParseError(err, resp)
//
//	assert.NoError(err)
//	assert.NotNil(result)
//	util.EqualPointerValues(assert, pointer.ToString("OSS access deny, please check your OSS access key and uri."), result.GetBaseResp().ErrMsg)
//}
