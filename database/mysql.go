package database

import (
	"database/sql"
	"turbo-parakeet/configurations"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL banco de dados relacional
type MySQL struct {
	connectionAddress string
	client            *sql.DB
}

// NewMySQL retorna uma nova instancia do banco de dados MySQL
func NewMySQL(confs *configurations.Database) *MySQL {
	return &MySQL{
		connectionAddress: confs.FullAddress(),
		client:            nil,
	}
}

// Open abre a conexao com o banco
func (m *MySQL) Open() (err error) {
	m.client, err = sql.Open(`mysql`, m.connectionAddress)
	return
}

// Close fecha a conexao com o banco
func (m *MySQL) Close() (err error) {
	err = m.client.Close()
	return
}

// QueryOne realiza uma busca com o retorno de um unico resultado
func (m *MySQL) QueryOne(query string, args []interface{}, entity Entity) (err error) {
	m.Open()
	defer m.Close()
	return entity.GetOne(m.client.QueryRow(query, args...))
}

// QueryMulti realiza uma busca com o retorno de muitos resultados
func (m *MySQL) QueryMulti(query string, args []interface{}, entity Entity) (err error) {
	m.Open()
	defer m.Close()
	return
}

// InsertOne insere um unico registro
func (m *MySQL) InsertOne(query string, args []interface{}, entity Entity) (err error) {
	m.Open()
	defer m.Close()
	return
}

// InsertMulti insere muitos registros
func (m *MySQL) InsertMulti(query string, args []interface{}, entity Entity) (err error) {
	m.Open()
	defer m.Close()
	return
}
