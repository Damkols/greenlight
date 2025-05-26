package main

import(
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)

}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": movies}

	err:= app.writeJson(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}
