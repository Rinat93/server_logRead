package server

import (
	"air_crug/log_view"
	"fmt"
	"net"
)

// Client пользователь
type Client struct {
	Connect net.Conn
	Mess    []string
}

// Тут регистрируем события/команды
func (c *Client) registryEvent(mess <-chan string, allClients *Connected) {
	files := new(log_view.LogFiles)
	files.Init()
	for {
		select {
		case message := <-mess:
			fmt.Println(message)
			c.Mess = append(c.Mess, message)
			if message == "#log" {
				for _, file := range files.Files {
					c.Connect.Write([]byte(file.Body))
				}
			} else if message == "#view_clients" {
				for _, client := range allClients.Clients {
					c.Connect.Write([]byte(client.Connect.RemoteAddr().String()))
				}
			} else if message == "#my_info" {
				c.Connect.Write([]byte(c.Connect.RemoteAddr().String() + "\n"))
				for _, m := range c.Mess {
					c.Connect.Write([]byte(m + "\n"))
				}
			}
		}
	}
}

// Connected Все пользователи
type Connected struct {
	Clients []Client
}

// Add Добавляем нового клиента
func (c *Connected) Add(client net.Conn, mess <-chan string) {
	clients := Client{Connect: client}
	c.Clients = append(c.Clients, clients)
	go clients.registryEvent(mess, c)
}
