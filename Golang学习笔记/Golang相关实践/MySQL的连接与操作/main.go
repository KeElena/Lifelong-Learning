package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	conn *sql.DB
)

type Result struct {
	Sno   string
	Sname string
}

func initDB() {
	var err error
	dsn := "GoLang:passwd@tcp(192.168.2.11:3306)/edt_db"
	conn, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func query() {
	var rst []Result
	rows, err := conn.Query("SELECT Sno,Sname FROM student WHERE CLno=101")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var val Result
		rows.Scan(&val.Sno, &val.Sname)
		rst = append(rst, val)
	}
	fmt.Println(rst)
}

func insert() {

	res, err := conn.Exec(`INSERT into student(Sname,CLno) VALUES("李６","101")`)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ := res.LastInsertId()
	fmt.Println("id", id)
}

func update() {
	Clno := "102"
	Sno := 2019103
	res, err := conn.Exec(`UPDATE student SET CLno=? WHERE Sno=?`, Clno, Sno)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, _ := res.RowsAffected()
	fmt.Println(n)
}

func delete() {
	res, err := conn.Exec("DELETE FROM student WHERE Sno=2019104")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, _ := res.RowsAffected()
	fmt.Println(n)
}

func prepareQuery() {

	var rst []Result
	sql := "SELECT Sno,Sname FROM student WHERE CLno=?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	rows, _ := stmt.Query(101)
	for rows.Next() {
		var val Result
		rows.Scan(&val.Sno, &val.Sname)
		rst = append(rst, val)
	}
	fmt.Println(rst)
}

func main() {
	initDB()
	//query()
	//insert()
	//update()
	//delete()
	prepareQuery()
	conn.Close()
}
