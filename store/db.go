package store

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var (
	dbInstance *sql.DB = nil
)

func Instance() (*sql.DB, error) {
	if dbInstance != nil {
		//TODO: check the connection is valid
		return dbInstance, nil
	}

	db, err := sql.Open("mysql", "root@/gogangbot")
	if err != nil {
		return nil, err
	}

	dbInstance = db
	return dbInstance, nil
}
