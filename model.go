package jsonerr

import (
	"encoding/json"
	"log"
	"github.com/pkg/errors"
	"fmt"
	"reflect"
)

type JSONError struct {
	Code      int    `json:"code"`
	ErrorType string `json:"errorType"`
	FuncName  string `json:"funcName"`
	Message   string `json:"message"`
}

func newJsonError(code int, sourceError error, msg, funcName string, doLog bool) error {
	if sourceError == nil {
		return nil
	}

	errorType := reflect.TypeOf(sourceError).String()
	if msg == "" && sourceError != nil {
		msg = sourceError.Error()
	}

	jsonErr := &JSONError{code, errorType, funcName, msg}

	serialized, err := json.Marshal(jsonErr)
	if err != nil {
		log.Println("JSONError serialization error:", err)
		return errors.New(fmt.Sprintf(`{"code": %d, "error_type": "%s", "message": "%s", "func_name": "%s"}`,
			code, errorType, msg, funcName))
	}

	text := string(serialized)
	if doLog {
		log.Println(text)
	}

	return errors.New(text)
}
