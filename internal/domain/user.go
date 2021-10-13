package domain

import mssql "github.com/denisenkom/go-mssqldb"

type User struct {
	Id       mssql.UniqueIdentifier
	Login    string
	LastName string
}
