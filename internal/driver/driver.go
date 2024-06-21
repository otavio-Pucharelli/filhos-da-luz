package driver

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

// DB holds the database connection pool.
type DB struct {
	SQL *pgx.Conn
}

var dbConn = &DB{}

const maxOpenDbConns = 10
const maxIdleDbConns = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates a connection pool to our PostgreSQL database.
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		log.Fatal("ConnectSQL: ", err)
	}

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err

	}

	return dbConn, nil
}

// CloseDB closes the database connection pool.
func CloseDB() {
	dbConn.SQL.Close(context.Background())
}

// testDB checks if the database is alive.
func testDB(d *pgx.Conn) error {
	ctx := context.Background()
	if err := d.Ping(ctx); err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database connection.
func NewDatabase(dsn string) (*pgx.Conn, error) {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
