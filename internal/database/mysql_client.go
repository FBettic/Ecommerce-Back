package database

import (
	"database/sql"
	"ecommerce-back/internal/logs"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	*sql.DB
}

func GetDB(ConnectionString string) *MySQL {

	log.Println("Connecting to database")

	db, err := sql.Open("mysql", ConnectionString)

	if err != nil {
		logs.Log().Errorf("Cannot create db tenant: %s", err.Error())
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to db: %s", err.Error())
	}

	return &MySQL{db}
}
