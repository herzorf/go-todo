package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./")
	viper.SetConfigName("mysql")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置错误", err)
	}

	username := viper.Get("mysql.username")
	password := viper.Get("mysql.password")
	fmt.Println(username, password)
}
