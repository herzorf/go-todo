package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/herzorf/go-todo/mysqlConnect"
	"github.com/herzorf/go-todo/type/mysql"
)

func main() {
	DB := mysqlConnect.ConnectMysql()
	rows, _ := DB.Query("select * from  goTodoTest")
	//循环读取结果
	var users []mysql.User
	for rows.Next() {
		//将每一行的结果都赋值到一个user对象中
		var user mysql.User
		err := rows.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	fmt.Println(users)
}
