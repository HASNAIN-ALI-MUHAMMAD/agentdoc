package database

import (

	"database/sql"
	"time"
)


type DocumentForm struct{
	Filename string 
	Path string 	
	Name string 	
	Size int64 	
	Type string 
	UploadedAt time.Time 
	FileId string
}

func (db *Database) Add_Document(doc DocumentForm) error{
	_,err := db.conn.Exec(AddDocumentQ,doc.FileId,doc.Filename,doc.Path,doc.Type,doc.UploadedAt,doc.UploadedAt,12,12)
	if err!=nil {
		return err		
	}
	return nil
}

func (db *Database) Get_Documents() (*sql.Rows,error){
	rows,err := db.conn.Query(Get_DocumentsAllQ)
	if err!=nil {
		return &sql.Rows{}, err		
	}
	return rows, nil
}


func (db *Database) Get_DocumentsByPath(path string) (*sql.Rows,error){
	rows,err := db.conn.Query(Get_DocumentsByPathQ,path)
	if err!=nil {
		return &sql.Rows{}, err		
	}
	return rows, nil
}

// func (db *Database) 