package server

import (
	"air_crug/config"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// ServerMain запуск сервера
func ServerMain() {
	l, err := net.Listen(config.TYPE, config.HOST+":"+config.PORT)
	allClients := new(Connected)
	if err != nil {
		fmt.Println("ERROR listening: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	// Новое соединение
	newConnections := make(chan net.Conn)

	fmt.Println("Listening on " + config.HOST + ":" + config.PORT)
	go func() {
		for {
			// канал для новых сообщении/комманд
			mess := make(chan string)
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
			}
			newConnections <- conn
			allClients.Add(conn, mess)
			go handleRequest(conn, mess)
		}
	}()

	// Ожидаем новых соединениии
	for {
		select {
		case conn := <-newConnections:
			log.Printf("Accepted new client: ", conn.RemoteAddr().String())
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
			break
		}
		mess <- res
	}

	//conn.Close()
}
