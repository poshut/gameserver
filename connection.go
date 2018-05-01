package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

const greeting string = "This is GameServer v0.1, (c) Richard Wohlbold 2018\n"
const prompt string = "Please type in a number. Type 'exit' to exit.\n"

func handleConnection(conn net.Conn, config []configItem) {
	defer conn.Close()

	/* Write a greeting to the screen */
	_, err := conn.Write([]byte(greeting))
	if err != nil {
		logWithConnection(conn, err.Error())
		return
	}

	/* Write all thhe configuration objects */
	for _, c := range config {
		toWrite := fmt.Sprintf("%s\n", c)
		_, err = conn.Write([]byte(toWrite))
		if err != nil {
			logWithConnection(conn, err.Error())
			return
		}
	}

	/* Write usage information */
	_, err = conn.Write([]byte(prompt))
	if err != nil {
		logWithConnection(conn, err.Error())
		return
	}

	/* Read existing number or 'exit' */
	reader := bufio.NewReader(conn)
	var input int64 = -1
	for {
		/* Read a string from connection */
		res, err := reader.ReadString('\n')
		if err != nil {
			logWithConnection(conn, err.Error())
			return
		}

		/* Remove CR and NL */
		res = strings.TrimRight(res, "\r\n")

		/* If 'exit', exit */
		if strings.Compare(res, "exit") == 0 {
			_, err = conn.Write([]byte("goodbye\n"))
			if err != nil {
				logWithConnection(conn, err.Error())
			}
			return
		}

		/* Parse int from input, if it is not convertible or too big, print information and loop again */
		input, err = strconv.ParseInt(res, 10, 32)
		if err != nil || int64(len(config)) <= input {
			logWithConnection(conn, err.Error())
			conn.Write([]byte(prompt))
			continue
		}
		break
	}

	/* Get the chosen configuration item */
	chosenConfig := config[input]

	/* Build a command object */
	cmd := chosenConfig.toCommand()

	/* Get pipes from command */
	pipe0, err := cmd.StdinPipe()
	if err != nil {
		logWithConnection(conn, err.Error())
		return
	}
	pipe1, err := cmd.StdoutPipe()
	if err != nil {
		logWithConnection(conn, err.Error())
		return
	}
	pipe2, err := cmd.StderrPipe()
	if err != nil {
		logWithConnection(conn, err.Error())
		return
	}

	/* Start command */
	err = cmd.Start()
	if err != nil {
		logWithConnection(conn, err.Error())
		return
	}

	/* Create a channel for synchronization */
	ch := make(chan bool)

	go connectStreams(conn, pipe0, nil, conn)
	go connectStreams(pipe1, conn, ch, conn)
	go connectStreams(pipe2, conn, ch, conn)
	/* Wait for the stdout and stderr goroutines to finish */
	<-ch
	<-ch

	logWithConnection(conn, "done")
}

/* Connects two streams, optionally receiving a message on completion */
func connectStreams(input io.Reader, output io.Writer, ch chan<- bool, conn net.Conn) {
	buffer := make([]byte, 512)
	n, err := input.Read(buffer)
	for err == nil {
		_, err = output.Write(buffer[:n])
		if err != nil {
			logWithConnection(conn, err.Error())
			return
		}
		n, err = input.Read(buffer)
	}
	if ch != nil {
		ch <- true
	}
}
