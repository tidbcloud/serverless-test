package tidbcloudlogin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
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
	HTTPClient        *http.Client
	CookieHeader      string

	mu sync.Mutex
}

func (ctx *WebApiLoginContext) Login(c context.Context) (error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	if ctx.Host == "" {
		return errors.New("login context is not complete: empty host")
	}

	if ctx.Auth0Domain == "" {
		return errors.New("login context is not complete: empty auth0 domain")
	}
	if ctx.Auth0ClientID == "" {
		return errors.New("login context is not complete: empty auth0 client id")
	}
	if ctx.Auth0ClientSecret == "" {
		return errors.New("login context is not complete: empty auth0 client secret")
	}
	if ctx.UserEmail == "" {
		return errors.New("login context is not complete: empty user email")
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
		return err
	}

	if ctx.HTTPClient == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return fmt.Errorf("failed to init cookie jar: %w", err)
		}
		ctx.HTTPClient = &http.Client{Jar: jar}
	}

	if err := ctx.loginConfirm(c, ctx.BearerToken); err != nil {
		return err
	}
	ctx.HTTPClient.Transport = util.NewBearerTransport(ctx.BearerToken)

	return nil
}

type loginConfirmReq struct {
	IDToken string `json:"id_token"`
}

func (ctx *WebApiLoginContext) loginConfirm(c context.Context, idToken string) error {
	loginConfirmURL := (&url.URL{
		Scheme: "https",
		Host:   ctx.Host,
		Path:   "/login_confirm",
	}).String()

	bodyBytes, err := json.Marshal(loginConfirmReq{IDToken: idToken})
	if err != nil {
		return fmt.Errorf("failed to marshal login_confirm payload: %w", err)
	}

	req, err := http.NewRequestWithContext(c, http.MethodPost, loginConfirmURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create login_confirm request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ctx.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("login_confirm request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("login_confirm failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(b)))
	}

	if ctx.HTTPClient.Jar == nil {
		return errors.New("login_confirm succeeded but http client has no cookie jar")
	}

	u, _ := url.Parse(loginConfirmURL)
	cookies := ctx.HTTPClient.Jar.Cookies(u)
	if len(cookies) == 0 {
		return errors.New("login_confirm succeeded but no cookies were set")
	}

	parts := make([]string, 0, len(cookies))
	for _, ck := range cookies {
		parts = append(parts, ck.Name+"="+ck.Value)

	}
	ctx.CookieHeader = strings.Join(parts, "; ")

	return nil
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
	cfg.Scheme = "https"
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
