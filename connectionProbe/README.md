# Dev guide

This project is used to probe the connection status of TiDB Cloud starter&essential in different regions.

The prject consists of:
1. Github Action: The github action will be triggered every 5 minutes, and run the probe code.
2. Probe Code: The probe code will read the db_config.yaml file to get the connection info of different regions.
   1. Try to connect and ping the TiDB Cloud cluster.
   2. Notify the slack channel if the connection fails.
   3. Reocrd the probe result in the meta db.
3. Frontend: The frontend will read the meta db and display the probe result in a table.
   1. can be deployed on Netlify.

# User Guide

### manage github action

The github action is defined in the `.github/workflows/connection_probe.yml` file.

Add you need to configure the following secrets in the github repo:
- PROBE_PASSWORD: the password of all the TiDB Cloud cluster (unify the password for simplicity).
- PROBE_LARK_WEBHOOK: The webhook url of the slack channel to notify.
- PROBE_META_DB_DSN: The DSN in "<user>:<PASSWORD>@tcp(<hosr>:<port>)/<db>" format of the meta db to record the probe result. (Optinal, if not provided, the probe result will not be recorded)

### manage meta db

The DDL of the meta db is defined in the `connectionProbe/storage/meta_db.sql` file.

The meta db needs open TLS in default.

### manage db config

The db config is defined in the `connectionProbe/db_config.yaml` file. Modify this file to add/remove the TiDB Cloud cluster to probe.

### manage frontend

The frontend code is in the `connectionProbe/frontend` folder. You can deploy it on Netlify or any other static site hosting service.

Needed env variables:
- MYSQL_HOST: The host of the meta db.
- MYSQL_PASSWORD: The password of the meta db.
- MYSQL_USER: The user of the meta db.


