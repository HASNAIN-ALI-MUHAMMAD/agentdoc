package main

import (
	"agentDoc/core/internals/database"
	"agentDoc/core/internals/filemanager"
	_ "agentDoc/core/internals/jobmanager"
	"context"
	"fmt"
	"path"

	// "strings"
	"time"
)


func main(){
	p:= "/home/hasnain/D/code/projects/agentDoc/agentDoc/README.md"
	fr := filemanager.NewFileReader("README.md",p)
	filedata,_ := fr.GetFileMetadata()
	fmt.Println(filedata)
	db,err := database.NewDatabase("test2.db")	
	if err!=nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()
	var docForm database.DocumentForm = database.DocumentForm{
		Filename: filedata["filename"].(string),
		Path:p,
		Name: path.Base(p),
		Size: filedata["size"].(int64),
		Type:fr.GetFileType(),
		UploadedAt:time.Now(),
		FileId: p,
	}
	err = db.Add_Document(docForm)
	if err!=nil {
		fmt.Println("Error adding document:", err)
		return
	}	
	rows,err := db.Get_Documents()
	if err!=nil {
		fmt.Println("Error retrieving document by path:", err)
		return
	}
	defer rows.Close()
	fmt.Printf("data from database\n")
	for i:=0; rows.Next();i++ {
		var id string
		var filename,filepath,file_type string
		var total_pages,total_chunks int
		var created_at time.Time
		var last_read time.Time
		err := rows.Scan(&id,&filename,&filepath,&file_type,&created_at,&last_read,&total_pages,&total_chunks)
		if err!=nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("Row %d: ID=%s, Filename=%s, Filepath=%s, Type=%s, Pages=%d, Chunks=%d, CreatedAt=%v, LastRead=%v\n",
			i, id, filename, filepath, file_type, total_pages, total_chunks, created_at, last_read)
	}
	// fmt.Println(fr.ReadFile())
	// jobmanager.TaskManager.RunTask("Bg-one",bgTest)
	// jobmanager.TaskManager.RunTask("Bg-two",bgTest)
	// jobmanager.TaskManager.RunCancellableTask("cancellable-bg-one",bgTestC)
	// jobmanager.TaskManager.WaitAll()
}


func bgTestC(ctx context.Context){
	select{
		case <- ctx.Done():
			fmt.Println("Cancellable background task received cancellation signal")
			return
		default:
			time.Sleep(5* time.Second)
	}
}

func bgTest(){
	time.Sleep(5* time.Second)
}