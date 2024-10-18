package config

import (
	"flag"
	"os"
)

var (
	S3URI             string
	S3SecretAccessKey string
	S3AccessKeyId     string
)

const (
	S3URIEnv             = "S3_URI"
	S3SecretAccessKeyEnv = "S3_SECRET_ACCESS_KEY"
	S3AccessKeyIdEnv     = "S3_ACCESS_KEY_ID"
)

func init() {
	flag.StringVar(&S3URI, "s3.uri", "", "")
	flag.StringVar(&S3SecretAccessKey, "s3.secret-access-key", "", "")
	flag.StringVar(&S3AccessKeyId, "s3.access-key-id", "", "")
}

func GetS3URI() string {
	if S3URI != "" {
		return S3URI
	}
	return os.Getenv(S3URIEnv)
}

func GetS3AccessKey() (string, string) {
	if S3SecretAccessKey != "" && S3AccessKeyId != "" {
		return S3AccessKeyId, S3SecretAccessKey
	}
	return os.Getenv(S3AccessKeyIdEnv), os.Getenv(S3SecretAccessKeyEnv)
}
