package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"statd/config"
	"statd/internal/cli"
	"statd/internal/server"
)

const (
	version = "0.0.0"
)

func usage() {
	fmt.Printf(`
StatD %s
Usage: %s <command> [options]

	Commands:
		server
		cli
	
	Environment Variables:
		%s
			defaults to %s
			is automatically created as well

	- use '<command> -h' for more information
`, version, filepath.Base(os.Args[0]),
		config.EnvVar, config.DefaultPath)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}
	operation := os.Args[1]
	if operation == "cli" {
		err := cli.Cmd()
		if err != nil {
			log.Fatalln(fmt.Errorf("Cli stopped. %w", err))
		}
	} else if operation == "server" {
		err := server.Serve()
		if err != nil {
			log.Fatalln(fmt.Errorf("Server failed. %w", err))
		}
	} else {
		log.Printf("Command should be either cli or server, not %s.\n", operation)
		os.Exit(0)
	}
}
