package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/glebateee/taskapp/internal/core/domain"
	core_errors "github.com/glebateee/taskapp/internal/core/errors"
	core_logger "github.com/glebateee/taskapp/internal/core/logger"
	core_http_request "github.com/glebateee/taskapp/internal/core/transport/http/request"
	core_http_response "github.com/glebateee/taskapp/internal/core/transport/http/response"
	core_http_types "github.com/glebateee/taskapp/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name" swaggertype:"string" example:"Максичка Кахановочка"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" swaggertype:"string" example:"+7688688809898"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("'full_name' can't be NULL: %w", core_errors.ErrInvalidArgument)
		}

		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("'full_name' must be between 3 and 100 symbols")
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("'phone_number' must be between 10 and 15 symbols")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("'phone_number' must start with '+' symbol")

			}
		}
	}
	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser 		godoc
// @Summary 		Patching user
// @Description 	Changing information about user in system
// @Description 	### Logic of fields updating
// @Description 	1.**Field not provided**: `phone_number` ignored, value in DB not changed
// @Description 	2.**Value provided, not null**: `phone_number : +1312313234` sets new value to record field in DB
// @Description 	3.**Value provided, null**: `phone_number : +null` sets NULL to record field in DB
// @Description 	Constraints: `full_name` value can't be set to null
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "ID of user to change"
// @Param 			request body PatchUserRequest true "Patch user request body"
// @Success 		200 {object} PatchUserResponse "Successfully patched user"
// @Failure 		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 		404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 		409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 			/users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContextMust(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return

	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}
	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(r PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		r.FullName.ToDomain(),
		r.PhoneNumber.ToDomain(),
	)
}
