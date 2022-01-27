package repository

import (
	"github.com/roistat/go-clickhouse"
)

type LogClickHouse struct {
	conn *clickhouse.Conn
}

func NewLogClickHouse(conn *clickhouse.Conn) *LogClickHouse {
	return &LogClickHouse{conn: conn}
}

func (l *LogClickHouse) WriteLog(message string) error {
	q := clickhouse.NewQuery("INSERT INTO default.logs VALUES (toDate(now()), ?, ?)", 1, message)
	err := q.Exec(l.conn)

	return err
}
