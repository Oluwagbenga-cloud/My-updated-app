package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	log.Fatal(err)

	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	logFatal(err)

	username = strings.Trim(username, "\r\n")

	welcomeMsg := fmt.Sprintf("Welcome %s, to the chat, Hi Deca Gophers.", username)

	fmt.Println(welcomeMsg)

}
