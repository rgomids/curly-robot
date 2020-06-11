package handlers

import (
	"net/http"
	"turbo-parakeet/server/middlewares"
)

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func loadContext(req *http.Request, msg string) *middlewares.LocalContext {
	ctxRaw := req.Context().Value(`ctx`)
	ctx := ctxRaw.(*middlewares.LocalContext)
	ctx.Log.Info(msg)
	return ctx
}

func HandleAlive(rsw http.ResponseWriter, req *http.Request) {
	ctx := loadContext(req, `Handle Alive`)
	ctx.Log.Info(`Server is Running correctly!`)
	rsw.Write([]byte(`{"status": "ok"}`))
}
