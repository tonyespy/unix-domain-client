// -*- Mode: Go; indent-tabs-mode: t -*-

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		_, err := r.Read(buf[:])
		if err != nil {
			log.Fatal("Read: ", err)
		}

		fmt.Println("Client read: ", string(buf[0:]))
	}
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage unix-domain-client <socketdir>")
	}

	dir := os.Args[1]
	path := dir + "/sock"

	fmt.Println("$SNAP is", dir)
	fmt.Println("socket path is ", path)

	c, err := net.Dial("unix", path)
	if err != nil {
		log.Fatal("Dial error: ", err)
	}

	defer c.Close()

	go reader(c)
	for {
		_, err := c.Write([]byte("hi"))
		if err != nil {
			log.Fatal("Write error: ", err)
		}

		time.Sleep(time.Duration(10)*time.Second)
	}

	fmt.Println("all done!")
}
