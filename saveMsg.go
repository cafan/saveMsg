package saveMsg
import(
	"os"
	"fmt"
	"runtime"
)
type SaveMsg struct{
	Path string 
	CallerFuncName string
}
func NewSaveMsg(path string)*SaveMsg{
	if path==""{
		path = "/tmp/mounter.log"
	}
	return &SaveMsg{Path:path}
	
}
func GetCallerInfo()(string,string){
	pc,_,_,_:=runtime.Caller(1)
	callerFunc:=runtime.FuncForPC(pc)
	callerFile,_:=callerFunc.FileLine(pc)
	callerName:=callerFunc.Name()
	return callerFile,callerName
}
// SaveMsgInPath is aim to save the message in the pointed path
// 1. get the caller function Name if setted before
// 2. saved information format :
//		<function name>:<key>:<value>
func (smg *SaveMsg)SaveMsgInPath(outputMsg,callerFile,callerName string){
	file,err:=os.OpenFile(smg.Path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	saveStr:=outputMsg
	defer file.Close()
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	if callerName!=""&&callerFile!=""{
		saveStr=fmt.Sprintf("Called from %s: by function %s.\nWith parameter: %s",callerFile,callerName,saveStr)
	}
	fmt.Fprintln(file,saveStr)
}
