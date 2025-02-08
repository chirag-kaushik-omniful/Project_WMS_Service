package main

import (
	"fmt"
	"wms/config"
	"wms/routes"
	"wms/utils/dbconn"
	"wms/utils/redisclient"

	"github.com/omniful/go_commons/db/sql/postgres"
	server "github.com/omniful/go_commons/http"
)

func main() {
	srv := server.InitializeServer(":8080", 0, 0, 0)

	routes.GetRouter(srv)

	redisclient.Connect(config.Redis_Config)

	dbconn.Connect(*config.Postgres_Config, &[]postgres.DBConfig{*config.Postgres_Config})

	err := srv.StartServer("wms")
	if err != nil {
		fmt.Println("Server error:", err)
	}

}
