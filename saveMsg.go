package saveMsg
import(
	"os"
	"fmt"
)
type SaveMsg struct{
	Path string 
	Msg string 
	CallerFuncName string 
}
func NewSaveMsg(path string)*SaveMsg{
	if path==""{
		path = "/tmp/mounter.log"
	}
	return &SaveMsg{Path:path}
	
}
func (smg *SaveMsg)SaveMsgInPath(outputMsg string){
	file,err:=os.OpenFile(smg.Path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	defer file.Close()
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	fmt.Fprintln(file,outputMsg)
}
