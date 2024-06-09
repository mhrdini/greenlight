package main

import (
	"fmt"
	"net/http"

	"github.com/mhrdini/greenlight/internal/data"
	"github.com/mhrdini/greenlight/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	// declare validator in handler for more flexibility in using it for more complex checks
	validator := validator.New()
	if movie.Validate(validator); !validator.Valid() {
		app.failedValidationResponse(w, r, validator.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
