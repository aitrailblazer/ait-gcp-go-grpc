package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models"

	"github.com/labstack/echo/v4"
)

const VERSION = "1.02"

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendPetStoreError(ctx echo.Context, code int, message string) error {
	code32 := int32(code)
	petErr := models.Error{
		Code:    &code32,
		Message: &message,
	}
	err := ctx.JSON(code, petErr)
	return err
}

// type PetStore struct {
// 	Lock   sync.Mutex
// 	Pets   map[int64]models.Pet
// 	NextId int64
// }

// func NewPetStore() *PetStore {
// 	return &PetStore{
// 		Pets:   make(map[int64]models.Pet),
// 		NextId: 1000,
// 	}
// }

// type Handler struct {}
type Handler struct {
	// Lock   sync.Mutex
	Pets   map[int64]models.Pet
	NextId int64
}

var dbmux sync.Mutex

//	func NewHandler() Handler {
//		// pong := Pong{
//		// 	Index:      1,
//		// 	Message:    "Pong!",
//		// 	ReceivedOn: time.Now(),
//		// }
//		handler := Handler{}
//		return handler
//	}
func NewHandler() *Handler {
	return &Handler{
		Pets:   make(map[int64]models.Pet),
		NextId: 1000,
	}
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
// func sendServiceError(ctx echo.Context, code int, message string) error {
// 	petErr := models.Error{
// 		Code:    int32(code),
// 		Message: message,
// 	}
// 	err := ctx.JSON(code, petErr)
// 	return err
// }

// curl localhost:8080/v1/ping
// curl localhost:8080/v1/ping?message=test
func (h *Handler) AitrailblazerServiceSendPing(ctx echo.Context, params models.AitrailblazerServiceSendPingParams) error {

	var i int32 = 369
	msec := time.Now()
	log.Printf("AitrailblazerServiceSendPing VERSION %s ", VERSION) // <3>

	message := params.Message
	var p string
	if params.Message != nil {
		fmt.Println("AitrailblazerServiceSendPing: message ", *message)
		p = "path parameter=" + *message + " Ping from AitrailblazerServiceSendPing"
	} else {
		p = " Ping from AitrailblazerServiceSendPing"
	}
	var v string = VERSION
	pong := models.Pong{
		Index:      &i,
		Message:    &p,
		ReceivedOn: &msec,
		Ver:        &v,
	}
	return ctx.JSON(http.StatusOK, pong)
}

// (GET /echo/:message)
// curl localhost:8080/echo/test
// curl localhost:8080/echo/test\?value\=test1234
func (h *Handler) AitrailblazerServiceEcho(ctx echo.Context, message string, params models.AitrailblazerServiceEchoParams) error {
	// ------------- Path parameter "message" -------------
	fmt.Println("AitrailblazerServiceEcho: Path parameter ", message)
	// paramsValue := params.Value

	if params.Value != nil {
		fmt.Println("AitrailblazerServiceEcho: params.Value ", *params.Value)
		message = "path parameter=" + message + " query=" + *params.Value + " Echo from AitrailblazerServiceEcho"
	} else {
		message = "path parameter=" + message + " Echo from AitrailblazerServiceEcho"
	}

	echoMessage := models.EchoMessage{
		Value: &message,
	}
	return ctx.JSON(http.StatusOK, echoMessage)
}

func (h *Handler) AitrailblazerServiceFindPetByID(ctx echo.Context, petId int64) error {
	dbmux.Lock()
	defer dbmux.Unlock()

	pet, found := h.Pets[petId]
	if !found {
		return sendPetStoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", petId))
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (h *Handler) AitrailblazerServiceDeletePet(ctx echo.Context, id int64) error {
	dbmux.Lock()
	defer dbmux.Unlock()

	_, found := h.Pets[id]
	if !found {
		return sendPetStoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", id))
	}
	delete(h.Pets, id)
	return ctx.NoContent(http.StatusNoContent)
}

//	rpc FindPets ( FindPetsParameters ) returns ( Pet ) {
//	    option (google.api.http) = {
//	      get:"/pets"
//	    };
//	 }
//
//	message FindPetsParameters {
//		repeated string tags = 1;
//		int32 limit = 2;
//	}
//
// FindPets implements all the handlers in the ServerInterface
func (h *Handler) AitrailblazerServiceFindPets(ctx echo.Context, params models.AitrailblazerServiceFindPetsParams) error {
	dbmux.Lock()
	defer dbmux.Unlock()

	var result []models.Pet
	for _, pet := range h.Pets {
		if params.Tags != nil {
			// If we have tags,  filter pets by tag
			for _, t := range *params.Tags {
				if pet.Tag != nil && (*pet.Tag == t) {
					result = append(result, pet)
				}
			}
		} else {
			// Add all pets if we're not filtering
			result = append(result, pet)
		}

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				// We're at the limit
				break
			}
		}
	}
	return ctx.JSON(http.StatusOK, result)
}

// post:"/pets" body:"new_pet"
func (h *Handler) AitrailblazerServiceAddPet(ctx echo.Context) error {
	// We expect a NewPet object in the request body.
	var newPet models.NewPet
	err := ctx.Bind(&newPet)
	if err != nil {
		return sendPetStoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}
	// We now have a pet, let's add it to our "database".

	// We're always asynchronous, so lock unsafe operations below
	dbmux.Lock()
	defer dbmux.Unlock()

	// We handle pets, not NewPets, which have an additional ID field
	var pet models.Pet
	pet.Name = newPet.Name
	pet.Tag = newPet.Tag
	pet.Id = &h.NextId
	h.NextId = h.NextId + 1

	// Insert into map
	h.Pets[*pet.Id] = pet

	// Now, we have to return the NewPet
	err = ctx.JSON(http.StatusCreated, pet)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}

	// Return no error. This refers to the handler. Even if we return an HTTP
	// error, but everything else is working properly, tell Echo that we serviced
	// the error. We should only return errors from Echo handlers if the actual
	// servicing of the error on the infrastructure level failed. Returning an
	// HTTP/400 or HTTP/500 from here means Echo/HTTP are still working, so
	// return nil.
	return nil
}

func (h *Handler) ListPets() []models.Pet {
	dbmux.Lock()
	defer dbmux.Unlock()
	var keys []int64        // Create a new slice of strings.  The slice is used to store the keys of the database map. The slice is created empty.
	for k := range h.Pets { // For each key in the database map.  The key is a string. The key is assigned to k.
		keys = append(keys, k) // The key is appended to the slice of keys.
	}
	// sort.Ints(keys)
	var pets []models.Pet
	for _, k := range keys {
		v := h.Pets[k]
		var pet models.Pet
		pet.Name = v.Name
		pet.Tag = v.Tag
		pet.Id = v.Id
		pets = append(pets, pet)
	}
	return pets
}

// (GET /listpets)
func (h *Handler) AitrailblazerServiceListPets(ctx echo.Context, params models.AitrailblazerServiceListPetsParams) error {
	pets := h.ListPets()

	return ctx.JSON(http.StatusOK, pets)
}

// (GET /v1/shelves/{shelf})
// func (h *Handler) AitrailblazerServiceGetShelf(ctx echo.Context, shelf int64) error {

// 	log.Printf("AitrailblazerServiceGetShelf VERSION %s ", VERSION) // <3>

// 	fmt.Println("AitrailblazerServiceGetShelf: shelf ", shelf)
// 	i := shelf

// 	theme := "My Bookshelf"
// 	bookshelf := models.Shelf{
// 		Id:    &i,
// 		Theme: &theme,
// 	}
// 	return ctx.JSON(http.StatusOK, bookshelf)
// }

// (GET /v1/ping)
// func (s *pingService) Send(ctx context.Context, req *pb.Request) (*pb.Response, error) {
// 	log.Print("sending ping response")
// 	return &pb.Response{
// 		Pong: &models.Pong{
// 			Index:      1,
// 			Message:    req.GetMessage(),
// 			ReceivedOn: ptypes.TimestampNow(),
// 		},
// 	}, nil
// }
