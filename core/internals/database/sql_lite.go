package database

import (
	"database/sql"
	"fmt"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
)

// a struct for the database
type Database struct{
	conn *sql.DB	
}
//go:embed schemas.sql
var schemasSql string

func NewDatabase(path string) (*Database, error){
	db,err := sql.Open("sqlite3",path)
	if err !=nil{
		return nil,err	
	}
	if err :=db.Ping();err!=nil{
		return nil,err
	}

	if _,err:=db.Exec(schemasSql);err!=nil{
		return nil,err
	}
	fmt.Printf("Database initialized.")
	return &Database{conn: db},nil
}
