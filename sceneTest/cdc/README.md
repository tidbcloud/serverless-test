# CDC Sync Test

This directory tests if changefeeds can sync data from TiDB Cloud clusters to downstreams.
The upstream TiDB Cloud clusters and downstreams are created manually before running the tests.

- MySQL
  - upstream: 10735492773134730885(alicloud-ap-southeast-1)
  - downstream: 10978086209882933443(alicloud-ap-southeast-1)
- Kafka
  - upstream: 10735492773134730885(alicloud-ap-southeast-1)
  - downstream: msk