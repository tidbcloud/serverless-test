package config

import (
	"bytes"
	"flag"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	PublicKey          string
	PrivateKey         string
	ServerlessEndpoint string
	IamEndpoint        string

	ConsoleApiHost    string
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
	UserEmail         string

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

// import vars
var (
	ImportClusterHost     string
	ImportClusterUser     string
	ImportClusterPassWord string

	ImportS3RoleArn                         string
	ImportS3ParquetURI                      string
	ImportS3SchemaCompressURI               string
	ImportS3SchemaTypeMisMatchedURI         string
	ImportS3SchemaColumnNumberMismatchedURI string

	ImportAzureURI                 string
	ImportAzureSASToken            string
	ImportAzureSASTokenNoPrivilege string

	ImportGCSURI                          string
	ImportGCSServiceAccountKey            string
	ImportGCSServiceAccountKeyNoPrivilege string

	ImportS3URI                        string
	ImportS3RoleArnNoPrivilege         string
	ImportS3RoleArnDiffExternalID      string
	ImportS3AccessKeyIdNoPrivilege     string
	ImportS3SecretAccessKeyNoPrivilege string

	ImportOSSURI                        string
	ImportOSSAccessKeyId                string
	ImportOSSSecretAccessKey            string
	ImportOSSAccessKeyIdNoPrivilege     string
	ImportOSSSecretAccessKeyNoPrivilege string
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

	flag.StringVar(&ConsoleApiHost, "console-api-host", "", "")
	flag.StringVar(&Auth0Domain, "auth0-domain", "", "")
	flag.StringVar(&Auth0ClientID, "auth0-client-id", "", "")
	flag.StringVar(&Auth0ClientSecret, "auth0-client-secret", "", "")
	flag.StringVar(&UserEmail, "user-email", "", "")

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
			println("Error reading config: ", err.Error())
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

	if ImportClusterHost == "" {
		ImportClusterHost = viper.GetString("import.cluster-host")
	}
	if ImportClusterUser == "" {
		ImportClusterUser = viper.GetString("import.cluster-user")
	}
	if ImportClusterPassWord == "" {
		ImportClusterPassWord = viper.GetString("import.cluster-password")
	}
	if ImportS3RoleArn == "" {
		ImportS3RoleArn = viper.GetString("import.s3.role-arn")
	}
	if ImportS3ParquetURI == "" {
		ImportS3ParquetURI = viper.GetString("import.s3.parquet-uri")
	}
	if ImportS3SchemaCompressURI == "" {
		ImportS3SchemaCompressURI = viper.GetString("import.s3.schema-compress-uri")
	}
	if ImportS3SchemaTypeMisMatchedURI == "" {
		ImportS3SchemaTypeMisMatchedURI = viper.GetString("import.s3.schema-type-mismatched-uri")
	}
	if ImportS3SchemaColumnNumberMismatchedURI == "" {
		ImportS3SchemaColumnNumberMismatchedURI = viper.GetString("import.s3.schema-column-number-mismatched-uri")
	}
	if ImportAzureURI == "" {
		ImportAzureURI = viper.GetString("import.azure.uri")
	}
	if ImportAzureSASToken == "" {
		ImportAzureSASToken = viper.GetString("import.azure.sas-token")
	}
	if ImportAzureSASTokenNoPrivilege == "" {
		ImportAzureSASTokenNoPrivilege = viper.GetString("import.azure.sas-token-no-privilege")
	}
	if ImportGCSURI == "" {
		ImportGCSURI = viper.GetString("import.gcs.uri")
	}
	if ImportGCSServiceAccountKey == "" {
		ImportGCSServiceAccountKey = viper.GetString("import.gcs.service-account-key")
	}
	if ImportGCSServiceAccountKeyNoPrivilege == "" {
		ImportGCSServiceAccountKeyNoPrivilege = viper.GetString("import.gcs.service-account-key-no-privilege")
	}
	if ImportS3URI == "" {
		ImportS3URI = viper.GetString("import.s3.uri")
	}
	if ImportS3RoleArnNoPrivilege == "" {
		ImportS3RoleArnNoPrivilege = viper.GetString("import.s3.role-arn-no-privilege")
	}
	if ImportS3RoleArnDiffExternalID == "" {
		ImportS3RoleArnDiffExternalID = viper.GetString("import.s3.role-arn-diff-external-id")
	}
	if ImportS3AccessKeyIdNoPrivilege == "" {
		ImportS3AccessKeyIdNoPrivilege = viper.GetString("import.s3.access-key-id-no-privilege")
	}
	if ImportS3SecretAccessKeyNoPrivilege == "" {
		ImportS3SecretAccessKeyNoPrivilege = viper.GetString("import.s3.secret-access-key-no-privilege")
	}
	if ImportOSSURI == "" {
		ImportOSSURI = viper.GetString("import.oss.uri")
	}
	if ImportOSSAccessKeyId == "" {
		ImportOSSAccessKeyId = viper.GetString("import.oss.access-key-id")
	}
	if ImportOSSSecretAccessKey == "" {
		ImportOSSSecretAccessKey = viper.GetString("import.oss.secret-access-key")
	}
	if ImportOSSAccessKeyIdNoPrivilege == "" {
		ImportOSSAccessKeyIdNoPrivilege = viper.GetString("import.oss.access-key-id-no-privilege")
	}
	if ImportOSSSecretAccessKeyNoPrivilege == "" {
		ImportOSSSecretAccessKeyNoPrivilege = viper.GetString("import.oss.secret-access-key-no-privilege")
	}

	if ConsoleApiHost == "" {
		ConsoleApiHost = viper.GetString("console-api-host")
	}
	if Auth0Domain == "" {
		Auth0Domain = viper.GetString("auth0-domain")
	}
	if Auth0ClientID == "" {
		Auth0ClientID = viper.GetString("auth0-client-id")
	}
	if Auth0ClientSecret == "" {
		Auth0ClientSecret = viper.GetString("auth0-client-secret")
	}
	if UserEmail == "" {
		UserEmail = viper.GetString("user-email")
	}
}

func GetRandomRegion() string {
	regionLists := []string{
		"regions/aws-us-west-2",
		"regions/aws-us-east-1",
		"regions/aws-ap-northeast-1",
		"regions/aws-ap-southeast-1",
		"regions/aws-eu-central-1",
		"regions/alicloud-ap-southeast-1",
	}
	size := len(regionLists)
	return regionLists[time.Now().Unix()%int64(size)]
}
