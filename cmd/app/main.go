package main

import (
	"github.com/mrkucher83/avito-shop/internal/routes"
	"github.com/mrkucher83/avito-shop/pkg/logger"
	"syscall"
)

const DefaultPort = "8080"

func main() {
	logger.InitLogger(logger.NewLogrusLogger())

	port := DefaultPort
	if value, ok := syscall.Getenv("AVITO_SHOP_PORT"); ok {
		port = value
	}

	routes.Start(port)
}
