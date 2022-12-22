//
// Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

package main

import (
	// "fmt"

	// "github.com/deepmap/oapi-codegen/pkg/middleware"

	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/api"
	"github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed ui
var embededFiles embed.FS

func getFileSystem() http.FileSystem {
	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "ui")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

const PROJECT_ID = "ait-gcp-go-grpc"

// curl -X GET http://localhost:8080/v1/ -v    HTTP/1.1 200 OK                     {status: "running"}
// curl -X GET http://localhost:8080/v1/1 -v   HTTP/1.1 404 Not Found              {"message":"Not Found"}
// curl -X POST http://localhost:8080/v1/  -v  HTTP/1.1 405 Method Not Allowed     {"message":"Method Not Allowed"}
// func (ctrl *controller) v1Handler(c echo.Context) error {
// 	log.Printf("v1Handler:Method %v c.Path %v ", c.Request().Method, c.Path())
// 	return c.String(http.StatusOK, "{\"status\": \"running\"}")
// }
// func (ctrl *controller) notFoundHandler(c echo.Context) error {
// 	log.Printf("notFoundHandler:Method %v c.Path %v ", c.Request().Method, c.Path())
// 	return c.String(http.StatusNotFound, "{\"status\": \"NotFound\"}")

// 	// return c.String(http.StatusNotFound, c.Request().Method+" "+c.Path())
// }

// type controller struct {
// 	logger echo.Logger //
// }

// go run -v main.go
// curl localhost:8080/v1/ping
// {"index":1,"message":"pong","receivedOn":"2022-12-07T12:35:24.465638-08:00"}

// curl localhost:8080/v1/shelves/{shelf}
// curl localhost:8080/v1/shelves/369
// {"id":369,"theme":"My Bookshelf"}

// routerSetup .
//
// .routerSetup
// [source,go]
// ----
// include::${gad:current:fq}[tag=routerSetup,indent=0]
// ----
// <1> create a new handler
// <2> log all requests
// <3> register the handlers
// tag::routerSetup[]

// func ListPets(h *api.Handler) []models.Pet {
// 	var keys []int64        // Create a new slice of strings.  The slice is used to store the keys of the database map. The slice is created empty.
// 	for k := range h.Pets { // For each key in the database map.  The key is a string. The key is assigned to k.
// 		keys = append(keys, k) // The key is appended to the slice of keys.
// 	}
// 	// sort.Ints(keys)
// 	var pets []models.Pet
// 	for _, k := range keys {
// 		v := h.Pets[k]
// 		var pet models.Pet
// 		pet.Name = v.Name
// 		pet.Tag = v.Tag
// 		pet.Id = v.Id
// 		pets = append(pets, pet)
// 	}
// 	return pets
// }

var dbmux sync.Mutex

func routerSetup(e *echo.Echo) *echo.Echo {
	h := api.NewHandler() // <1>
	dbmux.Lock()
	defer dbmux.Unlock()
	// Example #6: name = "Max", tag = "TagOfMax"
	name1 := "Max"
	tag1 := "TagOfMax"
	// Example #7: name = "Charlie", tag = "TagOfCharlie"
	name2 := "Charlie"
	tag2 := "TagOfCharlie"
	// Example #8: name = "Buster", tag = "TagOfBuster"
	name3 := "Buster"
	tag3 := "TagOfBuster"
	var pet1 models.Pet
	//
	pet1.Name = &name1
	pet1.Tag = &tag1
	var Id1 int64 = 100
	pet1.Id = &Id1
	// h.NextId = h.NextId + 1
	// Insert into map
	h.Pets[*pet1.Id] = pet1
	//
	var pet2 models.Pet

	pet2.Name = &name2
	pet2.Tag = &tag2
	var Id2 int64 = 110
	pet2.Id = &Id2
	// h.NextId = h.NextId + 1
	// Insert into map
	h.Pets[*pet2.Id] = pet2
	//
	//
	var pet3 models.Pet

	pet3.Name = &name3
	pet3.Tag = &tag3
	var Id3 int64 = 120
	// h.NextId = h.NextId + 1
	pet3.Id = &Id3
	// Insert into map
	h.Pets[*pet3.Id] = pet3

	pets := h.ListPets()
	for k := range pets {
		log.Printf("pets Id %v", *pets[k].Id)
		log.Printf("pets Name %v", *pets[k].Name)
		log.Printf("pets Name %v", *pets[k].Tag)
	}
	// log.Printf("pets %v", pets)
	// Log all requests
	e.Use(middleware.Logger()) // <2>
	api.RegisterHandlers(e, h) // <3>

	return e
}

// end::routerSetup[]

// func init() {

// use embed
// http://localhost:8080/ui/
// filehandler := http.FileServer(http.FS(src))
// use http.Handle
// shows the ui page. needs a slash (/) at the end of the name
// http.Handle("/ui/", http.StripPrefix("/", filehandler))

// }

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
	log.Printf("port %s ", port)           // <3>
	log.Printf("VERSION %s ", api.VERSION) //

	e := echo.New() // <4>
	assetHandler := http.FileServer(getFileSystem())

	e = routerSetup(e) // <5>

	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	log.Println(PROJECT_ID, "REST API listening on port", port) // <6>
	e.Logger.Fatal(e.Start(":" + port))
}

// end::funcmain[]
// Handler
// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }

// func main() {
// 	log.Print("starting server...")
// 	http.HandleFunc("/", handler)

// 	// Determine port for HTTP service.
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 		log.Printf("defaulting to port %s", port)
// 	}

// 	// Start HTTP server.
// 	log.Printf("listening on port %s", port)
// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	name := os.Getenv("NAME")
// 	if name == "" {
// 		name = "World"
// 	}
// 	fmt.Fprintf(w, "Hello %s!\n", name)
// }
