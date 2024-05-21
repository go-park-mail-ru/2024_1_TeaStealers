package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type respSuccess struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message,omitempty"`
	Payload    interface{} `json:"payload"`
}

type respError struct {
	Message string `json:"message"`
}

type responser interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

// WriteError prints error in json
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := respError{
		Message: message,
	}
	resp, err := errorResponse.MarshalJSON()
	if err != nil {
		return
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}

// WriteResponse writes a JSON response with the specified status code and data.
func WriteResponse(w http.ResponseWriter, statusCode int, response interface{}) error {
	respSuccess := respSuccess{
		StatusCode: statusCode,
		Payload:    response,
	}

	resp, err := respSuccess.MarshalJSON()
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

/*
	func GetIdUserByRequest(r *http.Request) uuid.UUID {
		id := r.Context().Value(middleware.CookieName)
		UUID, ok := id.(uuid.UUID)
		if !ok {
			return uuid.Nil
		}
		return UUID
	}
*/
func StringToTime(layout, value string) (time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TruncSlash(methodName string, count int) (string, error) {
	if count < 0 {
		return "", errors.New("count must be non-negative")
	}

	slashes := strings.Count(methodName, `/`)
	if slashes < count {
		return "", fmt.Errorf("methodName contains %d slashes, but count is %d", slashes, count)
	}

	// Split the methodName string into a slice of strings using `/` as the separator
	parts := strings.Split(methodName, `?`)

	parts = strings.Split(parts[0], `/`)

	// parts = parts[:len(parts)-count-1]

	// Join the remaining elements of the slice back into a string using `/` as the separator
	newMethodName := strings.Join(parts, `/`)
	trSlash := `/`
	newMethodName = newMethodName + trSlash

	return newMethodName, nil
}
