# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/service.proto](#api_v1_service-proto)
    - [AddPetParameters](#aitrailblazer-service-v1-AddPetParameters)
    - [DeletePetParameters](#aitrailblazer-service-v1-DeletePetParameters)
    - [EchoMessage](#aitrailblazer-service-v1-EchoMessage)
    - [Error](#aitrailblazer-service-v1-Error)
    - [FindPetByIDParameters](#aitrailblazer-service-v1-FindPetByIDParameters)
    - [FindPetsParameters](#aitrailblazer-service-v1-FindPetsParameters)
    - [ListPetsRequest](#aitrailblazer-service-v1-ListPetsRequest)
    - [ListPetsResponse](#aitrailblazer-service-v1-ListPetsResponse)
    - [NewPet](#aitrailblazer-service-v1-NewPet)
    - [Pet](#aitrailblazer-service-v1-Pet)
    - [PingRequest](#aitrailblazer-service-v1-PingRequest)
    - [PingResponse](#aitrailblazer-service-v1-PingResponse)
    - [Pong](#aitrailblazer-service-v1-Pong)
  
    - [AitrailblazerService](#aitrailblazer-service-v1-AitrailblazerService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/service.proto



<a name="aitrailblazer-service-v1-AddPetParameters"></a>

### AddPetParameters
The AddPetParameters message is used as input for 
the AddPet RPC. It contains a field called 
&#34;new_pet&#34; of type NewPet, which holds the 
information about the new pet being added to the 
system. This includes the name and tag of the pet.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| new_pet | [NewPet](#aitrailblazer-service-v1-NewPet) |  |  |






<a name="aitrailblazer-service-v1-DeletePetParameters"></a>

### DeletePetParameters
The Error message is used to communicate any
 errors that may have occurred during the 
delete operation. It contains a code field, 
which is an integer representing the error 
code, and a message field, which is a string 
describing the error in more detail. 
This message is returned by the DeletePet 
RPC if there is an issue with deleting the 
specified pet from the database.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="aitrailblazer-service-v1-EchoMessage"></a>

### EchoMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="aitrailblazer-service-v1-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  |  |
| message | [string](#string) |  |  |






<a name="aitrailblazer-service-v1-FindPetByIDParameters"></a>

### FindPetByIDParameters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="aitrailblazer-service-v1-FindPetsParameters"></a>

### FindPetsParameters
The FindPetsParameters message contains the 
parameters for the FindPets RPC. It has two 
fields: tags and limit. The tags field is a 
repeated string field, which means it can 
contain multiple strings. These strings 
represent the tags used to search for pets. 
The limit field is an integer field that 
represents the maximum number of pets to be 
returned in the result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [string](#string) | repeated |  |
| limit | [int32](#int32) |  |  |






<a name="aitrailblazer-service-v1-ListPetsRequest"></a>

### ListPetsRequest
The parent, which owns this collection of books.
Format: publishers/{publisher}
string parent = 1 [
  (google.api.field_behavior) = REQUIRED,
  (google.api.resource_reference) = {
    child_type: &#34;library.googleapis.com/Book&#34;
  }];


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [int32](#int32) |  | The maximum number of books to return. The service may return fewer than this value. If unspecified, at most 50 books will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000. |
| page_token | [string](#string) |  | A page token, received from a previous `ListBooks` call. Provide this to retrieve the subsequent page.

When paginating, all other parameters provided to `ListBooks` must match the call that provided the page token. |






<a name="aitrailblazer-service-v1-ListPetsResponse"></a>

### ListPetsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pets | [Pet](#aitrailblazer-service-v1-Pet) | repeated | The Pets. |
| next_page_token | [string](#string) |  | A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages. |






<a name="aitrailblazer-service-v1-NewPet"></a>

### NewPet
The NewPet message contains information about a 
new pet being added to the system. It includes 
the name of the pet (a string) and a tag 
(also a string) that can be used to identify 
the pet. This message is used as the &#34;new_pet&#34; 
field in the AddPetParameters message, which is 
passed as input to the AddPet RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| tag | [string](#string) |  |  |






<a name="aitrailblazer-service-v1-Pet"></a>

### Pet
The message Pet is a data structure that 
represents a pet in this gRPC function. 
It has three fields: name, tag, and id. 
The name field is a string that represents 
the name of the pet. The tag field is a 
string that represents a tag or label 
associated with the pet. The id field 
is an integer that represents a unique 
identifier for the pet. These three 
fields are defined as required fields, 
meaning that they must be provided when 
creating a new Pet object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the pet |
| tag | [string](#string) |  | The tag of the pet |
| id | [int64](#int64) |  | The id of the pet |






<a name="aitrailblazer-service-v1-PingRequest"></a>

### PingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="aitrailblazer-service-v1-PingResponse"></a>

### PingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pong | [Pong](#aitrailblazer-service-v1-Pong) |  |  |






<a name="aitrailblazer-service-v1-Pong"></a>

### Pong



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [int32](#int32) |  | index |
| message | [string](#string) |  | message |
| ver | [string](#string) |  | version |
| received_on | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | received_on |





 

 

 


<a name="aitrailblazer-service-v1-AitrailblazerService"></a>

### AitrailblazerService
This API represents ...

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendPing | [PingRequest](#aitrailblazer-service-v1-PingRequest) | [PingResponse](#aitrailblazer-service-v1-PingResponse) | SendPing |
| Echo | [EchoMessage](#aitrailblazer-service-v1-EchoMessage) | [EchoMessage](#aitrailblazer-service-v1-EchoMessage) | echo |
| FindPets | [FindPetsParameters](#aitrailblazer-service-v1-FindPetsParameters) | [Pet](#aitrailblazer-service-v1-Pet) | FindPets

The FindPets RPC (Remote Procedure Call) allows a client to search for pets based on certain parameters. The parameters include tags (such as breed, color, etc.) and a limit on the number of pets to be returned. The RPC uses a GET HTTP request to the &#34;/pets&#34; endpoint to retrieve the desired pets. The result of the RPC is a Pet message, which contains information about the found pets. |
| AddPet | [AddPetParameters](#aitrailblazer-service-v1-AddPetParameters) | [Pet](#aitrailblazer-service-v1-Pet) | AddPet

The AddPet RPC allows for the addition of a new pet to the system. It takes in AddPetParameters as input, which includes a NewPet message that contains the name and tag of the new pet. The RPC returns a Pet message, which likely includes information about the newly added pet. The RPC can be accessed via an HTTP POST request to the &#34;/pets&#34; endpoint, with the new pet information being sent in the body of the request.

Example #1: name = &#34;Spot&#34;, tag = &#34;TagOfSpot&#34;

Example #2: name = &#34;Fido&#34;, tag = &#34;TagOfFido&#34;

Example #3: name = &#34;Rover&#34;, tag = &#34;TagOfRover&#34;

Example #4: name = &#34;Spike&#34;, tag = &#34;TagOfSpike&#34;

Example #5: name = &#34;Buddy&#34;, tag = &#34;TagOfBuddy&#34; |
| FindPetByID | [FindPetByIDParameters](#aitrailblazer-service-v1-FindPetByIDParameters) | [Pet](#aitrailblazer-service-v1-Pet) | FindPetByID This is a function signature for a Remote Procedure Call (RPC) in the Google API specification language. The function is called &#34;FindPetByID&#34; and it takes in a single input parameter called &#34;FindPetByIDParameters&#34;. The function returns a single output of type &#34;Pet&#34;.

The &#34;option&#34; statement specifies additional options for the function, in this case specifying that the function can be called using an HTTP GET request to the URL &#34;/pets/{id}&#34;, where &#34;{id}&#34; is a placeholder for a specific pet ID.

This function allows a client to retrieve information about a specific pet by providing the pet&#39;s ID as a parameter in the HTTP request.

Returns a pet by ID |
| DeletePet | [DeletePetParameters](#aitrailblazer-service-v1-DeletePetParameters) | [Error](#aitrailblazer-service-v1-Error) | DeletePet

The DeletePet RPC (Remote Procedure Call) is used to delete a pet from a database. It takes in a DeletePetParameters message as input, which contains the ID of the pet to be deleted. The RPC returns an Error message, which includes a code and a message describing any errors that may have occurred during the delete operation. The RPC uses the HTTP DELETE method to delete the pet from the database, with the ID of the pet being specified in the URL as a path parameter. |
| ListPets | [ListPetsRequest](#aitrailblazer-service-v1-ListPetsRequest) | [ListPetsResponse](#aitrailblazer-service-v1-ListPetsResponse) | ListPets |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

