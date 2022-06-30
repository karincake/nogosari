package nogosari

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	hj "github.com/karincake/nogosari/httpjson"
	t "github.com/karincake/nogosari/types"
)

type httpConf struct {
	Host string
	Port int
}

var wg sync.WaitGroup

func (a *app) initHttp(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", a.HttpConf.Host, a.HttpConf.Port),
		Handler:      a.recoverPanic(a.rateLimit(router)),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		Logger.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		Logger.PrintInfo("completing background tasks", map[string]string{
			"addr": srv.Addr,
		})

		wg.Wait()
		shutdownError <- nil
	}()

	Logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
	})
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		Logger.PrintFatal(err, nil)
	}

	err = <-shutdownError
	if err != nil {
		Logger.PrintFatal(err, nil)
	}

	Logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})
}

// bonus health check
func (a *app) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := t.II{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.Env,
			"version":     a.Version,
		},
	}
	err := hj.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
