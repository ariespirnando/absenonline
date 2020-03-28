package config

import (
    "database/sql" 
    _ "github.com/go-sql-driver/mysql" 
    "os"
)

func Connect() *sql.DB{
	DB, err := sql.Open("mysql", os.Getenv("DB_SQL"))
	if err != nil { 
        panic("failed to connect database")
    }
    return DB
}