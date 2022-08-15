package fileop
import(
	"fmt"
	"mypro/getfuncinfo"
	"os"
	"io"
	"bufio"
	"io/ioutil"
)

func TestFileRead(){
	getfuncinfo.PrintFuncName()
	//打开文件
	fileobj,err := os.Open("./file.txt")
	if err != nil {
		fmt.Printf("打开文件失败!\n")
		return
	}

	//延迟执行,在全部返回之前再关闭文件
	defer fileobj.Close()
	//读文件
	for{
		var tmp [128]byte
		n,err := fileobj.Read(tmp[:])
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕!")
				return
			}else{
				fmt.Println("读文件失败,err:",err)
				return
			}
		}
		//到文件尾部，及时读取不到tmp满长度，err也不会报错，只是n返回的读取长度小于len(tmp)就是了
		fmt.Printf("读取%d字节\n",n)
		fmt.Println(string(tmp[:]))
	}
}

func TestBufio(){
	//bufio可按照行读取
	getfuncinfo.PrintFuncName()
	fileobj,err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("打开文件失败!")
		return
	}
	defer fileobj.Close()

	reader := bufio.NewReader(fileobj)
	for{
		//设定按照行读取
		line,err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("到达文件末尾!")
				return
			}else{
				fmt.Printf("read line 错误，err:%v\n",err)
				return
			}
		}
		fmt.Print(line)
	}
}

func TestIoutil(){
//ioutil 一次读取整个文件
	getfuncinfo.PrintFuncName()
	ret,err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Printf("read err:%v\n",err)
		return
	}
	fmt.Println(string(ret))
}
