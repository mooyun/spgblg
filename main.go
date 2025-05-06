package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	_ "path"
	"text/template"
)

type IndexDate struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var indexDate IndexDate
	indexDate.Title = "spgblg"
	indexDate.Desc = "刚刚开始"

	t := template.New("index.html")
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexDate)

	jsonStr, _ := json.Marshal(indexDate)
	w.Write([]byte(jsonStr))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
