package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"turbo-parakeet/configurations"
	"turbo-parakeet/database"
	"turbo-parakeet/utils"
)

var ctx *LocalContext

func init() {
	ctx = &LocalContext{}
}

type LocalContext struct {
	Log  *utils.Logger
	DB   *database.Default
	Conf *configurations.Default
}

func (ctx *LocalContext) newSessionContext() *LocalContext {
	return &LocalContext{
		Log:  ctx.Log.NewSession(),
		DB:   ctx.DB,
		Conf: ctx.Conf,
	}
}

func LoadContext(log *utils.Logger, db *database.Default, conf *configurations.Default) {
	ctx.Log = log
	ctx.DB = db
	ctx.Conf = conf
}

func loggingMiddleware(next http.Handler, newCTX *LocalContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newCTX.Log.Info(fmt.Sprintf("REQUEST RECEIVED FROM %s - URL REQUESTED %s", r.RemoteAddr, r.RequestURI))

		ctx := context.WithValue(r.Context(), `ctx`, newCTX)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loadSession(next http.Handler) http.Handler {
	return loggingMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		}),
		ctx.newSessionContext(),
	)
}

// Default ...
func Default(next http.Handler) http.Handler {
	return loadSession(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}),
	)
}

func Private(next http.Handler) http.Handler {
	return loadSession(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				c, err := r.Cookie("token")
				if err != nil {
					if err == http.ErrNoCookie {
						w.WriteHeader(http.StatusUnauthorized)
						ctx.Log.Error(err)
						return
					}
					w.WriteHeader(http.StatusBadRequest)
					ctx.Log.Error(err)
					return
				}

				jwt, err := utils.NewJWT(ctx.Conf.JWT.ExpirationTime, ctx.Conf.JWT.SecretRaw)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					ctx.Log.Error(err)
					return
				}

				contextID, statusCode, err := jwt.ValidateToken(c.Value)
				if err != nil {
					w.WriteHeader(statusCode)
					ctx.Log.Error(err)
					return
				}

				if contextID != ctx.Log.GetServerContextID() {
					w.WriteHeader(http.StatusUnauthorized)
					ctx.Log.Error(err)
					return
				}

				renewToken, err := jwt.RefreshToken(c.Value)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					ctx.Log.Error(err)
					return
				}

				http.SetCookie(w, &http.Cookie{
					Name:    "token",
					Value:   renewToken,
					Expires: jwt.Expiration,
				})
			},
		),
	)
}
