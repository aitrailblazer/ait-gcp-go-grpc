// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// EchoMessage defines model for EchoMessage.
type EchoMessage struct {
	Value *string `json:"value,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// GoogleProtobufAny Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type GoogleProtobufAny struct {
	// Type The type of the serialized message.
	Type                 *string                `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPet The NewPet message contains information about a  new pet being added to the system. It includes  the name of the pet (a string) and a tag  (also a string) that can be used to identify  the pet. This message is used as the "new_pet"  field in the AddPetParameters message, which is  passed as input to the AddPet RPC.
type NewPet struct {
	Name *string `json:"name,omitempty"`
	Tag  *string `json:"tag,omitempty"`
}

// Pet The message Pet is a data structure that  represents a pet in this gRPC function.  It has three fields: name, tag, and id.  The name field is a string that represents  the name of the pet. The tag field is a  string that represents a tag or label  associated with the pet. The id field  is an integer that represents a unique  identifier for the pet. These three  fields are defined as required fields,  meaning that they must be provided when  creating a new Pet object.
type Pet struct {
	// Id The id of the pet
	Id *int64 `json:"id,omitempty"`

	// Name The name of the pet
	Name *string `json:"name,omitempty"`

	// Tag The tag of the pet
	Tag *string `json:"tag,omitempty"`
}

// PingResponse defines model for PingResponse.
type PingResponse struct {
	Pong *Pong `json:"pong,omitempty"`
}

// Pong defines model for Pong.
type Pong struct {
	Index      *int32     `json:"index,omitempty"`
	Message    *string    `json:"message,omitempty"`
	ReceivedOn *time.Time `json:"receivedOn,omitempty"`
	Ver        *string    `json:"ver,omitempty"`
}

// Status The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	// Code The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code *int32 `json:"code,omitempty"`

	// Details A list of messages that carry the error details.  There is a common set of message types for APIs to use.
	Details *[]GoogleProtobufAny `json:"details,omitempty"`

	// Message A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message *string `json:"message,omitempty"`
}

// AitrailblazerServiceEchoParams defines parameters for AitrailblazerServiceEcho.
type AitrailblazerServiceEchoParams struct {
	Value *string `form:"value,omitempty" json:"value,omitempty"`
}

// AitrailblazerServiceFindPetsParams defines parameters for AitrailblazerServiceFindPets.
type AitrailblazerServiceFindPetsParams struct {
	Tags  *[]string `form:"tags,omitempty" json:"tags,omitempty"`
	Limit *int32    `form:"limit,omitempty" json:"limit,omitempty"`
}

// AitrailblazerServiceSendPingParams defines parameters for AitrailblazerServiceSendPing.
type AitrailblazerServiceSendPingParams struct {
	Message *string `form:"message,omitempty" json:"message,omitempty"`
}

// AitrailblazerServiceAddPetJSONRequestBody defines body for AitrailblazerServiceAddPet for application/json ContentType.
type AitrailblazerServiceAddPetJSONRequestBody = NewPet

// Getter for additional properties for GoogleProtobufAny. Returns the specified
// element and whether it was found
func (a GoogleProtobufAny) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GoogleProtobufAny
func (a *GoogleProtobufAny) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GoogleProtobufAny to handle AdditionalProperties
func (a *GoogleProtobufAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["@type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading '@type': %w", err)
		}
		delete(object, "@type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshalling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for GoogleProtobufAny to handle AdditionalProperties
func (a GoogleProtobufAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Type != nil {
		object["@type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '@type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
