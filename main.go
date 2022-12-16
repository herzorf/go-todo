package main

import (
	"context"
	"database/sql"
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
	var todo mysql.Todo
	route := gin.Default()
	route.POST("/api/v1/addTodo", func(c *gin.Context) {
		err := c.ShouldBind(&todo)
		if err != nil {
			panic(err)
		}
		query := "INSERT INTO `goTodoTest` (`name`, `done`) VALUES (?, ?)"
		insertResult, err := DB.ExecContext(context.Background(), query, todo.Name, todo.Done)
		if err != nil {
			fmt.Println("数据库插入错误", err)
		}
		fmt.Println(insertResult)
		id, err := insertResult.LastInsertId()
		fmt.Println("inserted id:", id)

		c.JSON(http.StatusOK, gin.H{"result": "请求成功"})
	})
	route.GET("/api/v1/getTodos", func(c *gin.Context) {
		query := "SELECT * from  goTodoTest"
		result, err := DB.Query(query)
		if err != nil {
			fmt.Println("数据库查询错误", err)
		}
		var todos []mysql.Todo

		for result.Next() {
			var todo mysql.Todo
			err := result.Scan(&todo.Id, &todo.Name, &todo.Done)
			if err != nil {
				fmt.Println("查询结果扫描错误", err)
			}
			todos = append(todos, todo)
		}
		c.JSON(http.StatusOK, gin.H{"data": todos})
	})

	err := route.Run()
	if err != nil {
		panic("gin 启动失败")
	}
}
