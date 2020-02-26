package main

import (
	"air_crug/config"
	"air_crug/log_view"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen(config.TYPE, config.HOST+":"+config.PORT)
	files := new(log_view.LogFiles)
	files.Init()
	if err != nil {
		fmt.Println("ERROR listening: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	// Новое соединение
	newConnections := make(chan net.Conn)
	// Новое сообщение
	mess := make(chan string)
	// Все пользователи
	allClients := make(map[net.Conn]int)
	fmt.Println("Listening on " + config.HOST + ":" + config.PORT)
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
			}
			newConnections <- conn
			allClients[conn] = 1

		}
	}()

	// Ожидаем событий
	for {
		select {
		case conn := <-newConnections:
			log.Printf("Accepted new client")
			go handleRequest(conn, mess)
		case message := <-mess:
			fmt.Println(message)
			if err != nil {
				fmt.Printf("parsing failed: %s\n", err)
				continue
			}
			if message == "#log" {
				for _, file := range files.Files {
					for user, _ := range allClients {
						user.Write([]byte(file.Body))
					}
				}
			}
		}
	}

}

func handleRequest(conn net.Conn, mess chan<- string) {

	for {
		buf := make([]byte, 2048)
		_, err := conn.Read(buf)
		buf = bytes.Trim(buf, "\x00")
		res := strings.TrimRight(string(buf), "\r\n")
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		mess <- res
	}

	//conn.Close()
}
