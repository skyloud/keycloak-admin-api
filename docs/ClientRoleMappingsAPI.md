# \ClientRoleMappingsAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGroupRoleMappingsClient**](ClientRoleMappingsAPI.md#DeleteGroupRoleMappingsClient) | **Delete** /{realm}/groups/{id}/role-mappings/clients/{client} | 
[**DeleteUserRoleMappingsClient**](ClientRoleMappingsAPI.md#DeleteUserRoleMappingsClient) | **Delete** /{realm}/users/{id}/role-mappings/clients/{client} | 
[**GetGroupRoleMappingsClient**](ClientRoleMappingsAPI.md#GetGroupRoleMappingsClient) | **Get** /{realm}/groups/{id}/role-mappings/clients/{client} | 
[**GetGroupRoleMappingsClientAvailable**](ClientRoleMappingsAPI.md#GetGroupRoleMappingsClientAvailable) | **Get** /{realm}/groups/{id}/role-mappings/clients/{client}/available | 
[**GetGroupRoleMappingsClientComposite**](ClientRoleMappingsAPI.md#GetGroupRoleMappingsClientComposite) | **Get** /{realm}/groups/{id}/role-mappings/clients/{client}/composite | 
[**GetUserRoleMappingsClient**](ClientRoleMappingsAPI.md#GetUserRoleMappingsClient) | **Get** /{realm}/users/{id}/role-mappings/clients/{client} | 
[**GetUserRoleMappingsClientAvailable**](ClientRoleMappingsAPI.md#GetUserRoleMappingsClientAvailable) | **Get** /{realm}/users/{id}/role-mappings/clients/{client}/available | 
[**GetUserRoleMappingsClientComposite**](ClientRoleMappingsAPI.md#GetUserRoleMappingsClientComposite) | **Get** /{realm}/users/{id}/role-mappings/clients/{client}/composite | 
[**PostGroupRoleMappingsClient**](ClientRoleMappingsAPI.md#PostGroupRoleMappingsClient) | **Post** /{realm}/groups/{id}/role-mappings/clients/{client} | 
[**PostUserRoleMappingsClient**](ClientRoleMappingsAPI.md#PostUserRoleMappingsClient) | **Post** /{realm}/users/{id}/role-mappings/clients/{client} | 



## DeleteGroupRoleMappingsClient

> DeleteGroupRoleMappingsClient(ctx, realm, id, client).RoleRepresentation(roleRepresentation).Execute()





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
	client := "client_example" // string | 
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.DeleteGroupRoleMappingsClient(context.Background(), realm, id, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.DeleteGroupRoleMappingsClient``: %v\n", err)
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
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRoleMappingsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) | RoleRepresentation | 

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


## DeleteUserRoleMappingsClient

> DeleteUserRoleMappingsClient(ctx, realm, id, client).RoleRepresentation(roleRepresentation).Execute()





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
	client := "client_example" // string | 
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.DeleteUserRoleMappingsClient(context.Background(), realm, id, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.DeleteUserRoleMappingsClient``: %v\n", err)
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
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRoleMappingsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) | RoleRepresentation | 

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


## GetGroupRoleMappingsClient

> []RoleRepresentation GetGroupRoleMappingsClient(ctx, realm, id, client).Execute()





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
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.GetGroupRoleMappingsClient(context.Background(), realm, id, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.GetGroupRoleMappingsClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleMappingsClient`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.GetGroupRoleMappingsClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleMappingsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRoleMappingsClientAvailable

> []RoleRepresentation GetGroupRoleMappingsClientAvailable(ctx, realm, id, client).Execute()





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
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.GetGroupRoleMappingsClientAvailable(context.Background(), realm, id, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.GetGroupRoleMappingsClientAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleMappingsClientAvailable`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.GetGroupRoleMappingsClientAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleMappingsClientAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRoleMappingsClientComposite

> []RoleRepresentation GetGroupRoleMappingsClientComposite(ctx, realm, id, client).BriefRepresentation(briefRepresentation).Execute()





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
	client := "client_example" // string | 
	briefRepresentation := "briefRepresentation_example" // string | if false, return roles with their attributes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.GetGroupRoleMappingsClientComposite(context.Background(), realm, id, client).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.GetGroupRoleMappingsClientComposite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleMappingsClientComposite`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.GetGroupRoleMappingsClientComposite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleMappingsClientCompositeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **string** | if false, return roles with their attributes | 

### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRoleMappingsClient

> []RoleRepresentation GetUserRoleMappingsClient(ctx, realm, id, client).Execute()





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
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.GetUserRoleMappingsClient(context.Background(), realm, id, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.GetUserRoleMappingsClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRoleMappingsClient`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.GetUserRoleMappingsClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleMappingsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRoleMappingsClientAvailable

> []RoleRepresentation GetUserRoleMappingsClientAvailable(ctx, realm, id, client).Execute()





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
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.GetUserRoleMappingsClientAvailable(context.Background(), realm, id, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.GetUserRoleMappingsClientAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRoleMappingsClientAvailable`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.GetUserRoleMappingsClientAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleMappingsClientAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRoleMappingsClientComposite

> []RoleRepresentation GetUserRoleMappingsClientComposite(ctx, realm, id, client).BriefRepresentation(briefRepresentation).Execute()





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
	client := "client_example" // string | 
	briefRepresentation := "briefRepresentation_example" // string | if false, return roles with their attributes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.GetUserRoleMappingsClientComposite(context.Background(), realm, id, client).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.GetUserRoleMappingsClientComposite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRoleMappingsClientComposite`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.GetUserRoleMappingsClientComposite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleMappingsClientCompositeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **string** | if false, return roles with their attributes | 

### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGroupRoleMappingsClient

> PostGroupRoleMappingsClient(ctx, realm, id, client).RoleRepresentation(roleRepresentation).Execute()





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
	client := "client_example" // string | 
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.PostGroupRoleMappingsClient(context.Background(), realm, id, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.PostGroupRoleMappingsClient``: %v\n", err)
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
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGroupRoleMappingsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) | RoleRepresentation | 

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


## PostUserRoleMappingsClient

> PostUserRoleMappingsClient(ctx, realm, id, client).RoleRepresentation(roleRepresentation).Execute()





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
	client := "client_example" // string | 
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.PostUserRoleMappingsClient(context.Background(), realm, id, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.PostUserRoleMappingsClient``: %v\n", err)
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
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUserRoleMappingsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) | RoleRepresentation | 

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

