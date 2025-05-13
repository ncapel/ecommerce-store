package main

import (
	"net/http"

	"github.com/ncapel/ecommerce-store/config"
	"github.com/ncapel/ecommerce-store/routes"
)

func main() {
	config.ConnectDB()
	mux := routes.InitRoutes()
	http.ListenAndServe(":8080", mux)

}
