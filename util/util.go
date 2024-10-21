package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/icholy/digest"
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
