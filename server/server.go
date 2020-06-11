package server

import (
	"net/http"
	"time"
	"turbo-parakeet/configurations"
	"turbo-parakeet/database"
	"turbo-parakeet/server/middlewares"
	"turbo-parakeet/server/routes"
	"turbo-parakeet/utils"
)

var runServer *Default

// Default sao as configuracoes padroes do servidor
type Default struct {
	StartDate int64
	Context   *Context
	http.Server
}

type Context struct {
	Log  *utils.Logger
	DB   *database.Default
	Conf *configurations.Default
}

// NewServer retorna uma nova configuracao de servidor
func NewServer(serverConfs *configurations.Server, ctxLog *utils.Logger, ctxDBs *database.Default, ctxConf *configurations.Default) *Default {
	runServer = &Default{
		StartDate: time.Now().Unix(),
		Context: &Context{
			Log:  ctxLog,
			DB:   ctxDBs,
			Conf: ctxConf,
		},
	}

	runServer.Addr = serverConfs.FullAddress()
	runServer.Handler = routes.NewRoutes()

	return runServer
}

// Run inicia o servidor
func (d *Default) Run() {
	populateContext(d.Context)
	d.Context.Log.Info(`Start listen and serve on: ` + d.Addr)
	d.Context.Log.Fatal(d.ListenAndServe())
}

func populateContext(ctx *Context) {
	middlewares.LoadContext(ctx.Log, ctx.DB, ctx.Conf)
}
