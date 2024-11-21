package main

import (
	"api_mysql/configs"
)

func main() {

	configs, err := configs.Loadconfig(".")
	if err != nil {
		panic(err)
	}
	println(configs.DBDriver)
}
