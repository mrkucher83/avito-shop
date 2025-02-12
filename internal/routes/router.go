package routes

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mrkucher83/avito-shop/internal/godb"
	"github.com/mrkucher83/avito-shop/internal/handlers"
	"github.com/mrkucher83/avito-shop/internal/middlewares"
	"github.com/mrkucher83/avito-shop/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(port string, repo *godb.Instance) {
	handler := handlers.NewRepo(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Welcome to Avito Shop!")); err != nil {
			logger.Warn("failed to write response: ", err)
		}
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/auth", handler.SignUp)

		r.With(middlewares.AuthMiddleware).Group(func(r chi.Router) {
			r.Get("/info", handler.GetEmployeeInfo)
			r.Post("/sendCoin", handler.SendCoin)
			r.Get("/buy/{item:[a-zA-Z-]+}", handler.BuyItem)
		})
	})

	logger.Info("starting server on %s", port)
	server := &http.Server{Addr: ":" + port, Handler: r}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})
	go gracefulShutdown(server, stop, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("failed to start service: %v", err)
	}

	<-done
}

func gracefulShutdown(srv *http.Server, stop chan os.Signal, done chan struct{}) {
	<-stop
	logger.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown failed: %v", err)
	} else {
		logger.Info("server gracefully stopped")
	}
	close(done)
}
