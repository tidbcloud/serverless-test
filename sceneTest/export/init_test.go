package export

import (
	"log"
	"net/http"

	"github.com/spf13/pflag"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/export"
)

func init() {
	pflag.StringVar(&clusterId, "cid", "", "")
}

func setup() {
	var err error
	exportClient, err = NewExportClient(config.LoadConfig())
	if err != nil {
		log.Panicf("failed to create export client: %v", err)
	}
}

// NewExportClient creates a new export API client with the given configuration
func NewExportClient(cfg *config.Config) (*export.APIClient, error) {
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(cfg.ServerlessEndpoint)
	if err != nil {
		return nil, err
	}
	exportCfg := export.NewConfiguration()
	exportCfg.HTTPClient = httpClient
	exportCfg.Host = serverlessURL.Host
	exportCfg.UserAgent = util.UserAgent
	return export.NewAPIClient(exportCfg), nil
}
