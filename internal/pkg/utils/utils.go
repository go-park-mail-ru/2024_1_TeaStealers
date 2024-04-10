package utils

import (
	"2024_1_TeaStealers/internal/pkg/middleware"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/satori/uuid"
)

// WriteError prints error in json
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	resp, err := json.Marshal(errorResponse)
	if err != nil {
		return
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}

// WriteResponse writes a JSON response with the specified status code and data.
func WriteResponse(w http.ResponseWriter, statusCode int, response interface{}) error {
	respSuccess := struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message,omitempty"`
		Payload    interface{} `json:"payload"`
	}{
		StatusCode: statusCode,
		Payload:    response,
	}
	resp, err := json.Marshal(respSuccess)
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

// GenerateHashString generate hash string
func GenerateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetIdUserByRequest(r *http.Request) uuid.UUID {
	id := r.Context().Value(middleware.CookieName)
	UUID, ok := id.(uuid.UUID)
	if !ok {
		return uuid.Nil
	}
	return UUID
}
