package server

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func GetHTML(w http.ResponseWriter, r *http.Request, path string, data map[string]any) {
	tmpl, err := template.ParseFiles(TemplateFilesPath + path + ".html")

	if err != nil {
		fmt.Println("error loading template: ", err.Error())
		http.NotFound(w, r)
		return
	}

	if data == nil {
		data = make(map[string]any)
	}

	// generate unique token for form
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	data["Token"] = token

	tmpl.Execute(w, data)
}
