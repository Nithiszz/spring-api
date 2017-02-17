package app

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Nithiszz/sprint-api/pkg/api"
)

// BindJSON parses json body to var
func bindJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// RenderJSON renders var to response write
func renderJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)

}

func renderString(w http.ResponseWriter, status int, v string) {
	w.WriteHeader(status)
	w.Write([]byte(v))
}

func renderNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func renderError(w http.ResponseWriter, err error) {
	switch err {
	case api.ErrNotFound:
		renderString(w, http.StatusNotFound, err.Error())
	default:
		renderString(w, http.StatusInternalServerError, err.Error())
	}

}
