# \RolesAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientRole**](RolesAPI.md#DeleteClientRole) | **Delete** /{realm}/clients/{id}/roles/{role-name} | 
[**DeleteClientRoleComposites**](RolesAPI.md#DeleteClientRoleComposites) | **Delete** /{realm}/clients/{id}/roles/{role-name}/composites | 
[**DeleteRoleByRealmByRoleName**](RolesAPI.md#DeleteRoleByRealmByRoleName) | **Delete** /{realm}/roles/{role-name} | 
[**DeleteRoleCompositesByRealmByRoleName**](RolesAPI.md#DeleteRoleCompositesByRealmByRoleName) | **Delete** /{realm}/roles/{role-name}/composites | 
[**GetClientRole**](RolesAPI.md#GetClientRole) | **Get** /{realm}/clients/{id}/roles/{role-name} | 
[**GetClientRoleComposites**](RolesAPI.md#GetClientRoleComposites) | **Get** /{realm}/clients/{id}/roles/{role-name}/composites | 
[**GetClientRoleCompositesClient**](RolesAPI.md#GetClientRoleCompositesClient) | **Get** /{realm}/clients/{id}/roles/{role-name}/composites/clients/{clientUuid} | 
[**GetClientRoleCompositesRealm**](RolesAPI.md#GetClientRoleCompositesRealm) | **Get** /{realm}/clients/{id}/roles/{role-name}/composites/realm | 
[**GetClientRoleGroups**](RolesAPI.md#GetClientRoleGroups) | **Get** /{realm}/clients/{id}/roles/{role-name}/groups | 
[**GetClientRoleManagementPermissions**](RolesAPI.md#GetClientRoleManagementPermissions) | **Get** /{realm}/clients/{id}/roles/{role-name}/management/permissions | 
[**GetClientRoleUsers**](RolesAPI.md#GetClientRoleUsers) | **Get** /{realm}/clients/{id}/roles/{role-name}/users | 
[**GetClientRoles**](RolesAPI.md#GetClientRoles) | **Get** /{realm}/clients/{id}/roles | 
[**GetRoleByRealmByRoleName**](RolesAPI.md#GetRoleByRealmByRoleName) | **Get** /{realm}/roles/{role-name} | 
[**GetRoleCompositesByRealmByRoleName**](RolesAPI.md#GetRoleCompositesByRealmByRoleName) | **Get** /{realm}/roles/{role-name}/composites | 
[**GetRoleCompositesClientByRealmByRoleNameByClientUuid**](RolesAPI.md#GetRoleCompositesClientByRealmByRoleNameByClientUuid) | **Get** /{realm}/roles/{role-name}/composites/clients/{clientUuid} | 
[**GetRoleCompositesRealmByRealmByRoleName**](RolesAPI.md#GetRoleCompositesRealmByRealmByRoleName) | **Get** /{realm}/roles/{role-name}/composites/realm | 
[**GetRoleGroupsByRealmByRoleName**](RolesAPI.md#GetRoleGroupsByRealmByRoleName) | **Get** /{realm}/roles/{role-name}/groups | 
[**GetRoleManagementPermissionsByRealmByRoleName**](RolesAPI.md#GetRoleManagementPermissionsByRealmByRoleName) | **Get** /{realm}/roles/{role-name}/management/permissions | 
[**GetRoleUsersByRealmByRoleName**](RolesAPI.md#GetRoleUsersByRealmByRoleName) | **Get** /{realm}/roles/{role-name}/users | 
[**GetRolesByRealm**](RolesAPI.md#GetRolesByRealm) | **Get** /{realm}/roles | 
[**PostClientRoleComposites**](RolesAPI.md#PostClientRoleComposites) | **Post** /{realm}/clients/{id}/roles/{role-name}/composites | 
[**PostClientRoles**](RolesAPI.md#PostClientRoles) | **Post** /{realm}/clients/{id}/roles | 
[**PostRoleCompositesByRealmByRoleName**](RolesAPI.md#PostRoleCompositesByRealmByRoleName) | **Post** /{realm}/roles/{role-name}/composites | 
[**PostRolesByRealm**](RolesAPI.md#PostRolesByRealm) | **Post** /{realm}/roles | 
[**PutClientRole**](RolesAPI.md#PutClientRole) | **Put** /{realm}/clients/{id}/roles/{role-name} | 
[**PutClientRoleManagementPermissions**](RolesAPI.md#PutClientRoleManagementPermissions) | **Put** /{realm}/clients/{id}/roles/{role-name}/management/permissions | 
[**PutRoleByRealmByRoleName**](RolesAPI.md#PutRoleByRealmByRoleName) | **Put** /{realm}/roles/{role-name} | 
[**PutRoleManagementPermissionsByRealmByRoleName**](RolesAPI.md#PutRoleManagementPermissionsByRealmByRoleName) | **Put** /{realm}/roles/{role-name}/management/permissions | 



## DeleteClientRole

> DeleteClientRole(ctx, realm, id, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteClientRole(context.Background(), realm, id, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteClientRole``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRoleRequest struct via the builder pattern


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


## DeleteClientRoleComposites

> DeleteClientRoleComposites(ctx, realm, id, roleName).RoleRepresentation(roleRepresentation).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteClientRoleComposites(context.Background(), realm, id, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteClientRoleComposites``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRoleCompositesRequest struct via the builder pattern


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


## DeleteRoleByRealmByRoleName

> DeleteRoleByRealmByRoleName(ctx, realm, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteRoleByRealmByRoleName(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRoleByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleByRealmByRoleNameRequest struct via the builder pattern


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


## DeleteRoleCompositesByRealmByRoleName

> DeleteRoleCompositesByRealmByRoleName(ctx, realm, roleName).RoleRepresentation(roleRepresentation).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteRoleCompositesByRealmByRoleName(context.Background(), realm, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRoleCompositesByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleCompositesByRealmByRoleNameRequest struct via the builder pattern


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


## GetClientRole

> RoleRepresentation GetClientRole(ctx, realm, id, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRole(context.Background(), realm, id, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRole`: RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRoleRequest struct via the builder pattern


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


## GetClientRoleComposites

> []RoleRepresentation GetClientRoleComposites(ctx, realm, id, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRoleComposites(context.Background(), realm, id, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRoleComposites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRoleComposites`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRoleComposites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRoleCompositesRequest struct via the builder pattern


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


## GetClientRoleCompositesClient

> []RoleRepresentation GetClientRoleCompositesClient(ctx, realm, id, roleName, clientUuid).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	clientUuid := "clientUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRoleCompositesClient(context.Background(), realm, id, roleName, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRoleCompositesClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRoleCompositesClient`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRoleCompositesClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** | role&#39;s name (not id!) | 
**clientUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRoleCompositesClientRequest struct via the builder pattern


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


## GetClientRoleCompositesRealm

> []RoleRepresentation GetClientRoleCompositesRealm(ctx, realm, id, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRoleCompositesRealm(context.Background(), realm, id, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRoleCompositesRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRoleCompositesRealm`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRoleCompositesRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRoleCompositesRealmRequest struct via the builder pattern


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


## GetClientRoleGroups

> []GroupRepresentation GetClientRoleGroups(ctx, realm, id, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()





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
	roleName := "roleName_example" // string | the role name.
	briefRepresentation := "briefRepresentation_example" // string | if false, return a full representation of the {@code GroupRepresentation} objects. (optional)
	first := "first_example" // string | first result to return. Ignored if negative or {@code null}. (optional)
	max := "max_example" // string | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRoleGroups(context.Background(), realm, id, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRoleGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRoleGroups`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRoleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRoleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **string** | if false, return a full representation of the {@code GroupRepresentation} objects. | 
 **first** | **string** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **string** | maximum number of results to return. Ignored if negative or {@code null}. | 

### Return type

[**[]GroupRepresentation**](GroupRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientRoleManagementPermissions

> ManagementPermissionReference GetClientRoleManagementPermissions(ctx, realm, id, roleName).Execute()





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
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRoleManagementPermissions(context.Background(), realm, id, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRoleManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRoleManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRoleManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRoleManagementPermissionsRequest struct via the builder pattern


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


## GetClientRoleUsers

> []UserRepresentation GetClientRoleUsers(ctx, realm, id, roleName).First(first).Max(max).Execute()





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
	roleName := "roleName_example" // string | the role name.
	first := "first_example" // string | first result to return. Ignored if negative or {@code null}. (optional)
	max := "max_example" // string | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRoleUsers(context.Background(), realm, id, roleName).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRoleUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRoleUsers`: []UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **first** | **string** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **string** | maximum number of results to return. Ignored if negative or {@code null}. | 

### Return type

[**[]UserRepresentation**](UserRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientRoles

> []RoleRepresentation GetClientRoles(ctx, realm, id).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()





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
	briefRepresentation := "briefRepresentation_example" // string |  (optional)
	first := "first_example" // string |  (optional)
	max := "max_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetClientRoles(context.Background(), realm, id).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetClientRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRoles`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetClientRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **string** |  | 
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


## GetRoleByRealmByRoleName

> RoleRepresentation GetRoleByRealmByRoleName(ctx, realm, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleByRealmByRoleName(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleByRealmByRoleName`: RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleByRealmByRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleByRealmByRoleNameRequest struct via the builder pattern


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


## GetRoleCompositesByRealmByRoleName

> []RoleRepresentation GetRoleCompositesByRealmByRoleName(ctx, realm, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleCompositesByRealmByRoleName(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleCompositesByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleCompositesByRealmByRoleName`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleCompositesByRealmByRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleCompositesByRealmByRoleNameRequest struct via the builder pattern


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


## GetRoleCompositesClientByRealmByRoleNameByClientUuid

> []RoleRepresentation GetRoleCompositesClientByRealmByRoleNameByClientUuid(ctx, realm, roleName, clientUuid).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	clientUuid := "clientUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleCompositesClientByRealmByRoleNameByClientUuid(context.Background(), realm, roleName, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleCompositesClientByRealmByRoleNameByClientUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleCompositesClientByRealmByRoleNameByClientUuid`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleCompositesClientByRealmByRoleNameByClientUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 
**clientUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleCompositesClientByRealmByRoleNameByClientUuidRequest struct via the builder pattern


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


## GetRoleCompositesRealmByRealmByRoleName

> []RoleRepresentation GetRoleCompositesRealmByRealmByRoleName(ctx, realm, roleName).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleCompositesRealmByRealmByRoleName(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleCompositesRealmByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleCompositesRealmByRealmByRoleName`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleCompositesRealmByRealmByRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleCompositesRealmByRealmByRoleNameRequest struct via the builder pattern


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


## GetRoleGroupsByRealmByRoleName

> []GroupRepresentation GetRoleGroupsByRealmByRoleName(ctx, realm, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()





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
	roleName := "roleName_example" // string | the role name.
	briefRepresentation := "briefRepresentation_example" // string | if false, return a full representation of the {@code GroupRepresentation} objects. (optional)
	first := "first_example" // string | first result to return. Ignored if negative or {@code null}. (optional)
	max := "max_example" // string | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleGroupsByRealmByRoleName(context.Background(), realm, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleGroupsByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleGroupsByRealmByRoleName`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleGroupsByRealmByRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleGroupsByRealmByRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **string** | if false, return a full representation of the {@code GroupRepresentation} objects. | 
 **first** | **string** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **string** | maximum number of results to return. Ignored if negative or {@code null}. | 

### Return type

[**[]GroupRepresentation**](GroupRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleManagementPermissionsByRealmByRoleName

> ManagementPermissionReference GetRoleManagementPermissionsByRealmByRoleName(ctx, realm, roleName).Execute()





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
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleManagementPermissionsByRealmByRoleName(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleManagementPermissionsByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleManagementPermissionsByRealmByRoleName`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleManagementPermissionsByRealmByRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleManagementPermissionsByRealmByRoleNameRequest struct via the builder pattern


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


## GetRoleUsersByRealmByRoleName

> []UserRepresentation GetRoleUsersByRealmByRoleName(ctx, realm, roleName).First(first).Max(max).Execute()





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
	roleName := "roleName_example" // string | the role name.
	first := "first_example" // string | first result to return. Ignored if negative or {@code null}. (optional)
	max := "max_example" // string | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleUsersByRealmByRoleName(context.Background(), realm, roleName).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleUsersByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleUsersByRealmByRoleName`: []UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleUsersByRealmByRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleUsersByRealmByRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **first** | **string** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **string** | maximum number of results to return. Ignored if negative or {@code null}. | 

### Return type

[**[]UserRepresentation**](UserRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRolesByRealm

> []RoleRepresentation GetRolesByRealm(ctx, realm).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()





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
	briefRepresentation := "briefRepresentation_example" // string |  (optional)
	first := "first_example" // string |  (optional)
	max := "max_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRolesByRealm(context.Background(), realm).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRolesByRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolesByRealm`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRolesByRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesByRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **briefRepresentation** | **string** |  | 
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


## PostClientRoleComposites

> PostClientRoleComposites(ctx, realm, id, roleName).RoleRepresentation(roleRepresentation).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.PostClientRoleComposites(context.Background(), realm, id, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PostClientRoleComposites``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientRoleCompositesRequest struct via the builder pattern


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


## PostClientRoles

> PostClientRoles(ctx, realm, id).RoleRepresentation(roleRepresentation).Execute()





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
	r, err := apiClient.RolesAPI.PostClientRoles(context.Background(), realm, id).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PostClientRoles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostClientRolesRequest struct via the builder pattern


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


## PostRoleCompositesByRealmByRoleName

> PostRoleCompositesByRealmByRoleName(ctx, realm, roleName).RoleRepresentation(roleRepresentation).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.PostRoleCompositesByRealmByRoleName(context.Background(), realm, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PostRoleCompositesByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRoleCompositesByRealmByRoleNameRequest struct via the builder pattern


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


## PostRolesByRealm

> PostRolesByRealm(ctx, realm).RoleRepresentation(roleRepresentation).Execute()





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
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.PostRolesByRealm(context.Background(), realm).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PostRolesByRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostRolesByRealmRequest struct via the builder pattern


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


## PutClientRole

> PutClientRole(ctx, realm, id, roleName).RoleRepresentation(roleRepresentation).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.PutClientRole(context.Background(), realm, id, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PutClientRole``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientRoleRequest struct via the builder pattern


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


## PutClientRoleManagementPermissions

> ManagementPermissionReference PutClientRoleManagementPermissions(ctx, realm, id, roleName).ManagementPermissionReference(managementPermissionReference).Execute()





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
	roleName := "roleName_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference | ManagementPermissionReference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.PutClientRoleManagementPermissions(context.Background(), realm, id, roleName).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PutClientRoleManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientRoleManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.PutClientRoleManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientRoleManagementPermissionsRequest struct via the builder pattern


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


## PutRoleByRealmByRoleName

> PutRoleByRealmByRoleName(ctx, realm, roleName).RoleRepresentation(roleRepresentation).Execute()





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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation | RoleRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.PutRoleByRealmByRoleName(context.Background(), realm, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PutRoleByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRoleByRealmByRoleNameRequest struct via the builder pattern


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


## PutRoleManagementPermissionsByRealmByRoleName

> ManagementPermissionReference PutRoleManagementPermissionsByRealmByRoleName(ctx, realm, roleName).ManagementPermissionReference(managementPermissionReference).Execute()





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
	roleName := "roleName_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference | ManagementPermissionReference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.PutRoleManagementPermissionsByRealmByRoleName(context.Background(), realm, roleName).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PutRoleManagementPermissionsByRealmByRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRoleManagementPermissionsByRealmByRoleName`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.PutRoleManagementPermissionsByRealmByRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRoleManagementPermissionsByRealmByRoleNameRequest struct via the builder pattern


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

