package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
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

func ReplaceURLPart(url string, replacement string, index int) string {
	parts := strings.Split(url, "/")

	if index >= 0 && index < len(parts) {
		parts[index] = replacement
	}

	replacedURL := strings.Join(parts, "/")
	return replacedURL
}

func GetValueFromInterface(i interface{}, fieldName string) (interface{}, error) {
	value := reflect.ValueOf(i)

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, fmt.Errorf("nil pointer value")
		}
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		fieldValue := value.FieldByName(fieldName)
		if fieldValue.IsValid() {
			return fieldValue.Interface(), nil
		} else {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}
	}

	if value.Kind() == reflect.Map && value.Type().Key().Kind() == reflect.String {
		mapKey := reflect.ValueOf(fieldName)
		fieldValue := value.MapIndex(mapKey)
		if fieldValue.IsValid() {
			return fieldValue.Interface(), nil
		} else {
			return nil, fmt.Errorf("key %s not found in map", fieldName)
		}
	}

	return nil, fmt.Errorf("unsupported value type: %s", value.Type().String())
}
