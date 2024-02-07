package users

import "net/http"

type Endpoint struct {
	UsersService UsersService
}

func NewEndpoint(usersService UsersService) *Endpoint {
	return &Endpoint{
		UsersService: usersService,
	}
}

func (e *Endpoint) PatternHandlers() map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"POST /users":     e.Post,
		"GET /users/{id}": e.GetID,
	}
}

type UsersService interface {
	CreateUser(name string) (string, error)
	GetUser(id uint32) (string, error)
}
