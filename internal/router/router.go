package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/side-hub-be/side-hub-be/internal/api"
	"github.com/side-hub-be/side-hub-be/internal/log"
)

// Route
func Route(e *echo.Echo) {
	e.GET("/health", api.GetHealth)
}

func Run(e *echo.Echo, address string) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			log.Fatalf("%+v", errors.Wrap(err, "invalid server"))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	wg.Wait()
}
