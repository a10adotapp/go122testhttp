package users

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type GetIDResponseData struct {
	User string `json:"user"`
}

func (e *Endpoint) GetID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		slog.Error("strconv.ParseUint", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := e.UsersService.GetUser(uint32(id))
	if err != nil {
		slog.Error("e.UsersService.GetUser", slog.Any("err", err))
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
