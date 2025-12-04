package main

import (
	"agentDoc/core/internals/database"
	"agentDoc/core/internals/filemanager"
	"context"
	"fmt"
	"path"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	db *database.Database
	fm *filemanager.FileManager
}

// NewApp creates a new App application struct
func NewApp() *App {
	// initialize the startups
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	// database setup at runtime
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
	a.fm = filemanager.NewFileManager(db)

	// on filedrop
	runtime.OnFileDrop(ctx,func (x,y int,paths []string){
		runtime.EventsEmit(ctx, "drop", paths)

		fmt.Println("File dropped at:",x,y)
		fmt.Println("Paths:",paths)
		runtime.MessageDialog(ctx,runtime.MessageDialogOptions{
			Type: runtime.InfoDialog,
			Title: "Files read",
			Message: "Files dropped: \n" + fmt.Sprint(paths),
		})		
	})

}

func (a *App) SelectFile() string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "All Files",
				Pattern:     "*.pdf;*.docx;*.txt;*.md;",
			},
		},
	})
	if err != nil {
		return ""
	}

	filename := path.Base(file)
	fr := a.fm.OpenFile(file,filename)
	fr.GetFileMetadata()	
	return file 
}