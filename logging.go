package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var logChan = make(chan string)

/* Just log */
func logOnly(s ...string) {
	full := strings.Join(s, " ")
	time := time.Now()
	timeString := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
	logChan <- fmt.Sprintf("%s %s\n", timeString, full)
}

/* Log with connection information */
func logConn(c net.Conn, s ...string) {
	full := strings.Join(s, " ")
	info := fmt.Sprintf("%s: %s", c.RemoteAddr(), full)
	logOnly(info)
}

/* Logging daemon */
func startLogger(args *serverArgs) {

	var file *os.File

	info, err := os.Stat(args.logFile)
	if err != nil || !info.Mode().IsRegular() {
		file, err = os.Create(args.logFile)
		if err != nil {
			fmt.Println("failed to create log file, logging to file will not be available")
			fmt.Println(err)
		}
	} else {
		file, err = os.OpenFile(args.logFile, os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			fmt.Println("failed to start logger, logging to file will not be available")
			fmt.Println(err)
			file = nil
		}
	}

	for {
		data := <-logChan
		fmt.Print(data)
		if file != nil {
			_, err := file.Write([]byte(data))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
