package main

import (
	"enigma-news/config"
	"enigma-news/main/master"
)

func main() {
	db, _ := config.Connection()    //ok
	router := config.CreateRouter() //ok
	master.Init(router, db)
	config.RunServer(router)
}
