package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/fbettic/ecommerce-back/internal/logs"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	*sql.DB
}

func GetDB() *MySQL {

	log.Println("Connecting to database")

	db, err := sql.Open("mysql", os.Getenv("CLEARDB_DATABASE_URL"))

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
