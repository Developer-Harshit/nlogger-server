package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func assignHandlers() {
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", serveHome)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/msg", msgHandler)
}

func serveHome(res http.ResponseWriter, req *http.Request) {
	sqlQuery := fmt.Sprintf("SELECT * FROM \"%s\"",XATA_TABLE)

	_,body,_ := sendQueryRequest(sqlQuery)
	var n QueryResponse
	json.Unmarshal(body, &n)
	sort.Slice(n.Records, func(i, j int) bool {
		return n.Records[i].Data.SystemTime > n.Records[j].Data.PostTime
	})
	compo := notificationComponent(n)
	//fmt.Println(n.Records)
	compo.Render(context.Background(), res)
}
func insertHandler(res http.ResponseWriter, req *http.Request) {
	jsonBody, _ := io.ReadAll(req.Body)
	sqlQuery := fmt.Sprintf("INSERT INTO \"%s\" (data) VALUES ('%s')",XATA_TABLE,jsonBody)
	_,_,err := sendQueryRequest(sqlQuery)
	if err != nil {
		fmt.Println(Red,err,Reset)
	}
}

func deleteHandler(res http.ResponseWriter, req *http.Request) {

}

func msgHandler(res http.ResponseWriter, req *http.Request) {
	var n MessageData
	err := json.NewDecoder(req.Body).Decode(&n)
	if err != nil {
		fmt.Println("ERROR PARSING JSON", err)
	}
	fmt.Println(Blue, "INFO MESSAGE_> ", n, Reset)
	fmt.Fprint(res, "WTF BRO")
}

func errorHandler(res http.ResponseWriter, req *http.Request) {
	var n ErrorData
	err := json.NewDecoder(req.Body).Decode(&n)
	if err != nil {
		fmt.Println("ERROR PARSING JSON", err)
	}
	fmt.Println(Red, "ERROR MESSAGE_> ", n, Reset)
	fmt.Fprint(res, "WTF BRO")
}
