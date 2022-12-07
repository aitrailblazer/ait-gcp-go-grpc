//
// Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

package main

import (
	"github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/api"
	"log"
	"net/http"
	"os"

	// "github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const PROJECT_ID = "ait-gcp-go-grpc"

// curl -X GET http://localhost:8080/v1/ -v    HTTP/1.1 200 OK                     {status: "running"}
// curl -X GET http://localhost:8080/v1/1 -v   HTTP/1.1 404 Not Found              {"message":"Not Found"}
// curl -X POST http://localhost:8080/v1/  -v  HTTP/1.1 405 Method Not Allowed     {"message":"Method Not Allowed"}
func (ctrl *controller) v1Handler(c echo.Context) error {
	log.Printf("v1Handler:Method %v c.Path %v ", c.Request().Method, c.Path())
	return c.String(http.StatusOK, "{\"status\": \"running\"}")
}
func (ctrl *controller) notFoundHandler(c echo.Context) error {
	log.Printf("notFoundHandler:Method %v c.Path %v ", c.Request().Method, c.Path())
	return c.String(http.StatusNotFound, "{\"status\": \"NotFound\"}")

	// return c.String(http.StatusNotFound, c.Request().Method+" "+c.Path())
}

type controller struct {
	logger echo.Logger //
}

func routerSetup(e *echo.Echo) *echo.Echo {
	// Create an instance of our handler which satisfies the generated interface
	// PongStore := api.NewPongStore()

	// Middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	ctrl := controller{logger: e.Logger}
	e.GET("/v1/", ctrl.v1Handler) // handler method with controller struct

	e.RouteNotFound("/*", ctrl.notFoundHandler) // any

	// We now register our petStore above as the handler for the interface
	// api.RegisterHandlers(e, PongStore)
	return e
}

// funcmain .
//
// .funcmain
// [source,go]
// ----
// include::${gad:current:fq}[tag=funcmain,indent=0]
// ----
// <1> PORT is set by the environment
// <2> if PORT is not set, use 8080
// <3> log the port
// <4> create a new echo instance
// <5> setup the router
// <6> log the port
// <7> start the server
// tag::funcmain[]
func main() {

	port := os.Getenv("PORT") // <1>
	if port == "" {
		port = "8080" // <2>
	}
	log.Printf("port %s ", port) // <3>

	e := echo.New() // <4>

	e = routerSetup(e) // <5>
	// Start server
	log.Println(PROJECT_ID, "REST API listening on port", port) // <6>
	e.Logger.Fatal(e.Start(":" + port))                         // <7>
}

// end::funcmain[]
