package main

import (
	"log"
	"voidspace-api/cmd/api"
	"voidspace-api/config"
	"voidspace-api/database"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func main() { 	
	err := godotenv.Load() //load .env
	if err != nil {
		log.Println("No .env file found, using default/fallback values")
	}

	cfg := config.GetConfig()

	var dbConfig = mysql.Config{ //dari config mysql env.go
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Addr:                 cfg.DBAddress,
		DBName:              	cfg.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	} //config untuk mysql

	//konek ke db
	db, err := database.MySqlDatabase(dbConfig) 
	if err != nil {
		log.Fatal(err)
	}

	//server jalan di Port config
	server := api.NewAPIServer(":" + cfg.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}