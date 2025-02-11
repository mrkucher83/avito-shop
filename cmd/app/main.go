package main

import (
	"github.com/mrkucher83/avito-shop/internal/routes"
	"github.com/mrkucher83/avito-shop/pkg/helpers/pg"
	"github.com/mrkucher83/avito-shop/pkg/logger"
	"syscall"
)

const DefaultPort = "8080"

func main() {
	logger.InitLogger(logger.NewLogrusLogger())

	port := DefaultPort
	if value, ok := syscall.Getenv("SERVER_PORT"); ok {
		port = value
	}

	repo, err := pg.NewDbInstance()
	if err != nil {
		logger.Fatal("failed connecting to DB: %w", err)
	}
	defer repo.Close()

	//routes.Start(port)
	srv := new(routes.Router)
	srv.Start(port)
}
