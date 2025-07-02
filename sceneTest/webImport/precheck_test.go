package imp

import (
	"context"
	"testing"

	"github.com/pingcap/log"
	"github.com/tidbcloud/serverless-test/config"
	consoleimportapi "github.com/tidbcloud/serverless-test/pkg/console/import"
	"github.com/tidbcloud/serverless-test/util"
	"go.uber.org/zap"
)

func TestPrecheck(t *testing.T) {
	ctx := context.Background()

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
	i, resp, err := r.Execute()
	err = util.ParseError(err, resp)
	if err != nil {
		t.Fatal(err)
	}
	log.L().Info("", zap.Any("importId", i))
}
