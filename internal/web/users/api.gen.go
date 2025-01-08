// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Id     *uint   `json:"id,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
	Task   *string `json:"task,omitempty"`
	UserId *uint   `json:"user_id,omitempty"`
}

// User defines model for User.
type User struct {
	Id       *uint   `json:"id,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = User

// PatchUsersIdJSONRequestBody defines body for PatchUsersId for application/json ContentType.
type PatchUsersIdJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Create a new user
	// (POST /users)
	PostUsers(ctx echo.Context) error
	// Delete user by ID
	// (DELETE /users/{id})
	DeleteUsersId(ctx echo.Context, id int) error
	// Partially update a user by ID
	// (PATCH /users/{id})
	PatchUsersId(ctx echo.Context, id int) error
	// Get tasks by user ID
	// (GET /users/{user_id}/task)
	GetUsersUserIdTask(ctx echo.Context, userId int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// DeleteUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsersId(ctx, id)
	return err
}

// PatchUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUsersId(ctx, id)
	return err
}

// GetUsersUserIdTask converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersUserIdTask(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "user_id" -------------
	var userId int

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", ctx.Param("user_id"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersUserIdTask(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.DELETE(baseURL+"/users/:id", wrapper.DeleteUsersId)
	router.PATCH(baseURL+"/users/:id", wrapper.PatchUsersId)
	router.GET(baseURL+"/users/:user_id/task", wrapper.GetUsersUserIdTask)

}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUsersRequestObject struct {
	Body *PostUsersJSONRequestBody
}

type PostUsersResponseObject interface {
	VisitPostUsersResponse(w http.ResponseWriter) error
}

type PostUsers201JSONResponse User

func (response PostUsers201JSONResponse) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUsersIdRequestObject struct {
	Id int `json:"id"`
}

type DeleteUsersIdResponseObject interface {
	VisitDeleteUsersIdResponse(w http.ResponseWriter) error
}

type DeleteUsersId204Response struct {
}

func (response DeleteUsersId204Response) VisitDeleteUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type PatchUsersIdRequestObject struct {
	Id   int `json:"id"`
	Body *PatchUsersIdJSONRequestBody
}

type PatchUsersIdResponseObject interface {
	VisitPatchUsersIdResponse(w http.ResponseWriter) error
}

type PatchUsersId200JSONResponse User

func (response PatchUsersId200JSONResponse) VisitPatchUsersIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersUserIdTaskRequestObject struct {
	UserId int `json:"user_id"`
}

type GetUsersUserIdTaskResponseObject interface {
	VisitGetUsersUserIdTaskResponse(w http.ResponseWriter) error
}

type GetUsersUserIdTask200JSONResponse []Task

func (response GetUsersUserIdTask200JSONResponse) VisitGetUsersUserIdTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Create a new user
	// (POST /users)
	PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error)
	// Delete user by ID
	// (DELETE /users/{id})
	DeleteUsersId(ctx context.Context, request DeleteUsersIdRequestObject) (DeleteUsersIdResponseObject, error)
	// Partially update a user by ID
	// (PATCH /users/{id})
	PatchUsersId(ctx context.Context, request PatchUsersIdRequestObject) (PatchUsersIdResponseObject, error)
	// Get tasks by user ID
	// (GET /users/{user_id}/task)
	GetUsersUserIdTask(ctx context.Context, request GetUsersUserIdTaskRequestObject) (GetUsersUserIdTaskResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUsers operation middleware
func (sh *strictHandler) PostUsers(ctx echo.Context) error {
	var request PostUsersRequestObject

	var body PostUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsers(ctx.Request().Context(), request.(PostUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersResponseObject); ok {
		return validResponse.VisitPostUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUsersId operation middleware
func (sh *strictHandler) DeleteUsersId(ctx echo.Context, id int) error {
	var request DeleteUsersIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUsersId(ctx.Request().Context(), request.(DeleteUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUsersIdResponseObject); ok {
		return validResponse.VisitDeleteUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUsersId operation middleware
func (sh *strictHandler) PatchUsersId(ctx echo.Context, id int) error {
	var request PatchUsersIdRequestObject

	request.Id = id

	var body PatchUsersIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUsersId(ctx.Request().Context(), request.(PatchUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUsersIdResponseObject); ok {
		return validResponse.VisitPatchUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsersUserIdTask operation middleware
func (sh *strictHandler) GetUsersUserIdTask(ctx echo.Context, userId int) error {
	var request GetUsersUserIdTaskRequestObject

	request.UserId = userId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersUserIdTask(ctx.Request().Context(), request.(GetUsersUserIdTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersUserIdTask")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersUserIdTaskResponseObject); ok {
		return validResponse.VisitGetUsersUserIdTaskResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
