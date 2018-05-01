package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type serverArgs struct {
	configFile string
	logFile    string
	port       string
	separator  string
}

func defaultArgs() serverArgs {
	return serverArgs{"server.conf", "out.log", ":8080", ","}
}

func getArgs() (serverArgs, error) {
	osArgs := os.Args[1:]
	args := defaultArgs()

	for i := 0; i < len(osArgs); i++ {
		switch osArgs[i] {
		case "-p", "--port":
			i++
			if i >= len(osArgs) {
				return defaultArgs(), errors.New("no port given")
			}
			p, err := strconv.Atoi(osArgs[i])
			if err != nil {
				return defaultArgs(), err
			}
			args.port = fmt.Sprintf(":%d", p)
		case "-c", "--config":
			i++
			if i >= len(osArgs) {
				return defaultArgs(), errors.New("no config file given")
			}
			args.configFile = osArgs[i]
		case "-o":
			i++
			if i >= len(osArgs) {
				return defaultArgs(), errors.New("no output file given")
			}
			args.logFile = osArgs[i]
		case "-s":
			i++
			if i >= len(osArgs) {
				return defaultArgs(), errors.New("no separator given")
			}
			args.separator = osArgs[i]
		default:
			return defaultArgs(), errors.New(fmt.Sprint("unrecognized option:", osArgs[i]))
		}
	}
	return args, nil

}
