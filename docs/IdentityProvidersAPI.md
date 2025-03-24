# \IdentityProvidersAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInstance**](IdentityProvidersAPI.md#DeleteInstance) | **Delete** /{realm}/identity-provider/instances/{alias} | 
[**DeleteMapper**](IdentityProvidersAPI.md#DeleteMapper) | **Delete** /{realm}/identity-provider/instances/{alias}/mappers/{id} | 
[**GetExport**](IdentityProvidersAPI.md#GetExport) | **Get** /{realm}/identity-provider/instances/{alias}/export | 
[**GetIdentityProviderProvider**](IdentityProvidersAPI.md#GetIdentityProviderProvider) | **Get** /{realm}/identity-provider/providers/{provider_id} | 
[**GetInstance**](IdentityProvidersAPI.md#GetInstance) | **Get** /{realm}/identity-provider/instances/{alias} | 
[**GetInstanceManagementPermissions**](IdentityProvidersAPI.md#GetInstanceManagementPermissions) | **Get** /{realm}/identity-provider/instances/{alias}/management/permissions | 
[**GetInstances**](IdentityProvidersAPI.md#GetInstances) | **Get** /{realm}/identity-provider/instances | 
[**GetMapper**](IdentityProvidersAPI.md#GetMapper) | **Get** /{realm}/identity-provider/instances/{alias}/mappers/{id} | 
[**GetMapperTypes**](IdentityProvidersAPI.md#GetMapperTypes) | **Get** /{realm}/identity-provider/instances/{alias}/mapper-types | 
[**GetMappers**](IdentityProvidersAPI.md#GetMappers) | **Get** /{realm}/identity-provider/instances/{alias}/mappers | 
[**PostImportConfig**](IdentityProvidersAPI.md#PostImportConfig) | **Post** /{realm}/identity-provider/import-config | 
[**PostInstances**](IdentityProvidersAPI.md#PostInstances) | **Post** /{realm}/identity-provider/instances | 
[**PostMappers**](IdentityProvidersAPI.md#PostMappers) | **Post** /{realm}/identity-provider/instances/{alias}/mappers | 
[**PutInstance**](IdentityProvidersAPI.md#PutInstance) | **Put** /{realm}/identity-provider/instances/{alias} | 
[**PutInstanceManagementPermissions**](IdentityProvidersAPI.md#PutInstanceManagementPermissions) | **Put** /{realm}/identity-provider/instances/{alias}/management/permissions | 
[**PutMapper**](IdentityProvidersAPI.md#PutMapper) | **Put** /{realm}/identity-provider/instances/{alias}/mappers/{id} | 



## DeleteInstance

> DeleteInstance(ctx, realm, alias).Execute()





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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.DeleteInstance(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## DeleteMapper

> DeleteMapper(ctx, realm, alias, id).Execute()





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
	alias := "alias_example" // string | 
	id := "id_example" // string | Mapper id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.DeleteMapper(context.Background(), realm, alias, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.DeleteMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 
**id** | **string** | Mapper id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMapperRequest struct via the builder pattern


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


## GetExport

> GetExport(ctx, realm, alias).Format(format).Execute()





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
	alias := "alias_example" // string | 
	format := "format_example" // string | Format to use (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.GetExport(context.Background(), realm, alias).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Format to use | 

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


## GetIdentityProviderProvider

> map[string]interface{} GetIdentityProviderProvider(ctx, realm, providerId).Execute()





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
	providerId := "providerId_example" // string | The provider id to get the factory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.GetIdentityProviderProvider(context.Background(), realm, providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetIdentityProviderProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityProviderProvider`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.GetIdentityProviderProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**providerId** | **string** | The provider id to get the factory | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> IdentityProviderRepresentation GetInstance(ctx, realm, alias).Execute()





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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.GetInstance(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityProviderRepresentation**](IdentityProviderRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceManagementPermissions

> ManagementPermissionReference GetInstanceManagementPermissions(ctx, realm, alias).Execute()





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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.GetInstanceManagementPermissions(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetInstanceManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.GetInstanceManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceManagementPermissionsRequest struct via the builder pattern


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


## GetInstances

> []IdentityProviderRepresentation GetInstances(ctx, realm).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()





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
	briefRepresentation := "briefRepresentation_example" // string | Boolean which defines whether brief representations are returned (default: false) (optional)
	first := "first_example" // string | Pagination offset (optional)
	max := "max_example" // string | Maximum results size (defaults to 100) (optional)
	search := "search_example" // string | Filter specific providers by name. Search can be prefix (name*), contains (name) or exact (&quot;name&quot;). Default prefixed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.GetInstances(context.Background(), realm).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstances`: []IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.GetInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **briefRepresentation** | **string** | Boolean which defines whether brief representations are returned (default: false) | 
 **first** | **string** | Pagination offset | 
 **max** | **string** | Maximum results size (defaults to 100) | 
 **search** | **string** | Filter specific providers by name. Search can be prefix (name*), contains (name) or exact (&amp;quot;name&amp;quot;). Default prefixed. | 

### Return type

[**[]IdentityProviderRepresentation**](IdentityProviderRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapper

> IdentityProviderMapperRepresentation GetMapper(ctx, realm, alias, id).Execute()





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
	alias := "alias_example" // string | 
	id := "id_example" // string | Mapper id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.GetMapper(context.Background(), realm, alias, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMapper`: IdentityProviderMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.GetMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 
**id** | **string** | Mapper id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapperTypes

> map[string]IdentityProviderMapperTypeRepresentation GetMapperTypes(ctx, realm, alias).Execute()





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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.GetMapperTypes(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetMapperTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMapperTypes`: map[string]IdentityProviderMapperTypeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.GetMapperTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMapperTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]IdentityProviderMapperTypeRepresentation**](IdentityProviderMapperTypeRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMappers

> []IdentityProviderMapperRepresentation GetMappers(ctx, realm, alias).Execute()





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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.GetMappers(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.GetMappers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMappers`: []IdentityProviderMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.GetMappers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImportConfig

> map[string]string PostImportConfig(ctx, realm).Body(body).Execute()





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
	body := map[string]interface{}{ ... } // map[string]interface{} | [AnyType] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.PostImportConfig(context.Background(), realm).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.PostImportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImportConfig`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.PostImportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostImportConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | [AnyType] | 

### Return type

**map[string]string**

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstances

> PostInstances(ctx, realm).IdentityProviderRepresentation(identityProviderRepresentation).Execute()





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
	identityProviderRepresentation := *openapiclient.NewIdentityProviderRepresentation() // IdentityProviderRepresentation | IdentityProviderRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.PostInstances(context.Background(), realm).IdentityProviderRepresentation(identityProviderRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.PostInstances``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityProviderRepresentation** | [**IdentityProviderRepresentation**](IdentityProviderRepresentation.md) | IdentityProviderRepresentation | 

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


## PostMappers

> PostMappers(ctx, realm, alias).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()





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
	alias := "alias_example" // string | 
	identityProviderMapperRepresentation := *openapiclient.NewIdentityProviderMapperRepresentation() // IdentityProviderMapperRepresentation | IdentityProviderMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.PostMappers(context.Background(), realm, alias).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.PostMappers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityProviderMapperRepresentation** | [**IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md) | IdentityProviderMapperRepresentation | 

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


## PutInstance

> PutInstance(ctx, realm, alias).IdentityProviderRepresentation(identityProviderRepresentation).Execute()





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
	alias := "alias_example" // string | 
	identityProviderRepresentation := *openapiclient.NewIdentityProviderRepresentation() // IdentityProviderRepresentation | IdentityProviderRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.PutInstance(context.Background(), realm, alias).IdentityProviderRepresentation(identityProviderRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.PutInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityProviderRepresentation** | [**IdentityProviderRepresentation**](IdentityProviderRepresentation.md) | IdentityProviderRepresentation | 

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


## PutInstanceManagementPermissions

> ManagementPermissionReference PutInstanceManagementPermissions(ctx, realm, alias).ManagementPermissionReference(managementPermissionReference).Execute()





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
	alias := "alias_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference | ManagementPermissionReference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.PutInstanceManagementPermissions(context.Background(), realm, alias).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.PutInstanceManagementPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutInstanceManagementPermissions`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.PutInstanceManagementPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutInstanceManagementPermissionsRequest struct via the builder pattern


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


## PutMapper

> PutMapper(ctx, realm, alias, id).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()





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
	alias := "alias_example" // string | 
	id := "id_example" // string | Mapper id
	identityProviderMapperRepresentation := *openapiclient.NewIdentityProviderMapperRepresentation() // IdentityProviderMapperRepresentation | IdentityProviderMapperRepresentation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.PutMapper(context.Background(), realm, alias, id).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.PutMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 
**id** | **string** | Mapper id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **identityProviderMapperRepresentation** | [**IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md) | IdentityProviderMapperRepresentation | 

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

