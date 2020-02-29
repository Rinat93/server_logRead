package orm

import (
	"database/sql"
	"fmt"
	"log"
)

// Query Запросы
type Query struct {
	Db    *sql.DB
	Table Table
}

// FindOne Поиск одного элемента
func (q *Query) Find(find map[string]interface{}) *sql.Rows {
	var query string
	for key, value := range find {
		query += fmt.Sprintf("%s=%s and", key, value)
	}
	rows, err := q.Db.Query("select * from $1 where $2", q.Table.Name, query)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

// Update Обновление
func (q *Query) Update() {

}

// Remove Удаление
func (q *Query) Remove() {

}
