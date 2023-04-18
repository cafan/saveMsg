package saveMsg
import(
	"os"
	"fmt"
	"runtime"
)
type SaveMsg struct{
	Path string 
	CallerFuncName string
	CallerFileName string
}
func NewSaveMsg(path string)*SaveMsg{
	if path==""{
		path = "/tmp/mounter.log"
	}
	return &SaveMsg{Path:path}
	
}
func GetCallerInfo(sve *SaveMsg)(){
	pc,_,_,_:=runtime.Caller(1)
	callerFunc:=runtime.FuncForPC(pc)
	sve.CallerFileName,_=callerFunc.FileLine(pc)
	sve.CallerFuncName=callerFunc.Name()
}
// SaveMsgInPath is aim to save the message in the pointed path
// 1. get the caller function Name if setted before
// 2. saved information format :
//		<function name>:<key>:<value>
func (sve *SaveMsg)SaveMsgInPath(outputMsg string){
	file,err:=os.OpenFile(sve.Path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	saveStr:=outputMsg
	defer file.Close()
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	if sve.CallerFuncName!=""&&sve.CallerFileName!=""{
		saveStr=fmt.Sprintf("Called from %s: by function %s.\nWith parameter: %s",sve.CallerFileName,sve.CallerFuncName,saveStr)
	}
	fmt.Fprintln(file,saveStr)
}
