package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func MySqlDatabase(config mysql.Config) (*sql.DB, error){
	db, err := sql.Open("mysql", config.FormatDSN()) //convert jadi format dsn
	// open koneksi db
	if err != nil {
		return nil, err // kalau error throw
	}

	if err:= db.Ping(); err != nil{
		db.Close() //coba ping kalau error bakal throw
		return nil, err
	}
	log.Println("database connected")
	return db, nil //liat di expected return (db dan error)
}