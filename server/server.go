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
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		fmt.Println("token: ", r.Form["Token"])

		err := ""
		if len(r.Form["username"]) != 1 || len(r.Form["password"]) != 1 {
			err = "Wrong format of username or password."
		} else if db.Accounts[r.Form["username"][0]] == "" {
			err = "Username: " + r.Form["username"][0] + " doesn't exist in database."
		} else if db.Accounts[r.Form["username"][0]] != r.Form["password"][0] {
			err = "Wrong Password."
		} else {
			http.Redirect(w, r, "/chatroom", http.StatusSeeOther)
			return
		}

		data := make(map[string]any)
		data["ErrorLogin"] = err
		GetHTML(w, r, r.URL.Path, data)

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

var db database

func StartServer() {
	db = NewDatabase()

	db.Accounts["admin"] = "admin"

	//http.HandleFunc("/", SayHello)           // set router
	mux := &UrlMux{}
	err := http.ListenAndServe(":6969", mux) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
