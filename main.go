package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gorcon/rcon"
)

type flags struct {
	rconHost     string
	rconPassword string
	commands     commandFlag
}

var options = &flags{}

func parseFlags() {

	flag.StringVar(&options.rconHost, "rcon-host", "127.0.0.1:27015", "RCON host and port")
	flag.StringVar(&options.rconPassword, "rcon-password", "", "RCON password")
	flag.Var(&options.commands, "command", "Command to execute")

	flag.Parse()
}

func main() {

	parseFlags()

	conn, err := rcon.Dial(options.rconHost, options.rconPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, cmd := range options.commands {
		response, err := conn.Execute(cmd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
	}
}

// Define a custom flag type to store multiple command values
type commandFlag []string

func (c *commandFlag) String() string {
	return strings.Join(*c, ", ")
}

func (c *commandFlag) Set(value string) error {
	*c = append(*c, value)
	return nil
}
