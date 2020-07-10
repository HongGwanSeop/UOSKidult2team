package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
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
	for {
		n, err := c.Read(buf)
		fmt.Println(buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
		if buf[0] == 0xFF {
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
	for i := 0; true; i++ {
		n, err := f.Read(buf)
		fmt.Println(i, ": ", n)
		if err != nil {
			fmt.Println(err)
			c.Write([]byte{0xFF})
			return
		}
		c.Write(buf[:n])
	}
}

func list(c net.Conn) {
	data := make([]byte, BSIZE)

	_, err := c.Write([]byte("list"))
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := c.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	var files []string
	json.Unmarshal(data[:n], &files)
	fmt.Println(string(data))
	for i, file := range files {
		fmt.Println(i, ": ", file)
	}
}

func get(c net.Conn, args ...string) {
	data := make([]byte, BSIZE)

	_, err := c.Write([]byte("get"))
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Read(data)

	m, _ := json.Marshal(args)
	_, err = c.Write([]byte(m))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range args {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if string(data[:n]) != "0" {
			fmt.Println(string(data[:n]))
		}
		freceive(c, file)
	}
}

func put(c net.Conn, args ...string) {
	data := make([]byte, BSIZE)

	_, err := c.Write([]byte("put"))
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Read(data)

	m, _ := json.Marshal(args)
	_, err = c.Write([]byte(m))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range args {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if string(data[:n]) != "0" {
			fmt.Println(string(data[:n]))
		}
		fsend(c, file)
	}
}

func connect(args ...string) {
	conn, err := net.Dial("tcp", "localhost:2015")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	switch args[0] {
	case "list":
		list(conn)
	case "get":
		if len(args) == 1 {
			invalid()
		} else {
			get(conn, args[1:]...)
		}
	case "put":
		if len(args) == 1 {
			invalid()
		} else {
			put(conn, args[1:]...)
		}
	default:
		invalid()
	}
}

func invalid() {
	fmt.Println("Invalid argument\nex) -------------------\n$ ./client list\n$ ./client put <filename>\n$ ./client get <filename>")
}

func main() {
	if len(os.Args) == 1 {
		invalid()
		return
	}
	args := os.Args[1:]
	connect(args...)
}
