package database

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"

	"golang.org/x/crypto/openpgp/errors"
)

type MyWebDB struct {
}

func (mw *MyWebDB) generateToken() string {
	crun := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crun, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	return token
}

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return false
}

func (mw *MyWebDB) CreateToken() (string, error) {
	t := mw.generateToken()
	stmt, err := db.Prepare("exec AddToken ?, ?")

	if CheckError(err) {
		return "", err
	}

	res, err := stmt.Exec(t, 1)

	if CheckError(err) {
		return "", err
	}

	nums, err := res.RowsAffected()

	if CheckError(err) {
		return "", err
	}

	fmt.Println(nums, "token created")

	return t, nil
}

func (mw *MyWebDB) RemoveToken(token string) error {
	stmt, err := db.Prepare("exec RemoveToken ?, ?")

	if CheckError(err) {
		return err
	}

	res, err := stmt.Exec(token, 1)

	if CheckError(err) {
		return err
	}

	nums, err := res.RowsAffected()

	if CheckError(err) {
		return err
	}

	fmt.Println(nums, " token removed")

	return nil
}

func (mw *MyWebDB) GetToken(t string, tabid int) (DBToken, error) {
	q := fmt.Sprintf("select * from CheckToken('%s', %v)", t, tabid)
	rows, err := db.Query(q)

	if CheckError(err) {
		return DBToken{}, err
	}

	for rows.Next() {
		var tmpt string
		var tabid int
		var date string

		err = rows.Scan(&tmpt, &tabid, &date)

		if CheckError(err) {
			return DBToken{}, err
		}

		ttime, terr := time.Parse(time.RFC3339, date)

		CheckError(terr)

		return DBToken{tmpt, tabid, ttime}, nil
	}

	fmt.Println("cant find ur token")
	return DBToken{}, errors.ErrUnknownIssuer
}

func (mw *MyWebDB) CreateAccount(acc DBAccount) (DBAccount, error) {
	acc.Token = mw.generateToken()
	// stmt, err := db.Prepare("exec AddAccount '?', '?', '?', '?'")

	// if CheckError(err) {
	// 	return DBAccount{}, err
	// }

	res, err := db.Exec("exec AddAccount ?, ?, ?, ?", acc.Token, acc.Fname, acc.Uname, acc.Passw)

	if CheckError(err) {
		return DBAccount{}, err
	}

	nums, err := res.RowsAffected()

	if CheckError(err) {
		return DBAccount{}, err
	}

	fmt.Println(nums, " account has been created")
	return acc, nil
}

func (mw *MyWebDB) RemoveAccount(token string) error {
	stmt, err := db.Prepare("exec RemoveAccount ?")

	if CheckError(err) {
		return err
	}

	res, err := stmt.Exec(token)

	if CheckError(err) {
		return err
	}

	nums, err := res.RowsAffected()

	if CheckError(err) {
		return err
	}

	fmt.Println(nums, " account removed")

	return nil
}

func (mw *MyWebDB) GetAccount(uname string, passw string) (DBAccount, error) {
	q := fmt.Sprintf("select * from CheckAccount('%s', '%s')", uname, passw)
	rows, err := db.Query(q)

	if CheckError(err) {
		return DBAccount{}, err
	}

	for rows.Next() {
		var id int
		var token string
		var fname string
		var uname string
		var passw string
		var cdate string

		err = rows.Scan(&id, &token, &fname, &uname, &passw, &cdate)

		if CheckError(err) {
			return DBAccount{}, err
		}

		return DBAccount{token, fname, uname, passw}, nil
	}

	fmt.Println("cant find ur account")
	return DBAccount{}, errors.ErrUnknownIssuer
}

func (mw *MyWebDB) CheckAccountUname(uname string) (bool, error) {
	q := fmt.Sprintf("select * from Accounts where uname = '%s'", uname)
	rows, err := db.Query(q)

	if CheckError(err) {
		return true, err
	}

	if rows.Next() {
		fmt.Println("found username")
		return true, nil
	}

	fmt.Println("cant find ur account")
	return false, nil
}

func (mw *MyWebDB) GetAccountByToken(token string) (DBAccount, error) {
	tt, err := mw.GetToken(token, 2)

	if CheckError(err) {
		return DBAccount{}, err
	}

	if tt.Tabid != 2 {
		return DBAccount{}, errors.ErrKeyIncorrect
	}

	q := fmt.Sprintf("select * from CheckTkAccount('%s')", token)
	rows, err := db.Query(q)

	if CheckError(err) {
		return DBAccount{}, err
	}

	for rows.Next() {
		var id int
		var tmpt string
		var fname string
		var uname string
		var passw string
		var date string

		err = rows.Scan(&id, &tmpt, &fname, &uname, &passw, &date)

		if CheckError(err) {
			return DBAccount{}, err
		}

		_, terr := time.Parse(time.RFC3339, date)

		if CheckError(terr) {
			return DBAccount{}, terr
		}

		return DBAccount{tmpt, fname, uname, passw}, nil
	}

	fmt.Println("cannot find ur account")
	return DBAccount{}, errors.ErrKeyIncorrect
}

func (mw *MyWebDB) UpdateAccount(old DBAccount, new DBAccount) error {
	odb, err := mw.GetAccountByToken(old.Token)

	if CheckError(err) {
		return err
	}

	if odb.Uname != old.Uname || odb.Passw != old.Passw {
		return errors.ErrKeyIncorrect
	}

	stmt, err := db.Prepare("exec UpdateAccount ?, ?, ?, ?")

	if CheckError(err) {
		return err
	}

	res, err := stmt.Exec(new.Token, new.Fname, new.Uname, new.Passw)

	if CheckError(err) {
		return err
	}

	nums, err := res.RowsAffected()

	if CheckError(err) {
		return err
	}

	fmt.Println(nums, " account updated")

	return nil
}
