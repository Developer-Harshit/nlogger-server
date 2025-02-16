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
	http.HandleFunc("/error", serveError)

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
	Messages 	string
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
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
func LogColor(field string , value any) {
	fmt.Printf("\033[35m>>>\033[0m    %v%v: %v%v%v\n",Cyan,field,Yellow,value,Reset);
}
func LogTime(field string,t long) {
	tt,err := t.Int64();
	if err != nil {
		log.Println("ERROR WHILE GETTING TIME",err)
	}
	ttt := time.Unix(0, 1000000* int64(tt))
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err == nil {
		LogColor(field,ttt.In(loc))
	}
}

func (n Notification) Log() {
	fmt.Println(Red,"--------------Printing Notification-----------------",Reset)
	LogColor("PackageName: ", n.PackageName);
	LogTime("PostTime: ",n.PostTime);
	LogTime("SystemTime", n.SystemTime)
	LogColor("IsOngoing: ", n.IsOngoing);
	LogColor("TickerText: ", n.TickerText);
	LogColor("Title: ", n.Title);
	LogColor("TitleBig: ", n.TitleBig);
	LogColor("Text: ", n.Text);
	LogColor("Messages: ", n.Messages);
	LogColor("TextBig: ", n.TextBig);
	LogColor("TextInfo: ", n.TextInfo);
	LogColor("TextSub: ", n.TextSub);
	LogColor("TextLines: ", n.TextLines);
	LogColor("TextSummary: ", n.TextSummary);
	fmt.Println(Red,"----------------------------------------------------",Reset)
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
type ErrorMsg struct {
	Message string
}
func serveError(rw http.ResponseWriter, req *http.Request) {
	var n ErrorMsg
	err := json.NewDecoder(req.Body).Decode(&n)
	if err != nil {
		log.Println("ERROR PARSING JSON",err)
		fmt.Fprint(rw, "hi from sender")
		return;
	}
	fmt.Println(Red,"ERROR MESSAGE_> ",n,Reset);
	fmt.Fprint(rw, "hi from sender")
}
func serveHello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, page)
}
