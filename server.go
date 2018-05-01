package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("getting arguments")

	args, err := getArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("starting logger")
	go startLogger(&args)

	fmt.Println("parsing config")
	config, err := parseConfig(&args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("starting socket")
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
