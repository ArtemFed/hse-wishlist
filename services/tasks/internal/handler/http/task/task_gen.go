// Package task provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// TaskCreate defines model for TaskCreate.
type TaskCreate struct {
	CreatedBy openapi_types.UUID `json:"CreatedBy"`
	EndedAt   time.Time          `json:"EndedAt"`
	Percent   float32            `json:"Percent"`
	StartedAt time.Time          `json:"StartedAt"`
	Status    string             `json:"Status"`
}

// TaskFilter defines model for TaskFilter.
type TaskFilter struct {
	CreatedBy *openapi_types.UUID `json:"CreatedBy,omitempty"`
	Percent   *float32            `json:"Percent,omitempty"`
	Status    *string             `json:"Status,omitempty"`
	UUID      *openapi_types.UUID `json:"UUID,omitempty"`
}

// TaskGet defines model for TaskGet.
type TaskGet struct {
	CreatedAt    time.Time          `json:"CreatedAt"`
	CreatedBy    openapi_types.UUID `json:"CreatedBy"`
	EndedAt      time.Time          `json:"EndedAt"`
	LastUpdateAt time.Time          `json:"LastUpdateAt"`
	Percent      float32            `json:"Percent"`
	StartedAt    time.Time          `json:"StartedAt"`
	Status       string             `json:"Status"`
	UUID         openapi_types.UUID `json:"UUID"`
}

// TaskUpdate defines model for TaskUpdate.
type TaskUpdate struct {
	CreatedBy openapi_types.UUID `json:"CreatedBy"`
	EndedAt   time.Time          `json:"EndedAt"`
	Percent   float32            `json:"Percent"`
	StartedAt time.Time          `json:"StartedAt"`
	Status    string             `json:"Status"`
	UUID      openapi_types.UUID `json:"UUID"`
}

// ModelUUID defines model for ModelUUID.
type ModelUUID struct {
	UUID openapi_types.UUID `json:"UUID"`
}

// GetTasksJSONRequestBody defines body for GetTasks for application/json ContentType.
type GetTasksJSONRequestBody = TaskFilter

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = TaskCreate

// PutTasksJSONRequestBody defines body for PutTasks for application/json ContentType.
type PutTasksJSONRequestBody = TaskUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List tasks
	// (GET /tasks)
	GetTasks(c *gin.Context)
	// Create a task
	// (POST /tasks)
	PostTasks(c *gin.Context)
	// Update a task
	// (PUT /tasks)
	PutTasks(c *gin.Context)
	// Get task by ID
	// (GET /tasks/{id})
	GetTasksId(c *gin.Context, id openapi_types.UUID)
	// End a task
	// (PATCH /tasks/{id}/end)
	PatchTasksIdEnd(c *gin.Context, id openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetTasks operation middleware
func (siw *ServerInterfaceWrapper) GetTasks(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTasks(c)
}

// PostTasks operation middleware
func (siw *ServerInterfaceWrapper) PostTasks(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostTasks(c)
}

// PutTasks operation middleware
func (siw *ServerInterfaceWrapper) PutTasks(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutTasks(c)
}

// GetTasksId operation middleware
func (siw *ServerInterfaceWrapper) GetTasksId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTasksId(c, id)
}

// PatchTasksIdEnd operation middleware
func (siw *ServerInterfaceWrapper) PatchTasksIdEnd(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchTasksIdEnd(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/tasks", wrapper.GetTasks)
	router.POST(options.BaseURL+"/tasks", wrapper.PostTasks)
	router.PUT(options.BaseURL+"/tasks", wrapper.PutTasks)
	router.GET(options.BaseURL+"/tasks/:id", wrapper.GetTasksId)
	router.PATCH(options.BaseURL+"/tasks/:id/end", wrapper.PatchTasksIdEnd)
}
