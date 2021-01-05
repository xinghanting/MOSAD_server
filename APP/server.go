package app

import (
	"log"
	"net/http"
)

//RunServer 开始运行服务
func RunServer() {
	//首先初始化和服务器相连接的服务
	MyService := NewService{}
	err := MyService.InitNewService()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Get Data!")

	//启动服务器
	log.Println("Server started at localhost:9090")
	router := NewRouter(&MyService)
	log.Fatal(http.ListenAndServe("127.0.0.1:9090", router))
}
