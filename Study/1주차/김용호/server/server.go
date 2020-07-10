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

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Println("Waiting for connections...")

	if _, err := os.Stat("./files"); os.IsNotExist(err) {
		err = os.MkdirAll("./files", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Client connected")
		go connectionHandler(conn)
	}
}

func connectionHandler(conn net.Conn) {

	defer conn.Close()
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	buffer = bytes.Trim(buffer, "\x00")
	commandString := strings.TrimSpace(string(buffer))
	commandArray := strings.Split(commandString, " ")
	command := commandArray[0]

	switch command {
	case "get":
		sendFileToClient(conn, commandArray[1])
	case "put":
		getFileFromClient(conn, commandArray[1])
	case "list":
		sendFileList(conn)
	}
}

func getFileFromClient(conn net.Conn, fileName string) {

	if _, err := os.Stat("./files/" + fileName); err == nil {
		fmt.Println("./files/", fileName, " 가 이미 존재하므로 덮어씁니다.")
	}

	file, err := os.Create("./files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var current int64 = 0
	fileBuffer := make([]byte, 4096)
	for err == nil || err != io.EOF {
		conn.Read(fileBuffer)
		buff := bytes.Trim(fileBuffer, "\x00")

		_, err = file.WriteAt(buff, current)
		if len(string(fileBuffer)) != len(string(buff)) {
			break
		}
		current += 4096
	}
}

func sendFileToClient(conn net.Conn, fileName string) {

	file, err := os.Open("./files/" + fileName)
	if err != nil {
		conn.Write([]byte("-1"))
	}

	n, err := io.Copy(conn, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send", fileName, "to client [", n, "bytes ]")

	defer file.Close()
}

func sendFileList(conn net.Conn) {

	files, err := ioutil.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}

	var fileListStr string
	for _, file := range files {
		fileListStr = fileListStr + file.Name() + "\n"
	}

	fileList := []byte(fileListStr)
	_, err = conn.Write(fileList)
	if err != nil {
		log.Fatal(err)
	}
}
