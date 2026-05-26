package core_http_response

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

const statusCodeUninitialized = -1

func NewResponseWriter(
	w http.ResponseWriter,
) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     statusCodeUninitialized,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ResponseWriter) GetStatusCode() int {
	if w.statusCode == statusCodeUninitialized {
		w.statusCode = http.StatusOK
	}
	return w.statusCode
}
