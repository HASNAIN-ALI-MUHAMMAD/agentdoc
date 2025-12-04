package filemanager

import (
	"agentDoc/core/internals/database"
	"os"
	"strings"
)


type FileManager struct{
	db *database.Database
	files []*FileReader
	// mutex later for concurrency
}

type FileReader struct {
	filename string
	filepath string
}

func NewFileReader(filename, filepath string) *FileReader {
	return &FileReader{
		filename: filename,
		filepath: filepath,
	}
}

func NewFileManager(db *database.Database) *FileManager {
	return &FileManager{
		db: db,
		files: make([]*FileReader, 0),
	}
}

func (fr *FileReader) GetFileType() string {
	ext := strings.Split(fr.filename, ".")
	if len(ext) < 2 {
		return "unknown"
	}
	return strings.ToLower(ext[len(ext)-1])
}

func (fr *FileReader) ReadFile() (string, error){
	filedata,err := os.ReadFile(fr.filename)
	if err!=nil{
		return "",err 
	}
	return  string(filedata),nil
}

func (fr *FileReader) GetFileMetadata() (map[string]interface{},error){
	fileInfo,err := os.Stat(fr.filepath)
	if err != nil{
		return nil,err
	}
	return map[string]interface{}{
		"filename": fr.filename,
		"filepath": fr.filepath,
		"size": fileInfo.Size(),
		"mod_time": fileInfo.ModTime(),
		"is_dir": fileInfo.IsDir(),
	},nil
}

// func (fr *FileReader) DBFileRecord() {
// 	docs ,err := database.Get_DocumentsByPath(fr.filepath)
// 	if err!=nil{
// 		fmt.Println("Error fetching document:", err)
// 		return
// 	}
// 	defer docs.Close()
// }


func (fm *FileManager) OpenFile(filename,filepath string) *FileReader{
	fileReader := NewFileReader(filename, filepath)
	fm.files = append(fm.files, fileReader)
	return fileReader	
}

func (fm *FileManager) CloseFile(filepath string){
	// remove file from fm.files
	for i, file := range fm.files {
		if file.filepath == filepath {
			fm.files = append(fm.files[:i], fm.files[i+1:]...)
			break
		}
	}
}