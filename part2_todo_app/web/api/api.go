package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ListenAndServe() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/get/{itemId}", handleGETToDoItem)
	mux.HandleFunc("GET /api/get-all", handleGETAllToDoItems)
	mux.HandleFunc("POST /api/create", handlePOSTCreateToDoItem)
	mux.HandleFunc("PATCH /api/edit", handlePATCHEditToDoItem)

	fmt.Println("Starting api server at http://localhost:8085")
	err := http.ListenAndServe("localhost:8085", mux)
	if err != nil {
		log.Fatalln("there's an error with the server:", err)
	}
}

func encodeJson(writer http.ResponseWriter, data any) {
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		fmt.Println("error encoding json:", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func decodeJSONBody(writer http.ResponseWriter, request *http.Request, decodeTarget any) error {
	ct := request.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	request.Body = http.MaxBytesReader(writer, request.Body, 1048576)

	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&decodeTarget)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}

// sample decoder credit Alex Edwards blog
