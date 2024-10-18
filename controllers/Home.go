package controllers

import (
	"log"
	"net/http"
	"os"
)

  /**
  * Attempts to get the IP from the request.
  * Falls back to RemoteAddr if IP can not be deduced from headers.
  */
func IpFromRequest(request *http.Request) (string, string) {
    var ip string

    if ip = request.Header.Get("X-Real-IP"); ip != "" {
        return "X-Real-IP", ip
    }

    if ip = request.Header.Get("X-Forwarded-For"); ip != "" {
        return "X-Forwarded-For", ip
    }

    return "RemoteAddr", request.RemoteAddr
}

func Home(writer http.ResponseWriter, request *http.Request) {
    data, err := os.ReadFile("./resources/views/home.html")
    cwd, _ := os.Getwd()
    ipSource, ip := IpFromRequest(request)

    log.Printf("Received %s from %s (%s)", request.Method, ip, ipSource)

    if err != nil {
        log.Printf("Could not read resources/views/home.html: %s (%s)", err.Error(), cwd) 
        writer.WriteHeader(500)
        writer.Write([]byte("An internal server error happened."))
        
        return
    }

    if request.Method != "GET" {
        writer.WriteHeader(405)
        writer.Write([]byte("Method not allowed."))

        return
    }

	writer.Write(data)
}
