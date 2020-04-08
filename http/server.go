package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// ListenAndServe starts a new web server of a provided addr
func ListenAndServe(addr string) {
	srv := &http.Server{
		Handler: router(),
		Addr: addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15* time.Second,
		IdleTimeout: 10 * time.Second,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatalf("failed to start http server: %v", err)
		}
	}()

	<-interrupt
	log.Println("Got interrupt signal, going to gracefully shutdown the server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Failed to gracefully shutdwon the server; %v", err)
	}
	log.Println("Server gracefully shutdown")
}
