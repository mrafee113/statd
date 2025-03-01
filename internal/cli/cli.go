package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"statd/internal/server"
)

func Cmd() error {
	fs := flag.NewFlagSet("cli", flag.ExitOnError)
	flagMap := make(map[string]*bool)
	flagMap["battery-charge"] = fs.Bool("battery-charge", false, "Device battery charge")
	flagMap["date"] = fs.Bool("date", false, "Current date and time")
	flagMap["reload-conf"] = fs.Bool("reload-conf", false, "Reload configurations")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s cli [options]\n- Note that one and only one option can be selected.\n", os.Args[0])
		fs.PrintDefaults()
	}
	fs.Parse(os.Args[2:])

	var count int
	var cmd string
	for key, value := range flagMap {
		if *value {
			count++
			cmd = key
		}
	}
	if count != 1 {
		return fmt.Errorf("ArgumentError: One and only one option must be provided alongside the cli command.")
	}

	conn, err := net.Dial("unix", server.SocketPath)
	if err != nil {
		return fmt.Errorf("DialError: %w", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(cmd))
	if err != nil {
		return fmt.Errorf("WriteError: %w", err)
	}

	var response []byte
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil && errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}
		response = append(response, buf[:n]...)
	}
	fmt.Println(string(response))
	return nil
}
