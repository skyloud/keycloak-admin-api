# \RoleMapperAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGroupRoleMappingsRealm**](RoleMapperAPI.md#DeleteGroupRoleMappingsRealm) | **Delete** /{realm}/groups/{id}/role-mappings/realm | 
[**DeleteUserRoleMappingsRealm**](RoleMapperAPI.md#DeleteUserRoleMappingsRealm) | **Delete** /{realm}/users/{id}/role-mappings/realm | 
[**GetGroupRoleMappings**](RoleMapperAPI.md#GetGroupRoleMappings) | **Get** /{realm}/groups/{id}/role-mappings | 
[**GetGroupRoleMappingsRealm**](RoleMapperAPI.md#GetGroupRoleMappingsRealm) | **Get** /{realm}/groups/{id}/role-mappings/realm | 
[**GetGroupRoleMappingsRealmAvailable**](RoleMapperAPI.md#GetGroupRoleMappingsRealmAvailable) | **Get** /{realm}/groups/{id}/role-mappings/realm/available | 
[**GetGroupRoleMappingsRealmComposite**](RoleMapperAPI.md#GetGroupRoleMappingsRealmComposite) | **Get** /{realm}/groups/{id}/role-mappings/realm/composite | 
[**GetUserRoleMappings**](RoleMapperAPI.md#GetUserRoleMappings) | **Get** /{realm}/users/{id}/role-mappings | 
[**GetUserRoleMappingsRealm**](RoleMapperAPI.md#GetUserRoleMappingsRealm) | **Get** /{realm}/users/{id}/role-mappings/realm | 
[**GetUserRoleMappingsRealmAvailable**](RoleMapperAPI.md#GetUserRoleMappingsRealmAvailable) | **Get** /{realm}/users/{id}/role-mappings/realm/available | 
[**GetUserRoleMappingsRealmComposite**](RoleMapperAPI.md#GetUserRoleMappingsRealmComposite) | **Get** /{realm}/users/{id}/role-mappings/realm/composite | 
[**PostGroupRoleMappingsRealm**](RoleMapperAPI.md#PostGroupRoleMappingsRealm) | **Post** /{realm}/groups/{id}/role-mappings/realm | 
[**PostUserRoleMappingsRealm**](RoleMapperAPI.md#PostUserRoleMappingsRealm) | **Post** /{realm}/users/{id}/role-mappings/realm | 



## DeleteGroupRoleMappingsRealm

> DeleteGroupRoleMappingsRealm(ctx, realm, id).RoleRepresentation(roleRepresentation).Execute()





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
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.DeleteGroupRoleMappingsRealm(context.Background(), realm, id).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.DeleteGroupRoleMappingsRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGroupRoleMappingsRealmRequest struct via the builder pattern


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


## DeleteUserRoleMappingsRealm

> DeleteUserRoleMappingsRealm(ctx, realm, id).RoleRepresentation(roleRepresentation).Execute()





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
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.DeleteUserRoleMappingsRealm(context.Background(), realm, id).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.DeleteUserRoleMappingsRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserRoleMappingsRealmRequest struct via the builder pattern


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


## GetGroupRoleMappings

> MappingsRepresentation GetGroupRoleMappings(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.RoleMapperAPI.GetGroupRoleMappings(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetGroupRoleMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleMappings`: MappingsRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetGroupRoleMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappingsRepresentation**](MappingsRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRoleMappingsRealm

> []RoleRepresentation GetGroupRoleMappingsRealm(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.RoleMapperAPI.GetGroupRoleMappingsRealm(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetGroupRoleMappingsRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleMappingsRealm`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetGroupRoleMappingsRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleMappingsRealmRequest struct via the builder pattern


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


## GetGroupRoleMappingsRealmAvailable

> []RoleRepresentation GetGroupRoleMappingsRealmAvailable(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.RoleMapperAPI.GetGroupRoleMappingsRealmAvailable(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetGroupRoleMappingsRealmAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleMappingsRealmAvailable`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetGroupRoleMappingsRealmAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleMappingsRealmAvailableRequest struct via the builder pattern


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


## GetGroupRoleMappingsRealmComposite

> []RoleRepresentation GetGroupRoleMappingsRealmComposite(ctx, realm, id).BriefRepresentation(briefRepresentation).Execute()





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
	briefRepresentation := "briefRepresentation_example" // string | if false, return roles with their attributes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMapperAPI.GetGroupRoleMappingsRealmComposite(context.Background(), realm, id).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetGroupRoleMappingsRealmComposite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleMappingsRealmComposite`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetGroupRoleMappingsRealmComposite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleMappingsRealmCompositeRequest struct via the builder pattern


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


## GetUserRoleMappings

> MappingsRepresentation GetUserRoleMappings(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.RoleMapperAPI.GetUserRoleMappings(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetUserRoleMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRoleMappings`: MappingsRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetUserRoleMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappingsRepresentation**](MappingsRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRoleMappingsRealm

> []RoleRepresentation GetUserRoleMappingsRealm(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.RoleMapperAPI.GetUserRoleMappingsRealm(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetUserRoleMappingsRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRoleMappingsRealm`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetUserRoleMappingsRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleMappingsRealmRequest struct via the builder pattern


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


## GetUserRoleMappingsRealmAvailable

> []RoleRepresentation GetUserRoleMappingsRealmAvailable(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.RoleMapperAPI.GetUserRoleMappingsRealmAvailable(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetUserRoleMappingsRealmAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRoleMappingsRealmAvailable`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetUserRoleMappingsRealmAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleMappingsRealmAvailableRequest struct via the builder pattern


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


## GetUserRoleMappingsRealmComposite

> []RoleRepresentation GetUserRoleMappingsRealmComposite(ctx, realm, id).BriefRepresentation(briefRepresentation).Execute()





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
	briefRepresentation := "briefRepresentation_example" // string | if false, return roles with their attributes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMapperAPI.GetUserRoleMappingsRealmComposite(context.Background(), realm, id).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.GetUserRoleMappingsRealmComposite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRoleMappingsRealmComposite`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.GetUserRoleMappingsRealmComposite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleMappingsRealmCompositeRequest struct via the builder pattern


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


## PostGroupRoleMappingsRealm

> PostGroupRoleMappingsRealm(ctx, realm, id).RoleRepresentation(roleRepresentation).Execute()





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
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.PostGroupRoleMappingsRealm(context.Background(), realm, id).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.PostGroupRoleMappingsRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostGroupRoleMappingsRealmRequest struct via the builder pattern


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


## PostUserRoleMappingsRealm

> PostUserRoleMappingsRealm(ctx, realm, id).RoleRepresentation(roleRepresentation).Execute()





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
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.PostUserRoleMappingsRealm(context.Background(), realm, id).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.PostUserRoleMappingsRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUserRoleMappingsRealmRequest struct via the builder pattern


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

