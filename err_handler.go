package main

import "net/http"

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "something went terribly wrong")
}
