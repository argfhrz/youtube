package model

import "database/sql"

type BaseData struct {
	DB *sql.DB `json:",omitempty"`
}

func (baseData BaseData) Exec(sql string, args ...interface{}) (sql.Result, error) {

	return baseData.DB.Exec(sql, args...)

}

func (baseData BaseData) Query(sql string, args ...interface{}) (*sql.Rows, error) {

	return baseData.DB.Query(sql, args...)

}
