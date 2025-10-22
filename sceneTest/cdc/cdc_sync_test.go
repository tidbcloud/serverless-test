package cdc

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/go-sql-driver/mysql"
	"github.com/tidbcloud/serverless-test/config"
	"github.com/tidbcloud/serverless-test/util"
	"github.com/tidbcloud/tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cdc"
	"github.com/xdg-go/scram"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

const (
	mysqlSyncTable = "test.cdc"
	kafkaSyncTable = "kafka.cdc"
	kafkaTopic     = "cdc-test"
	kafkaEndpoints = "b-2-public.cdctest.pht5v6.c3.kafka.ap-southeast-1.amazonaws.com:9196,b-1-public.cdctest.pht5v6.c3.kafka.ap-southeast-1.amazonaws.com:9196"
)

// TestMySQLSync tests if the MySQL changefeed can sync within 1 minute on alicloud-ap-southeast-1
func TestMySQLSync(t *testing.T) {
	ctx := context.Background()
	log.Println("test mysql sync start")

	cfg := config.LoadConfig().Changefeed.MySQL
	t.Logf("start to test mysql changefeed sync, changefeed_id: %s, cluster_id:%s, region:%s",
		cfg.ChangefeedID,
		cfg.ClusterID,
		cfg.Region)
	cf, err := getChangefeed(ctx, cfg.ClusterID, cfg.ChangefeedID)
	if err != nil {
		t.Fatalf("failed to get changefeed: %v", err)
	}
	if *cf.State != cdc.CHANGEFEEDSTATEENUM_RUNNING {
		t.Fatalf("changefeed is not running, current state: %s", *cf.State)
	}

	ts := time.Now().UnixMilli()
	t.Log("start to insert into upstream tidb cloud cluster")
	err = executeDB(ctx, cfg.ClusterDSN, fmt.Sprintf("insert into %s (id, name) values (%d, 'cdc')", mysqlSyncTable, ts))
	if err != nil {
		t.Fatalf("failed to insert into upstream tidb cloud cluster: %v", err)
	}

	t.Log("wait for 2 minute for data to sync")
	st := time.Now()
	for {
		if time.Since(st) > 2*time.Minute {
			t.Fatalf("data not synced to downstream mysql after 2 minute")
		}
		exist, err := queryDB(ctx, cfg.MySQLDSN, fmt.Sprintf("select id, name from test.cdc where id = %d", ts))
		if err != nil {
			t.Fatalf("failed to query downstream mysql: %v", err)
		}
		if exist {
			log.Println("test mysql sync success")
			return
		}
		time.Sleep(20 * time.Second)
	}
}

// TestMySQLSync tests if the Kafka changefeed can sync within 1 minute on alicloud-ap-southeast-1
func TestKafkaSync(t *testing.T) {
	ctx := context.Background()
	log.Println("test kafka sync start")
	cfg := config.LoadConfig().Changefeed.Kafka
	t.Log(fmt.Sprintf("start to test kafka changefeed sync, changefeed_id: %s, cluster_id:%s, region:%s",
		cfg.ChangefeedID,
		cfg.ClusterID,
		cfg.Region))
	cf, err := getChangefeed(ctx, cfg.ClusterID, cfg.ChangefeedID)
	if err != nil {
		t.Fatalf("failed to get changefeed: %v", err)
	}
	if *cf.State != cdc.CHANGEFEEDSTATEENUM_RUNNING {
		t.Fatalf("changefeed is not running, current state: %s", *cf.State)
	}

	ts := time.Now().UnixMilli()

	t.Log("consume kafka topic to check if data exists")
	messages := make(chan string, 100)
	defer close(messages)
	endpoints := strings.Split(kafkaEndpoints, ",")
	consumer, err := kafkaConsume(endpoints, kafkaTopic, kafkaSASLScramAuth(cfg.KafkaSASLSCRAMUser, cfg.KafkaSASLSCRAMPassword), messages)
	if consumer != nil {
		defer consumer.Close()
	}
	if err != nil {
		t.Fatalf("failed to consume kafka topic: %v", err)
	}

	t.Log("start to insert into upstream tidb cloud cluster")
	err = executeDB(ctx, cfg.ClusterDSN, fmt.Sprintf("insert into %s (id, name) values (%d, 'cdc')", kafkaSyncTable, ts))
	if err != nil {
		t.Fatalf("failed to insert into upstream tidb cloud cluster: %v", err)
	}

	t.Log("wait for 2 minute for data to sync")
	consumeTimeout := time.After(2 * time.Minute)
	found := false
	for {
		select {
		case msg := <-messages:
			t.Logf("received kafka message: %s", msg)
			if strings.Contains(msg, fmt.Sprintf("%d", ts)) {
				found = true
				log.Printf("find kafka message: %s", msg)
				goto Done
			}
		case <-consumeTimeout:
			t.Log("stopping message consumption after timeout")
			goto Done
		}
	}
Done:
	if !found {
		t.Fatal("data not synced in 2 minutes")
	}
	log.Println("test kafka sync success")
}

func executeDB(ctx context.Context, dsn string, query string) (err error) {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
	})

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open mysql connection. error:%s\n", err.Error())
		return err
	}
	defer conn.Close()

	res, err := conn.ExecContext(ctx, query)
	if err != nil {
		fmt.Printf("Failed to execute query. error:%s\n", err.Error())
		return err
	}
	rowAffect, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get rows affected. error:%s\n", err.Error())
		return err
	}
	if rowAffect == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func queryDB(ctx context.Context, dsn string, query string) (exist bool, err error) {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
	})

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open mysql connection. error:%s\n", err.Error())
		return false, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		fmt.Printf("Failed to execute query. error:%s\n", err.Error())
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func getChangefeed(ctx context.Context, clusterId, changefeedId string) (*cdc.Changefeed, error) {
	r := cdcClient.ChangefeedServiceAPI.ChangefeedServiceGetChangefeed(ctx, clusterId, changefeedId)
	res, h, err := r.Execute()
	return res, util.ParseError(err, h)
}

func kafkaConsume(endpoints []string, topic string, config *sarama.Config, messages chan<- string) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer(endpoints, config)
	if err != nil {
		return nil, fmt.Errorf("Failed to start consumer: %v", err)
	}

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return consumer, fmt.Errorf("Failed to get the list of partitions: %v", err)
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return consumer, fmt.Errorf("Failed to get the list of partitions: %v", err)
		}
		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()
			for {
				select {
				case msg := <-pc.Messages():
					if messages != nil {
						messages <- string(msg.Value)
					}
				}
			}
		}(pc)
	}
	return consumer, nil
}

func kafkaSASLScramAuth(username, password string) *sarama.Config {
	config := sarama.NewConfig()
	config.Net.SASL.Enable = true
	config.Net.SASL.User = username
	config.Net.SASL.Password = password
	config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
		return &XDGSCRAMClient{HashGeneratorFcn: scram.SHA512}
	}
	config.Net.TLS.Enable = true
	return config
}

type XDGSCRAMClient struct {
	*scram.Client
	conv             *scram.ClientConversation
	HashGeneratorFcn scram.HashGeneratorFcn
}

func (x *XDGSCRAMClient) Begin(user, password, authzID string) error {
	c, err := scram.SHA512.NewClient(user, password, authzID)
	if err != nil {
		return err
	}
	x.Client = c
	x.conv = c.NewConversation()
	return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (string, error) {
	return x.conv.Step(challenge)
}

func (x *XDGSCRAMClient) Done() bool {
	return x.conv.Done()
}
