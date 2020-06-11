package main

import (
	"turbo-parakeet/configurations"
	"turbo-parakeet/database"
	"turbo-parakeet/server"
	"turbo-parakeet/utils"
)

var (
	runtime *Runtime
)

// Runtime carrega o escopo printpal do sistema
type Runtime struct {
	Configurations *configurations.Default
	Log            *utils.Logger
	Server         *server.Default
	DB             *database.Default
}

func init() {
	runtime = newRuntime()
	runtime.Log.Info(`Starting Service`)
}

func main() {
	defer runtime.Log.CloseLogFile()

	runtime.DB = database.NewDatabase()
	runtime.DB.RegisterNewDB(`mysql`, database.NewMySQL(runtime.Configurations.DBs[`mysql`]))

	runtime.Server = server.NewServer(runtime.Configurations.Server, runtime.Log, runtime.DB, runtime.Configurations)
	runtime.Server.Run()

}

func newRuntime() *Runtime {
	rt := &Runtime{
		Configurations: configurations.NewConfigurations(),
	}
	rt.Log = rt.Configurations.Log

	return rt
}
