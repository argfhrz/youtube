package connection

import (
	"database/sql"
	"youtube/config"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Driver string
}

func (connection *Connection) connectionString(driver string, user string, password string,
	dbName string) string {

	connection.Driver = driver

	connString := user + ":" + password + "@/" + dbName

	return connString

}

func (connection Connection) OpenConnection(driver string, user string, password string,
	dbName string) *sql.DB {

	connString := connection.connectionString(driver, user, password, dbName)
	db, err := sql.Open(driver, connString)
	if err != nil {
		panic(err)
	}

	return db
}

func OpenConnection(phase string) *sql.DB {

	connection := Connection{}
	dbConf := config.DB_CONFIGS[phase]

	db := connection.OpenConnection(dbConf.Driver,
		dbConf.User,
		dbConf.Password,
		dbConf.DbName,
	)

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
