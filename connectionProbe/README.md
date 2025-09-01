# User guide

This project is used to probe the connection status of TiDB Cloud Starter & Essential clusters in different regions.

1. To manage probe clusters, edit the `connectionProbe/db_config.yaml` file to add or remove TiDB Cloud clusters.
2. View the [probe dashboard](https://tidbcloud-connection-probe.netlify.app/)

# Dev Guide

The project consists of:

1. GitHub Action: The GitHub Action will be triggered every 5 minutes to run the probe code.
2. Probe Code: The probe code will read the `db_config.yaml` file to get the connection info of different regions.
   1. Try to connect and ping the TiDB Cloud cluster.
   2. Notify the Lark channel if the connection fails.
   3. Record the probe result in the meta DB.
3. Frontend: The frontend will read the meta DB and display the probe results in a table.


### Manage GitHub Action

The GitHub Action is defined in the `.github/workflows/connection-probe.yml` file.

You need to configure the following secrets in the GitHub repository:
- PROBE_PASSWORD: The password for all TiDB Cloud clusters (unified for simplicity).
- PROBE_LARK_WEBHOOK: The webhook URL of the Lark channel for notifications.
- PROBE_META_DB_DSN: The DSN in `<user>:<PASSWORD>@tcp(<host>:<port>)/<db>` format of the meta DB to record probe results. (Optional, if not provided, the probe results will not be recorded)

### Manage Meta DB

The DDL of the meta DB is defined in the `connectionProbe/storage/meta_db.sql` file.

The meta DB requires TLS. This project use TiDB Cloud Starter cluster as the meta DB, you can also use any MySQL-compatible database with TLS enabled.

### Manage DB Config

The DB config is defined in the `connectionProbe/db_config.yaml` file. Modify this file to add/remove TiDB Cloud clusters to probe.

### Manage Frontend

The frontend code is in the `connectionProbe/frontend` folder. You can deploy it on Netlify or any other static site hosting service.

Required environment variables:
- MYSQL_HOST: The host of the meta DB.
- MYSQL_PASSWORD: The password of the meta DB.
- MYSQL_USER: The user of the meta DB.

You can also deploy locally for testing, note keep the connection info the same as the meta DB used in GitHub Action:

```bash
cd connectionProbe/frontend
export MYSQL_HOST=<meta_db_host>
export MYSQL_USER=<meta_db_user>
export MYSQL_PASSWORD=<meta_db_password>
npm install
npm run dev
```