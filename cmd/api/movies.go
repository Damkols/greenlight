package main

import(
	"net/http"
	"time"
	"greenlight.damkols.net/internal/data"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r, err)
		return
	}

	movie := data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "Casablanca",
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
		Version: 1,
	}

	err = app.writeJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
