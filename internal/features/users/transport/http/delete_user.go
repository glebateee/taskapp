package users_transport_http

import (
	"net/http"

	core_logger "github.com/glebateee/taskapp/internal/core/logger"
	core_http_request "github.com/glebateee/taskapp/internal/core/transport/http/request"
	core_http_response "github.com/glebateee/taskapp/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContextMust(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	logger.Debug("invoke DeleteUser handler")
	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}
	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}
	responseHandler.NoContentResponse()
	//response := GetUserResponse(userDTOFromDomain(userDomain))
	//responseHandler.JSONResponse(response, http.StatusOK)
}
