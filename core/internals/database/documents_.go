package database

import (
	"database/sql"
	"time"
)


type DocumentForm struct{
	Filename string `json:"filename"`
	Path string 	`json:"filepath"`
	Name string 	`json:"name"`
	Size int64 	`json:"size"`
	Type string 	`json:"type"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func (db *Database) Add_Document(doc DocumentForm) error{
	return nil
}

func (db *Database) Get_Documents() (sql.Rows,error){
	rows,err := db.conn.Query(Get_Documents)
	if err!=nil {
		return sql.Rows{}, err		
	}
	return *rows, nil
}