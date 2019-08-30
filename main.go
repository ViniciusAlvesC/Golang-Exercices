package main

import (
		 "net/http"
		 "time"
		 "fmt"  )

const base64GifPixel = "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"

func main() {

	duration, _ := time.ParseDuration("1000ns")

	server := &http.Server{
			Addr       : "192.168.0.5:8082",
			IdleTimeout: duration,
		//	Handler    : 
		}

	http.HandleFunc("/", getINFO)

	server.ListenAndServe()
}

func getINFO(w http.ResponseWriter, r *http.Request) {
    qurl := r.URL.String()
    fmt.Fprintf(w, r.Method + " " + qurl)
}
