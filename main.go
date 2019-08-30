package main


import (
    "net/http"
    "time"
    "fmt")

func main() {
    StartServer()

}

func getINFO(w http.ResponseWriter, r * http.Request) {
    qurl: = r.URL.String()
    fmt.Fprintf(w, r.Method + " " + qurl)
}

func StartServer() {
    duration, _: = time.ParseDuration("1000ns")

    server: = & http.Server {
        Addr: "192.168.0.5:8082",
        IdleTimeout: duration,
        //	Handler    : 
    }

    http.HandleFunc("/", getINFO)

    server.ListenAndServe()
}