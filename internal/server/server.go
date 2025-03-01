package server

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"statd/config"
	"statd/internal/scripts"
	"strings"
	"syscall"
	"time"
)

const SocketPath = "/tmp/statd.socket"

func initSocketFile() error {
	if _, err := os.Stat(SocketPath); err == nil {
		conn, err := net.Dial("unix", SocketPath)
		if err == nil {
			conn.Close()
			return fmt.Errorf("ServerAlreadyRunningError")
		}
		if err := os.Remove(SocketPath); err != nil {
			return fmt.Errorf("FileDeletionError: %w", err)
		}
	}
	return nil
}

func Serve() error {
	fs := flag.NewFlagSet("server", flag.ExitOnError)
	configPath := fs.String("conf", "", "Path to store configurations")
	fs.Usage = func() {
		fmt.Printf("Usage: %s server [options]\n\n", os.Args[0])
		fs.PrintDefaults()
	}
	fs.Parse(os.Args[2:])

	config.SelectConfigFile(*configPath)
	config.LoadConfig()

	if err := initSocketFile(); err != nil {
		return err
	}

	listener, err := net.Listen("unix", SocketPath)
	if err != nil {
		return fmt.Errorf("Failed to open unix socket at %s. %w", SocketPath, err)
	}

	ctxt, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go tearDown(sigChan, listener, cancel)

	fmt.Printf("Server listening on %s\n", SocketPath)

	for {
		select {
		case <-ctxt.Done():
			break
		default:
			conn, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					break
				}
				fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("AcceptError: %w", err))
				continue
			}
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil && errors.Is(err, io.EOF) {
		fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("ReadError: %w", err))
		return
	}
	cmd := string(buf[:n])
	response, err := handleCmd(cmd)
	if err != nil {
		response = fmt.Sprintf("%v, stdout=%s", err, response)
		_, err = conn.Write([]byte(response))
		return
	}
	_, err = conn.Write([]byte(response))
}

func handleCmd(cmd string) (string, error) {
	var response string
	var err error

	switch strings.TrimSpace(cmd) {
	case "date":
		response, err = scripts.Date()
	case "battery-charge":
		response, err = scripts.BatteryCharge()
	case "reload-conf":
		err = config.LoadConfig()
		if err != nil {
			response = "done"
		}
	default:
		return "", fmt.Errorf("UnknownCommandError: %s", cmd)
	}
	return response, err
}

func tearDown(sigChan chan os.Signal, listener net.Listener, ctxtCancel context.CancelFunc) {
	sig := <-sigChan
	log.Printf("Received signal %s. Shutting down...", sig)
	ctxtCancel()
	listener.Close()
	time.Sleep(100 * time.Millisecond)
	if err := os.Remove(SocketPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Printf("FileDeletionError: %v", err)
	}
	os.Exit(0)
}
