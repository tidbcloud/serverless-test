package cdc

import "github.com/tidbcloud/serverless-test/config"

var clusterId string

func init() {
	cfg := config.LoadConfig()
	println(cfg.Changefeed)
}
