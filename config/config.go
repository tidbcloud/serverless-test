package config

import (
	"bytes"
	"flag"
	"sync"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Endpoint struct {
	Serverless string `mapstructure:"serverless"`
	IAM        string `mapstructure:"iam"`
}

type Azure struct {
	URI      string `mapstructure:"uri"`
	SASToken string `mapstructure:"sas-token"`
}

type S3 struct {
	URI             string `mapstructure:"uri"`
	SecretAccessKey string `mapstructure:"secret-access-key"`
	AccessKeyID     string `mapstructure:"access-key-id"`
	RoleARN         string `mapstructure:"role-arn"`
}

type GCS struct {
	URI               string `mapstructure:"uri"`
	ServiceAccountKey string `mapstructure:"service-account-key"`
}

type ImportS3 struct {
	RoleARN                         string `mapstructure:"role-arn"`
	ParquetURI                      string `mapstructure:"parquet-uri"`
	SchemaCompressURI               string `mapstructure:"schema-compress-uri"`
	SchemaTypeMismatchedURI         string `mapstructure:"schema-type-mismatched-uri"`
	SchemaColumnNumberMismatchedURI string `mapstructure:"schema-column-number-mismatched-uri"`
	URI                             string `mapstructure:"uri"`
	RoleARNNoPrivilege              string `mapstructure:"role-arn-no-privilege"`
	RoleARNDiffExternalID           string `mapstructure:"role-arn-diff-external-id"`
	AccessKeyIDNoPrivilege          string `mapstructure:"access-key-id-no-privilege"`
	SecretAccessKeyNoPrivilege      string `mapstructure:"secret-access-key-no-privilege"`
}

type ImportAzure struct {
	URI                 string `mapstructure:"uri"`
	SASToken            string `mapstructure:"sas-token"`
	SASTokenNoPrivilege string `mapstructure:"sas-token-no-privilege"`
}

type ImportGCS struct {
	URI                          string `mapstructure:"uri"`
	ServiceAccountKey            string `mapstructure:"service-account-key"`
	ServiceAccountKeyNoPrivilege string `mapstructure:"service-account-key-no-privilege"`
}

type ImportOSS struct {
	URI                        string `mapstructure:"uri"`
	AccessKeyID                string `mapstructure:"access-key-id"`
	SecretAccessKey            string `mapstructure:"secret-access-key"`
	AccessKeyIDNoPrivilege     string `mapstructure:"access-key-id-no-privilege"`
	SecretAccessKeyNoPrivilege string `mapstructure:"secret-access-key-no-privilege"`
	RoleARN                    string `mapstructure:"role-arn"`
	RoleARNNoPrivilege         string `mapstructure:"role-arn-no-privilege"`
	RoleARNDiffExternalID      string `mapstructure:"role-arn-diff-external-id"`
}

type Import struct {
	ClusterHost     string      `mapstructure:"cluster-host"`
	ClusterUser     string      `mapstructure:"cluster-user"`
	ClusterPassword string      `mapstructure:"cluster-password"`
	S3              ImportS3    `mapstructure:"s3"`
	Azure           ImportAzure `mapstructure:"azure"`
	GCS             ImportGCS   `mapstructure:"gcs"`
	OSS             ImportOSS   `mapstructure:"oss"`
}

type Config struct {
	PublicKey  string   `mapstructure:"public-key"`
	PrivateKey string   `mapstructure:"private-key"`
	Endpoint   Endpoint `mapstructure:"endpoint"`

	ConsoleAPIHost    string `mapstructure:"console-api-host"`
	Auth0Domain       string `mapstructure:"auth0-domain"`
	Auth0ClientID     string `mapstructure:"auth0-client-id"`
	Auth0ClientSecret string `mapstructure:"auth0-client-secret"`
	UserEmail         string `mapstructure:"user-email"`

	Azure     Azure  `mapstructure:"azure"`
	S3        S3     `mapstructure:"s3"`
	GCS       GCS    `mapstructure:"gcs"`
	ProjectID string `mapstructure:"project-id"`
	Import    Import `mapstructure:"import"`
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
	flag.StringVar(&configContent, "config", "", "")
	flag.StringVar(&configAddress, "config-address", ".", "")
	flag.StringVar(&cfg.PublicKey, "public-key", "", "")
	flag.StringVar(&cfg.PrivateKey, "private-key", "", "")
	flag.StringVar(&cfg.Endpoint.Serverless, "endpoint.serverless", defaultServerlessEndpoint, "")
	flag.StringVar(&cfg.Endpoint.IAM, "endpoint.iam", defaultIamEndpoint, "")
	flag.StringVar(&cfg.ConsoleAPIHost, "console-api-host", "", "")
	flag.StringVar(&cfg.Auth0Domain, "auth0-domain", "", "")
	flag.StringVar(&cfg.Auth0ClientID, "auth0-client-id", "", "")
	flag.StringVar(&cfg.Auth0ClientSecret, "auth0-client-secret", "", "")
	flag.StringVar(&cfg.UserEmail, "user-email", "", "")
	flag.StringVar(&cfg.Azure.URI, "azure.uri", "", "")
	flag.StringVar(&cfg.Azure.SASToken, "azure.sas-token", "", "")
	flag.StringVar(&cfg.S3.URI, "s3.uri", "", "")
	flag.StringVar(&cfg.S3.SecretAccessKey, "s3.secret-access-key", "", "")
	flag.StringVar(&cfg.S3.AccessKeyID, "s3.access-key-id", "", "")
	flag.StringVar(&cfg.S3.RoleARN, "s3.role-arn", "", "")
	flag.StringVar(&cfg.GCS.URI, "gcs.uri", "", "")
	flag.StringVar(&cfg.GCS.ServiceAccountKey, "gcs.service-account-key", "", "")
	flag.StringVar(&cfg.ProjectID, "project-id", "", "")
	flag.StringVar(&cfg.Import.ClusterHost, "import.cluster-host", "", "")
	flag.StringVar(&cfg.Import.ClusterUser, "import.cluster-user", "", "")
	flag.StringVar(&cfg.Import.ClusterPassword, "import.cluster-password", "", "")
	flag.StringVar(&cfg.Import.S3.RoleARN, "import.s3.role-arn", "", "")
	flag.StringVar(&cfg.Import.S3.ParquetURI, "import.s3.parquet-uri", "", "")
	flag.StringVar(&cfg.Import.S3.SchemaCompressURI, "import.s3.schema-compress-uri", "", "")
	flag.StringVar(&cfg.Import.S3.SchemaTypeMismatchedURI, "import.s3.schema-type-mismatched-uri", "", "")
	flag.StringVar(&cfg.Import.S3.SchemaColumnNumberMismatchedURI, "import.s3.schema-column-number-mismatched-uri", "", "")
	flag.StringVar(&cfg.Import.Azure.URI, "import.azure.uri", "", "")
	flag.StringVar(&cfg.Import.Azure.SASToken, "import.azure.sas-token", "", "")
	flag.StringVar(&cfg.Import.Azure.SASTokenNoPrivilege, "import.azure.sas-token-no-privilege", "", "")
	flag.StringVar(&cfg.Import.GCS.URI, "import.gcs.uri", "", "")
	flag.StringVar(&cfg.Import.GCS.ServiceAccountKey, "import.gcs.service-account-key", "", "")
	flag.StringVar(&cfg.Import.GCS.ServiceAccountKeyNoPrivilege, "import.gcs.service-account-key-no-privilege", "", "")
	flag.StringVar(&cfg.Import.S3.URI, "import.s3.uri", "", "")
	flag.StringVar(&cfg.Import.S3.RoleARNNoPrivilege, "import.s3.role-arn-no-privilege", "", "")
	flag.StringVar(&cfg.Import.S3.RoleARNDiffExternalID, "import.s3.role-arn-diff-external-id", "", "")
	flag.StringVar(&cfg.Import.S3.AccessKeyIDNoPrivilege, "import.s3.access-key-id-no-privilege", "", "")
	flag.StringVar(&cfg.Import.S3.SecretAccessKeyNoPrivilege, "import.s3.secret-access-key-no-privilege", "", "")
	flag.StringVar(&cfg.Import.OSS.URI, "import.oss.uri", "", "")
	flag.StringVar(&cfg.Import.OSS.AccessKeyID, "import.oss.access-key-id", "", "")
	flag.StringVar(&cfg.Import.OSS.SecretAccessKey, "import.oss.secret-access-key", "", "")
	flag.StringVar(&cfg.Import.OSS.AccessKeyIDNoPrivilege, "import.oss.access-key-id-no-privilege", "", "")
	flag.StringVar(&cfg.Import.OSS.SecretAccessKeyNoPrivilege, "import.oss.secret-access-key-no-privilege", "", "")
	flag.StringVar(&cfg.Import.OSS.RoleARN, "import.oss.role-arn", "", "")
	flag.StringVar(&cfg.Import.OSS.RoleARNNoPrivilege, "import.oss.role-arn-no-privilege", "", "")
	flag.StringVar(&cfg.Import.OSS.RoleARNDiffExternalID, "import.oss.role-arn-diff-external-id", "", "")

	// need to act this like since testing.Run will call flag.Parse() if not parsed
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
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
