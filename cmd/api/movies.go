package main

import (
	"fmt"
	"net/http"

	"github.com/mhrdini/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:      id,
		Title:   "Eternal Sunshine of the Spotless Mind",
		Year:    2004,
		Runtime: 108,
		Genres:  []string{"scifi", "romance", "drama"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.errorLog.Print(err)
		app.serverErrorResponse(w, r, err)
	}
}
