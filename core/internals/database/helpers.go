package database

import(
	"os"
	"path/filepath"
)

func Get_Path_DB(filename string) (string,error){
	config,_ := os.UserConfigDir()
	path:=filepath.Join(config,"agentDoc")
	if err:= os.MkdirAll(path,0755);err!=nil{
		return "",err
	}
	return filepath.Join(path,filename),nil
}