# \RolesByIDAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRolesById**](RolesByIDAPI.md#DeleteRolesById) | **Delete** /{realm}/roles-by-id/{role-id} | 
[**DeleteRolesByIdComposites**](RolesByIDAPI.md#DeleteRolesByIdComposites) | **Delete** /{realm}/roles-by-id/{role-id}/composites | 
[**GetRolesById**](RolesByIDAPI.md#GetRolesById) | **Get** /{realm}/roles-by-id/{role-id} | 
[**GetRolesByIdComposites**](RolesByIDAPI.md#GetRolesByIdComposites) | **Get** /{realm}/roles-by-id/{role-id}/composites | 
[**GetRolesByIdCompositesClient**](RolesByIDAPI.md#GetRolesByIdCompositesClient) | **Get** /{realm}/roles-by-id/{role-id}/composites/clients/{clientUuid} | 
[**GetRolesByIdCompositesRealm**](RolesByIDAPI.md#GetRolesByIdCompositesRealm) | **Get** /{realm}/roles-by-id/{role-id}/composites/realm | 
[**GetRolesByIdManagementPermissions**](RolesByIDAPI.md#GetRolesByIdManagementPermissions) | **Get** /{realm}/roles-by-id/{role-id}/management/permissions | 
[**PostRolesByIdComposites**](RolesByIDAPI.md#PostRolesByIdComposites) | **Post** /{realm}/roles-by-id/{role-id}/composites | 
[**PutRolesById**](RolesByIDAPI.md#PutRolesById) | **Put** /{realm}/roles-by-id/{role-id} | 
[**PutRolesByIdManagementPermissions**](RolesByIDAPI.md#PutRolesByIdManagementPermissions) | **Put** /{realm}/roles-by-id/{role-id}/management/permissions | 



## DeleteRolesById

> DeleteRolesById(ctx, realm, roleId).Execute()





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
	roleId := "roleId_example" // string | id of role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.DeleteRolesById(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.DeleteRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** | id of role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRolesByIdRequest struct via the builder pattern


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


## DeleteRolesByIdComposites

> DeleteRolesByIdComposites(ctx, realm, roleId).RoleRepresentation(roleRepresentation).Execute()





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
	roleId := "roleId_example" // string | 
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.DeleteRolesByIdComposites(context.Background(), realm, roleId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.DeleteRolesByIdComposites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRolesByIdCompositesRequest struct via the builder pattern


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


## GetRolesById

> RoleRepresentation GetRolesById(ctx, realm, roleId).Execute()





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
	roleId := "roleId_example" // string | id of role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.GetRolesById(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.GetRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolesById`: RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.GetRolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** | id of role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleRepresentation**](RoleRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRolesByIdComposites

> []RoleRepresentation GetRolesByIdComposites(ctx, realm, roleId).First(first).Max(max).Search(search).Execute()





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
	roleId := "roleId_example" // string | 
	first := "first_example" // string |  (optional)
	max := "max_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.GetRolesByIdComposites(context.Background(), realm, roleId).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.GetRolesByIdComposites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolesByIdComposites`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.GetRolesByIdComposites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesByIdCompositesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **first** | **string** |  | 
 **max** | **string** |  | 
 **search** | **string** |  | 

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


## GetRolesByIdCompositesClient

> []RoleRepresentation GetRolesByIdCompositesClient(ctx, realm, roleId, clientUuid).Execute()





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
	roleId := "roleId_example" // string | 
	clientUuid := "clientUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.GetRolesByIdCompositesClient(context.Background(), realm, roleId, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.GetRolesByIdCompositesClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolesByIdCompositesClient`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.GetRolesByIdCompositesClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 
**clientUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesByIdCompositesClientRequest struct via the builder pattern


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


## GetRolesByIdCompositesRealm

> []RoleRepresentation GetRolesByIdCompositesRealm(ctx, realm, roleId).Execute()





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
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.GetRolesByIdCompositesRealm(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.GetRolesByIdCompositesRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolesByIdCompositesRealm`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.GetRolesByIdCompositesRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesByIdCompositesRealmRequest struct via the builder pattern


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


## GetRolesByIdManagementPermissions

> ManagementPermissionReference GetRolesByIdManagementPermissions(ctx, realm, roleId).Execute()





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
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.GetRolesByIdManagementPermissions(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.GetRolesByIdManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolesByIdManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.GetRolesByIdManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesByIdManagementPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManagementPermissionReference**](ManagementPermissionReference.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRolesByIdComposites

> PostRolesByIdComposites(ctx, realm, roleId).RoleRepresentation(roleRepresentation).Execute()





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
	roleId := "roleId_example" // string | 
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.PostRolesByIdComposites(context.Background(), realm, roleId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.PostRolesByIdComposites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRolesByIdCompositesRequest struct via the builder pattern


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


## PutRolesById

> PutRolesById(ctx, realm, roleId).RoleRepresentation(roleRepresentation).Execute()





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
	roleId := "roleId_example" // string | id of role
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.PutRolesById(context.Background(), realm, roleId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.PutRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** | id of role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRolesByIdRequest struct via the builder pattern


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


## PutRolesByIdManagementPermissions

> ManagementPermissionReference PutRolesByIdManagementPermissions(ctx, realm, roleId).ManagementPermissionReference(managementPermissionReference).Execute()





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
	roleId := "roleId_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference | ManagementPermissionReference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.PutRolesByIdManagementPermissions(context.Background(), realm, roleId).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.PutRolesByIdManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRolesByIdManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.PutRolesByIdManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRolesByIdManagementPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managementPermissionReference** | [**ManagementPermissionReference**](ManagementPermissionReference.md) | ManagementPermissionReference | 

### Return type

[**ManagementPermissionReference**](ManagementPermissionReference.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

