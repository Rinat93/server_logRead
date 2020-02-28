package core

import (
	"air_crug/config"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net"
	"os"
)

// SCore Главная структура соединения
type SCore struct {
	Net net.Listener
	Db  *sql.DB // хранить соединение с базой
}

// CreateDb Создание таблиц БД
func (c *SCore) CreateDb() error {
	f, err := os.Open("./databases/user.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b1 := make([]byte, 5000)
	_, err = f.Read(b1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b1))
	res, err := c.Db.Exec(string(b1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return nil

}

// Connect Запускаем сервер
func (c *SCore) Connect() {
	l, err := net.Listen(config.TYPE, config.HOST+":"+config.PORT)
	if err != nil {
		log.Fatal(err)
	}
	(*c).Net = l
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	(*c).Db = db
	c.CreateDb()
	fmt.Println("Listening on " + config.HOST + ":" + config.PORT)
}
