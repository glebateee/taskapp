package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/glebateee/taskapp/internal/core/errors"
	core_logger "github.com/glebateee/taskapp/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	logger *core_logger.Logger
	w      http.ResponseWriter
}

func NewHTTPResponseHandler(
	logger *core_logger.Logger,
	w http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		logger: logger,
		w:      w,
	}
}

func (h *HTTPResponseHandler) HTMLResponse(file []byte) {
	h.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.w.WriteHeader(http.StatusOK)
	_, err := h.w.Write(file)
	if err != nil {
		h.logger.Error("write html HTTP response", zap.Error(err))
	}
}
func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.logger.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.logger.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.logger.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.logger.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("Unexpected panic: %v", p)
	h.logger.Error(msg, zap.Error(err))
	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {

	response := ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}
	h.JSONResponse(response, statusCode)
}

func (h *HTTPResponseHandler) JSONResponse(
	body any,
	statusCode int,
) {
	h.w.Header().Set("Content-Type", "application/json")
	h.w.WriteHeader(statusCode)
	if err := json.NewEncoder(h.w).Encode(body); err != nil {
		h.logger.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.w.WriteHeader(http.StatusNoContent)
}
