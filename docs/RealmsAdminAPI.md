# \RealmsAdminAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdminEvents**](RealmsAdminAPI.md#DeleteAdminEvents) | **Delete** /{realm}/admin-events | 
[**DeleteByRealm**](RealmsAdminAPI.md#DeleteByRealm) | **Delete** /{realm} | 
[**DeleteDefaultDefaultClientScope**](RealmsAdminAPI.md#DeleteDefaultDefaultClientScope) | **Delete** /{realm}/default-default-client-scopes/{clientScopeId} | 
[**DeleteDefaultGroup**](RealmsAdminAPI.md#DeleteDefaultGroup) | **Delete** /{realm}/default-groups/{groupId} | 
[**DeleteDefaultOptionalClientScope**](RealmsAdminAPI.md#DeleteDefaultOptionalClientScope) | **Delete** /{realm}/default-optional-client-scopes/{clientScopeId} | 
[**DeleteEvents**](RealmsAdminAPI.md#DeleteEvents) | **Delete** /{realm}/events | 
[**DeleteLocalizationByRealmByLocale**](RealmsAdminAPI.md#DeleteLocalizationByRealmByLocale) | **Delete** /{realm}/localization/{locale} | 
[**DeleteLocalizationByRealmByLocaleByKey**](RealmsAdminAPI.md#DeleteLocalizationByRealmByLocaleByKey) | **Delete** /{realm}/localization/{locale}/{key} | 
[**DeleteSession**](RealmsAdminAPI.md#DeleteSession) | **Delete** /{realm}/sessions/{session} | 
[**Get**](RealmsAdminAPI.md#Get) | **Get** / | 
[**GetAdminEvents**](RealmsAdminAPI.md#GetAdminEvents) | **Get** /{realm}/admin-events | 
[**GetByRealm**](RealmsAdminAPI.md#GetByRealm) | **Get** /{realm} | 
[**GetClientSessionStats**](RealmsAdminAPI.md#GetClientSessionStats) | **Get** /{realm}/client-session-stats | 
[**GetCredentialRegistrators**](RealmsAdminAPI.md#GetCredentialRegistrators) | **Get** /{realm}/credential-registrators | 
[**GetDefaultDefaultClientScopes**](RealmsAdminAPI.md#GetDefaultDefaultClientScopes) | **Get** /{realm}/default-default-client-scopes | 
[**GetDefaultGroups**](RealmsAdminAPI.md#GetDefaultGroups) | **Get** /{realm}/default-groups | 
[**GetDefaultOptionalClientScopes**](RealmsAdminAPI.md#GetDefaultOptionalClientScopes) | **Get** /{realm}/default-optional-client-scopes | 
[**GetEvents**](RealmsAdminAPI.md#GetEvents) | **Get** /{realm}/events | 
[**GetEventsConfig**](RealmsAdminAPI.md#GetEventsConfig) | **Get** /{realm}/events/config | 
[**GetGroupByPath**](RealmsAdminAPI.md#GetGroupByPath) | **Get** /{realm}/group-by-path/{path} | 
[**GetLocalizationByRealm**](RealmsAdminAPI.md#GetLocalizationByRealm) | **Get** /{realm}/localization | 
[**GetLocalizationByRealmByLocale**](RealmsAdminAPI.md#GetLocalizationByRealmByLocale) | **Get** /{realm}/localization/{locale} | 
[**GetLocalizationByRealmByLocaleByKey**](RealmsAdminAPI.md#GetLocalizationByRealmByLocaleByKey) | **Get** /{realm}/localization/{locale}/{key} | 
[**GetPolicies**](RealmsAdminAPI.md#GetPolicies) | **Get** /{realm}/client-policies/policies | 
[**GetProfiles**](RealmsAdminAPI.md#GetProfiles) | **Get** /{realm}/client-policies/profiles | 
[**GetUsersManagementPermissions**](RealmsAdminAPI.md#GetUsersManagementPermissions) | **Get** /{realm}/users-management-permissions | 
[**Post**](RealmsAdminAPI.md#Post) | **Post** / | 
[**PostClientDescriptionConverter**](RealmsAdminAPI.md#PostClientDescriptionConverter) | **Post** /{realm}/client-description-converter | 
[**PostLocalization**](RealmsAdminAPI.md#PostLocalization) | **Post** /{realm}/localization/{locale} | 
[**PostLogoutAll**](RealmsAdminAPI.md#PostLogoutAll) | **Post** /{realm}/logout-all | 
[**PostPartialExport**](RealmsAdminAPI.md#PostPartialExport) | **Post** /{realm}/partial-export | 
[**PostPartialImport**](RealmsAdminAPI.md#PostPartialImport) | **Post** /{realm}/partialImport | 
[**PostPushRevocationByRealm**](RealmsAdminAPI.md#PostPushRevocationByRealm) | **Post** /{realm}/push-revocation | 
[**PostTestSmtpConnection**](RealmsAdminAPI.md#PostTestSmtpConnection) | **Post** /{realm}/testSMTPConnection | 
[**PutByRealm**](RealmsAdminAPI.md#PutByRealm) | **Put** /{realm} | 
[**PutDefaultDefaultClientScope**](RealmsAdminAPI.md#PutDefaultDefaultClientScope) | **Put** /{realm}/default-default-client-scopes/{clientScopeId} | 
[**PutDefaultGroup**](RealmsAdminAPI.md#PutDefaultGroup) | **Put** /{realm}/default-groups/{groupId} | 
[**PutDefaultOptionalClientScope**](RealmsAdminAPI.md#PutDefaultOptionalClientScope) | **Put** /{realm}/default-optional-client-scopes/{clientScopeId} | 
[**PutEventsConfig**](RealmsAdminAPI.md#PutEventsConfig) | **Put** /{realm}/events/config | 
[**PutLocalization**](RealmsAdminAPI.md#PutLocalization) | **Put** /{realm}/localization/{locale}/{key} | 
[**PutPolicies**](RealmsAdminAPI.md#PutPolicies) | **Put** /{realm}/client-policies/policies | 
[**PutProfiles**](RealmsAdminAPI.md#PutProfiles) | **Put** /{realm}/client-policies/profiles | 
[**PutUsersManagementPermissions**](RealmsAdminAPI.md#PutUsersManagementPermissions) | **Put** /{realm}/users-management-permissions | 



## DeleteAdminEvents

> DeleteAdminEvents(ctx, realm).Execute()





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
	r, err := apiClient.RealmsAdminAPI.DeleteAdminEvents(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteAdminEvents``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAdminEventsRequest struct via the builder pattern


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


## DeleteByRealm

> DeleteByRealm(ctx, realm).Execute()





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
	r, err := apiClient.RealmsAdminAPI.DeleteByRealm(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteByRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteByRealmRequest struct via the builder pattern


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


## DeleteDefaultDefaultClientScope

> DeleteDefaultDefaultClientScope(ctx, realm, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.DeleteDefaultDefaultClientScope(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteDefaultDefaultClientScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultDefaultClientScopeRequest struct via the builder pattern


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


## DeleteDefaultGroup

> DeleteDefaultGroup(ctx, realm, groupId).Execute()



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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.DeleteDefaultGroup(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteDefaultGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultGroupRequest struct via the builder pattern


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


## DeleteDefaultOptionalClientScope

> DeleteDefaultOptionalClientScope(ctx, realm, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.DeleteDefaultOptionalClientScope(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteDefaultOptionalClientScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultOptionalClientScopeRequest struct via the builder pattern


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


## DeleteEvents

> DeleteEvents(ctx, realm).Execute()





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
	r, err := apiClient.RealmsAdminAPI.DeleteEvents(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteEvents``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEventsRequest struct via the builder pattern


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


## DeleteLocalizationByRealmByLocale

> DeleteLocalizationByRealmByLocale(ctx, realm, locale).Execute()



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
	locale := "locale_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.DeleteLocalizationByRealmByLocale(context.Background(), realm, locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteLocalizationByRealmByLocale``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalizationByRealmByLocaleRequest struct via the builder pattern


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


## DeleteLocalizationByRealmByLocaleByKey

> DeleteLocalizationByRealmByLocaleByKey(ctx, realm, locale, key).Execute()



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
	locale := "locale_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.DeleteLocalizationByRealmByLocaleByKey(context.Background(), realm, locale, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteLocalizationByRealmByLocaleByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalizationByRealmByLocaleByKeyRequest struct via the builder pattern


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


## DeleteSession

> DeleteSession(ctx, realm, session).Execute()





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
	session := "session_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.DeleteSession(context.Background(), realm, session).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.DeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**session** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern


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


## Get

> []RealmRepresentation Get(ctx).BriefRepresentation(briefRepresentation).Execute()





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
	briefRepresentation := "briefRepresentation_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.Get(context.Background()).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: []RealmRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **briefRepresentation** | **string** |  | 

### Return type

[**[]RealmRepresentation**](RealmRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminEvents

> []AdminEventRepresentation GetAdminEvents(ctx, realm).AuthClient(authClient).AuthIpAddress(authIpAddress).AuthRealm(authRealm).AuthUser(authUser).DateFrom(dateFrom).DateTo(dateTo).First(first).Max(max).OperationTypes(operationTypes).ResourcePath(resourcePath).ResourceTypes(resourceTypes).Execute()





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
	authClient := "authClient_example" // string |  (optional)
	authIpAddress := "authIpAddress_example" // string |  (optional)
	authRealm := "authRealm_example" // string |  (optional)
	authUser := "authUser_example" // string | user id (optional)
	dateFrom := "dateFrom_example" // string |  (optional)
	dateTo := "dateTo_example" // string |  (optional)
	first := "first_example" // string |  (optional)
	max := "max_example" // string | Maximum results size (defaults to 100) (optional)
	operationTypes := "operationTypes_example" // string | [String] (optional)
	resourcePath := "resourcePath_example" // string |  (optional)
	resourceTypes := "resourceTypes_example" // string | [String] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.GetAdminEvents(context.Background(), realm).AuthClient(authClient).AuthIpAddress(authIpAddress).AuthRealm(authRealm).AuthUser(authUser).DateFrom(dateFrom).DateTo(dateTo).First(first).Max(max).OperationTypes(operationTypes).ResourcePath(resourcePath).ResourceTypes(resourceTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetAdminEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminEvents`: []AdminEventRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetAdminEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authClient** | **string** |  | 
 **authIpAddress** | **string** |  | 
 **authRealm** | **string** |  | 
 **authUser** | **string** | user id | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **first** | **string** |  | 
 **max** | **string** | Maximum results size (defaults to 100) | 
 **operationTypes** | **string** | [String] | 
 **resourcePath** | **string** |  | 
 **resourceTypes** | **string** | [String] | 

### Return type

[**[]AdminEventRepresentation**](AdminEventRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByRealm

> RealmRepresentation GetByRealm(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.GetByRealm(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetByRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByRealm`: RealmRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetByRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RealmRepresentation**](RealmRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientSessionStats

> []string GetClientSessionStats(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.GetClientSessionStats(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetClientSessionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientSessionStats`: []string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetClientSessionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSessionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentialRegistrators

> []string GetCredentialRegistrators(ctx, realm).Execute()



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
	resp, r, err := apiClient.RealmsAdminAPI.GetCredentialRegistrators(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetCredentialRegistrators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialRegistrators`: []string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetCredentialRegistrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialRegistratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultDefaultClientScopes

> []ClientScopeRepresentation GetDefaultDefaultClientScopes(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.GetDefaultDefaultClientScopes(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetDefaultDefaultClientScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultDefaultClientScopes`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetDefaultDefaultClientScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultDefaultClientScopesRequest struct via the builder pattern


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


## GetDefaultGroups

> []GroupRepresentation GetDefaultGroups(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.GetDefaultGroups(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetDefaultGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultGroups`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetDefaultGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDefaultOptionalClientScopes

> []ClientScopeRepresentation GetDefaultOptionalClientScopes(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.GetDefaultOptionalClientScopes(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetDefaultOptionalClientScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultOptionalClientScopes`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetDefaultOptionalClientScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultOptionalClientScopesRequest struct via the builder pattern


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


## GetEvents

> []EventRepresentation GetEvents(ctx, realm).Client(client).DateFrom(dateFrom).DateTo(dateTo).First(first).IpAddress(ipAddress).Max(max).Type_(type_).User(user).Execute()





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
	client := "client_example" // string | App or oauth client name (optional)
	dateFrom := "dateFrom_example" // string | From date (optional)
	dateTo := "dateTo_example" // string | To date (optional)
	first := "first_example" // string | Paging offset (optional)
	ipAddress := "ipAddress_example" // string | IP Address (optional)
	max := "max_example" // string | Maximum results size (defaults to 100) (optional)
	type_ := "type__example" // string | The types of events to return [String] (optional)
	user := "user_example" // string | User id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.GetEvents(context.Background(), realm).Client(client).DateFrom(dateFrom).DateTo(dateTo).First(first).IpAddress(ipAddress).Max(max).Type_(type_).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: []EventRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | **string** | App or oauth client name | 
 **dateFrom** | **string** | From date | 
 **dateTo** | **string** | To date | 
 **first** | **string** | Paging offset | 
 **ipAddress** | **string** | IP Address | 
 **max** | **string** | Maximum results size (defaults to 100) | 
 **type_** | **string** | The types of events to return [String] | 
 **user** | **string** | User id | 

### Return type

[**[]EventRepresentation**](EventRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsConfig

> RealmEventsConfigRepresentation GetEventsConfig(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.GetEventsConfig(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetEventsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsConfig`: RealmEventsConfigRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetEventsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RealmEventsConfigRepresentation**](RealmEventsConfigRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupByPath

> GroupRepresentation GetGroupByPath(ctx, realm, path).Execute()



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
	path := "path_example" // string | PathSegment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.GetGroupByPath(context.Background(), realm, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetGroupByPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupByPath`: GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetGroupByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**path** | **string** | PathSegment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupByPathRequest struct via the builder pattern


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


## GetLocalizationByRealm

> []string GetLocalizationByRealm(ctx, realm).Execute()



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
	resp, r, err := apiClient.RealmsAdminAPI.GetLocalizationByRealm(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetLocalizationByRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationByRealm`: []string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetLocalizationByRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationByRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationByRealmByLocale

> map[string]string GetLocalizationByRealmByLocale(ctx, realm, locale).UseRealmDefaultLocaleFallback(useRealmDefaultLocaleFallback).Execute()



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
	locale := "locale_example" // string | 
	useRealmDefaultLocaleFallback := "useRealmDefaultLocaleFallback_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.GetLocalizationByRealmByLocale(context.Background(), realm, locale).UseRealmDefaultLocaleFallback(useRealmDefaultLocaleFallback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetLocalizationByRealmByLocale``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationByRealmByLocale`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetLocalizationByRealmByLocale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationByRealmByLocaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **useRealmDefaultLocaleFallback** | **string** |  | 

### Return type

**map[string]string**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationByRealmByLocaleByKey

> string GetLocalizationByRealmByLocaleByKey(ctx, realm, locale, key).Execute()



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
	locale := "locale_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.GetLocalizationByRealmByLocaleByKey(context.Background(), realm, locale, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetLocalizationByRealmByLocaleByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationByRealmByLocaleByKey`: string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetLocalizationByRealmByLocaleByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationByRealmByLocaleByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicies

> ClientPoliciesRepresentation GetPolicies(ctx, realm).Execute()



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
	resp, r, err := apiClient.RealmsAdminAPI.GetPolicies(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicies`: ClientPoliciesRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientPoliciesRepresentation**](ClientPoliciesRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfiles

> ClientProfilesRepresentation GetProfiles(ctx, realm).IncludeGlobalProfiles(includeGlobalProfiles).Execute()



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
	includeGlobalProfiles := "includeGlobalProfiles_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.GetProfiles(context.Background(), realm).IncludeGlobalProfiles(includeGlobalProfiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfiles`: ClientProfilesRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeGlobalProfiles** | **string** |  | 

### Return type

[**ClientProfilesRepresentation**](ClientProfilesRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersManagementPermissions

> ManagementPermissionReference GetUsersManagementPermissions(ctx, realm).Execute()



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
	resp, r, err := apiClient.RealmsAdminAPI.GetUsersManagementPermissions(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.GetUsersManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.GetUsersManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersManagementPermissionsRequest struct via the builder pattern


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


## Post

> Post(ctx).Body(body).Execute()





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
	body := os.NewFile(1234, "some_file") // *os.File | [file] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.Post(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | [file] | 

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


## PostClientDescriptionConverter

> ClientRepresentation PostClientDescriptionConverter(ctx, realm).Body(body).Execute()





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
	body := "body_example" // string | [string] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.PostClientDescriptionConverter(context.Background(), realm).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PostClientDescriptionConverter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientDescriptionConverter`: ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.PostClientDescriptionConverter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientDescriptionConverterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | [string] | 

### Return type

[**ClientRepresentation**](ClientRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLocalization

> PostLocalization(ctx, realm, locale).Body(body).Execute()





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
	locale := "locale_example" // string | 
	body := "body_example" // string | [string] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PostLocalization(context.Background(), realm, locale).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PostLocalization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLocalizationRequest struct via the builder pattern


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


## PostLogoutAll

> GlobalRequestResult PostLogoutAll(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.PostLogoutAll(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PostLogoutAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLogoutAll`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.PostLogoutAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLogoutAllRequest struct via the builder pattern


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


## PostPartialExport

> PostPartialExport(ctx, realm).ExportClients(exportClients).ExportGroupsAndRoles(exportGroupsAndRoles).Execute()





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
	exportClients := "exportClients_example" // string |  (optional)
	exportGroupsAndRoles := "exportGroupsAndRoles_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PostPartialExport(context.Background(), realm).ExportClients(exportClients).ExportGroupsAndRoles(exportGroupsAndRoles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PostPartialExport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostPartialExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportClients** | **string** |  | 
 **exportGroupsAndRoles** | **string** |  | 

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


## PostPartialImport

> PostPartialImport(ctx, realm).Body(body).Execute()





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
	body := os.NewFile(1234, "some_file") // *os.File | [file] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PostPartialImport(context.Background(), realm).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PostPartialImport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostPartialImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | [file] | 

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


## PostPushRevocationByRealm

> GlobalRequestResult PostPushRevocationByRealm(ctx, realm).Execute()





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
	resp, r, err := apiClient.RealmsAdminAPI.PostPushRevocationByRealm(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PostPushRevocationByRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPushRevocationByRealm`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.PostPushRevocationByRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPushRevocationByRealmRequest struct via the builder pattern


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


## PostTestSmtpConnection

> PostTestSmtpConnection(ctx, realm).Body(body).Execute()





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
	body := "body_example" // string | [string] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PostTestSmtpConnection(context.Background(), realm).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PostTestSmtpConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostTestSmtpConnectionRequest struct via the builder pattern


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


## PutByRealm

> PutByRealm(ctx, realm).RealmRepresentation(realmRepresentation).Execute()





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
	realmRepresentation := *openapiclient.NewRealmRepresentation() // RealmRepresentation | RealmRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutByRealm(context.Background(), realm).RealmRepresentation(realmRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutByRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutByRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realmRepresentation** | [**RealmRepresentation**](RealmRepresentation.md) | RealmRepresentation | 

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


## PutDefaultDefaultClientScope

> PutDefaultDefaultClientScope(ctx, realm, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutDefaultDefaultClientScope(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutDefaultDefaultClientScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDefaultDefaultClientScopeRequest struct via the builder pattern


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


## PutDefaultGroup

> PutDefaultGroup(ctx, realm, groupId).Execute()



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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutDefaultGroup(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutDefaultGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDefaultGroupRequest struct via the builder pattern


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


## PutDefaultOptionalClientScope

> PutDefaultOptionalClientScope(ctx, realm, clientScopeId).Execute()



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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutDefaultOptionalClientScope(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutDefaultOptionalClientScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDefaultOptionalClientScopeRequest struct via the builder pattern


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


## PutEventsConfig

> PutEventsConfig(ctx, realm).RealmEventsConfigRepresentation(realmEventsConfigRepresentation).Execute()



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
	realmEventsConfigRepresentation := *openapiclient.NewRealmEventsConfigRepresentation() // RealmEventsConfigRepresentation | RealmEventsConfigRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutEventsConfig(context.Background(), realm).RealmEventsConfigRepresentation(realmEventsConfigRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutEventsConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutEventsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realmEventsConfigRepresentation** | [**RealmEventsConfigRepresentation**](RealmEventsConfigRepresentation.md) | RealmEventsConfigRepresentation | 

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


## PutLocalization

> PutLocalization(ctx, realm, locale, key).Body(body).Execute()



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
	locale := "locale_example" // string | 
	key := "key_example" // string | 
	body := "body_example" // string | [string] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutLocalization(context.Background(), realm, locale, key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutLocalization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLocalizationRequest struct via the builder pattern


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


## PutPolicies

> PutPolicies(ctx, realm).ClientPoliciesRepresentation(clientPoliciesRepresentation).Execute()



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
	clientPoliciesRepresentation := *openapiclient.NewClientPoliciesRepresentation() // ClientPoliciesRepresentation | ClientPoliciesRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutPolicies(context.Background(), realm).ClientPoliciesRepresentation(clientPoliciesRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutPolicies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientPoliciesRepresentation** | [**ClientPoliciesRepresentation**](ClientPoliciesRepresentation.md) | ClientPoliciesRepresentation | 

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


## PutProfiles

> PutProfiles(ctx, realm).ClientProfilesRepresentation(clientProfilesRepresentation).Execute()



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
	clientProfilesRepresentation := *openapiclient.NewClientProfilesRepresentation() // ClientProfilesRepresentation | ClientProfilesRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.PutProfiles(context.Background(), realm).ClientProfilesRepresentation(clientProfilesRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutProfiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientProfilesRepresentation** | [**ClientProfilesRepresentation**](ClientProfilesRepresentation.md) | ClientProfilesRepresentation | 

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


## PutUsersManagementPermissions

> ManagementPermissionReference PutUsersManagementPermissions(ctx, realm).ManagementPermissionReference(managementPermissionReference).Execute()



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
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference | ManagementPermissionReference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.PutUsersManagementPermissions(context.Background(), realm).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.PutUsersManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUsersManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.PutUsersManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUsersManagementPermissionsRequest struct via the builder pattern


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

