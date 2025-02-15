package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "embed"
)

//go:embed page.html
var page string

func main() {
	http.HandleFunc("/", serveHello)
	http.HandleFunc("/test", test)
		http.HandleFunc("/send", serveSend)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server:", err)
	}
}

type long = json.Number
type Notification struct {
	PackageName string
	PostTime    long
	SystemTime  long
	IsOngoing   bool
	When        long
	TickerText  string
	Title       string
	TitleBig    string
	Text        string
	TextBig     string
	TextInfo    string
	TextSub     string
	TextLines   string
	TextSummary string
}
type test_struct struct {
	Name  string
	Email string
	Phone string
}

func serveSend(rw http.ResponseWriter, req *http.Request) {
	var n Notification
	err := json.NewDecoder(req.Body).Decode(&n)
	if err != nil {
		panic(err)
	}
	log.Printf("%#v", n)
	fmt.Fprint(rw, "hi from sender")
}

func test(rw http.ResponseWriter, req *http.Request) {
	var t test_struct
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Printf("%#v", t)
}

func serveHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}
