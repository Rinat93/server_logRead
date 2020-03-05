package server

import (
	"air_crug/core"
	"air_crug/log_view"
	"errors"
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
	User    *core.User
}
type error interface {
	Error() string
}

func (c *Client) AuthIfUser() error {
	if c.User == nil {
		c.Connect.Write([]byte(fmt.Sprint("Авторизуйтесь или зарегистрируйтесь(#login или #register)")))
		return errors.New("Не авторизирован пользователь")
	}
	return nil
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
			// Авторизация
			case "#login":
				if c.AuthIfUser() != nil {
					if len(command) < 3 {
						c.Connect.Write([]byte(fmt.Sprint("Введите имя и пароль пользователя через пробел \n не хватает аргументов \n", command)))
						continue
					}
					allClients.db.GetUser(command[1], command[2])
					c.User = allClients.db.User
					c.Connect.Write([]byte(fmt.Sprintf("Вы авторизованы ваша информация login: %s\t пароль: %s\t ID: %d \n", c.User.Login, c.User.Password, c.User.Id)))
				} else {
					c.Connect.Write([]byte(fmt.Sprint("Вы уже авторизированы")))
				}
				// Регистрация новых пользователей(можно установить ограничения)
			case "#register":
				if c.AuthIfUser() != nil {
					if len(command) < 3 {
						c.Connect.Write([]byte(fmt.Sprint("Введите имя и пароль пользователя через пробел \n не хватает аргументов", command)))
						continue
					}
					allClients.db.SetUser(command[1], command[2])
					c.User = allClients.db.User
					c.Connect.Write([]byte(fmt.Sprintf("Вы зарегистрированы ваша информация login: %s\t пароль: %s\t ID: %d \n", c.User.Login, c.User.Password, c.User.Id)))
				} else {
					c.Connect.Write([]byte(fmt.Sprint("Вы уже авторизированы")))
				}
			// Вывод лог файлы
			case "#log":
				if c.AuthIfUser() != nil {
					continue
				}
				for _, file := range files.Files {
					c.Connect.Write([]byte(fmt.Sprintf("Имя: %s\t Дата ред: %s\t Путь: %s\t Размер: %d\n", file.Name, file.ModTime, file.Path, file.Size)))
					// c.Connect.Write([]byte(fmt.Sprintln(string(file.Body))))
				}
			// Чтение лог файла
			case "#read":
				if c.AuthIfUser() != nil {
					continue
				}
				if len(command) > 1 {
					for _, f := range command[1:] {
						for _, af := range files.Files {
							if f == af.Path+af.Name {
								c.Connect.Write([]byte(fmt.Sprintln(string(af.Body))))
							}
						}
					}
				}

			// Добавить новый путь логов
			case "+log":
				if c.AuthIfUser() != nil {
					continue
				}
				if len(command) > 1 {
					for _, path := range command[1:] {
						files.Add(path)
					}
				}
			// Вывести всех активных пользователей
			case "#view_clients":
				if c.AuthIfUser() != nil {
					continue
				}
				// показываем все соедиенния
				for _, client := range allClients.Clients {
					c.Connect.Write([]byte(client.Ip.String()))
				}
			// Моя информация
			case "#my_info":
				if c.AuthIfUser() != nil {
					continue
				}
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
			// Выйти
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
	db      *core.SCore
}

// Add Добавляем нового клиента
func (c *Connected) Add(client net.Conn, mess chan string) {
	clients := Client{Connect: client, Ip: client.RemoteAddr()}
	c.Clients = append(c.Clients, clients)
	go clients.registryEvent(mess, c)
}
