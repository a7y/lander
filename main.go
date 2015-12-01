package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

var csvPath string
var addr string

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "%s <csv-file> <addr>\n", os.Args[0])
		os.Exit(2)
	}

	csvPath = os.Args[1]
	addr = os.Args[2]

	setupCsv(Subscriber{})

	fmt.Println("CSV file located at", csvPath)
	fmt.Println("listening on", addr)

	http.Handle("/subscribe", cors.Default().Handler(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method == "POST" {
					handlePost(w, r)
				} else {
					http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				}
			},
		),
	))
	http.ListenAndServe(addr, nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil || r.Form.Get("email") == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	s := Subscriber{
		Email: r.Form.Get("email"),
		Host:  getIp(r),
		When:  time.Now().String(),
	}

	err = s.Save(csvPath)
	if err != nil {
		fmt.Println(r.URL.Path, "POST", "Couldn't save subscriber: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json, e := json.Marshal(s)
	if e != nil {
		fmt.Println(r.URL.Path, "POST", "Couldn't marshal JSON: ", e)
		http.Error(w, fmt.Sprintf("Internal server error"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func getIp(r *http.Request) string {
	// this is actually not reliable but ¯\_(ツ)_/¯
	if r.Header.Get("X-FORWARDED-FOR") != "" {
		return r.Header.Get("X-FORWARDED-FOR")
	}

	return r.RemoteAddr
}
