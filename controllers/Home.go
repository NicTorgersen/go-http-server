package controllers

import (
	"log"
	"net/http"
	"os"
)

func Home(writer http.ResponseWriter, request *http.Request) {
    data, err := os.ReadFile("./resources/views/home.html")
    cwd, _ := os.Getwd()

    if err != nil {
        log.Printf("Could not read resources/views/home.html: %s (%s)", err.Error(), cwd) 
        writer.WriteHeader(500)
        writer.Write([]byte("An internal server error happened."))
    }

	writer.Write(data)
}
