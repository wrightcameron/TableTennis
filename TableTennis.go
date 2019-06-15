package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//Get arguments from command line
	args := os.Args[1:]
	if len(args) != 3 {
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
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("Connection established.")
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := "Pong"
		fmt.Print("Text to send: " + newmessage + "\n")
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
	//conn.Close()
}

func client(serverHost string, port string) {

	serverHost = serverHost + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverHost)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	for {
		message := "Ping"
		fmt.Print("Text to send: " + message + "\n")
		//text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, message+"\n")
		// listen for reply
		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + response)
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
