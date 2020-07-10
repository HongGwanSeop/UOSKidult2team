package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

const BSIZE = 4096

func freceive(c net.Conn, file string) {
	buf := make([]byte, BSIZE)
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		c.Write([]byte("File already exist!!\n"))
		return
	}
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		c.Write([]byte(err.Error()))
		return
	}
	c.Write([]byte("0"))
	n, err := c.Read(buf)
	if err != nil || string(buf[:n]) != "0" {
		fmt.Println("freceive: ", err)
		return
	}

	for {
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		if buf[0] == 0xFF && n == 1 {
			if n != 1 {
				f.Write(buf[:n-1])
			}
			break
		}
		f.Write(buf[:n])
	}
}

func fsend(c net.Conn, file string) {
	buf := make([]byte, BSIZE)
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		c.Write([]byte(err.Error()))
		return
	}
	c.Write([]byte("0"))
	n, err := c.Read(buf)
	if err != nil || string(buf[:n]) != "0" {
		fmt.Println("fsend: ", err)
		return
	}

	for {
		n, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
			c.Write([]byte{0xFF})
			return
		}
		c.Write(buf[:n])
	}
}

func list(c net.Conn) {
	var files []string
	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	m, _ := json.Marshal(files[1:])
	_, err = c.Write([]byte(m))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func get(c net.Conn) {
	c.Write([]byte{0})
	var files []string
	data := make([]byte, BSIZE)
	n, err := c.Read(data)
	if err != nil {
		fmt.Println("get1")
		return
	}
	json.Unmarshal(data[:n], &files)
	for i, file := range files {
		fmt.Println(i, ": ", file)
		fsend(c, file)
	}
}

func put(c net.Conn) {
	c.Write([]byte{0})
	var files []string
	data := make([]byte, BSIZE)
	n, err := c.Read(data)
	if err != nil {
		fmt.Println("put1")
		return
	}
	json.Unmarshal(data[:n], &files)
	for i, file := range files {
		fmt.Println(i, ": ", file)
		freceive(c, file)
	}
}

func requestHandler(c net.Conn) {
	data := make([]byte, BSIZE)
	n, err := c.Read(data)
	if err != nil {
		fmt.Println("err) requestHandler: ", err)
		return
	}

	fmt.Println(string(data[:n]))
	switch string(data[:n]) {
	case "list":
		list(c)
	case "get":
		get(c)
	case "put":
		put(c)
	}
}

func main() {
	l, err := net.Listen("tcp", ":2015")
	if err != nil {
		fmt.Println("err) net.Listen: ", err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("err) l.Accept(): ", err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}
