package export

import (
	"flag"
	"net/http"

	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/export"
)

func init() {
	flag.StringVar(&clusterId, "cid", "", "")
}

func setup() {
	var err error
	exportClient, err = NewExportClient(config.LoadConfig())
	if err != nil {
		panic(err)
	}
}

func NewExportClient(cfg *config.Config) (*export.APIClient, error) {
	httpclient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(cfg.ServerlessEndpoint)
	if err != nil {
		return nil, err
	}
	exportCfg := export.NewConfiguration()
	exportCfg.HTTPClient = httpclient
	exportCfg.Host = serverlessURL.Host
	exportCfg.UserAgent = util.UserAgent
	return export.NewAPIClient(exportCfg), nil
}
