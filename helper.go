package nogosari

import (
	// "fmt"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// to simplify casting
type I map[string]interface{} // similar to gin.H
type S map[string]string

// others
func StructToMap(in interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Struct {
		return nil
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		out[typ.Field(i).Name] = v.Field(i).Interface()
	}
	return out
}

func writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	return nil
}
