package tidbcloudlogin

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/tidbcloud/serverless-test/pkg/coreportalapi"
	"github.com/tidbcloud/serverless-test/util"
)

type WebApiLoginContext struct {
	Host              string
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
	UserEmail         string
	BearerToken       string

	mu sync.Mutex
}

func (ctx *WebApiLoginContext) LoginAndGetToken(c context.Context) (string, error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	if ctx.Host == "" {
		return "", errors.New("login context is not complete: empty host")
	}

	if ctx.Auth0Domain == "" {
		return "", errors.New("login context is not complete: empty auth0 domain")
	}
	if ctx.Auth0ClientID == "" {
		return "", errors.New("login context is not complete: empty auth0 client id")
	}
	if ctx.Auth0ClientSecret == "" {
		return "", errors.New("login context is not complete: empty auth0 client secret")
	}
	if ctx.UserEmail == "" {
		return "", errors.New("login context is not complete: empty user email")
	}

	var err error
	ctx.BearerToken, err = ctx.getToken(
		c,
		ctx.Host,
		ctx.Auth0Domain,
		ctx.Auth0ClientID,
		ctx.Auth0ClientSecret,
		ctx.UserEmail,
	)
	if err != nil {
		return "", err
	}

	return ctx.BearerToken, nil
}

// getToken requests a token from tidbcloud test api `POST /api/v1/test-token`.
// The returned token is used as bearer token to access tidbcloud api.
// Parameters:
//
//   - host: tidbcloud api host, e.g. "us-west-2.staging.shared.aws.tidbcloud.com"
//   - auth0Domain: auth0 domain, e.g. "tidb-staging.us.auth0.com"
//   - auth0ClientID: auth0 client id, e.g. "xxxxx"
//   - auth0ClientSecret: auth0 client secret, e.g. "xxxxx"
//   - email: which user is logging into tidbcloud, e.g. "xuyifan02@pingcap.com"
func (ctx *WebApiLoginContext) getToken(c context.Context,
	host, auth0Domain, auth0ClientID, auth0ClientSecret, email string,
) (string, error) {
	cfg := coreportalapi.NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Transport: http.DefaultTransport,
	}
	cfg.Host = host
	cfg.UserAgent = "serverless-test"
	cli := coreportalapi.NewAPIClient(cfg)

	r := cli.TestTokenServiceAPI.TestTokenServiceGenerateTestToken(c)
	r = r.Body(coreportalapi.CentralGenerateTestTokenReq{
		Email:             &email,
		Auth0Domain:       &auth0Domain,
		Auth0ClientId:     &auth0ClientID,
		Auth0ClientSecret: &auth0ClientSecret,
	})
	token, resp, err := r.Execute()
	if err != nil {
		return "", err
	}
	err = util.ParseError(err, resp)
	if err != nil {
		return "", err
	}

	if token.Token == nil {
		return "", nil
	}

	return *token.Token, nil
}
