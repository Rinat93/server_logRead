package orm

import (
	"database/sql"
	"fmt"
)

// Table Таблицы
type Table struct {
	Name   string
	Column []Column
	Db     *sql.DB
}

// Create Добавление таблицы
func (t *Table) Create() {
	var table string = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, ", t.Name)
	for _, columns := range t.Column {
		table += fmt.Sprintf("%s  %s,", columns.Name, columns.Types)
	}
	table += ");"
	t.Db.Exec(table)
}

// Delete Удаление таблицы
func (t *Table) Delete() {
	t.Db.Exec("DROP TABLE %s", t.Name)
}
