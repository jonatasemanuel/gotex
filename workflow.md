# Go, Postgres + SQLx

## Dependecies
```shell
go install github.com/joho/godotenv
go install github.com/lib/pq
go install github.com/jackc/pgx
```

```shell
/gote-api-server
../cmd/api
    main.go
../database
    database.go
Makefile
.env
```

```go
//  project/
//  ../gote-api-server/cmd/api/main.go
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonatasemanuel/gote-server/database"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	dbConn, err := database.ConnectDatabase(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer dbConn.DB.Close()
}
```
```go
//  project/
//      gote-api-server/database/database.go
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
```
## SQL Table and Structs

- Using ```sqlx``` to initialize migrations:
```shell
sqlx migrate add -r init
```
- The sqlx script will be create a migration folder at root path project with two sql files, init up and down. 

```
migrations/
    00000000000000_init.up.sql
    11111111111111_init.down.sql
```
