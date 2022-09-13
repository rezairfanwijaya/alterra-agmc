package main

import (
	"encoding/json"
	"net/http"
)

// struct untuk menampung data article
type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// handler
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte("Hallo"))
}

// handler
func article(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-type", "application/json")

	// add data
	news := []Article{
		{
			Id:      1,
			Title:   "Article 1",
			Content: "Article 1 has been post",
		}, {
			Id:      2,
			Title:   "Article 2",
			Content: "Article 2 has been post",
		},
	}

	// cek method
	if r.Method == "GET" {
		resp, err := json.Marshal(news)
		if err != nil {
			http.Error(w, "Failed in Marshal", http.StatusInternalServerError)
			return
		}

		w.Write(resp)
		return
	}

	// error jika method bukan get
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func main() {
	// ser route
	http.HandleFunc("/", index)
	http.HandleFunc("/articles", article)

	// run server
	http.ListenAndServe(":8080", nil)
}
