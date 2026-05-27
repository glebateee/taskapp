package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/glebateee/taskapp/internal/core/logger"
	core_http_request "github.com/glebateee/taskapp/internal/core/transport/http/request"
	core_http_response "github.com/glebateee/taskapp/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContextMust(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	logger.Debug("invoke GetTask handler")
	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task_id path value")
		return
	}
	tasksDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}
	response := GetTaskResponse(taskDTOFromDomain(tasksDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
