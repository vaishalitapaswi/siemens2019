package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func loadPage(title string) (*[]byte, error) {
	filename := "static/" + title + ".html"
	fmt.Println(filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	body, _ := loadPage(title)
	fmt.Fprintln(w, string(*body))
}
func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
