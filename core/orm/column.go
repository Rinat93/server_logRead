package orm

import "database/sql"

// Field Работа с столбца БД
type Column struct {
	Name      string
	Types     string
	MaxLength int16
	Table     string
	Db        *sql.DB
}

// Add - добавление столбца
func (c *Column) Add() {
	c.Db.Exec("ALTER TABLE $1 ADD COLUMN $2 $3", c.Table, c.Name, c.Types)
}

// Rename - переименование столбца
func (c *Column) Rename() {
	c.Db.Exec("ALTER TABLE $1 ADD COLUMN $2 $3", c.Table, c.Name, c.Types)
}
