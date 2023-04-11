package http

import (
	"net/http"
)

func (handler Handler) welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
