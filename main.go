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
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
func LogColor(field string , value any) string {
	return fmt.Sprintf("\033[35m>>>\033[0m    %v%v: %v%v%v\n",Cyan,field,Yellow,value,Reset);
}
func LogTime(field string,t long) string {
	tt,err := t.Int64();
	if err != nil {
		log.Println("ERROR WHILE GETTING TIME",err)
	}
	ttt := time.Unix(0, 1000000* int64(tt))
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err == nil {
		return LogColor(field,ttt.In(loc))
	}
        return ""

}

func (n Notification) Log() {
	str := fmt.Sprintln("\n",Red,"--------------Printing Notification-----------------",Reset)
	str += LogColor("PackageName: ", n.PackageName);
	str += LogTime("PostTime: ",n.PostTime);
	str += LogTime("SystemTime", n.SystemTime)
	str += LogColor("IsOngoing: ", n.IsOngoing);
	str += LogColor("TickerText: ", n.TickerText);
	str += LogColor("Title: ", n.Title);
	str += LogColor("TitleBig: ", n.TitleBig);
	str += LogColor("Text: ", n.Text);
	str += LogColor("TextBig: ", n.TextBig);
	str += LogColor("TextInfo: ", n.TextInfo);
	str += LogColor("TextSub: ", n.TextSub);
	str += LogColor("TextLines: ", n.TextLines);
	str += LogColor("TextSummary: ", n.TextSummary);
	str += fmt.Sprintln(Red,"----------------------------------------------------",Reset)
	fmt.Println(str)
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
