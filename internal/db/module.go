package db

import (
	"database/sql"
	"github.com/denisenkom/go-mssqldb"
	"log"
)

func MakeNewConnector(connString string) *mssql.Connector {
	connector, err := mssql.NewConnector(connString)
	if err != nil {
		log.Println(err)
		return nil
	}
	return connector
}

func OpenDB(conn *mssql.Connector) *sql.DB {
	db := sql.OpenDB(conn)
	return db
}
