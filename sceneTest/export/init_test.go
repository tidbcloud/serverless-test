package export

import (
	"flag"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/export"
	"net/http"
	"os"
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
