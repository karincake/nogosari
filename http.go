package nogosari

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (a *app) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := I{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.env,
			"version":     a.version,
		},
	}
	err := writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) initHttp(routerIn *httprouter.Router) {
	routerIn.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthcheckHandler)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", a.httpConf.port),
		// Handler:      routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		loggerX.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Call Shutdown() on the server like before, but now we only send on the
		// shutdownError channel if it returns an error.
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		// Log a message to say that we're waiting for any background goroutines to
		// complete their tasks.
		loggerX.PrintInfo("completing background tasks", map[string]string{
			"addr": srv.Addr,
		})

		// Call Wait() to block until our WaitGroup counter is zero --- essentially
		// blocking until the background goroutines have finished. Then we return nil on
		// the shutdownError channel, to indicate that the shutdown completed without
		// any issues.

		wgX.Wait()
		shutdownError <- nil
	}()

	loggerX.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
	})

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		loggerX.PrintFatal(err, nil)
	}

	err = <-shutdownError
	if err != nil {
		loggerX.PrintFatal(err, nil)
	}

	loggerX.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})
}
