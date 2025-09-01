import { NextRequest, NextResponse } from 'next/server';
import mysql from 'mysql2/promise';

export async function GET(req: NextRequest) {
  const pool = mysql.createPool({
    host: process.env.MYSQL_HOST,
    port: 4000,
    user: process.env.MYSQL_USER,
    password: process.env.MYSQL_PASSWORD,
    database: "test",
    waitForConnections: true,
    connectionLimit: 5,
    ssl: {
      minVersion: 'TLSv1.2',
      rejectUnauthorized: true,
    },
  });
  
  const [availabilityRows] = await pool.query(
    `SELECT region, plan, utc8_date as date,
      COUNT(*) as total,
      SUM(success) as success_count,
      ROUND(100*SUM(success)/COUNT(*),2) as availability,
      APPROX_PERCENTILE(IF(success = 1, latency_ms, NULL), 99) AS p99
    FROM connection_probe_result
    WHERE create_time >= CURDATE() - INTERVAL 61 DAY
    GROUP BY region, plan, utc8_date`
  );
  
  await pool.end();
  return NextResponse.json(availabilityRows);
}
