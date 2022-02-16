package main
import(
	"fmt"
	"os"
	"io"
)
func main(){
	fileReader()
}
func fileReader(){
	f, err := os.Open("sample2.txt")
	if err != nil{
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 3)
	for {
		n, err := f.Read(buf)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
			break
		}
		if n>0{
			fmt.Println(buf[:n])
		}
	}
}