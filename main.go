package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

func LogTime(title string,t long) {
	tt,err := t.Int64();
	if err != nil {
		log.Println("ERROR WHILE GETTING TIME",err)
	}
	ttt := time.Unix(0, 1000000* int64(tt))
	fmt.Println(title,ttt)
}
func (n Notification) Log() {
	fmt.Println("--------------Printing Notification-----------------")
	fmt.Println("PackageName: ", n.PackageName);
	LogTime("PostTime: ",n.PostTime);
	LogTime("SystemTime", n.SystemTime)
	fmt.Println("IsOngoing: ", n.IsOngoing);
	fmt.Println("TickerText: ", n.TickerText);
	fmt.Println("Title: ", n.Title);
	fmt.Println("TitleBig: ", n.TitleBig);
	fmt.Println("Text: ", n.Text);
	fmt.Println("TextBig: ", n.TextBig);
	fmt.Println("TextInfo: ", n.TextInfo);
	fmt.Println("TextSub: ", n.TextSub);
	fmt.Println("TextLines: ", n.TextLines);
	fmt.Println("TextSummary: ", n.TextSummary);
	fmt.Println("----------------------------------------------------")
}

func serveSend(rw http.ResponseWriter, req *http.Request) {
	var n Notification
	err := json.NewDecoder(req.Body).Decode(&n)
	if err != nil {
		log.Println("ERROR PARSING JSON",err)
	}
	n.Log();
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

func serveHello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, page)
}
