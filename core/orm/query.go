package orm

import "database/sql"

// Query Запросы
type Query struct {
	Db *sql.DB
}

// FindOne Поиск одного элемента
func (q *Query) FindOne() {

}

// FindAll Поиск множества элементов
func (q *Query) FindAll() {

}

// Update Обновление
func (q *Query) Update() {

}

// Remove Удаление
func (q *Query) Remove() {

}
