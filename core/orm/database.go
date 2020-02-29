package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database struct {
	Name string
	Db   *sql.DB // хранить соединение с базой
}

// Create - Создание базы
func (d *Database) Create() error {
	f, err := os.Open(fmt.Sprintf("../databases/$1", d.Name))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b1 := make([]byte, 5000)
	_, err = f.Read(b1)
	if err != nil {
		log.Fatal(err)
	}
	_, err = d.Db.Exec(string(b1))
	if err != nil {
		log.Fatal(err)
	}
	return nil

}
