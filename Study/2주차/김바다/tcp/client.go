package main
import (
	"fmt"
	"io/ioutil"
	"log"
	//"os"
	"net"
)
func main(){
	conn,err := net.Dial("tcp","localhost:8080")
	if nil != err{
		log.Println(err);
	}

	recvBuf := make([]byte, 4096)
	n,err := conn.Read(recvBuf)
	if nil != err{
		log.Println(err);
		return 
	}
	fmt.Println(recvBuf[:n])
	err = ioutil.WriteFile("receive.txt",recvBuf[:n],0)
	if err != nil {
		panic(err)
	}
}