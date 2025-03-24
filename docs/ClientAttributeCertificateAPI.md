# \ClientAttributeCertificateAPI

All URIs are relative to *https://keycloak.example.com/admin/realms*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCertificate**](ClientAttributeCertificateAPI.md#GetCertificate) | **Get** /{realm}/clients/{id}/certificates/{attr} | 
[**PostDownload**](ClientAttributeCertificateAPI.md#PostDownload) | **Post** /{realm}/clients/{id}/certificates/{attr}/download | 
[**PostGenerate**](ClientAttributeCertificateAPI.md#PostGenerate) | **Post** /{realm}/clients/{id}/certificates/{attr}/generate | 
[**PostGenerateAndDownload**](ClientAttributeCertificateAPI.md#PostGenerateAndDownload) | **Post** /{realm}/clients/{id}/certificates/{attr}/generate-and-download | 
[**PostUpload**](ClientAttributeCertificateAPI.md#PostUpload) | **Post** /{realm}/clients/{id}/certificates/{attr}/upload | 
[**PostUploadCertificate**](ClientAttributeCertificateAPI.md#PostUploadCertificate) | **Post** /{realm}/clients/{id}/certificates/{attr}/upload-certificate | 



## GetCertificate

> CertificateRepresentation GetCertificate(ctx, realm, id, attr).Execute()





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
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.GetCertificate(context.Background(), realm, id, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.GetCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificate`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDownload

> *os.File PostDownload(ctx, realm, id, attr).KeyStoreConfig(keyStoreConfig).Execute()





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
	attr := "attr_example" // string | 
	keyStoreConfig := *openapiclient.NewKeyStoreConfig() // KeyStoreConfig | KeyStoreConfig (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.PostDownload(context.Background(), realm, id, attr).KeyStoreConfig(keyStoreConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.PostDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDownload`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.PostDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **keyStoreConfig** | [**KeyStoreConfig**](KeyStoreConfig.md) | KeyStoreConfig | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGenerate

> CertificateRepresentation PostGenerate(ctx, realm, id, attr).Execute()





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
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.PostGenerate(context.Background(), realm, id, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.PostGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGenerate`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.PostGenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGenerateAndDownload

> *os.File PostGenerateAndDownload(ctx, realm, id, attr).KeyStoreConfig(keyStoreConfig).Execute()





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
	attr := "attr_example" // string | 
	keyStoreConfig := *openapiclient.NewKeyStoreConfig() // KeyStoreConfig | KeyStoreConfig (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.PostGenerateAndDownload(context.Background(), realm, id, attr).KeyStoreConfig(keyStoreConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.PostGenerateAndDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGenerateAndDownload`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.PostGenerateAndDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGenerateAndDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **keyStoreConfig** | [**KeyStoreConfig**](KeyStoreConfig.md) | KeyStoreConfig | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUpload

> CertificateRepresentation PostUpload(ctx, realm, id, attr).Execute()





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
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.PostUpload(context.Background(), realm, id, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.PostUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUpload`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.PostUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUploadCertificate

> CertificateRepresentation PostUploadCertificate(ctx, realm, id, attr).Execute()





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
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.PostUploadCertificate(context.Background(), realm, id, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.PostUploadCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUploadCertificate`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.PostUploadCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUploadCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

