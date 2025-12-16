package respond

import (
	"encoding/json"
	"net/http"

	kiterrors "github.com/nshinoks/go-webkit/errors"
)

func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func Error(w http.ResponseWriter, err error) {
	p := kiterrors.ToProblem(err)
	p.Write(w)
}
