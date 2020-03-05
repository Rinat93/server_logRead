package server

import (
	"air_crug/core"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

// ServerMain запуск сервера команд
func ServerCommand() {
	l := new(core.SCore)
	l.Connect()
	allClients := new(Connected)
	defer l.Net.Close()
	// Новое соединение
	newConnections := make(chan net.Conn)

	go func() {
		for {
			// канал для новых сообщении/комманд
			mess := make(chan string)
			conn, err := l.Net.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
			}
			newConnections <- conn
			allClients.db = l
			allClients.Add(conn, mess)
			go handleRequest(conn, mess)
		}
	}()

	// Ожидаем новых соединениии
	for {
		select {
		case conn := <-newConnections:
			log.Printf("Accepted new client: ", conn.RemoteAddr().String())
			conn.Write([]byte(fmt.Sprintf("Введите логи и пароль(#login login password)!")))
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
