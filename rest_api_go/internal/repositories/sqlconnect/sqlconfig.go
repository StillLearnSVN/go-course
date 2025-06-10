package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectDb() (*sql.DB, error) {
	
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	
	// connectionString := "root:Mypass123@tcp(localhost:3306)/" + dbname
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		// panic(err)
		return nil, err
	}

	return db, nil
}