package logger

import (
	"net/http"
)

type ResponseBody struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

type StatusWriter struct {
	http.ResponseWriter
	Status int
}

func (w *StatusWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *StatusWriter) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = 200
	}
	n, e := w.ResponseWriter.Write(b)
	return n, e
}
