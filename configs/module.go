package configs

import "fmt"

type DatabaseConfiguration struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
}

func (dbCfg *DatabaseConfiguration) CreateConnStr() string {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.Port, dbCfg.Database)

	return connString
}
