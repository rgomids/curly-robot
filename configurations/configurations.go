package configurations

import (
	"encoding/json"
	"fmt"
	"os"
	"turbo-parakeet/utils"
)

type dbName string

var (
	runtimeLogger         *utils.Logger
	configurationsDefault *Default
)

// Default sao as configuracoes padroes do sistema
type Default struct {
	DBs    map[dbName]*Database `json:"databases"`
	Server *Server              `json:"server"`
	JWT    *JWT                 `json:"jwt"`
	Log    *utils.Logger
}

// Server configuracoes do servidor
type Server struct {
	Port string `json:"port"`
	Host string `json:"host"`
	SSL  bool   `json:"ssl"`
}

// JWT e a estrutura base do componente JTW
type JWT struct {
	SecretRaw      string `json:"secret"`
	ExpirationTime int    `json:"expiration_time"`
}

// Database configuracoes do banco de dados
type Database struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
}

func init() {
	var err error
	runtimeLogger = utils.NewLogger()
	runtimeLogger.Info(`Loading Configurations`)

	defer func() {
		if err != nil {
			runtimeLogger.Fatal(err)
		}
	}()

	var configFile *os.File
	var projectName = os.Getenv("TURBO_PARAKEET_DEFAULT_NAME")
	if configFile, err = os.Open(fmt.Sprintf("/etc/%s/configs.json", projectName)); err == nil {
		confDecoded := json.NewDecoder(configFile)
		err = confDecoded.Decode(&configurationsDefault)
		if configurationsDefault.JWT.SecretRaw == "" {
			configurationsDefault.JWT.SecretRaw = fmt.Sprintf("/opt/%s/.key/%s", projectName, projectName)
		}
	}

	// TODO: O que fazer quando as configuracoes nao estiverem preenchidas...
}

// NewConfigurations carrega as configuracoes
func NewConfigurations() *Default {
	configurationsDefault.Log = runtimeLogger
	return configurationsDefault
}

// FullAddress retorna a string de configuracao
func (d *Database) FullAddress() string {
	return fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, d.User, d.Password, d.Host, d.Port, d.Schema)
}

// FullAddress retorna a string de configuracao
func (s *Server) FullAddress() string {
	return fmt.Sprintf(`%s:%s`, s.Host, s.Port)
}
