package users

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type PostRequestData struct {
	Name string `json:"name"`
}

type PostResponseData struct {
	User string `json:"user"`
}

func (e *Endpoint) Post(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("io.ReadAll", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	requestData := PostRequestData{}
	if err := json.Unmarshal(requestBody, &requestData); err != nil {
		slog.Error("json.Unmarshal", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := e.UsersService.CreateUser(requestData.Name)
	if err != nil {
		slog.Error("e.UsersService.CreateUser", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := GetIDResponseData{
		User: user,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		slog.Error("json.Marshal", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}
