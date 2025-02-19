package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	handlers2 "github.com/nglmq/ozon-test/internal/app/http/handlers"
	"github.com/nglmq/ozon-test/internal/app/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartHTTPServer(urlService *service.URLService) {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)

	router.Route("/", func(r chi.Router) {
		r.Post("/", handlers2.HandlePost(urlService))
		r.Get("/{id}", handlers2.HandleGet(urlService))
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
