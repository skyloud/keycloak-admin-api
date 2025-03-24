# \ComponentAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponent**](ComponentAPI.md#DeleteComponent) | **Delete** /{realm}/components/{id} | 
[**GetComponent**](ComponentAPI.md#GetComponent) | **Get** /{realm}/components/{id} | 
[**GetComponents**](ComponentAPI.md#GetComponents) | **Get** /{realm}/components | 
[**GetSubComponentTypes**](ComponentAPI.md#GetSubComponentTypes) | **Get** /{realm}/components/{id}/sub-component-types | 
[**PostComponents**](ComponentAPI.md#PostComponents) | **Post** /{realm}/components | 
[**PutComponent**](ComponentAPI.md#PutComponent) | **Put** /{realm}/components/{id} | 



## DeleteComponent

> DeleteComponent(ctx, realm, id).Execute()



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
	r, err := apiClient.ComponentAPI.DeleteComponent(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.DeleteComponent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteComponentRequest struct via the builder pattern


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


## GetComponent

> ComponentRepresentation GetComponent(ctx, realm, id).Execute()



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
	resp, r, err := apiClient.ComponentAPI.GetComponent(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.GetComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponent`: ComponentRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.GetComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ComponentRepresentation**](ComponentRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponents

> []ComponentRepresentation GetComponents(ctx, realm).Name(name).Parent(parent).Type_(type_).Execute()



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
	name := "name_example" // string |  (optional)
	parent := "parent_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentAPI.GetComponents(context.Background(), realm).Name(name).Parent(parent).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.GetComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponents`: []ComponentRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.GetComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **parent** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**[]ComponentRepresentation**](ComponentRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubComponentTypes

> []ComponentTypeRepresentation GetSubComponentTypes(ctx, realm, id).Type_(type_).Execute()





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
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentAPI.GetSubComponentTypes(context.Background(), realm, id).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.GetSubComponentTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubComponentTypes`: []ComponentTypeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.GetSubComponentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubComponentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** |  | 

### Return type

[**[]ComponentTypeRepresentation**](ComponentTypeRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComponents

> PostComponents(ctx, realm).ComponentRepresentation(componentRepresentation).Execute()



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
	componentRepresentation := *openapiclient.NewComponentRepresentation() // ComponentRepresentation | ComponentRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentAPI.PostComponents(context.Background(), realm).ComponentRepresentation(componentRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.PostComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentRepresentation** | [**ComponentRepresentation**](ComponentRepresentation.md) | ComponentRepresentation | 

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


## PutComponent

> PutComponent(ctx, realm, id).ComponentRepresentation(componentRepresentation).Execute()



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
	componentRepresentation := *openapiclient.NewComponentRepresentation() // ComponentRepresentation | ComponentRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentAPI.PutComponent(context.Background(), realm, id).ComponentRepresentation(componentRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.PutComponent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentRepresentation** | [**ComponentRepresentation**](ComponentRepresentation.md) | ComponentRepresentation | 

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

