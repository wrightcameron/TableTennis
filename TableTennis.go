package main

import (
	"fmt"
	"os"
	"strconv"
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
	portNum, err := strconv.Atoi(args[2])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(-1)
	}

	//Either launch the server or the client
	if isServer {
		server(serverHost, portNum)
	} else {
		client(serverHost, portNum)
	}
}

func server(serverHost string, port int) {

}

func client(serverHost string, port int) {

}
