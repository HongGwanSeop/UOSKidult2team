package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("./client <command> <fileName>")
	}

	if _, err := os.Stat("./files"); os.IsNotExist(err) {
		err = os.MkdirAll("./files", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	var command string = os.Args[1]
	log.Println("Connected to server!")

	switch command {
	case "put":
		putFileToServer(conn, os.Args[2])
	case "get":
		getFileFromServer(conn, os.Args[2])
	case "list":
		getFileList(conn)
	default:
		log.Println("Invalid Command...")
	}
}

func putFileToServer(conn net.Conn, fileName string) {

	file, err := os.Open("./files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = conn.Write([]byte("put " + fileName))

	n, err := io.Copy(conn, file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Send", fileName, "to Server [", n, "bytes ]")
}

func getFileFromServer(conn net.Conn, fileName string) {

	if _, err := os.Stat("./files/" + fileName); err == nil {
		log.Println("./files/", fileName, " 가 이미 존재하므로 덮어씁니다.")
	}

	file, err := os.Create("./files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = conn.Write([]byte("get " + fileName))

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

func getFileList(conn net.Conn) {

	_, err := conn.Write([]byte("list"))
	if err != nil {
		log.Fatal(err)
	}

	received := make([]byte, 4096)
	_, err = conn.Read(received)
	if err != nil {
		log.Fatal(err)
	}

	fileLIst := string(received)
	fmt.Println(fileLIst)
}
