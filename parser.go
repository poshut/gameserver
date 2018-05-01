package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ConfigItem struct {
	id   int
	name string
	prog string
	args []string
}

func (item ConfigItem) String() string {
	return fmt.Sprintf("%d %s", item.id, item.name)
}

func (item ConfigItem) toCommand() *exec.Cmd {
	return exec.Command(item.prog, item.args...)
}

func parseConfig(args *serverArgs) ([]ConfigItem, error) {
	config := make([]ConfigItem, 0)
	file, err := os.Open(args.configFile)
	if err != nil {
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 0
	for ; scanner.Scan(); id++ {
		line := scanner.Text()
		parsed, err := parseLine(line, id, args)
		if err != nil {
			return nil, err
		}
		config = append(config, parsed)
	}
	return config, nil
}

func parseLine(line string, id int, args *serverArgs) (ConfigItem, error) {
	items := strings.Split(line, args.separator)

	if len(items) < 2 {
		return ConfigItem{-1, "", "", make([]string, 0)}, errors.New(fmt.Sprint("line too short:", line))
	}
	name := items[0]
	prog := items[1]

	return ConfigItem{id, name, prog, items[2:]}, nil
}
