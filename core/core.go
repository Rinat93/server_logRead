package core

import (
	"air_crug/config"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// SCore Главная структура соединения
type SCore struct {
	Net  net.Listener
	Db   *sql.DB // хранить соединение с базой
	User *User
}

//
type User struct {
	Id        int32
	Login     string
	Password  string
	Dates_reg time.Time
	Hystory   []Hystory
}

type Hystory struct {
	Id       int32
	Dates    time.Time
	User_id  int32
	Commands []Commands
}

type Commands struct {
	Id         int32
	Dates      time.Time
	Commands   string
	Command_id int32
}

// CreateDb Создание таблиц БД
func (c *SCore) CreateDb() error {
	f, err := os.Open("../databases/user.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b1 := make([]byte, 5000)
	_, err = f.Read(b1)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.Db.Exec(string(b1))
	if err != nil {
		log.Fatal(err)
	}
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

// GetUser Получить пользователей
func (c *SCore) GetUser(login string, password string) {
	rows, err := c.Db.Query("select * from User where login=$1 password=$2", login, password)
	if err != nil {
		log.Fatal(err)
	}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Dates_reg)
		if err != nil {
			log.Fatal("error getuser", err)
		}
		fmt.Println(user.Id)
	}
	c.User = &user
	c.SetHistory()
}

// SetUser Создание пользователя
func (c *SCore) SetUser(login string, password string) {
	_, err := c.Db.Exec("insert into User (login, password, dates_reg) values ($1, $2, $3)", login, password, time.Now().String())
	if err != nil {
		log.Fatal("Уже есть ", err)
	}
	c.GetUser(login, password)
}

// SetPassword Заменить пароль
func (c *SCore) SetPassword(password string) {
	c.User.Password = password
	_, err := c.Db.Exec("insert into User (password) values ($1)", password)
	if err != nil {
		log.Fatal(err)
	}
}

// SetHistory  Записать историю
func (c *SCore) SetHistory() {
	hystory := Hystory{
		Dates:   time.Now(),
		User_id: c.User.Id,
	}
	c.User.Hystory = append(c.User.Hystory, hystory)
	_, err := c.Db.Exec("insert into History (dates,user_id) values ($1,$2)", hystory.Dates.String(), hystory.User_id)
	if err != nil {
		log.Fatal(err)
	}
}
