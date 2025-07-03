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

func TestCheckImportPermission(t *testing.T) {
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
