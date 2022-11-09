package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnecDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	configDB := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
	db, err := sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}
