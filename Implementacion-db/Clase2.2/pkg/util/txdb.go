package util

import (
	"database/sql"
	"log"
	"os"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		panic(err)
	}
	configDB := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
	txdb.Register("txdb", "mysql", configDB.FormatDSN())
}

// instancia txdb
func InitTxDB() *sql.DB {
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
