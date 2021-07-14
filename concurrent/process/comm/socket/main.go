package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go tcpClient()
	go tcpServer()

	wg.Wait()
}

func tcpClient() {
	/*
		1.initiate tcp dialing and get conn
		2.sent data by conn.Write()
	*/

	defer wg.Done()

	//net.DialTCP()
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}

	// if not closing the connection
	// client -> ESTABLISHED
	// server -> ESTABLISHED
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, err := inputReader.ReadString('\n')

		if err != nil {
			log.Println(err)
			break
		}

		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			log.Println(err)
			break
		}
	}

}

func tcpServer() {
	/*
		1.lisnten tcp address
		2.get client request and create connection
		3.receive connection and create a channel for processing
	*/

	// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	// not close
	// server -> CLOSE-WAIT
	// client -> FIN-WAIT-2
	defer conn.Close()

	// get the address of the previous request
	//conn.RemoteAddr()

	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err.Error())
			break
		}

		log.Printf("the data received from the client is: %v\n", string(buf[:n]))
	}

	// 演示
	wg.Done()
}
