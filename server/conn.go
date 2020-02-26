package server

import (
	"air_crug/log_view"
	"fmt"
	"log"
	"net"
)

// Client пользователь
type Client struct {
	Connect net.Conn
	Mess    []string
	Ip      net.Addr
}

// Тут регистрируем события/команды
func (c *Client) registryEvent(mess chan string, allClients *Connected) {
	files := new(log_view.LogFiles)
	files.Init()
	closed := make(chan bool, 0)
	for {
		select {
		case message := <-mess:
			c.Mess = append(c.Mess, message)
			fmt.Println("Написал: ", message)
			// Отдаем логи
			if message == "#log" {
				for _, file := range files.Files {
					c.Connect.Write([]byte(file.Body))
				}
			} else if message == "#view_clients" {
				// показываем все соедиенния
				for _, client := range allClients.Clients {
					c.Connect.Write([]byte(client.Ip.String()))
				}
			} else if message == "#my_info" {
				// Информация о данном пользователе
				_, err := c.Connect.Write([]byte(fmt.Sprintln("Ваш IP адресс:", c.Connect.RemoteAddr().String()+"\n")))
				if err != nil {
					log.Fatal(err)
				}
				for _, m := range c.Mess {
					_, err := c.Connect.Write([]byte(m + "\n"))
					if err != nil {
						log.Fatal(err)
					}
				}
			} else if message == "exit" {
				newClient := new(Connected)
				for _, client := range allClients.Clients {
					if client.Connect != c.Connect {
						newClient.Clients = append(newClient.Clients, client)
					}
				}
				allClients.Clients = newClient.Clients
				c.Connect.Close()
				closed <- true
				break
			}
		// Если была введена команда выхода то необходимо закрыть каналы
		case exit := <-closed:
			if exit {
				close(mess)
				close(closed)
			}

		}
	}
}

// Connected Все пользователи
type Connected struct {
	Clients []Client
}

// Add Добавляем нового клиента
func (c *Connected) Add(client net.Conn, mess chan string) {
	clients := Client{Connect: client, Ip: client.RemoteAddr()}
	c.Clients = append(c.Clients, clients)
	go clients.registryEvent(mess, c)
}
