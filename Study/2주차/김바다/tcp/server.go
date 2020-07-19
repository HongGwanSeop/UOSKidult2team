package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)
func main(){
	l,err := net.Listen("tcp","localhost:8080")
	if nil != err{
		log.Println(err);
	}
	defer l.Close()

	bytes,err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	conn,err:=l.Accept()
		if nil != err {
			log.Println(err);
		}
		defer conn.Close()
	fmt.Println(bytes)
	_,err=conn.Write(bytes)
	if err!= nil{
		log.Println(err)
		return
	}
}
