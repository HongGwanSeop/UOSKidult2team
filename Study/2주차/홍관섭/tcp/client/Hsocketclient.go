package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if nil != err {
		log.Println(err)
	}
	var s string
	var err2 error
	for {
		//fmt.Scanln(&s) 이걸 사용할경우에는 공백을 못받는다겠네~
		fmt.Println("명령어를 입력하세요")
		str := bufio.NewReader(os.Stdin)
		s, err2 = str.ReadString('\n') //공백을 제대로 받아먹으려면 이렇게 해야함ㅇㅇ
		if err2 != nil {
			log.Println(err2)
		}
		temp := strings.Fields(s)
		if temp[0] == "/client" {
			switch temp[1] {
			case "list":
				{
					getList(conn)
				}
			case "put":
				{
					if len(temp) == 2 {
						fmt.Println("파일이름을 제대로 입력하세요.")
					} else {
						putFile(temp[2], conn)
					}
				}
			case "get":
				{
					if len(temp) == 2 {
						fmt.Println("파일이름을 제대로 입력하세요.")
					} else {
						getFile(temp[2], conn)
					}
				}
			default:
				fmt.Println("잘못된 명령어입니다.")
			}
		} else {
			fmt.Println("잘못된 명령어입니다.")
		}
	}
}

func getList(connection net.Conn) {
	data := make([]byte, 4096)
	fmt.Println("리스트를 얻는 중입니다..")
	connection.Write([]byte("list"))
	n, err := connection.Read(data)
	if err != nil {
		log.Println(err)
	}
	log.Println("Filelist : \n" + string(data[:n]))

}
func putFile(fileName string, connection net.Conn) {

	fmt.Println("서버로 파일을 보냅니다.")

	var err error

	//file to read
	file, err := os.Open(strings.TrimSpace(fileName)) // For read access.
	if err != nil {
		connection.Write([]byte("-1"))
		log.Fatal(err)
	}
	connection.Write([]byte("send " + fileName))
	n, err := io.Copy(connection, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send", fileName, "byte:", n)
	file.Close()

}

func getFile(fileName string, connection net.Conn) {
	var currentByte int64 = 0

	fileBuffer := make([]byte, 4096)

	var err error
	file, err := os.Create(strings.TrimSpace(fileName))
	if err != nil {
		log.Fatal(err)
	}

	connection.Write([]byte("get " + fileName))
	for {

		connection.Read(fileBuffer)
		cleanedFileBuffer := bytes.Trim(fileBuffer, "\x00")

		_, err = file.WriteAt(cleanedFileBuffer, currentByte)

		currentByte += 4096

		if err == io.EOF || err == nil {
			break
		}

	}

	file.Close()
	return
}
