package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func loadPage(title string) (*[]byte, error) {
	filename := "static/" + title
	log.Println(filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error Loading File" + filename)
		return nil, err
	}
	return &body, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	body, _ := loadPage(title)
	fmt.Fprintln(w, string(*body))
}
func savehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("empno"))
	e := emp{}
	e.Empno, _ = strconv.Atoi(r.FormValue("empno"))
	e.Ename = r.FormValue("ename")
	e.Salary, _ = strconv.Atoi(r.FormValue("salary"))
	fmt.Println(e)
	body, _ := loadPage("index.html")
	fmt.Fprintln(w, string(*body))
}
func main() {
	http.HandleFunc("/insertdata", savehandler)
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
