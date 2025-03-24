# \ProtocolMappersAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientProtocolMappersModel**](ProtocolMappersAPI.md#DeleteClientProtocolMappersModel) | **Delete** /{realm}/clients/{id1}/protocol-mappers/models/{id2} | 
[**DeleteClientScopeProtocolMappersModel**](ProtocolMappersAPI.md#DeleteClientScopeProtocolMappersModel) | **Delete** /{realm}/client-scopes/{id1}/protocol-mappers/models/{id2} | 
[**DeleteClientTemplateProtocolMappersModel**](ProtocolMappersAPI.md#DeleteClientTemplateProtocolMappersModel) | **Delete** /{realm}/client-templates/{id1}/protocol-mappers/models/{id2} | 
[**GetClientProtocolMappersModel**](ProtocolMappersAPI.md#GetClientProtocolMappersModel) | **Get** /{realm}/clients/{id1}/protocol-mappers/models/{id2} | 
[**GetClientProtocolMappersModels**](ProtocolMappersAPI.md#GetClientProtocolMappersModels) | **Get** /{realm}/clients/{id}/protocol-mappers/models | 
[**GetClientProtocolMappersProtocol**](ProtocolMappersAPI.md#GetClientProtocolMappersProtocol) | **Get** /{realm}/clients/{id}/protocol-mappers/protocol/{protocol} | 
[**GetClientScopeProtocolMappersModel**](ProtocolMappersAPI.md#GetClientScopeProtocolMappersModel) | **Get** /{realm}/client-scopes/{id1}/protocol-mappers/models/{id2} | 
[**GetClientScopeProtocolMappersModels**](ProtocolMappersAPI.md#GetClientScopeProtocolMappersModels) | **Get** /{realm}/client-scopes/{id}/protocol-mappers/models | 
[**GetClientScopeProtocolMappersProtocol**](ProtocolMappersAPI.md#GetClientScopeProtocolMappersProtocol) | **Get** /{realm}/client-scopes/{id}/protocol-mappers/protocol/{protocol} | 
[**GetClientTemplateProtocolMappersModel**](ProtocolMappersAPI.md#GetClientTemplateProtocolMappersModel) | **Get** /{realm}/client-templates/{id1}/protocol-mappers/models/{id2} | 
[**GetClientTemplateProtocolMappersModels**](ProtocolMappersAPI.md#GetClientTemplateProtocolMappersModels) | **Get** /{realm}/client-templates/{id}/protocol-mappers/models | 
[**GetClientTemplateProtocolMappersProtocol**](ProtocolMappersAPI.md#GetClientTemplateProtocolMappersProtocol) | **Get** /{realm}/client-templates/{id}/protocol-mappers/protocol/{protocol} | 
[**PostClientProtocolMappersAddModels**](ProtocolMappersAPI.md#PostClientProtocolMappersAddModels) | **Post** /{realm}/clients/{id}/protocol-mappers/add-models | 
[**PostClientProtocolMappersModels**](ProtocolMappersAPI.md#PostClientProtocolMappersModels) | **Post** /{realm}/clients/{id}/protocol-mappers/models | 
[**PostClientScopeProtocolMappersAddModels**](ProtocolMappersAPI.md#PostClientScopeProtocolMappersAddModels) | **Post** /{realm}/client-scopes/{id}/protocol-mappers/add-models | 
[**PostClientScopeProtocolMappersModels**](ProtocolMappersAPI.md#PostClientScopeProtocolMappersModels) | **Post** /{realm}/client-scopes/{id}/protocol-mappers/models | 
[**PostClientTemplateProtocolMappersAddModels**](ProtocolMappersAPI.md#PostClientTemplateProtocolMappersAddModels) | **Post** /{realm}/client-templates/{id}/protocol-mappers/add-models | 
[**PostClientTemplateProtocolMappersModels**](ProtocolMappersAPI.md#PostClientTemplateProtocolMappersModels) | **Post** /{realm}/client-templates/{id}/protocol-mappers/models | 
[**PutClientProtocolMappersModel**](ProtocolMappersAPI.md#PutClientProtocolMappersModel) | **Put** /{realm}/clients/{id1}/protocol-mappers/models/{id2} | 
[**PutClientScopeProtocolMappersModel**](ProtocolMappersAPI.md#PutClientScopeProtocolMappersModel) | **Put** /{realm}/client-scopes/{id1}/protocol-mappers/models/{id2} | 
[**PutClientTemplateProtocolMappersModel**](ProtocolMappersAPI.md#PutClientTemplateProtocolMappersModel) | **Put** /{realm}/client-templates/{id1}/protocol-mappers/models/{id2} | 



## DeleteClientProtocolMappersModel

> DeleteClientProtocolMappersModel(ctx, realm, id1, id2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.DeleteClientProtocolMappersModel(context.Background(), realm, id1, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.DeleteClientProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientScopeProtocolMappersModel

> DeleteClientScopeProtocolMappersModel(ctx, realm, id1, id2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.DeleteClientScopeProtocolMappersModel(context.Background(), realm, id1, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.DeleteClientScopeProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientScopeProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientTemplateProtocolMappersModel

> DeleteClientTemplateProtocolMappersModel(ctx, realm, id1, id2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.DeleteClientTemplateProtocolMappersModel(context.Background(), realm, id1, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.DeleteClientTemplateProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientTemplateProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientProtocolMappersModel

> ProtocolMapperRepresentation GetClientProtocolMappersModel(ctx, realm, id1, id2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientProtocolMappersModel(context.Background(), realm, id1, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientProtocolMappersModel`: ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientProtocolMappersModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientProtocolMappersModels

> []ProtocolMapperRepresentation GetClientProtocolMappersModels(ctx, realm, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientProtocolMappersModels(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientProtocolMappersModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientProtocolMappersModels`: []ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientProtocolMappersModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientProtocolMappersModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientProtocolMappersProtocol

> []ProtocolMapperRepresentation GetClientProtocolMappersProtocol(ctx, realm, id, protocol).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocol := "protocol_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientProtocolMappersProtocol(context.Background(), realm, id, protocol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientProtocolMappersProtocol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientProtocolMappersProtocol`: []ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientProtocolMappersProtocol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**protocol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientProtocolMappersProtocolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientScopeProtocolMappersModel

> ProtocolMapperRepresentation GetClientScopeProtocolMappersModel(ctx, realm, id1, id2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientScopeProtocolMappersModel(context.Background(), realm, id1, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientScopeProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientScopeProtocolMappersModel`: ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientScopeProtocolMappersModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientScopeProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientScopeProtocolMappersModels

> []ProtocolMapperRepresentation GetClientScopeProtocolMappersModels(ctx, realm, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientScopeProtocolMappersModels(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientScopeProtocolMappersModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientScopeProtocolMappersModels`: []ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientScopeProtocolMappersModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientScopeProtocolMappersModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientScopeProtocolMappersProtocol

> []ProtocolMapperRepresentation GetClientScopeProtocolMappersProtocol(ctx, realm, id, protocol).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocol := "protocol_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientScopeProtocolMappersProtocol(context.Background(), realm, id, protocol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientScopeProtocolMappersProtocol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientScopeProtocolMappersProtocol`: []ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientScopeProtocolMappersProtocol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**protocol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientScopeProtocolMappersProtocolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientTemplateProtocolMappersModel

> ProtocolMapperRepresentation GetClientTemplateProtocolMappersModel(ctx, realm, id1, id2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientTemplateProtocolMappersModel(context.Background(), realm, id1, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientTemplateProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientTemplateProtocolMappersModel`: ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientTemplateProtocolMappersModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientTemplateProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientTemplateProtocolMappersModels

> []ProtocolMapperRepresentation GetClientTemplateProtocolMappersModels(ctx, realm, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientTemplateProtocolMappersModels(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientTemplateProtocolMappersModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientTemplateProtocolMappersModels`: []ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientTemplateProtocolMappersModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientTemplateProtocolMappersModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientTemplateProtocolMappersProtocol

> []ProtocolMapperRepresentation GetClientTemplateProtocolMappersProtocol(ctx, realm, id, protocol).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocol := "protocol_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolMappersAPI.GetClientTemplateProtocolMappersProtocol(context.Background(), realm, id, protocol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.GetClientTemplateProtocolMappersProtocol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientTemplateProtocolMappersProtocol`: []ProtocolMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ProtocolMappersAPI.GetClientTemplateProtocolMappersProtocol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**protocol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientTemplateProtocolMappersProtocolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientProtocolMappersAddModels

> PostClientProtocolMappersAddModels(ctx, realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PostClientProtocolMappersAddModels(context.Background(), realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PostClientProtocolMappersAddModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientProtocolMappersAddModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientProtocolMappersModels

> PostClientProtocolMappersModels(ctx, realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PostClientProtocolMappersModels(context.Background(), realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PostClientProtocolMappersModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientProtocolMappersModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientScopeProtocolMappersAddModels

> PostClientScopeProtocolMappersAddModels(ctx, realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PostClientScopeProtocolMappersAddModels(context.Background(), realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PostClientScopeProtocolMappersAddModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientScopeProtocolMappersAddModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientScopeProtocolMappersModels

> PostClientScopeProtocolMappersModels(ctx, realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PostClientScopeProtocolMappersModels(context.Background(), realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PostClientScopeProtocolMappersModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientScopeProtocolMappersModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientTemplateProtocolMappersAddModels

> PostClientTemplateProtocolMappersAddModels(ctx, realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PostClientTemplateProtocolMappersAddModels(context.Background(), realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PostClientTemplateProtocolMappersAddModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientTemplateProtocolMappersAddModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientTemplateProtocolMappersModels

> PostClientTemplateProtocolMappersModels(ctx, realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PostClientTemplateProtocolMappersModels(context.Background(), realm, id).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PostClientTemplateProtocolMappersModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientTemplateProtocolMappersModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientProtocolMappersModel

> PutClientProtocolMappersModel(ctx, realm, id1, id2).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PutClientProtocolMappersModel(context.Background(), realm, id1, id2).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PutClientProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientScopeProtocolMappersModel

> PutClientScopeProtocolMappersModel(ctx, realm, id1, id2).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PutClientScopeProtocolMappersModel(context.Background(), realm, id1, id2).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PutClientScopeProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientScopeProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientTemplateProtocolMappersModel

> PutClientTemplateProtocolMappersModel(ctx, realm, id1, id2).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id1 := "id1_example" // string | 
	id2 := "id2_example" // string | 
	protocolMapperRepresentation := *openapiclient.NewProtocolMapperRepresentation() // ProtocolMapperRepresentation | ProtocolMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProtocolMappersAPI.PutClientTemplateProtocolMappersModel(context.Background(), realm, id1, id2).ProtocolMapperRepresentation(protocolMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMappersAPI.PutClientTemplateProtocolMappersModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id1** | **string** |  | 
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientTemplateProtocolMappersModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **protocolMapperRepresentation** | [**ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) | ProtocolMapperRepresentation | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

