package server

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

// rounter
type UrlMux struct{}

func MuxGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.URL.Path == "/hello" {
		SayHello(w, r)
		return
	}

	if r.URL.Path == "/login" || r.URL.Path == "/register" {
		GetHTML(w, r, r.URL.Path, nil)
		return
	}

	// 404
	http.NotFound(w, r)
}

func MuxPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.URL.Path == "/login" {
		PostLogin(w, r)
		return
	}

	if r.URL.Path == "/register" {
		PostRegister(w, r)
		return
	}
	// 404
	http.NotFound(w, r)
}

func (mux *UrlMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		MuxGet(w, r)
		return
	case "POST":
		MuxPost(w, r)
		return
	}

	// 404
	http.NotFound(w, r)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k, reflect.TypeOf(k))
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello user!")
}

func StartServer() {

	//http.HandleFunc("/", SayHello)           // set router
	mux := &UrlMux{}
	err := http.ListenAndServe(":6969", mux) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
