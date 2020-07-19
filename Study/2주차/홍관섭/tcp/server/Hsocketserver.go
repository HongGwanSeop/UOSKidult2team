package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000") //소켓을 열기
	if nil != err {
		log.Println(err)
	}
	defer l.Close() //메인프로세스가 종료되면 소켓도 종료

	for {
		conn, err := l.Accept() //소켓 연결
		if nil != err {
			log.Println(err)
			continue
		}
		defer conn.Close() //프로세스 종료시 연결도 종료
		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf) //대기를 타다가 값을 주면 이제 읽는다.
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		data := recvBuf[:n]
		InputString := strings.TrimSpace(string(data))
		Command := strings.Fields(InputString)
		if Command[0] == "get" {
			fmt.Println("send file")
			sendFile(Command[1], conn)
		} else if Command[0] == "send" {
			fmt.Println("getting a file")

			getFile(Command[1], conn)

		} else if Command[0] == "list" {
			fmt.Println("send list")
			sendList(conn)
		} else {
			_, err = conn.Write([]byte("잘못된 명령어입니다."))
		}
	}
}
func sendFile(fileName string, connection net.Conn) {
	file, err := os.Open(fileName)
	if err != nil {
		connection.Write([]byte("-1"))
	}

	_, err = io.Copy(connection, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send", fileName)

	defer file.Close()
}

func getFile(fileName string, connection net.Conn) {

	var currentByte int64 = 0

	fileBuffer := make([]byte, 4096)

	var err error
	file, err := os.Create(strings.TrimSpace(fileName))
	if err != nil {
		log.Fatal(err)
	}
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

func sendList(connection net.Conn) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	var filelist string
	for _, file := range files {
		// 파일명
		filelist += (file.Name() + "\n")
	}
	connection.Write([]byte(filelist))
}
