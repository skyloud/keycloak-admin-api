# \ClientsAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientByRealmById**](ClientsAPI.md#DeleteClientByRealmById) | **Delete** /{realm}/clients/{id} | 
[**DeleteDefaultClientScope**](ClientsAPI.md#DeleteDefaultClientScope) | **Delete** /{realm}/clients/{id}/default-client-scopes/{clientScopeId} | 
[**DeleteNode**](ClientsAPI.md#DeleteNode) | **Delete** /{realm}/clients/{id}/nodes/{node} | 
[**DeleteOptionalClientScope**](ClientsAPI.md#DeleteOptionalClientScope) | **Delete** /{realm}/clients/{id}/optional-client-scopes/{clientScopeId} | 
[**DeleteRotated**](ClientsAPI.md#DeleteRotated) | **Delete** /{realm}/clients/{id}/client-secret/rotated | 
[**GetClientByRealmById**](ClientsAPI.md#GetClientByRealmById) | **Get** /{realm}/clients/{id} | 
[**GetClientManagementPermissions**](ClientsAPI.md#GetClientManagementPermissions) | **Get** /{realm}/clients/{id}/management/permissions | 
[**GetClientSecret**](ClientsAPI.md#GetClientSecret) | **Get** /{realm}/clients/{id}/client-secret | 
[**GetClientUserSessions**](ClientsAPI.md#GetClientUserSessions) | **Get** /{realm}/clients/{id}/user-sessions | 
[**GetClients**](ClientsAPI.md#GetClients) | **Get** /{realm}/clients | 
[**GetDefaultClientScopes**](ClientsAPI.md#GetDefaultClientScopes) | **Get** /{realm}/clients/{id}/default-client-scopes | 
[**GetGenerateExampleAccessToken**](ClientsAPI.md#GetGenerateExampleAccessToken) | **Get** /{realm}/clients/{id}/evaluate-scopes/generate-example-access-token | 
[**GetGenerateExampleIdToken**](ClientsAPI.md#GetGenerateExampleIdToken) | **Get** /{realm}/clients/{id}/evaluate-scopes/generate-example-id-token | 
[**GetGenerateExampleUserinfo**](ClientsAPI.md#GetGenerateExampleUserinfo) | **Get** /{realm}/clients/{id}/evaluate-scopes/generate-example-userinfo | 
[**GetGranted**](ClientsAPI.md#GetGranted) | **Get** /{realm}/clients/{id}/evaluate-scopes/scope-mappings/{roleContainerId}/granted | 
[**GetInstallationProvider**](ClientsAPI.md#GetInstallationProvider) | **Get** /{realm}/clients/{id}/installation/providers/{providerId} | 
[**GetNotGranted**](ClientsAPI.md#GetNotGranted) | **Get** /{realm}/clients/{id}/evaluate-scopes/scope-mappings/{roleContainerId}/not-granted | 
[**GetOfflineSessionCount**](ClientsAPI.md#GetOfflineSessionCount) | **Get** /{realm}/clients/{id}/offline-session-count | 
[**GetOfflineSessions**](ClientsAPI.md#GetOfflineSessions) | **Get** /{realm}/clients/{id}/offline-sessions | 
[**GetOptionalClientScopes**](ClientsAPI.md#GetOptionalClientScopes) | **Get** /{realm}/clients/{id}/optional-client-scopes | 
[**GetProtocolMappers**](ClientsAPI.md#GetProtocolMappers) | **Get** /{realm}/clients/{id}/evaluate-scopes/protocol-mappers | 
[**GetRotated**](ClientsAPI.md#GetRotated) | **Get** /{realm}/clients/{id}/client-secret/rotated | 
[**GetServiceAccountUser**](ClientsAPI.md#GetServiceAccountUser) | **Get** /{realm}/clients/{id}/service-account-user | 
[**GetSessionCount**](ClientsAPI.md#GetSessionCount) | **Get** /{realm}/clients/{id}/session-count | 
[**GetTestNodesAvailable**](ClientsAPI.md#GetTestNodesAvailable) | **Get** /{realm}/clients/{id}/test-nodes-available | 
[**PostClientPushRevocation**](ClientsAPI.md#PostClientPushRevocation) | **Post** /{realm}/clients/{id}/push-revocation | 
[**PostClientSecret**](ClientsAPI.md#PostClientSecret) | **Post** /{realm}/clients/{id}/client-secret | 
[**PostClients**](ClientsAPI.md#PostClients) | **Post** /{realm}/clients | 
[**PostNodes**](ClientsAPI.md#PostNodes) | **Post** /{realm}/clients/{id}/nodes | 
[**PostRegistrationAccessToken**](ClientsAPI.md#PostRegistrationAccessToken) | **Post** /{realm}/clients/{id}/registration-access-token | 
[**PutClient**](ClientsAPI.md#PutClient) | **Put** /{realm}/clients/{id} | 
[**PutClientManagementPermissions**](ClientsAPI.md#PutClientManagementPermissions) | **Put** /{realm}/clients/{id}/management/permissions | 
[**PutDefaultClientScope**](ClientsAPI.md#PutDefaultClientScope) | **Put** /{realm}/clients/{id}/default-client-scopes/{clientScopeId} | 
[**PutOptionalClientScope**](ClientsAPI.md#PutOptionalClientScope) | **Put** /{realm}/clients/{id}/optional-client-scopes/{clientScopeId} | 



## DeleteClientByRealmById

> DeleteClientByRealmById(ctx, realm, id).Execute()





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
	r, err := apiClient.ClientsAPI.DeleteClientByRealmById(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.DeleteClientByRealmById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteClientByRealmByIdRequest struct via the builder pattern


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


## DeleteDefaultClientScope

> DeleteDefaultClientScope(ctx, realm, id, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.DeleteDefaultClientScope(context.Background(), realm, id, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.DeleteDefaultClientScope``: %v\n", err)
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
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultClientScopeRequest struct via the builder pattern


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


## DeleteNode

> DeleteNode(ctx, realm, id, node).Execute()





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
	node := "node_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.DeleteNode(context.Background(), realm, id, node).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.DeleteNode``: %v\n", err)
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
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


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


## DeleteOptionalClientScope

> DeleteOptionalClientScope(ctx, realm, id, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.DeleteOptionalClientScope(context.Background(), realm, id, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.DeleteOptionalClientScope``: %v\n", err)
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
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionalClientScopeRequest struct via the builder pattern


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


## DeleteRotated

> DeleteRotated(ctx, realm, id).Execute()





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
	r, err := apiClient.ClientsAPI.DeleteRotated(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.DeleteRotated``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRotatedRequest struct via the builder pattern


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


## GetClientByRealmById

> ClientRepresentation GetClientByRealmById(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetClientByRealmById(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClientByRealmById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientByRealmById`: ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClientByRealmById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientByRealmByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientRepresentation**](ClientRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientManagementPermissions

> ManagementPermissionReference GetClientManagementPermissions(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetClientManagementPermissions(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClientManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClientManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientManagementPermissionsRequest struct via the builder pattern


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


## GetClientSecret

> CredentialRepresentation GetClientSecret(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetClientSecret(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientSecret`: CredentialRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialRepresentation**](CredentialRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientUserSessions

> []UserSessionRepresentation GetClientUserSessions(ctx, realm, id).First(first).Max(max).Execute()





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
	first := "first_example" // string | Paging offset (optional)
	max := "max_example" // string | Maximum results size (defaults to 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetClientUserSessions(context.Background(), realm, id).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClientUserSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientUserSessions`: []UserSessionRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClientUserSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientUserSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **first** | **string** | Paging offset | 
 **max** | **string** | Maximum results size (defaults to 100) | 

### Return type

[**[]UserSessionRepresentation**](UserSessionRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> []ClientRepresentation GetClients(ctx, realm).ClientId(clientId).First(first).Max(max).Q(q).Search(search).ViewableOnly(viewableOnly).Execute()





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
	clientId := "clientId_example" // string | filter by clientId (optional)
	first := "first_example" // string | the first result (optional)
	max := "max_example" // string | the max results to return (optional)
	q := "q_example" // string |  (optional)
	search := "search_example" // string | whether this is a search query or a getClientById query (optional)
	viewableOnly := "viewableOnly_example" // string | filter clients that cannot be viewed in full by admin (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetClients(context.Background(), realm).ClientId(clientId).First(first).Max(max).Q(q).Search(search).ViewableOnly(viewableOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: []ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** | filter by clientId | 
 **first** | **string** | the first result | 
 **max** | **string** | the max results to return | 
 **q** | **string** |  | 
 **search** | **string** | whether this is a search query or a getClientById query | 
 **viewableOnly** | **string** | filter clients that cannot be viewed in full by admin | 

### Return type

[**[]ClientRepresentation**](ClientRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultClientScopes

> []ClientScopeRepresentation GetDefaultClientScopes(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetDefaultClientScopes(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetDefaultClientScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultClientScopes`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetDefaultClientScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultClientScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ClientScopeRepresentation**](ClientScopeRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenerateExampleAccessToken

> AccessToken GetGenerateExampleAccessToken(ctx, realm, id).Scope(scope).UserId(userId).Execute()





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
	scope := "scope_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetGenerateExampleAccessToken(context.Background(), realm, id).Scope(scope).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetGenerateExampleAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenerateExampleAccessToken`: AccessToken
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetGenerateExampleAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenerateExampleAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 
 **userId** | **string** |  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenerateExampleIdToken

> IDToken GetGenerateExampleIdToken(ctx, realm, id).Scope(scope).UserId(userId).Execute()





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
	scope := "scope_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetGenerateExampleIdToken(context.Background(), realm, id).Scope(scope).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetGenerateExampleIdToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenerateExampleIdToken`: IDToken
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetGenerateExampleIdToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenerateExampleIdTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 
 **userId** | **string** |  | 

### Return type

[**IDToken**](IDToken.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenerateExampleUserinfo

> map[string]map[string]interface{} GetGenerateExampleUserinfo(ctx, realm, id).Scope(scope).UserId(userId).Execute()





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
	scope := "scope_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetGenerateExampleUserinfo(context.Background(), realm, id).Scope(scope).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetGenerateExampleUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenerateExampleUserinfo`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetGenerateExampleUserinfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenerateExampleUserinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 
 **userId** | **string** |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGranted

> []RoleRepresentation GetGranted(ctx, realm, id, roleContainerId).Scope(scope).Execute()





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
	roleContainerId := "roleContainerId_example" // string | either realm name OR client UUID
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetGranted(context.Background(), realm, id, roleContainerId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetGranted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGranted`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetGranted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleContainerId** | **string** | either realm name OR client UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGrantedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scope** | **string** |  | 

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


## GetInstallationProvider

> GetInstallationProvider(ctx, realm, id, providerId).Execute()



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
	providerId := "providerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.GetInstallationProvider(context.Background(), realm, id, providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetInstallationProvider``: %v\n", err)
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
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstallationProviderRequest struct via the builder pattern


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


## GetNotGranted

> []RoleRepresentation GetNotGranted(ctx, realm, id, roleContainerId).Scope(scope).Execute()





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
	roleContainerId := "roleContainerId_example" // string | either realm name OR client UUID
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetNotGranted(context.Background(), realm, id, roleContainerId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNotGranted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotGranted`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNotGranted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**roleContainerId** | **string** | either realm name OR client UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotGrantedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scope** | **string** |  | 

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


## GetOfflineSessionCount

> map[string]int64 GetOfflineSessionCount(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetOfflineSessionCount(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOfflineSessionCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfflineSessionCount`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOfflineSessionCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfflineSessionCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetOfflineSessions

> []UserSessionRepresentation GetOfflineSessions(ctx, realm, id).First(first).Max(max).Execute()





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
	first := "first_example" // string | Paging offset (optional)
	max := "max_example" // string | Maximum results size (defaults to 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetOfflineSessions(context.Background(), realm, id).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOfflineSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfflineSessions`: []UserSessionRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOfflineSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfflineSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **first** | **string** | Paging offset | 
 **max** | **string** | Maximum results size (defaults to 100) | 

### Return type

[**[]UserSessionRepresentation**](UserSessionRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptionalClientScopes

> []ClientScopeRepresentation GetOptionalClientScopes(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetOptionalClientScopes(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOptionalClientScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionalClientScopes`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOptionalClientScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionalClientScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ClientScopeRepresentation**](ClientScopeRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtocolMappers

> []ProtocolMapperEvaluationRepresentation GetProtocolMappers(ctx, realm, id).Scope(scope).Execute()





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
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetProtocolMappers(context.Background(), realm, id).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetProtocolMappers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtocolMappers`: []ProtocolMapperEvaluationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetProtocolMappers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtocolMappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 

### Return type

[**[]ProtocolMapperEvaluationRepresentation**](ProtocolMapperEvaluationRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRotated

> CredentialRepresentation GetRotated(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetRotated(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetRotated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRotated`: CredentialRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetRotated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRotatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialRepresentation**](CredentialRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountUser

> UserRepresentation GetServiceAccountUser(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetServiceAccountUser(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetServiceAccountUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccountUser`: UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetServiceAccountUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserRepresentation**](UserRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionCount

> map[string]int64 GetSessionCount(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetSessionCount(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetSessionCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionCount`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetSessionCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetTestNodesAvailable

> GlobalRequestResult GetTestNodesAvailable(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.GetTestNodesAvailable(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetTestNodesAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestNodesAvailable`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetTestNodesAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestNodesAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GlobalRequestResult**](GlobalRequestResult.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientPushRevocation

> GlobalRequestResult PostClientPushRevocation(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.PostClientPushRevocation(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PostClientPushRevocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientPushRevocation`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.PostClientPushRevocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientPushRevocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GlobalRequestResult**](GlobalRequestResult.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientSecret

> CredentialRepresentation PostClientSecret(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.PostClientSecret(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PostClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientSecret`: CredentialRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.PostClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialRepresentation**](CredentialRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClients

> PostClients(ctx, realm).ClientRepresentation(clientRepresentation).Execute()





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
	clientRepresentation := *openapiclient.NewClientRepresentation() // ClientRepresentation | ClientRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.PostClients(context.Background(), realm).ClientRepresentation(clientRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PostClients``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRepresentation** | [**ClientRepresentation**](ClientRepresentation.md) | ClientRepresentation | 

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


## PostNodes

> PostNodes(ctx, realm, id).Body(body).Execute()





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
	body := "body_example" // string | [string] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.PostNodes(context.Background(), realm, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PostNodes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | [string] | 

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


## PostRegistrationAccessToken

> ClientRepresentation PostRegistrationAccessToken(ctx, realm, id).Execute()





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
	resp, r, err := apiClient.ClientsAPI.PostRegistrationAccessToken(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PostRegistrationAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRegistrationAccessToken`: ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.PostRegistrationAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRegistrationAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientRepresentation**](ClientRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClient

> PutClient(ctx, realm, id).ClientRepresentation(clientRepresentation).Execute()





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
	clientRepresentation := *openapiclient.NewClientRepresentation() // ClientRepresentation | ClientRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.PutClient(context.Background(), realm, id).ClientRepresentation(clientRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PutClient``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRepresentation** | [**ClientRepresentation**](ClientRepresentation.md) | ClientRepresentation | 

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


## PutClientManagementPermissions

> ManagementPermissionReference PutClientManagementPermissions(ctx, realm, id).ManagementPermissionReference(managementPermissionReference).Execute()





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
	resp, r, err := apiClient.ClientsAPI.PutClientManagementPermissions(context.Background(), realm, id).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PutClientManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.PutClientManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientManagementPermissionsRequest struct via the builder pattern


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


## PutDefaultClientScope

> PutDefaultClientScope(ctx, realm, id, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.PutDefaultClientScope(context.Background(), realm, id, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PutDefaultClientScope``: %v\n", err)
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
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDefaultClientScopeRequest struct via the builder pattern


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


## PutOptionalClientScope

> PutOptionalClientScope(ctx, realm, id, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.PutOptionalClientScope(context.Background(), realm, id, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PutOptionalClientScope``: %v\n", err)
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
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOptionalClientScopeRequest struct via the builder pattern


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

