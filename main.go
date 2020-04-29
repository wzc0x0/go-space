package main

import (
	"fmt"

	"go-space/config"
	"go-space/routers"
	"go-space/servers"
)

func main() {
	if err := config.InitConfig("config/config.yaml"); err != nil {
		fmt.Println(err)
		return
	}
	db, err := servers.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	r := routers.InitRouter()
	r.Run(config.GetConfig().Server.Port)
}
