package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

func ConnectDatabase(dns string) (*DB, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	dbConn.DB = db

	return dbConn, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	fmt.Println("***Pinged database successfully!***")

	return nil
}
