package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[97m"

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
	Messages    string
	Text        string
	TextBig     string
	TextInfo    string
	TextSub     string
	TextLines   string
	TextSummary string
}

type NotificationModel struct {
	Data Notification
	Id string
}
type QueryResponse struct {
	Records []NotificationModel
}
type MessageData struct {
	Message string
}
type ErrorData struct {
	Error string
}

func LogColor(field string, value any) {
	fmt.Printf("\033[35m>>>\033[0m    %v%v: %v%v%v\n", Cyan, field, Yellow, value, Reset)
}

func LogTime(field string, t long) {
	tt, err := t.Int64()
	if err != nil {
		LogColor(field,"ERROR WHILE PARSING TIME")
		return
	}
	ttt := time.Unix(0, 1000000*int64(tt))
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		LogColor(field,"ERROR WHILE PARSING TIME")
		return
	}
	LogColor(field, ttt.In(loc))
}


func stringLineSlice(lines string) []string {
	return strings.Split(strings.Trim(lines, "\n"), "\n")
}

func (n Notification) Log() {
	fmt.Println(Red, "--------------Printing Notification-----------------", Reset)
	LogColor("PackageName: ", n.PackageName)
	LogTime("PostTime: ", n.PostTime)
	LogTime("SystemTime", n.SystemTime)
	LogColor("IsOngoing: ", n.IsOngoing)
	LogColor("TickerText: ", n.TickerText)
	LogColor("Title: ", n.Title)
	LogColor("TitleBig: ", n.TitleBig)
	LogColor("Text: ", n.Text)
	LogColor("Messages: ", n.Messages)
	LogColor("TextBig: ", n.TextBig)
	LogColor("TextInfo: ", n.TextInfo)
	LogColor("TextSub: ", n.TextSub)
	LogColor("TextLines: ", n.TextLines)
	LogColor("TextSummary: ", n.TextSummary)
	fmt.Println(Red, "----------------------------------------------------", Reset)
}

func sendQueryRequest(sql string) (*http.Response,[]byte,error) {
	client := &http.Client{}
	jsonBody, _ := json.Marshal(map[string]string{
          "statement": sql,
    })
	req, _ := http.NewRequest("POST", XATA_URL + "/sql", bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer " + XATA_API_KEY)
	req.Header.Add("Content-Type", "application/json")

	res,err :=client.Do(req)
	if err != nil {
		return nil,nil ,err
	}
    body, err := io.ReadAll(res.Body)
    defer res.Body.Close()
    if err != nil {
    	return nil,nil,err
    }

    fmt.Println("resonse Status:", res.Status)
    fmt.Println("resonse Body:", string(body))
    return res,body,nil
}
