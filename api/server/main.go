//
// Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

package main

import (
	"log"
	"net/http"
	"os"

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
	// Middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	ctrl := controller{logger: e.Logger}
	e.GET("/v1/", ctrl.v1Handler) // handler method with controller struct

	e.RouteNotFound("/*", ctrl.notFoundHandler) // any

	return e
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("port %s ", port)

	e := echo.New()

	e = routerSetup(e)
	// Start server
	log.Println(PROJECT_ID, "REST API listening on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
