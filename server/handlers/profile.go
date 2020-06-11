package handlers

import (
	"encoding/json"
	"net/http"
	"turbo-parakeet/entities"
	"turbo-parakeet/server/middlewares"
	"turbo-parakeet/utils"
)

func HandleAuth(rsw http.ResponseWriter, req *http.Request) {
	var err error
	var ctx *middlewares.LocalContext = loadContext(req, `Handle Auth`)
	var res *response = &response{Status: `fail`}
	defer func() {
		if err != nil {
			ctx.Log.Error(err)
			res.Message = err.Error()
		}

		resByte, _ := json.Marshal(res)
		rsw.Write(resByte)
	}()

	profile := entities.Profile{}

	err = json.NewDecoder(req.Body).Decode(&profile)
	if err != nil {
		rsw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = profile.Auth()
	if err != nil {
		return
	}

	jwt, err := utils.NewJWT(ctx.Conf.JWT.ExpirationTime, ctx.Conf.JWT.SecretRaw)
	if err != nil {
		return
	}
	token, err := jwt.GenerateToken(profile.Username, profile.Email, ctx.Log.GetServerContextID())
	if err != nil {
		rsw.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(rsw, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: jwt.Expiration,
	})

	res.Status = `success`
	res.Message = `Autentication Success`
	res.Result = map[string]interface{}{
		`username`: profile.Username,
		`token`:    token,
	}

}
