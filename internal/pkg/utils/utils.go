package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, message)
}

func WriteResponse(w http.ResponseWriter, statusCode int, response interface{}) error {
	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	w.Write(resp)

	return nil
}

func ReadRequestData(r *http.Request, request interface{}) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(data, &request); err != nil {
		return err
	}
	return nil
}
