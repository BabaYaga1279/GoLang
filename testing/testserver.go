package testing

import (
	"fmt"
	"log"

	"github.com/BabaYaga1279/GoLang/httpserver"
	"github.com/BabaYaga1279/GoLang/server"
	"github.com/BabaYaga1279/GoLang/server/database"
	_ "github.com/BabaYaga1279/GoLang/server/database"
	"github.com/BabaYaga1279/GoLang/socket"
)

func TestSocketServer() {
	socket.StartServer()
}

func TestSocketClient() {
	socket.StartClient()
}

func TestHTTPServer() {
	httpserver.Test()
}

func TestServer() {
	server.StartServer()
}

func TestDatabase() {
	database.Test()
	database.Close()
}

func TestDatabase2() {
	db := database.GetDB()
	defer database.Close()

	token, err := db.CreateToken()

	if err != nil {
		fmt.Println("cannot create new token: ", err.Error())
		log.Fatal(err)
	}

	fmt.Println(token)

	var tb database.DBToken
	tb, err = db.GetToken(token, 1)

	if err != nil {
		fmt.Println("cannot get token: ", err.Error())
		log.Fatal(err)
	}

	fmt.Println(tb)

	err = db.RemoveToken(tb.Token)
	fmt.Println("token removed")

	admin, _ := db.GetAccount("bqm2709", "Quangminh270901")

	fmt.Println(admin)

	acc := database.DBAccount{
		"",
		"bui minh",
		"garan",
		"password",
	}

	acc, err = db.CreateAccount(acc)

	if err != nil {
		fmt.Println("cannot create account: ", err.Error())
		log.Fatal(err)
	}

	fmt.Println(acc)

	acc2, err := db.GetAccount(acc.Uname, acc.Passw)

	if err != nil {
		fmt.Println("cannot get account: ", err.Error())
		log.Fatal(err)
	}

	fmt.Println(acc2)

	acc2.Fname = "minh bui quang"
	err = db.UpdateAccount(acc, acc2)

	if err != nil {
		fmt.Println("cannot update accoutn: ", err.Error())
		log.Fatal(err)
	}
	fmt.Println("account updated")

	acc2, err = db.GetAccount(acc2.Uname, acc2.Passw)

	if err != nil {
		fmt.Println("cannot get account: ", err.Error())
		log.Fatal(err)
	}

	fmt.Println(acc2)

	err = db.RemoveAccount(acc2.Token)

	if err != nil {
		fmt.Println("cannot remove account: ", err.Error())
		log.Fatal(err)
	}
	fmt.Println("account removed")
	acc2, err = db.GetAccountByToken(acc2.Token)

	if err != nil {
		fmt.Println("cannot get account by token: ", err.Error())
	} else {
		fmt.Println(acc2)
	}

}
