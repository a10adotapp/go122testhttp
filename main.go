package main

import (
	"go122test/endpoint/users"
	"log/slog"
	"net/http"
	"os"
)

type Endpoint interface {
	PatternHandlers() map[string]func(http.ResponseWriter, *http.Request)
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})))

	mux := http.NewServeMux()

	for _, endpoint := range []Endpoint{
		users.NewEndpoint(&DummyUsersService{}),
	} {
		for pattern, handler := range endpoint.PatternHandlers() {
			mux.Handle(pattern, http.HandlerFunc(handler))
		}
	}

	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	}))
}

type DummyUsersService struct{}

func (d *DummyUsersService) CreateUser(name string) (string, error) {
	return name, nil
}

func (d *DummyUsersService) GetUser(id uint32) (string, error) {
	return "John Doe", nil
}
