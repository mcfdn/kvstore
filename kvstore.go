package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/mcfdn/kvstore/operations"
	"github.com/mcfdn/kvstore/store"
)

func main() {
	hostPtr := flag.String("h", "localhost", "The host to listen on")
	portPtr := flag.Int("p", 7777, "The port to listen on")
	flag.Parse()

	router := operations.NewRouter()
	operations.RegisterOperations(router)

	listen(*hostPtr, *portPtr, router, store.New())
}

func listen(host string, port int, r *operations.Router, s *store.Store) {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	defer ln.Close()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening on %s:%d\n", host, port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn, r, s)
	}
}

func handleConnection(conn net.Conn, r *operations.Router, s *store.Store) {
	defer conn.Close()
	defer fmt.Println("Client closed connection")

	fmt.Println("Client initiated connection")

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		routeMessage(message, conn, r, s)
	}
}

func routeMessage(message string, conn net.Conn, r *operations.Router, s *store.Store) {
	args := strings.Fields(message)
	if len(args) < 1 {
		return
	}

	ctx := context.WithValue(context.Background(), "args", args[1:])

	result, err := r.Route(ctx, s, args[0])
	if err != nil {
		conn.Write([]byte(fmt.Sprintln(err.Error())))

		return
	}

	if val := result.Value; val != "" {
		conn.Write([]byte(fmt.Sprintln(val)))
	}

	conn.Write([]byte(fmt.Sprintln("OK")))
}
