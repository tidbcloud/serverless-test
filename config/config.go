package config

import (
	"bytes"
	"sync"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	PublicKey          string `mapstructure:"public-key"`
	PrivateKey         string `mapstructure:"private-key"`
	ServerlessEndpoint string `mapstructure:"endpoint.serverless"`
	IamEndpoint        string `mapstructure:"endpoint.iam"`

	ConsoleAPIHost    string `mapstructure:"console-api-host"`
	Auth0Domain       string `mapstructure:"auth0-domain"`
	Auth0ClientID     string `mapstructure:"auth0-client-id"`
	Auth0ClientSecret string `mapstructure:"auth0-client-secret"`
	UserEmail         string `mapstructure:"user-email"`

	AzureURI      string `mapstructure:"azure.uri"`
	AzureSASToken string `mapstructure:"azure.sas-token"`

	S3URI             string `mapstructure:"s3.uri"`
	S3SecretAccessKey string `mapstructure:"s3.secret-access-key"`
	S3AccessKeyID     string `mapstructure:"s3.access-key-id"`
	S3RoleARN         string `mapstructure:"s3.role-arn"`

	GCSURI               string `mapstructure:"gcs.uri"`
	GCSServiceAccountKey string `mapstructure:"gcs.service-account-key"`

	ProjectID string `mapstructure:"project-id"`

	ImportClusterHost     string `mapstructure:"import.cluster-host"`
	ImportClusterUser     string `mapstructure:"import.cluster-user"`
	ImportClusterPassword string `mapstructure:"import.cluster-password"`

	ImportS3RoleARN                         string `mapstructure:"import.s3.role-arn"`
	ImportS3ParquetURI                      string `mapstructure:"import.s3.parquet-uri"`
	ImportS3SchemaCompressURI               string `mapstructure:"import.s3.schema-compress-uri"`
	ImportS3SchemaTypeMismatchedURI         string `mapstructure:"import.s3.schema-type-mismatched-uri"`
	ImportS3SchemaColumnNumberMismatchedURI string `mapstructure:"import.s3.schema-column-number-mismatched-uri"`

	ImportAzureURI                 string `mapstructure:"import.azure.uri"`
	ImportAzureSASToken            string `mapstructure:"import.azure.sas-token"`
	ImportAzureSASTokenNoPrivilege string `mapstructure:"import.azure.sas-token-no-privilege"`

	ImportGCSURI                          string `mapstructure:"import.gcs.uri"`
	ImportGCSServiceAccountKey            string `mapstructure:"import.gcs.service-account-key"`
	ImportGCSServiceAccountKeyNoPrivilege string `mapstructure:"import.gcs.service-account-key-no-privilege"`

	ImportS3URI                        string `mapstructure:"import.s3.uri"`
	ImportS3RoleARNNoPrivilege         string `mapstructure:"import.s3.role-arn-no-privilege"`
	ImportS3RoleARNDiffExternalID      string `mapstructure:"import.s3.role-arn-diff-external-id"`
	ImportS3AccessKeyIDNoPrivilege     string `mapstructure:"import.s3.access-key-id-no-privilege"`
	ImportS3SecretAccessKeyNoPrivilege string `mapstructure:"import.s3.secret-access-key-no-privilege"`

	ImportOSSURI                        string `mapstructure:"import.oss.uri"`
	ImportOSSAccessKeyID                string `mapstructure:"import.oss.access-key-id"`
	ImportOSSSecretAccessKey            string `mapstructure:"import.oss.secret-access-key"`
	ImportOSSAccessKeyIDNoPrivilege     string `mapstructure:"import.oss.access-key-id-no-privilege"`
	ImportOSSSecretAccessKeyNoPrivilege string `mapstructure:"import.oss.secret-access-key-no-privilege"`
	ImportOSSRoleARN                    string `mapstructure:"import.oss.role-arn"`
	ImportOSSRoleARNNoPrivilege         string `mapstructure:"import.oss.role-arn-no-privilege"`
	ImportOSSRoleARNDiffExternalID      string `mapstructure:"import.oss.role-arn-diff-external-id"`
}

var (
	configInstance *Config
	once           sync.Once
	configContent  string
	configAddress  string
)

const (
	defaultServerlessEndpoint = "https://serverless.tidbapi.com"
	defaultIamEndpoint        = "https://iam.tidbapi.com"
)

func init() {
	pflag.StringVar(&configContent, "config", "", "")
	pflag.StringVar(&configAddress, "config-address", ".", "")
}

func LoadConfig() *Config {
	once.Do(func() {
		configInstance = &Config{}
		if err := initializeConfig(configInstance); err != nil {
			panic(err)
		}
	})
	return configInstance
}

func initializeConfig(cfg *Config) error {
	pflag.StringVar(&cfg.PublicKey, "public-key", "", "")
	pflag.StringVar(&cfg.PrivateKey, "private-key", "", "")
	pflag.StringVar(&cfg.ServerlessEndpoint, "endpoint.serverless", "", "")
	pflag.StringVar(&cfg.IamEndpoint, "endpoint.iam", "", "")
	pflag.StringVar(&cfg.ConsoleAPIHost, "console-api-host", "", "")
	pflag.StringVar(&cfg.Auth0Domain, "auth0-domain", "", "")
	pflag.StringVar(&cfg.Auth0ClientID, "auth0-client-id", "", "")
	pflag.StringVar(&cfg.Auth0ClientSecret, "auth0-client-secret", "", "")
	pflag.StringVar(&cfg.UserEmail, "user-email", "", "")
	pflag.StringVar(&cfg.S3URI, "s3.uri", "", "")
	pflag.StringVar(&cfg.S3SecretAccessKey, "s3.secret-access-key", "", "")
	pflag.StringVar(&cfg.S3AccessKeyID, "s3.access-key-id", "", "")
	pflag.StringVar(&cfg.S3RoleARN, "s3.role-arn", "", "")
	pflag.StringVar(&cfg.AzureURI, "azure.uri", "", "")
	pflag.StringVar(&cfg.AzureSASToken, "azure.sas-token", "", "")
	pflag.StringVar(&cfg.GCSURI, "gcs.uri", "", "")
	pflag.StringVar(&cfg.GCSServiceAccountKey, "gcs.service-account-key", "", "")
	pflag.StringVar(&cfg.ProjectID, "project-id", "", "")
	pflag.StringVar(&cfg.ImportClusterHost, "import.cluster-host", "", "")
	pflag.StringVar(&cfg.ImportClusterUser, "import.cluster-user", "", "")
	pflag.StringVar(&cfg.ImportClusterPassword, "import.cluster-password", "", "")
	pflag.StringVar(&cfg.ImportS3RoleARN, "import.s3.role-arn", "", "")
	pflag.StringVar(&cfg.ImportS3ParquetURI, "import.s3.parquet-uri", "", "")
	pflag.StringVar(&cfg.ImportS3SchemaCompressURI, "import.s3.schema-compress-uri", "", "")
	pflag.StringVar(&cfg.ImportS3SchemaTypeMismatchedURI, "import.s3.schema-type-mismatched-uri", "", "")
	pflag.StringVar(&cfg.ImportS3SchemaColumnNumberMismatchedURI, "import.s3.schema-column-number-mismatched-uri", "", "")
	pflag.StringVar(&cfg.ImportAzureURI, "import.azure.uri", "", "")
	pflag.StringVar(&cfg.ImportAzureSASToken, "import.azure.sas-token", "", "")
	pflag.StringVar(&cfg.ImportAzureSASTokenNoPrivilege, "import.azure.sas-token-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportGCSURI, "import.gcs.uri", "", "")
	pflag.StringVar(&cfg.ImportGCSServiceAccountKey, "import.gcs.service-account-key", "", "")
	pflag.StringVar(&cfg.ImportGCSServiceAccountKeyNoPrivilege, "import.gcs.service-account-key-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportS3URI, "import.s3.uri", "", "")
	pflag.StringVar(&cfg.ImportS3RoleARNNoPrivilege, "import.s3.role-arn-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportS3RoleARNDiffExternalID, "import.s3.role-arn-diff-external-id", "", "")
	pflag.StringVar(&cfg.ImportS3AccessKeyIDNoPrivilege, "import.s3.access-key-id-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportS3SecretAccessKeyNoPrivilege, "import.s3.secret-access-key-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportOSSURI, "import.oss.uri", "", "")
	pflag.StringVar(&cfg.ImportOSSAccessKeyID, "import.oss.access-key-id", "", "")
	pflag.StringVar(&cfg.ImportOSSSecretAccessKey, "import.oss.secret-access-key", "", "")
	pflag.StringVar(&cfg.ImportOSSAccessKeyIDNoPrivilege, "import.oss.access-key-id-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportOSSSecretAccessKeyNoPrivilege, "import.oss.secret-access-key-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportOSSRoleARN, "import.oss.role-arn", "", "")
	pflag.StringVar(&cfg.ImportOSSRoleARNNoPrivilege, "import.oss.role-arn-no-privilege", "", "")
	pflag.StringVar(&cfg.ImportOSSRoleARNDiffExternalID, "import.oss.role-arn-diff-external-id", "", "")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigType("toml")
	if configContent != "" {
		err := viper.ReadConfig(bytes.NewBuffer([]byte(configContent)))
		if err != nil {
			return err
		}
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(configAddress)
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	viper.SetDefault("endpoint.serverless", defaultServerlessEndpoint)
	viper.SetDefault("endpoint.iam", defaultIamEndpoint)
	return viper.Unmarshal(cfg)
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
