package db

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "nihankhan"
	hostname = "127.0.0.1:3306"
	dbName   = "Nihan"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

var (
	dbc  *sql.DB
	err  error
)

func Connect() {
	dbc, err = sql.Open("mysql", dsn(""))

	if err != nil {
		panic(err)
	}
}