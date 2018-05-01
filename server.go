package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args, err := getArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go startLogger(&args)
	config, err := parseConfig(&args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", args.port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log(fmt.Sprint("server started on ", ln.Addr()))

	for {
		conn, err := ln.Accept()
		if err != nil {
			log(err.Error())
		} else {
			go handleConnection(conn, config)
		}
	}
}
