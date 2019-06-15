# Table Tennis

* Author: Cameron Wright

## Overview

Program consists of a server and a client sending messages of Ping and Pong back
and forth between each other.

Used for understanding server and client socket connections in golang

## Manifest

|Filename       | Description                                                       |
|---------------|-------------------------------------------------------------------|
|TableTennis.go | Main file containing client and server                            |
|README.md      | This file                                                         |

## Building the project

Go as a language has a built in build automation tool that makes life easier.  Assuming
that you have go installed, to build the main file, run the command 

`$ go build`

A binary file titled "TableTennis" will be created that can be executed with

`$ ./TableTennis`

Follow instructins of program and output will be pushed to console. To clean 
repository run following command

`$ go clean`

## Features and usage

This program requires two instances of the program to be run, so have either two computers or open two shell consoles.

When running the program it will ask for several command line arguments.

Usage:  `./TableTennis <server|client> <serverHost> <port#>`

The server will need to be run first and requires a port number.  ServerHost is not required but current argument handling wants something.  Client does need a server host and port number.

Once a server instance and client are running them will continue to send messages back and forth till the user kills both.
