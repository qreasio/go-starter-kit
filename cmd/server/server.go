package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/qreasio/go-starter-kit/internal/config"
	"github.com/qreasio/go-starter-kit/pkg/log"

	"github.com/pkg/errors"
)

// Start to start the server
func Start(cfg *config.Config, r http.Handler, logger log.Logger) error {
	serverErrors := make(chan error, 1)

	server := http.Server{
		Addr:    fmt.Sprintf("%s", cfg.Server.Port),
		Handler: r,
	}

	// Start the service listening for requests.
	go func() {
		logger.Infof("http server listening on :%s", cfg.Server.Port)
		serverErrors <- server.ListenAndServe()
	}()

	// Shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don'usertransport collect this error.
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	// Blocking main and waiting for shutdown.
	case sig := <-shutdown:
		logger.Infof("start shutdown: %v", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 10)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Infof("graceful shutdown did not complete in %v : %v", 10, err)
			err = server.Close()
			return err
		}

		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
