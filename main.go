package main

import (
	"github.com/GabrielLoureiroGomes/basket-collection/api"
	"github.com/GabrielLoureiroGomes/basket-collection/config"
)

func main() {
	config.InitConfig()
	server := api.NewServer()
	server.StartServer()
}
