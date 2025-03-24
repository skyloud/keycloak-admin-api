# \GroupsAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGroupByRealmById**](GroupsAPI.md#DeleteGroupByRealmById) | **Delete** /{realm}/groups/{id} | 
[**GetChildren**](GroupsAPI.md#GetChildren) | **Get** /{realm}/groups/{id}/children | 
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /{realm}/groups/{id} | 
[**GetGroupManagementPermissions**](GroupsAPI.md#GetGroupManagementPermissions) | **Get** /{realm}/groups/{id}/management/permissions | 
[**GetGroupsByRealm**](GroupsAPI.md#GetGroupsByRealm) | **Get** /{realm}/groups | 
[**GetGroupsCountByRealm**](GroupsAPI.md#GetGroupsCountByRealm) | **Get** /{realm}/groups/count | 
[**GetMembers**](GroupsAPI.md#GetMembers) | **Get** /{realm}/groups/{id}/members | 
[**PostChildren**](GroupsAPI.md#PostChildren) | **Post** /{realm}/groups/{id}/children | 
[**PostGroups**](GroupsAPI.md#PostGroups) | **Post** /{realm}/groups | 
[**PutGroupByRealmById**](GroupsAPI.md#PutGroupByRealmById) | **Put** /{realm}/groups/{id} | 
[**PutGroupManagementPermissions**](GroupsAPI.md#PutGroupManagementPermissions) | **Put** /{realm}/groups/{id}/management/permissions | 



## DeleteGroupByRealmById

> DeleteGroupByRealmById(ctx, realm, id).Execute()



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
	r, err := apiClient.GroupsAPI.DeleteGroupByRealmById(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupByRealmById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGroupByRealmByIdRequest struct via the builder pattern


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


## GetChildren

> []GroupRepresentation GetChildren(ctx, realm, id).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetChildren(context.Background(), realm, id).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChildren`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **string** |  | 
 **first** | **string** |  | 
 **max** | **string** |  | 

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


## GetGroup

> GroupRepresentation GetGroup(ctx, realm, id).Execute()



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
	resp, r, err := apiClient.GroupsAPI.GetGroup(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupRepresentation**](GroupRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupManagementPermissions

> ManagementPermissionReference GetGroupManagementPermissions(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.GroupsAPI.GetGroupManagementPermissions(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupManagementPermissionsRequest struct via the builder pattern


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


## GetGroupsByRealm

> []GroupRepresentation GetGroupsByRealm(ctx, realm).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).PopulateHierarchy(populateHierarchy).Q(q).Search(search).Execute()





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
	exact := "exact_example" // string |  (optional)
	first := "first_example" // string |  (optional)
	max := "max_example" // string |  (optional)
	populateHierarchy := "populateHierarchy_example" // string |  (optional)
	q := "q_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupsByRealm(context.Background(), realm).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).PopulateHierarchy(populateHierarchy).Q(q).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupsByRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsByRealm`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupsByRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsByRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **briefRepresentation** | **string** |  | 
 **exact** | **string** |  | 
 **first** | **string** |  | 
 **max** | **string** |  | 
 **populateHierarchy** | **string** |  | 
 **q** | **string** |  | 
 **search** | **string** |  | 

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


## GetGroupsCountByRealm

> map[string]int64 GetGroupsCountByRealm(ctx, realm).Search(search).Top(top).Execute()





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
	search := "search_example" // string |  (optional)
	top := "top_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupsCountByRealm(context.Background(), realm).Search(search).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupsCountByRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsCountByRealm`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupsCountByRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsCountByRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **top** | **string** |  | 

### Return type

**map[string]int64**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembers

> []UserRepresentation GetMembers(ctx, realm, id).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()





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
	briefRepresentation := "briefRepresentation_example" // string | Only return basic information (only guaranteed to return id, username, created, first and last name, email, enabled state, email verification state, federation link, and access. Note that it means that namely user attributes, required actions, and not before are not returned.) (optional)
	first := "first_example" // string | Pagination offset (optional)
	max := "max_example" // string | Maximum results size (defaults to 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetMembers(context.Background(), realm, id).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMembers`: []UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **string** | Only return basic information (only guaranteed to return id, username, created, first and last name, email, enabled state, email verification state, federation link, and access. Note that it means that namely user attributes, required actions, and not before are not returned.) | 
 **first** | **string** | Pagination offset | 
 **max** | **string** | Maximum results size (defaults to 100) | 

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


## PostChildren

> PostChildren(ctx, realm, id).GroupRepresentation(groupRepresentation).Execute()





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
	groupRepresentation := *openapiclient.NewGroupRepresentation() // GroupRepresentation | GroupRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.PostChildren(context.Background(), realm, id).GroupRepresentation(groupRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.PostChildren``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupRepresentation** | [**GroupRepresentation**](GroupRepresentation.md) | GroupRepresentation | 

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


## PostGroups

> PostGroups(ctx, realm).GroupRepresentation(groupRepresentation).Execute()





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
	groupRepresentation := *openapiclient.NewGroupRepresentation() // GroupRepresentation | GroupRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.PostGroups(context.Background(), realm).GroupRepresentation(groupRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.PostGroups``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRepresentation** | [**GroupRepresentation**](GroupRepresentation.md) | GroupRepresentation | 

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


## PutGroupByRealmById

> PutGroupByRealmById(ctx, realm, id).GroupRepresentation(groupRepresentation).Execute()





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
	groupRepresentation := *openapiclient.NewGroupRepresentation() // GroupRepresentation | GroupRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.PutGroupByRealmById(context.Background(), realm, id).GroupRepresentation(groupRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.PutGroupByRealmById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutGroupByRealmByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupRepresentation** | [**GroupRepresentation**](GroupRepresentation.md) | GroupRepresentation | 

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


## PutGroupManagementPermissions

> ManagementPermissionReference PutGroupManagementPermissions(ctx, realm, id).ManagementPermissionReference(managementPermissionReference).Execute()





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
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference | ManagementPermissionReference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.PutGroupManagementPermissions(context.Background(), realm, id).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.PutGroupManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutGroupManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.PutGroupManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutGroupManagementPermissionsRequest struct via the builder pattern


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

