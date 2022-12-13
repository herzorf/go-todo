package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"strings"
)

//var DB *sql.DB

func main() {
	viper.AddConfigPath("./")
	viper.SetConfigName("mysql")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取mysql配置错误", err)
	}

	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	ip := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	//dbName := viper.GetString("mysql.go-todo")

	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", "?charset=utf8"}, "")
	fmt.Println(path)
	DB, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println("open错误", err)
	} else {
		fmt.Println("open成功", DB)
	}

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	rows, err := DB.Query("show databases;")
	//循环读取结果
	for rows.Next() {

		var result []string
		fmt.Println(rows.Scan(result))
	}
}
