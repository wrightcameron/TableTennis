package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func main() {
	//Get arguments from command line
	args := os.Args[1:]
	if len(args) != 3 {
		//TODO: Make text red
		fmt.Println("Usage:  ./TableTennis <server|client> <serverHost> <port#>")
		os.Exit(-1)
	}

	//Parse arguments into variables
	isServer := false
	if strings.EqualFold("server", args[0]) {
		isServer = true
	}
	//TODO Add error handling if the user does not add the right types.
	serverHost := args[1]
	portNum := args[2]

	//Either launch the server or the client
	if isServer {
		server(portNum)
	} else {
		client(serverHost, portNum)
	}
}

func server(port string) {
	port = ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	// conn, err := listener.Accept()
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	fmt.Println("Connection established.")
	for {
		//Accept new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error")
		}

		//Send it a ping
		_, err = conn.Write([]byte("Ping"))
		fmt.Println("wrote Ping")
		//result, err := ioutil.ReadAll(conn)
		//checkError(err)
		//fmt.Println(string(result))

		conn.Close()
	}
	//conn.Close()
}

func client(serverHost string, port string) {

	serverHost = serverHost + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverHost)
	checkError(err)
	// conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// checkError(err)
	for {
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		checkError(err)

		result, err := ioutil.ReadAll(conn)
		checkError(err)
		fmt.Println(string(result))
		//_, err = conn.Write([]byte("Pong"))

		conn.Close()
	}
	//checkError(err)
	//os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
