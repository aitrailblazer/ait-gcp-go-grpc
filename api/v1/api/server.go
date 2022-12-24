package api

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models"
	"github.com/go-playground/validator/v10"

	// "github.com/go-playground/validator/v10"

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

// var ( // The following are used to validate the data that is being sent to the server.
//
//	uUIDRegex                 = regexp.MustCompile(uUIDRegexString)                 // The uUIDRegex is used to validate the uUID. The uUIDRegex is a regexp.Regexp. The uUIDRegex is created from the uUIDRegexString. The uUIDRegexString is a string. The uUIDRegexString is assigned to the uUIDRegex.
//	alphaNumericAndSpaceRegex = regexp.MustCompile(alphaNumericAndSpaceRegexString) // The alphaNumericAndSpaceRegex is used to validate the alphaNumericAndSpace. The alphaNumericAndSpaceRegex is a regexp.Regexp. The alphaNumericAndSpaceRegex is created from the alphaNumericAndSpaceRegexString. The alphaNumericAndSpaceRegexString is a string. The alphaNumericAndSpaceRegexString is assigned to the alphaNumericAndSpaceRegex.
//	unitPriceRegex            = regexp.MustCompile(unitPriceString)                 // The unitPriceRegex is used to validate the unitPrice. The unitPriceRegex is a regexp.Regexp. The unitPriceRegex is created from the unitPriceString. The unitPriceString is a string. The unitPriceString is assigned to the unitPriceRegex.
//
// )
// MinMax    string `validate:"min=1,max=10"`

type NewPetValid struct {
	Name *string `validate:"min=1,max=10"`
	Tag  *string `validate:"min=1,max=10"`
}

// use a single instance of Validate, it caches struct info

var validate *validator.Validate

// UserStructLevelValidation contains custom struct level validations that don't always

// make sense at the field validation level. For Example this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//

// NOTE: you may ask why wouldn't I just do this outside of validator, because doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.

func NewPetStructLevelValidation(sl validator.StructLevel) {

	newPetValid := sl.Current().Interface().(NewPetValid)

	if len(*newPetValid.Name) == 0 && len(*newPetValid.Tag) == 0 {
		sl.ReportError(*newPetValid.Name, "name", "Name", "nameortag", "")
		sl.ReportError(*newPetValid.Tag, "tag", "Tag", "nameortag", "")
	}
	// plus can do more, even with different tag than
	// "fnameorlname"
}

// post:"/pets" body:"new_pet"
func (h *Handler) AitrailblazerServiceAddPet(ctx echo.Context) error {
	// We expect a NewPet object in the request body.
	var newPet models.NewPet
	err := ctx.Bind(&newPet)
	if err != nil {
		return sendPetStoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}
	validate = validator.New()
	// register function to get tag name from json tags.

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// register validation for 'NewPet'
	// NOTE: only have to register a non-pointer type for 'NewPet', validator
	// internally dereferences during it's type checks.
	validate.RegisterStructValidation(NewPetStructLevelValidation, NewPetValid{})
	newPetValid := &NewPetValid{
		Name: newPet.Name,
		Tag:  newPet.Tag,
	}

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err = validate.Struct(newPetValid)
	if err != nil {
		return sendPetStoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}

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
