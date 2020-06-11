package database

import (
	"database/sql"
	"fmt"
)

type dbName string

// Entity e a interface padrao das entidades
type Entity interface {
	GetOne(*sql.Row) error
	GetAll() error
	InsertOne() error
	InsertAll() error
}

// Database e a interface padrao de banco de dados
type Database interface {
	Open() error
	Close() error
	QueryOne(string, []interface{}, Entity) error
	QueryMulti(string, []interface{}, Entity) error
	InsertOne(string, []interface{}, Entity) error
	InsertMulti(string, []interface{}, Entity) error
}

// Default tem as configuracoes padroes do banco de dados
type Default struct {
	DBs     map[dbName]Database
	Current Database
}

// NewDatabase retorna a estrutura padrao preparada para uso
func NewDatabase() *Default {
	return &Default{
		DBs:     make(map[dbName]Database),
		Current: nil,
	}
}

// RegisterNewDB registra um novo banco de dados para o sistema
func (d *Default) RegisterNewDB(name string, db Database) (err error) {
	d.DBs[dbName(name)] = db
	return
}

// LoadDB Carrega um banco de dados para sua utilizacao em Current
func (d *Default) LoadDB(name string) (err error) {
	db, err := d.GetDB(name)
	if err != nil {
		return
	}

	d.Current = db
	return
}

// GetDB retorna o banco de dados para sua utilizacao
func (d *Default) GetDB(name string) (db Database, err error) {
	db, thisExists := d.DBs[dbName(name)]
	if !thisExists {
		return nil, fmt.Errorf(`Database %s is not registered yet`, name)
	}
	return
}
