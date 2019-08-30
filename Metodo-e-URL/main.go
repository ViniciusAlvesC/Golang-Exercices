package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	done := make(chan bool)
	go StartServer()
	log.Println("[INFO] Servidor no ar!")
	<-done

}

func getINFO(w http.ResponseWriter, r *http.Request) {
	qurl := r.URL.String()
	fmt.Fprintf(w, r.Method+" "+qurl)
}

func StartServer() {
	duration, _ := time.ParseDuration("1000ns")

	server := &http.Server{
		Addr:        "192.168.0.5:8082",
		IdleTimeout: duration,
		//	Handler    :
	}

	http.HandleFunc("/", getINFO)

	server.ListenAndServe()
}
