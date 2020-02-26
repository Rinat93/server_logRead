package server

import (
	"air_crug/log_view"
	"fmt"
	"log"
	"net"
	"strings"
)

// Client пользователь
type Client struct {
	Connect net.Conn
	Mess    []string
	Ip      net.Addr
}

// Тут регистрируем события/команды
func (c *Client) registryEvent(mess chan string, allClients *Connected) {
	closed := make(chan bool, 0)
	files := new(log_view.LogFiles)
	files.Init()
	for {
		select {
		case message := <-mess:
			c.Mess = append(c.Mess, message)
			command := strings.Split(message, " ")
			fmt.Println("Написал: ", command[0])
			if len(command) == 0 {
				continue
			}
			// Отдаем логи
			switch command[0] {
			case "#log":
				for _, file := range files.Files {
					c.Connect.Write([]byte(fmt.Sprintf("Имя: %s\t Дата ред: %s\t Путь: %s\t Размер: %d\n", file.Name, file.ModTime, file.Path, file.Size)))
					// c.Connect.Write([]byte(fmt.Sprintln(string(file.Body))))
				}
			case "+log":
				if len(command) > 1 {
					for _, path := range command[1:] {
						files.Add(path)
					}
				}
			case "#view_clients":
				// показываем все соедиенния
				for _, client := range allClients.Clients {
					c.Connect.Write([]byte(client.Ip.String()))
				}
			case "#my_info":
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
			case "#exit":
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
