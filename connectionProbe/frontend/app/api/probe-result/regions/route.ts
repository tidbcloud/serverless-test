import { NextRequest, NextResponse } from 'next/server';
import mysql from 'mysql2/promise';

export async function GET() {
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
  const [rows] = await pool.query(
    `SELECT DISTINCT region FROM connection_probe_result ORDER BY region`
  );
  await pool.end();
  return NextResponse.json(rows.map((r: any) => r.region));
}
