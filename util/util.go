package util

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/icholy/digest"
	"github.com/pingcap/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func NewDigestTransport(publicKey, privateKey string) http.RoundTripper {
	return &digest.Transport{
		Username:  publicKey,
		Password:  privateKey,
		Transport: NewDebugTransport(http.DefaultTransport),
	}
}

func NewDebugTransport(inner http.RoundTripper) http.RoundTripper {
	return &DebugTransport{inner: inner}
}

type DebugTransport struct {
	inner http.RoundTripper
}

func (dt *DebugTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	debug := os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG") == "1"

	if debug {
		dump, err := httputil.DumpRequestOut(r, true)
		if err != nil {
			return nil, err
		}
		fmt.Printf("\n%s", string(dump))
	}

	resp, err := dt.inner.RoundTrip(r)
	if err != nil {
		return resp, err
	}

	if debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		fmt.Printf("%s\n", string(dump))
	}

	return resp, err
}

func ParseError(err error, resp *http.Response) error {
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err == nil {
		return nil
	}
	if resp == nil {
		return err
	}
	body, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		return err
	}
	path := "<path>"
	if resp.Request != nil {
		path = fmt.Sprintf("[%s %s]", resp.Request.Method, resp.Request.URL.Path)
	}
	traceId := "<trace_id>"
	if resp.Header.Get("X-Debug-Trace-Id") != "" {
		traceId = resp.Header.Get("X-Debug-Trace-Id")
		log.Info("trace id", zap.String("trace_id", traceId))
	}
	return fmt.Errorf("%s[%s][%s] %s", path, err.Error(), traceId, body)
}

func ValidateApiUrl(value string) (*url.URL, error) {
	u, err := url.ParseRequestURI(value)
	if err != nil {
		return nil, errors.New("invalid api url format: " + value)
	}
	return u, nil
}

func GetResponse(url string) (*http.Response, error) {
	httpClient := resty.New()
	resp, err := httpClient.GetClient().Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		// read the body to get the error message
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("receiving status of %d", resp.StatusCode)
		}
		type AwsError struct {
			Code    string `xml:"Code"`
			Message string `xml:"Message"`
		}
		v := AwsError{}
		err = xml.Unmarshal(body, &v)
		if err != nil {
			return nil, fmt.Errorf("receiving status of %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("receiving status of %d. code: %s, message: %s", resp.StatusCode, v.Code, v.Message)
	}
	if resp.ContentLength <= 0 {
		resp.Body.Close()
		return nil, fmt.Errorf("file is empty")
	}
	return resp, nil
}

// NewBearerTransport returns a RoundTripper that adds a Bearer token to the Authorization header.
func NewBearerTransport(token string) http.RoundTripper {
	return &bearerTransport{
		token: token,
		inner: NewDebugTransport(http.DefaultTransport),
	}
}

type bearerTransport struct {
	token string
	inner http.RoundTripper
}

func (bt *bearerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r2 := r.Clone(r.Context())
	r2.Header.Set("Authorization", "Bearer "+bt.token)
	return bt.inner.RoundTrip(r2)
}

// EqualPointerValues checks if two pointers point to the same value.
// see https://github.com/stretchr/testify/issues/1118
func EqualPointerValues(t *require.Assertions, expected, actual interface{}) {
	t.EqualValues(expected, actual, "exp=%v, got=%v", reflect.Indirect(reflect.ValueOf(expected)), reflect.Indirect(reflect.ValueOf(actual)))
}
