package config

import (
	"bytes"
	"flag"
	"os"

	"github.com/spf13/viper"
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

	ProjectId string
)

var (
	configContent string
	configAddress string
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

	ProjectIdEnv = "PROJECT_ID"

	GCSURIEnv               = "GCS_URI"
	GCSServiceAccountKeyEnv = "GCS_SERVICE_ACCOUNT_KEY"
)

func init() {
	flag.StringVar(&PublicKey, "public-key", "", "")
	flag.StringVar(&PrivateKey, "private-key", "", "")
	flag.StringVar(&ServerlessEndpoint, "endpoint.serverless", "", "")
	flag.StringVar(&IamEndpoint, "endpoint.iam", "", "")

	flag.StringVar(&S3URI, "s3.uri", "", "")
	flag.StringVar(&S3SecretAccessKey, "s3.secret-access-key", "", "")
	flag.StringVar(&S3AccessKeyId, "s3.access-key-id", "", "")
	flag.StringVar(&S3RoleArn, "s3.role-arn", "", "")

	flag.StringVar(&AzureURI, "azure.uri", "", "")
	flag.StringVar(&AzureSASToken, "azure.sas-token", "", "")

	flag.StringVar(&GCSURI, "gcs.uri", "", "")
	flag.StringVar(&GCSServiceAccountKey, "gcs.service-account-key", "", "")

	flag.StringVar(&ProjectId, "project-id", "", "")

	flag.StringVar(&configContent, "config", "", "")
	flag.StringVar(&configAddress, "config-address", ".", "")
}

func InitializeConfig() {
	// priority: flag > env > config > default
	flag.Parse()
	getEnvironment()
	getConfig()
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
	if ProjectId == "" {
		ProjectId = os.Getenv(ProjectIdEnv)
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

func getConfig() {
	viper.SetConfigType("toml")
	if configContent != "" {
		err := viper.ReadConfig(bytes.NewBuffer([]byte(configContent)))
		if err != nil {
			println("Error reading config file: ", err.Error())
			return
		}
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(configAddress)
		err := viper.ReadInConfig()
		if err != nil {
			println("Error reading config file: ", err.Error())
			return
		}
	}

	if PublicKey == "" {
		PublicKey = viper.GetString("public-key")
	}
	if PrivateKey == "" {
		PrivateKey = viper.GetString("private-key")
	}
	if ServerlessEndpoint == "" {
		ServerlessEndpoint = viper.GetString("endpoint.serverless")
	}
	if IamEndpoint == "" {
		IamEndpoint = viper.GetString("endpoint.iam")
	}
	if AzureURI == "" {
		AzureURI = viper.GetString("azure.uri")
	}
	if AzureSASToken == "" {
		AzureSASToken = viper.GetString("azure.sas-token")
	}
	if S3URI == "" {
		S3URI = viper.GetString("s3.uri")
	}
	if S3SecretAccessKey == "" {
		S3SecretAccessKey = viper.GetString("s3.secret-access-key")
	}
	if S3AccessKeyId == "" {
		S3AccessKeyId = viper.GetString("s3.access-key-id")
	}
	if S3RoleArn == "" {
		S3RoleArn = viper.GetString("s3.role-arn")
	}
	if GCSURI == "" {
		GCSURI = viper.GetString("gcs.uri")
	}
	if GCSServiceAccountKey == "" {
		GCSServiceAccountKey = viper.GetString("gcs.service-account-key")
	}
	if ProjectId == "" {
		ProjectId = viper.GetString("project-id")
	}
}
