package main

import (
	"log"
	"net/http"
	"strconv"

	"http-boilerplate/server/httphandler"
)

const port = 8000

func main() {
	httpHandler := httphandler.InitHTTPHandler()
	httpHandler.RegisterRoutes()

	server := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: httpHandler.Mux,
	}

	log.Println("Starting Server on Port ", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Fail serving: ", err)
	}
}
