package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/herzorf/go-todo/mysqlConnect"
	"github.com/herzorf/go-todo/type/mysql"
	"github.com/herzorf/go-todo/type/request"
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

		c.JSON(http.StatusOK, gin.H{"result": "添加成功"})
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
	route.POST("/api/v1/toggleTodo", func(c *gin.Context) {
		var id request.ToggleTodo
		err := c.ShouldBind(&id)
		if err != nil {
			panic(err)
		}
		selectQuery := "SELECT * FROM goTodoTest WHERE id=?"
		result, err := DB.Query(selectQuery, id.Id)
		var todo mysql.Todo
		for result.Next() {
			err := result.Scan(&todo.Id, &todo.Name, &todo.Done)
			if err != nil {
				panic(err)
			}
			fmt.Println(todo)
		}
		query := "UPDATE goTodoTest set done= ? WHERE id=?"
		_, err = DB.Exec(query, !todo.Done, todo.Id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "请求成功"})

	})
	err := route.Run()
	if err != nil {
		panic("gin 启动失败")
	}
}
