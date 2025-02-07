package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new Movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	//params := httprouter.ParamsFromContext(r.Context())

	// id, err := strconv.Atoi(params.ByName("id"))

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id)

}
