package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bramsedana/law-cots-1/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	handler := handler.NewHandler()

	router := httprouter.New()

	router.GET("/healthz", handler.Healthz)
	router.GET("/arithmetics", handler.Arithmetic)

	// Start server
	log.Println("Listening at port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
