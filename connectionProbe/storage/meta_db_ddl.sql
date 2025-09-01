CREATE TABLE test.connection_probe_result (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    region VARCHAR(64) NOT NULL,
    cluster_id VARCHAR(64) NOT NULL,
    plan VARCHAR(32) NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utc8_date varchar(64) NOT NULL,
    err_msg TEXT,
    port INT NOT NULL,
    latency_ms int(64) NOT NULL,
    success tinyint NOT NULL,
    KEY idx_region_date_cluster (region, plan, utc8_date),
    KEY idx_create_time (create_time),
    KEY idx_cluster (cluster_id)
) ;