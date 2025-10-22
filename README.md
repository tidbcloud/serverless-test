# serverless-test

The serverless-test is a collection of e2e tests based on the Open API for the TiDB Cloud serverless.

The goal of the serverless-test is:

1. Ensure the correctness of any features in TiDB Cloud serverless.
2. Provide examples for the usage of the TiDB Cloud Open API.

## Tested Features

The following features are tested in the serverless-test:

- Import
- Export
- Branch
- Cluster
- Changefeed

### Changefeed Scene Test

The cdc_sync_test.go tests if changefeeds can sync data from TiDB Cloud clusters to downstreams.
The upstream TiDB Cloud clusters and downstreams are created manually before running the tests.

- MySQL
  - upstream: 10735492773134730885(alicloud-ap-southeast-1)
  - downstream: TiDB Starter 10978086209882933443(alicloud-ap-southeast-1)
  - changefeed ID: cfd-lyzks35w7jf6db3vnt4p4is7du(filter test.*)
- Kafka
  - upstream: 10735492773134730885(alicloud-ap-southeast-1)
  - downstream: MSK cluster in ap-southeast-1
  - changefeed ID: cfd-lid7svbc3jeinjxpz3m7qwmhyu(filter kafka.*)

## How to Contribute

- SDK: The SDK is imported from [tidbcloud-cli](https://github.com/tidbcloud/tidbcloud-cli/tree/main/pkg). Open a pull request to the tidbcloud-cli repository if you want to modify the SDK. 
- Common configurations: you can add common configurations in the `config` package.
- Add tests: when you want to add tests for a new feature, create a new test file under sceneTest.
- Add GitHub actions: after adding tests, you can add GitHub actions to run the tests automatically. e.g., [export-scene-test.yml](.github/workflows/export-scene-test.yml).

## LICENSE

Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
