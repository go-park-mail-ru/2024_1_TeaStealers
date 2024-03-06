package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// WriteError writes an error response with the specified status code and message.
//func WriteError(w http.ResponseWriter, statusCode int, message string) {
//	w.WriteHeader(statusCode)
//	fmt.Fprintln(w, message)

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	errorResponse := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	resp, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// WriteResponse writes a JSON response with the specified status code and data.
func WriteResponse(w http.ResponseWriter, statusCode int, response interface{}) error {
	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)

	return nil
}

// ReadRequestData reads and parses the request body into the provided structure.
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
