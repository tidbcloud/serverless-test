package storage

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type ProbeResult struct {
	Region    string
	ClusterID string
	Plan      string
	UTC8Date  string
	ErrMsg    string
	Port      int
	LatencyMs int64
	Success   bool
}

type Storage struct {
	db *sql.DB
}

func NewStorage(dsn string) (*Storage, error) {
	if dsn == "" {
		return nil, errors.New("empty DSN for storage")
	}
	// Parse the DSN to extract the host
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN %s: %v", dsn, err)
	}
	// Register TLS config with the extracted host as ServerName
	host := cfg.Addr
	if idx := strings.Index(host, ":"); idx > 0 {
		host = host[:idx]
	}
	mysql.RegisterTLSConfig("meta", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: host,
	})
	dsn = fmt.Sprintf("%s?tls=meta", dsn)

	printDsn := dsn
	println(fmt.Sprintf("dns: %s", printDsn))

	// // Update DSN to use the TLS config
	// cfg.TLSConfig = "meta"
	// dsn = cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping mysql: %v", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	if s == nil {
		return nil
	}
	return s.db.Close()
}

func (s *Storage) InsertProbeResult(result *ProbeResult) {
	if s == nil {
		return
	}
	query := `
        INSERT INTO test.connection_probe_result 
        (region, cluster_id, plan, utc8_date, err_msg, port, latency_ms, success) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `

	success := 0
	if result.Success {
		success = 1
	}

	_, err := s.db.Exec(query,
		result.Region,
		result.ClusterID,
		result.Plan,
		result.UTC8Date,
		result.ErrMsg,
		result.Port,
		result.LatencyMs,
		success,
	)
	if err != nil {
		println("Failed to insert probe result:", err.Error())
	}
	fmt.Printf("Insert probe success: %s(%d)\n", result.ClusterID, result.Port)
	return
}

func (s *Storage) CleanProbeResults() {
	if s == nil {
		return
	}

	hour := time.Now().Hour()
	minutes := time.Now().Minute()
	// only clean between 3:00 and 3:30 AM
	if hour != 3 || minutes > 30 {
		return
	}

	query := `
		DELETE FROM test.connection_probe_result 
		WHERE create_time < DATE_SUB(NOW(), INTERVAL 100 DAY)
	`
	result, err := s.db.Exec(query)
	if err != nil {
		fmt.Printf("Failed to clean probe results: %v\n", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows: %v\n", err)
		return
	}

	fmt.Printf("Successfully cleaned %d probe results older than 100 days\n", rowsAffected)
}

func (s *Storage) InsertProbeResults(results []*ProbeResult) {
	if s == nil || len(results) == 0 {
		return
	}

	// 构建批量插入的 SQL
	valueStrings := make([]string, 0, len(results))
	valueArgs := make([]interface{}, 0, len(results)*8) // 8 是每行的字段数

	for _, result := range results {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?)")
		success := 0
		if result.Success {
			success = 1
		}
		valueArgs = append(valueArgs,
			result.Region,
			result.ClusterID,
			result.Plan,
			result.UTC8Date,
			result.ErrMsg,
			result.Port,
			result.LatencyMs,
			success,
		)
	}

	query := fmt.Sprintf(`
        INSERT INTO test.connection_probe_result 
        (region, cluster_id, plan, utc8_date, err_msg, port, latency_ms, success) 
        VALUES %s
    `, strings.Join(valueStrings, ","))

	_, err := s.db.Exec(query, valueArgs...)
	if err != nil {
		println("Failed to batch insert probe results:", err.Error())
		return
	}

	fmt.Printf("Successfully inserted %d probe results\n", len(results))
	return
}
