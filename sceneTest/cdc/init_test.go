package cdc

import (
	"log"
	"net/http"

	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cdc"
)

var cdcClient *cdc.APIClient

func setup() {
	var err error
	cdcClient, err = NewCDCClient(config.LoadConfig())
	if err != nil {
		log.Panicf("failed to create cdc client: %v", err)
	}
}

// NewCDCClient creates a new export API client with the given configuration
func NewCDCClient(cfg *config.Config) (*cdc.APIClient, error) {
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(cfg.Endpoint.Serverless)
	if err != nil {
		return nil, err
	}
	cdcCfg := cdc.NewConfiguration()
	cdcCfg.HTTPClient = httpClient
	cdcCfg.Host = serverlessURL.Host
	cdcCfg.UserAgent = util.UserAgent
	return cdc.NewAPIClient(cdcCfg), nil
}
