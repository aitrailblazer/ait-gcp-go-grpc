# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/service.proto](#api_v1_service-proto)
    - [GetShelfRequest](#aitrailblazer-service-v1-GetShelfRequest)
    - [PingRequest](#aitrailblazer-service-v1-PingRequest)
    - [PingResponse](#aitrailblazer-service-v1-PingResponse)
    - [Pong](#aitrailblazer-service-v1-Pong)
    - [Shelf](#aitrailblazer-service-v1-Shelf)
  
    - [AitrailblazerService](#aitrailblazer-service-v1-AitrailblazerService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/service.proto



<a name="aitrailblazer-service-v1-GetShelfRequest"></a>

### GetShelfRequest
Request message for GetShelf method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shelf | [int64](#int64) |  | The ID of the shelf resource to retrieve. |






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






<a name="aitrailblazer-service-v1-Shelf"></a>

### Shelf
A shelf resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | A unique shelf id. |
| theme | [string](#string) |  | A theme of the shelf (fiction, poetry, etc). |





 

 

 


<a name="aitrailblazer-service-v1-AitrailblazerService"></a>

### AitrailblazerService
## API Overview

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendPing | [PingRequest](#aitrailblazer-service-v1-PingRequest) | [PingResponse](#aitrailblazer-service-v1-PingResponse) |  |
| GetShelf | [GetShelfRequest](#aitrailblazer-service-v1-GetShelfRequest) | [Shelf](#aitrailblazer-service-v1-Shelf) | rpc GetBook(GetBookRequest) returns (Book) { option (google.api.http) = { get: &#34;/v1/{name=publishers/*/books/*}&#34; }; option (google.api.method_signature) = &#34;name&#34;; } Returns a specific bookstore shelf. |

 



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

