package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type data struct {
	rid    int
	reward string
}

var slice []data

//golang操作mysql实现curd基础操作
func main() {
	db, error := sql.open("mysql", "root:@tcp(127.0.0.1:3306)/pu_localhost?charset=utf8")
	if error != nil {
		log.fatal(error)
	}
	//select
	rows, error := db.query("select * from rewards")
	for rows.next() {
		var rid int
		var reward string
		error = rows.scan(&rid, &reward)
		a := data{rid, reward}
		slice = append(slice, a)
	}

	fmt.println(slice)
	for _, v := range slice {
		fmt.println(v)
	}
	//update
	update, error := db.prepare("update rewards set reward =? where rid = ?")
	res, _ := update.exec("tiger", 5)
	id, error := res.rowsaffected()
	fmt.println(id)
	//insert
	insert, error := db.prepare("insert rewards set reward = ?")
	insetres, _ := insert.exec("docker")
	fmt.println(insetres)
	insertid, _ := insetres.lastinsertid()
	fmt.println(insertid)
	//delete
	delete, error := db.prepare("delete from rewards where rid = ?")
	deleteres, _ := delete.exec(12)
	fmt.println(deleteres)
}
