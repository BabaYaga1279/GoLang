package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type DBToken struct {
	Token string
	Tabid int
	Cdate time.Time
}

func (t DBToken) String() string {
	return fmt.Sprintf("%v %v %v", t.Token, t.Tabid, t.Cdate)
}

type DBAccount struct {
	Token string
	Fname string
	Uname string
	Passw string
}

func (a DBAccount) String() string {
	return fmt.Sprintf("%v %v \n   %v %v", a.Token, a.Fname, a.Uname, a.Passw)
}

type MyDB interface {
	generateToken() string

	CreateToken() (string, error)
	RemoveToken(token string) error
	GetToken(token string, tabid int) (DBToken, error)

	CreateAccount(acc DBAccount) (DBAccount, error)
	RemoveAccount(token string) error
	GetAccount(uname, passw string) (DBAccount, error)
	CheckAccountUname(uname string) (bool, error)
	GetAccountByToken(token string) (DBAccount, error)
	UpdateAccount(old DBAccount, new DBAccount) error
}

func GetDB() MyDB {
	return &MyWebDB{}
}

var db *sql.DB

var (
	port     = "1433"
	user     = "bqm2709"
	password = "Quangminh270901"
	database = "mywebdb"
	address  = "192.168.1.6"
	protocol = "tcp"
)

func Ping(db *sql.DB) bool {
	err := db.Ping()
	if err != nil {
		fmt.Println("cannot Ping: ", err.Error())
		return false
	}

	fmt.Println("Pinged successful")
	return true
}

func Test() {
	Ping(db)

	rows, err := db.Query("Select * from Tokens")

	if err != nil {
		fmt.Println("cannot query: ", err.Error())
		log.Fatal(err)
	}

	for rows.Next() {
		var token string
		var tabid int
		var date string

		err = rows.Scan(&token, &tabid, &date)

		if err != nil {
			fmt.Println("cannot scan rows: ", err.Error())
			log.Fatal(err)
		}

		fmt.Println(token, tabid, date)
		t, terr := time.Parse(time.RFC3339, date)
		if terr != nil {
			fmt.Println("cannot convert time: ", terr.Error())
		} else {
			fmt.Println(t)
		}
	}

	stmt, err := db.Prepare("select * from CheckToken(?,?)")

	if err != nil {
		fmt.Println("cannot query procedure: ", err.Error())
		log.Fatal(err)
	}

	res, err := stmt.Exec("0000000000000001", 2)

	raf, err := res.RowsAffected()

	fmt.Println(raf)
}

func init() {

	connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)
	var err error
	db, err = sql.Open("mssql", connectionString)

	if err != nil {
		fmt.Println("cannot open database: ", err.Error())
		log.Fatal(err)
	}

	fmt.Println("Opened db connection")

	// defer func() {
	// 	db.Close()
	// 	fmt.Println("Closed db connection")
	// }()

}

func Close() {
	db.Close()
	fmt.Println("Closed db connection")
}
