package mysqlConnect

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func ConnectMysql() *sql.DB {
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
	dbName := viper.GetString("mysql.dbname")

	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println("数据库链接path：", path)

	DB, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println("数据库open错误", err)
	} else {
		fmt.Println("数据库open成功")
	}

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	return DB
}
