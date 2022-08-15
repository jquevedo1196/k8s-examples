package main

import (
    "log"
    "net/http"
	"os"
	"time"
	"encoding/json"
)

type server struct{}

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	resp := HandsOn{
		Time: time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
		

	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
	}

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResp)
}

func main() {
    s := &server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":9090", nil))
}