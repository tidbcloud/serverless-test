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
				Uri:      cfg.Import.S3.URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.Import.S3.RoleARN,
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
				Uri:      cfg.Import.S3.URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.Import.S3.RoleARNNoPrivilege,
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
				Uri:      cfg.Import.S3.URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ROLE_ARN,
				RoleArn:  &cfg.Import.S3.RoleARNDiffExternalID,
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
				Uri:      cfg.Import.S3.URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &consoleimportapi.S3SourceAccessKey{
					Id:     cfg.S3.AccessKeyID,
					Secret: cfg.S3.SecretAccessKey,
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
				Uri:      cfg.Import.S3.URI,
				AuthType: consoleimportapi.IMPORTS3AUTHTYPEENUM_ACCESS_KEY,
				AccessKey: &consoleimportapi.S3SourceAccessKey{
					Id:     cfg.Import.S3.AccessKeyIDNoPrivilege,
					Secret: cfg.Import.S3.SecretAccessKeyNoPrivilege,
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
				Uri:               cfg.Import.GCS.URI,
				AuthType:          consoleimportapi.IMPORTGCSAUTHTYPEENUM_SERVICE_ACCOUNT_KEY,
				ServiceAccountKey: &cfg.Import.GCS.ServiceAccountKey,
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
				Uri:               cfg.Import.GCS.URI,
				AuthType:          consoleimportapi.IMPORTGCSAUTHTYPEENUM_SERVICE_ACCOUNT_KEY,
				ServiceAccountKey: &cfg.Import.GCS.ServiceAccountKeyNoPrivilege,
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
				Uri:      cfg.Import.Azure.URI,
				AuthType: consoleimportapi.IMPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
				SasToken: &cfg.Import.Azure.SASToken,
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
				Uri:      cfg.Import.Azure.URI,
				AuthType: consoleimportapi.IMPORTAZUREBLOBAUTHTYPEENUM_SAS_TOKEN,
				SasToken: &cfg.Import.Azure.SASTokenNoPrivilege,
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
				Uri:      cfg.Import.OSS.URI,
				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ACCESS_KEY,
				AccessKey: consoleimportapi.NewOSSSourceAccessKey(
					cfg.Import.OSS.AccessKeyID, cfg.Import.OSS.SecretAccessKey),
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
				Uri:      cfg.Import.OSS.URI,
				AuthType: consoleimportapi.IMPORTOSSAUTHTYPEENUM_ACCESS_KEY,
				AccessKey: consoleimportapi.NewOSSSourceAccessKey(
					cfg.Import.OSS.AccessKeyIDNoPrivilege, cfg.Import.OSS.SecretAccessKeyNoPrivilege),
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
