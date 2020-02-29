package orm

import "database/sql"

// Table Таблицы
type Table struct {
	Name   string
	Fields []Field
	Db     *sql.DB
}

// Add Добавление таблицы
func (t *Table) Add() {

}

// Update Обновление таблицы
func (t *Table) Update() {

}

// Delete Удаление таблицы
func (t *Table) Delete() {

}
