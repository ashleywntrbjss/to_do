package api

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

const ServerProtocol = "http://"
const ServerAddress = "localhost:8085"

var activeRepo repo.Repo

func ListenAndServe(ctx context.Context, repo repo.Repo) {
	mux := http.NewServeMux()

	activeRepo = repo

	if activeRepo == nil {
		panic("no active repo")
	}

	mux.HandleFunc("GET /api/get/{itemId}", handleGETToDoItem)
	mux.HandleFunc("GET /api/get-all", handleGETAllToDoItems)
	mux.HandleFunc("POST /api/create", handlePOSTCreateToDoItem)
	mux.HandleFunc("PUT /api/edit", handlePUTEditToDoItem)
	mux.HandleFunc("PATCH /api/toggle-complete/{itemId}", handlePATCHToggleComplete)
	mux.HandleFunc("DELETE /api/delete/{itemId}", handleDELETEToDoItem)

	slog.InfoContext(ctx, fmt.Sprintf("Server listening on %s%s", ServerProtocol, ServerAddress))

	err := http.ListenAndServe(ServerAddress, middleware(ctx, mux))
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("There's an error with the server: %s", err.Error()))
	}
}

func middleware(ctx context.Context, existingHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestId := generateRequestID()

		handlerCtx := context.WithValue(request.Context(), "requestId", requestId)
		handlerCtx = context.WithValue(handlerCtx, "logger", ctx.Value("logger"))
		handlerCtx, cancel := context.WithTimeout(handlerCtx, 5*time.Second)
		defer cancel()

		request = request.WithContext(handlerCtx)

		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)
			return
		}

		slog.InfoContext(handlerCtx, fmt.Sprintf("%v - %v", request.Method, request.URL.Path))

		existingHandler.ServeHTTP(writer, request)
	})
}

func encodeJson(ctx context.Context, writer http.ResponseWriter, data any) {
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error encoding json: %s", err.Error()))
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

func generateRequestID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
