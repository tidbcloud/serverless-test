package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tidbcloud/serverless-test/connectionProbe/probe"
	"gopkg.in/yaml.v2"
)

const concurrency = 5

var (
	password    string
	larkWebhook string
	actionURL   string
)

func init() {
	flag.StringVar(&password, "password", "", "database password")
	flag.StringVar(&larkWebhook, "lark-webhook", "", "lark webhook url")
	flag.StringVar(&actionURL, "action-url", "", "the url for action button in lark notification")
	flag.Parse()
}

// loadConfig loads the database configurations from a YAML file.
// It also add password from command line flag and create two entries for each DB config with port 4000 and 3306 respectively.
func loadConfig(path string) ([]*probe.DBConfig, error) {
	if password == "" {
		log.Fatalf("Password must be provided via -password flag")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var regionDBs map[string][]probe.DBConfig
	if err := yaml.Unmarshal(data, &regionDBs); err != nil {
		return nil, err
	}

	var allDBs []*probe.DBConfig
	for region, dbs := range regionDBs {
		for i := range dbs {
			dbs[i].Region = region
			dbs[i].Password = password

			db1 := dbs[i]
			db2 := dbs[i]
			db1.Port = 4000
			db2.Port = 3306
			allDBs = append(allDBs, &db1)
			allDBs = append(allDBs, &db2)
		}
	}
	return allDBs, nil
}

func main() {
	allDBs, err := loadConfig("./db_config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	jobs := make(chan *probe.DBConfig, len(allDBs))
	notifyCh := make(chan *probe.NotifyInfo, len(allDBs))
	ctx := context.Background()

	for i := 0; i < concurrency; i++ {
		go func() {
			for db := range jobs {
				probe.ProbeDB(ctx, db, notifyCh)
			}
		}()
	}

	for _, db := range allDBs {
		jobs <- db
	}
	close(jobs)

	for i := 0; i < len(allDBs); i++ {
		res := <-notifyCh
		// TODO rocord the result
		if !res.Success {
			probe.NotifyFailure(res, larkWebhook, actionURL)
		}
	}
	close(notifyCh)

	println("All probes completed")
}
