# \ClientInitialAccessAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientsInitialAcces**](ClientInitialAccessAPI.md#DeleteClientsInitialAcces) | **Delete** /{realm}/clients-initial-access/{id} | 
[**GetClientsInitialAccess**](ClientInitialAccessAPI.md#GetClientsInitialAccess) | **Get** /{realm}/clients-initial-access | 
[**PostClientsInitialAccess**](ClientInitialAccessAPI.md#PostClientsInitialAccess) | **Post** /{realm}/clients-initial-access | 



## DeleteClientsInitialAcces

> DeleteClientsInitialAcces(ctx, realm, id).Execute()



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
	r, err := apiClient.ClientInitialAccessAPI.DeleteClientsInitialAcces(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInitialAccessAPI.DeleteClientsInitialAcces``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteClientsInitialAccesRequest struct via the builder pattern


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


## GetClientsInitialAccess

> []ClientInitialAccessPresentation GetClientsInitialAccess(ctx, realm).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInitialAccessAPI.GetClientsInitialAccess(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInitialAccessAPI.GetClientsInitialAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsInitialAccess`: []ClientInitialAccessPresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientInitialAccessAPI.GetClientsInitialAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsInitialAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ClientInitialAccessPresentation**](ClientInitialAccessPresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientsInitialAccess

> ClientInitialAccessPresentation PostClientsInitialAccess(ctx, realm).ClientInitialAccessCreatePresentation(clientInitialAccessCreatePresentation).Execute()





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
	clientInitialAccessCreatePresentation := *openapiclient.NewClientInitialAccessCreatePresentation() // ClientInitialAccessCreatePresentation | ClientInitialAccessCreatePresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInitialAccessAPI.PostClientsInitialAccess(context.Background(), realm).ClientInitialAccessCreatePresentation(clientInitialAccessCreatePresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInitialAccessAPI.PostClientsInitialAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientsInitialAccess`: ClientInitialAccessPresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientInitialAccessAPI.PostClientsInitialAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientsInitialAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientInitialAccessCreatePresentation** | [**ClientInitialAccessCreatePresentation**](ClientInitialAccessCreatePresentation.md) | ClientInitialAccessCreatePresentation | 

### Return type

[**ClientInitialAccessPresentation**](ClientInitialAccessPresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

