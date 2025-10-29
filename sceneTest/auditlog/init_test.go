package auditlg

import (
	"log"
	"net/http"

	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/auditlog"
)

var auditLogClient *auditlog.APIClient

func setup() {
	var err error
	auditLogClient, err = NewAuditLogClient(config.LoadConfig())
	if err != nil {
		log.Panicf("failed to create cdc client: %v", err)
	}
}

// NewAuditLogClient creates a new audit log API client with the given configuration
func NewAuditLogClient(cfg *config.Config) (*auditlog.APIClient, error) {
	httpClient := &http.Client{
		Transport: util.NewDigestTransport(cfg.PublicKey, cfg.PrivateKey),
	}
	serverlessURL, err := util.ValidateApiUrl(cfg.Endpoint.Serverless)
	if err != nil {
		return nil, err
	}
	alCfg := auditlog.NewConfiguration()
	alCfg.HTTPClient = httpClient
	alCfg.Host = serverlessURL.Host
	alCfg.UserAgent = util.UserAgent
	return auditlog.NewAPIClient(alCfg), nil
}
