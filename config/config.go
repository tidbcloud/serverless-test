package config

import (
	"flag"
	"os"
)

var (
	PublicKey          string
	PrivateKey         string
	ServerlessEndpoint string
	IamEndpoint        string

	AzureURI      string
	AzureSASToken string

	S3URI             string
	S3SecretAccessKey string
	S3AccessKeyId     string
	S3RoleArn         string

	GCSURI               string
	GCSServiceAccountKey string
)

const (
	PublicKeyEnv          = "TIDB_CLOUD_PUBLIC_KEY"
	PrivateKeyEnv         = "TIDB_CLOUD_PRIVATE_KEY"
	ServerlessEndpointEnv = "TIDB_CLOUD_SERVERLESS_ENDPOINT"
	IamEndpointEnv        = "TIDB_CLOUD_IAM_ENDPOINT"

	defaultServerlessEndpoint = "https://serverless.tidbapi.com"
	defaultIamEndpoint        = "https://iam.tidbapi.com"

	AzureURIEnv      = "Azure_URI"
	AzureSASTokenEnv = "Azure_SAS_Token"

	S3URIEnv             = "S3_URI"
	S3SecretAccessKeyEnv = "S3_SECRET_ACCESS_KEY"
	S3AccessKeyIdEnv     = "S3_ACCESS_KEY_ID"
	S3RoleArnEnv         = "S3_ROLE_ARN"

	GCSURIEnv               = "GCS_URI"
	GCSServiceAccountKeyEnv = "GCS_SERVICE_ACCOUNT_KEY"
)

func init() {
	flag.StringVar(&PublicKey, "public-key", "", "")
	flag.StringVar(&PrivateKey, "private-key", "", "")
	flag.StringVar(&ServerlessEndpoint, "serverless-endpoint", "", "")
	flag.StringVar(&IamEndpoint, "iam-endpoint", "", "")

	flag.StringVar(&S3URI, "s3.uri", "", "")
	flag.StringVar(&S3SecretAccessKey, "s3.secret-access-key", "", "")
	flag.StringVar(&S3AccessKeyId, "s3.access-key-id", "", "")
	flag.StringVar(&S3RoleArn, "s3.role-arn", "", "")

	flag.StringVar(&AzureURI, "azure.uri", "", "")
	flag.StringVar(&AzureSASToken, "azure.sas-token", "", "")

	flag.StringVar(&GCSURI, "gcs.uri", "", "")
	flag.StringVar(&GCSServiceAccountKey, "gcs.service-account-key", "", "")
}

func InitializeConfig() {
	// priority: flag > env > default
	flag.Parse()
	getEnvironment()
	getDefault()
}

func getEnvironment() {
	if PublicKey == "" {
		PublicKey = os.Getenv(PublicKeyEnv)
	}
	if PrivateKey == "" {
		PrivateKey = os.Getenv(PrivateKeyEnv)
	}
	if ServerlessEndpoint == "" {
		ServerlessEndpoint = os.Getenv(ServerlessEndpointEnv)
	}
	if IamEndpoint == "" {
		IamEndpoint = os.Getenv(IamEndpointEnv)
	}
	if AzureURI == "" {
		AzureURI = os.Getenv(AzureURIEnv)
	}
	if AzureSASToken == "" {
		AzureSASToken = os.Getenv(AzureSASTokenEnv)
	}
	if S3URI == "" {
		S3URI = os.Getenv(S3URIEnv)
	}
	if S3SecretAccessKey == "" {
		S3SecretAccessKey = os.Getenv(S3SecretAccessKeyEnv)
	}
	if S3AccessKeyId == "" {
		S3AccessKeyId = os.Getenv(S3AccessKeyIdEnv)
	}
	if S3RoleArn == "" {
		S3RoleArn = os.Getenv(S3RoleArnEnv)
	}
	if GCSURI == "" {
		GCSURI = os.Getenv(GCSURIEnv)
	}
	if GCSServiceAccountKey == "" {
		GCSServiceAccountKey = os.Getenv(GCSServiceAccountKeyEnv)
	}
}

func getDefault() {
	if ServerlessEndpoint == "" {
		ServerlessEndpoint = defaultServerlessEndpoint
	}
	if IamEndpoint == "" {
		IamEndpoint = defaultIamEndpoint
	}
}
