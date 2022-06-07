package nogosari

import (
	"fmt"
	"net/http"

	hj "github.com/karincake/nogosari/httpjson"
)

func (a *app) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := mi{"error": message}
	err := hj.WriteJSON(w, status, env, nil)
	if err != nil {
		a.logError(r, err)
		w.WriteHeader(500)
	}
}

func (a *app) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	a.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (a *app) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	a.errorResponse(w, r, http.StatusNotFound, message)
}

func (a *app) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	a.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (a *app) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (a *app) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	a.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (a *app) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	a.errorResponse(w, r, http.StatusConflict, message)
}

func (a *app) logError(r *http.Request, err error) {
	a.Logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (a *app) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	a.errorResponse(w, r, http.StatusTooManyRequests, message)
}
