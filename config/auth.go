package config

import (
	"errors"
	"flag"
	"os"

	"github.com/shiyuhang0/serverless-scene-test/client"
)

var (
	publicKey          string
	privateKey         string
	serverlessEndpoint string
	iamEndpoint        string
)

const (
	publicKeyEnv          = "TIDB_CLOUD_PUBLIC_KEY"
	privateKeyEnv         = "TIDB_CLOUD_PRIVATE_KEY"
	serverlessEndpointEnv = "TIDB_CLOUD_SERVERLESS_ENDPOINT"
	iamEndpointEnv        = "TIDB_CLOUD_IAM_ENDPOINT"

	defaultServerlessEndpoint = "https://serverless.tidbapi.com"
	defaultIamEndpoint        = "https://iam.tidbapi.com"
)

func init() {
	flag.StringVar(&publicKey, "public-key", "", "")
	flag.StringVar(&privateKey, "private-key", "", "")
	flag.StringVar(&serverlessEndpoint, "serverless-endpoint", "", "")
	flag.StringVar(&iamEndpoint, "iam-endpoint", "", "")
}

func GetClient() (*client.ClientDelegate, error) {
	public, private := GetOpenAPIKey()
	if public == "" || private == "" {
		return nil, errors.New("public key or private key is empty")
	}
	return client.NewClientDelegateWithApiKey(public, private, GetServerlessEndpoint(), GetIamEndpoint())
}

func GetOpenAPIKey() (string, string) {
	if publicKey != "" && privateKey != "" {
		return publicKey, privateKey
	}
	if os.Getenv(publicKeyEnv) != "" && os.Getenv(privateKeyEnv) != "" {
		return os.Getenv(publicKeyEnv), os.Getenv(privateKeyEnv)
	}
	return "", ""
}

func GetServerlessEndpoint() string {
	if serverlessEndpoint != "" {
		return serverlessEndpoint
	}
	if os.Getenv(serverlessEndpointEnv) != "" {
		return os.Getenv(serverlessEndpointEnv)
	}
	return defaultServerlessEndpoint
}

func GetIamEndpoint() string {
	if iamEndpoint != "" {
		return iamEndpoint
	}
	if os.Getenv(iamEndpointEnv) != "" {
		return os.Getenv(iamEndpointEnv)
	}
	return defaultIamEndpoint
}
