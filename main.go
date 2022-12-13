package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/herzorf/go-todo/mysqlConnect"
	"github.com/herzorf/go-todo/type/mysql"
	"net/http"
)

var DB *sql.DB

func main() {
	DB = mysqlConnect.ConnectMysql()

	route := gin.Default()
	route.POST("/api/v1/getTodo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"xxx": "xxxx"})
	})
	route.GET("/home", func(c *gin.Context) {

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
		marshal, _ := json.Marshal(users)

		c.String(http.StatusOK, string(marshal))
	})

	err := route.Run()
	if err != nil {
		panic("gin 启动失败")
	}
}
