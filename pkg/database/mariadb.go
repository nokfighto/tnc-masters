package database

import (
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	// connectionString = "tncadmin:IbO#25489@tcp(172.19.218.100:3306)/newscenter?parseTime=true"
	connectionString = "newscenter_rw:Passw0_rw@tcp(172.19.248.90:3524)/newscenter?parseTime=true"
	maxIdleConns     = 10
	maxOpenConns     = 20
)

var DB *sqlx.DB

func Connect() (func(), error) {
	var err error
	DB, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	DB.SetMaxIdleConns(maxIdleConns)
	DB.SetMaxOpenConns(maxOpenConns)

	if _, err := DB.Query("select 1"); err != nil {
		DB.Close()
		return nil, errors.New("database don't connect")
	}
	log.Println("database open")

	closeConnection := func() {
		if err := DB.Close(); err != nil {
			log.Panic(err)
		} else {
			log.Println("database closed")
		}
	}

	return closeConnection, nil
}

func IsOpen() bool {
	if DB.Stats().OpenConnections > 0 {
		return true
	} else {
		return false
	}
}
