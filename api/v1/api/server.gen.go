// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /echo/{message})
	AitrailblazerServiceEcho(ctx echo.Context, message string, params AitrailblazerServiceEchoParams) error

	// (GET /listpets)
	AitrailblazerServiceListPets(ctx echo.Context, params AitrailblazerServiceListPetsParams) error

	// (GET /pets)
	AitrailblazerServiceFindPets(ctx echo.Context, params AitrailblazerServiceFindPetsParams) error

	// (POST /pets)
	AitrailblazerServiceAddPet(ctx echo.Context) error

	// (DELETE /pets/{id})
	AitrailblazerServiceDeletePet(ctx echo.Context, id int64) error

	// (GET /pets/{id})
	AitrailblazerServiceFindPetByID(ctx echo.Context, id int64) error

	// (GET /v1/ping)
	AitrailblazerServiceSendPing(ctx echo.Context, params AitrailblazerServiceSendPingParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AitrailblazerServiceEcho converts echo context to params.
func (w *ServerInterfaceWrapper) AitrailblazerServiceEcho(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "message" -------------
	var message string

	err = runtime.BindStyledParameterWithLocation("simple", false, "message", runtime.ParamLocationPath, ctx.Param("message"), &message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter message: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params AitrailblazerServiceEchoParams
	// ------------- Optional query parameter "value" -------------

	err = runtime.BindQueryParameter("form", true, false, "value", ctx.QueryParams(), &params.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter value: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AitrailblazerServiceEcho(ctx, message, params)
	return err
}

// AitrailblazerServiceListPets converts echo context to params.
func (w *ServerInterfaceWrapper) AitrailblazerServiceListPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AitrailblazerServiceListPetsParams
	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageSize: %s", err))
	}

	// ------------- Optional query parameter "pageToken" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageToken", ctx.QueryParams(), &params.PageToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageToken: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AitrailblazerServiceListPets(ctx, params)
	return err
}

// AitrailblazerServiceFindPets converts echo context to params.
func (w *ServerInterfaceWrapper) AitrailblazerServiceFindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AitrailblazerServiceFindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AitrailblazerServiceFindPets(ctx, params)
	return err
}

// AitrailblazerServiceAddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AitrailblazerServiceAddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AitrailblazerServiceAddPet(ctx)
	return err
}

// AitrailblazerServiceDeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) AitrailblazerServiceDeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AitrailblazerServiceDeletePet(ctx, id)
	return err
}

// AitrailblazerServiceFindPetByID converts echo context to params.
func (w *ServerInterfaceWrapper) AitrailblazerServiceFindPetByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AitrailblazerServiceFindPetByID(ctx, id)
	return err
}

// AitrailblazerServiceSendPing converts echo context to params.
func (w *ServerInterfaceWrapper) AitrailblazerServiceSendPing(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AitrailblazerServiceSendPingParams
	// ------------- Optional query parameter "message" -------------

	err = runtime.BindQueryParameter("form", true, false, "message", ctx.QueryParams(), &params.Message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter message: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AitrailblazerServiceSendPing(ctx, params)
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

	router.GET(baseURL+"/echo/:message", wrapper.AitrailblazerServiceEcho)
	router.GET(baseURL+"/listpets", wrapper.AitrailblazerServiceListPets)
	router.GET(baseURL+"/pets", wrapper.AitrailblazerServiceFindPets)
	router.POST(baseURL+"/pets", wrapper.AitrailblazerServiceAddPet)
	router.DELETE(baseURL+"/pets/:id", wrapper.AitrailblazerServiceDeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.AitrailblazerServiceFindPetByID)
	router.GET(baseURL+"/v1/ping", wrapper.AitrailblazerServiceSendPing)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
