package config

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handleFunc func(http.ResponseWriter, *http.Request, httprouter.Params)

func apply(handle handleFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		handle(w, r, params)
	}
}
