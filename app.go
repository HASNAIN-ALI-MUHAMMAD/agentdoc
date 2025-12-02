package main

import (
	"agentDoc/core/internals/database"
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	db *database.Database
}

// NewApp creates a new App application struct
func NewApp() *App {
	// initialize the startups
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	path,err:= database.Get_Path_DB("sql_lite.db")
	if err !=nil{
		runtime.MessageDialog(ctx,runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Database Path Error",
			Message: err.Error(),
		})
		runtime.Quit(ctx)
		return
	}
	fmt.Println("Database Path:",path)
	db,err := database.NewDatabase(path)
	if err != nil{
		runtime.MessageDialog(ctx,runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Database Error",
			Message: err.Error(),
		})
		runtime.Quit(ctx)
		return
	}
	a.ctx = ctx
	a.db = db
}


func (a *App) Get_Documents() (string,error){
	rows,err := a.db.Get_Documents()
	if err !=nil{
		return "", err
	}
	defer rows.Close()
	var result string
	for rows.Next(){
		var id int
		var title string
		var content string
		err := rows.Scan(&id,&title,&content)
		if err !=nil{
			return "", err
		}
		result += fmt.Sprintf("ID: %d, Title: %s, Content: %s\n",id,title,content)
	}
	return result, nil
}