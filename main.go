package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Hello world",
			Desc:    "Hello world Article",
			Content: "This the content of the first article ever"},
	}

	fmt.Println("All articles endpoint")
	json.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	met := r.Method // you can find more info in the r attributes
	fmt.Fprintf(w, "Home page created by Alaref %s", met)
	/*
		The w is the response we send
		and the Fprintf append the text or the content we want to the w variable
	*/
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {

	io.Copy(os.Stdout, r.Body)

	fmt.Fprint(w, "The post endpoint workd nicely")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// maybe write a program that wait for the user to click a button then close the server
func main() {
	handleRequests()
}
