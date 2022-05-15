package server

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/BabaYaga1279/GoLang/server/database"
)

var (
	db = database.GetDB()
)

const (
	CO_reg_token_tabid = 1
	CO_acc_token_tabid = 2

	TP_varerrorlogin    = "ErrorLogin"
	TP_varerrorregister = "ErrorRegister"
	TP_Token            = "Token"
	TP_frmfullname      = "fullname"
	TP_frmusername      = "username"
	TP_frmpassword      = "password"
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
	token, err := db.CreateToken()

	if CheckError(err) {
		http.NotFound(w, r)
		return
	}

	data[TP_Token] = token

	tmpl.Execute(w, data)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostLogin")
	data := make(map[string]any)

	if CheckToken(w, r, r.FormValue(TP_Token), CO_reg_token_tabid) {
		return
	}

	var acc database.DBAccount
	acc.Uname = r.FormValue(TP_frmusername)
	acc.Passw = r.FormValue(TP_frmpassword)

	fmt.Println("checking account: ", acc)
	var aerr error
	acc, aerr = db.GetAccount(acc.Uname, acc.Passw)

	if CheckError(aerr) {
		data[TP_varerrorlogin] = "account or password is not valid"
		GetHTML(w, r, r.URL.Path, data)
		return
	}

	// pass acc into to data ...
	// ...
	//

	GetHTML(w, r, r.URL.Path, data)
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostRegister")
	data := make(map[string]any)

	if CheckToken(w, r, r.FormValue(TP_Token), CO_reg_token_tabid) {
		return
	}

	var acc database.DBAccount
	acc.Fname = r.FormValue(TP_frmfullname)
	acc.Uname = r.FormValue(TP_frmusername)
	acc.Passw = r.FormValue(TP_frmpassword)

	est, err := db.CheckAccountUname(acc.Uname)

	if CheckError(err) {
		http.NotFound(w, r)
		return
	}

	if est {
		data[TP_varerrorregister] = "username already exists"
		GetHTML(w, r, r.URL.Path, data)
		return
	}

	if acc.Fname == "" || acc.Passw == "" || acc.Uname == "" {
		data[TP_varerrorregister] = "all fields must be filled"
		GetHTML(w, r, r.URL.Path, data)
		return
	}

	acc, err = db.CreateAccount(acc)

	if CheckError(err) {
		data[TP_varerrorlogin] = "cannot create account"
		GetHTML(w, r, r.URL.Path, data)
		return
	}

	// pass acc into to data ...
	// ...
	//

	GetHTML(w, r, r.URL.Path, data)
}

func CheckToken(w http.ResponseWriter, r *http.Request, t string, tabid int) bool {
	err := VerifyToken(r.FormValue(TP_Token), CO_reg_token_tabid)

	if err != "" {
		data := make(map[string]any)
		data[TP_varerrorlogin] = err
		GetHTML(w, r, r.URL.Path, data)
		return true
	}

	return false
}

func VerifyToken(t string, tabid int) string {
	var token database.DBToken
	token.Token = t
	token.Tabid = tabid

	err := ""

	fmt.Println("checking token: ", token)
	// check token
	if token.Token == "" {
		err = "invalid token"
	} else {
		var terr error
		token, terr = db.GetToken(token.Token, token.Tabid)
		if CheckError(terr) {
			// token not exists in database
			err = "invalid token"
		} else {
			// expired token
			if time.Now().Sub(token.Cdate).Minutes() >= 5 {
				err = "expired token, pls reload page"
			}

			db.RemoveToken(token.Token)
		}
	}

	return err
}

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return false
}
