package routes

import (
	"turbo-parakeet/server/handlers"
	"turbo-parakeet/server/middlewares"

	"github.com/gorilla/mux"
)

/* TODO:	O sistema de rotas aqui esta ruim, pois, nao existe uma clara distincao
 * sobre as rotas privadas e as rotas publicas. Vou deixar dessa forma mas
 * a ideia e corrigir isso, garantindo que as rotas tenham seu alinhamento correto
 */

// NewRoutes retorna todas as rotas configuradas
func NewRoutes() *mux.Router {
	r := mux.NewRouter()

	// Rotas Publicas
	r = r.PathPrefix(`/api`).Subrouter()
	r.HandleFunc(`/alive`, handlers.HandleAlive).Methods(`GET`)
	r.HandleFunc(`/auth`, handlers.HandleAuth).Methods(`POST`)

	r.Use(middlewares.Default)

	// Rotas Privadas
	s := r.PathPrefix(``).Subrouter()
	s.HandleFunc(`/alive_sec`, handlers.HandleAlive).Methods(`GET`)

	s.Use(middlewares.Private)

	return r
}
