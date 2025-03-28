/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ClientScopesAPIService ClientScopesAPI service
type ClientScopesAPIService service

type ApiDeleteClientScopeRequest struct {
	ctx        context.Context
	ApiService *ClientScopesAPIService
	realm      string
	id         string
}

func (r ApiDeleteClientScopeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClientScopeExecute(r)
}

/*
DeleteClientScope Method for DeleteClientScope

Delete the client scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param id
	@return ApiDeleteClientScopeRequest
*/
func (a *ClientScopesAPIService) DeleteClientScope(ctx context.Context, realm string, id string) ApiDeleteClientScopeRequest {
	return ApiDeleteClientScopeRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		id:         id,
	}
}

// Execute executes the request
func (a *ClientScopesAPIService) DeleteClientScopeExecute(r ApiDeleteClientScopeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.DeleteClientScope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-scopes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteClientTemplateRequest struct {
	ctx        context.Context
	ApiService *ClientScopesAPIService
	realm      string
	id         string
}

func (r ApiDeleteClientTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClientTemplateExecute(r)
}

/*
DeleteClientTemplate Method for DeleteClientTemplate

Delete the client scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param id
	@return ApiDeleteClientTemplateRequest
*/
func (a *ClientScopesAPIService) DeleteClientTemplate(ctx context.Context, realm string, id string) ApiDeleteClientTemplateRequest {
	return ApiDeleteClientTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		id:         id,
	}
}

// Execute executes the request
func (a *ClientScopesAPIService) DeleteClientTemplateExecute(r ApiDeleteClientTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.DeleteClientTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetClientScopeRequest struct {
	ctx        context.Context
	ApiService *ClientScopesAPIService
	realm      string
	id         string
}

func (r ApiGetClientScopeRequest) Execute() (*ClientScopeRepresentation, *http.Response, error) {
	return r.ApiService.GetClientScopeExecute(r)
}

/*
GetClientScope Method for GetClientScope

Get representation of the client scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param id
	@return ApiGetClientScopeRequest
*/
func (a *ClientScopesAPIService) GetClientScope(ctx context.Context, realm string, id string) ApiGetClientScopeRequest {
	return ApiGetClientScopeRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ClientScopeRepresentation
func (a *ClientScopesAPIService) GetClientScopeExecute(r ApiGetClientScopeRequest) (*ClientScopeRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientScopeRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.GetClientScope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-scopes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClientScopesRequest struct {
	ctx        context.Context
	ApiService *ClientScopesAPIService
	realm      string
}

func (r ApiGetClientScopesRequest) Execute() ([]ClientScopeRepresentation, *http.Response, error) {
	return r.ApiService.GetClientScopesExecute(r)
}

/*
GetClientScopes Method for GetClientScopes

Get client scopes belonging to the realm Returns a list of client scopes belonging to the realm

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@return ApiGetClientScopesRequest
*/
func (a *ClientScopesAPIService) GetClientScopes(ctx context.Context, realm string) ApiGetClientScopesRequest {
	return ApiGetClientScopesRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
	}
}

// Execute executes the request
//
//	@return []ClientScopeRepresentation
func (a *ClientScopesAPIService) GetClientScopesExecute(r ApiGetClientScopesRequest) ([]ClientScopeRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ClientScopeRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.GetClientScopes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClientTemplateRequest struct {
	ctx        context.Context
	ApiService *ClientScopesAPIService
	realm      string
	id         string
}

func (r ApiGetClientTemplateRequest) Execute() (*ClientScopeRepresentation, *http.Response, error) {
	return r.ApiService.GetClientTemplateExecute(r)
}

/*
GetClientTemplate Method for GetClientTemplate

Get representation of the client scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param id
	@return ApiGetClientTemplateRequest
*/
func (a *ClientScopesAPIService) GetClientTemplate(ctx context.Context, realm string, id string) ApiGetClientTemplateRequest {
	return ApiGetClientTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ClientScopeRepresentation
func (a *ClientScopesAPIService) GetClientTemplateExecute(r ApiGetClientTemplateRequest) (*ClientScopeRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientScopeRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.GetClientTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClientTemplatesRequest struct {
	ctx        context.Context
	ApiService *ClientScopesAPIService
	realm      string
}

func (r ApiGetClientTemplatesRequest) Execute() ([]ClientScopeRepresentation, *http.Response, error) {
	return r.ApiService.GetClientTemplatesExecute(r)
}

/*
GetClientTemplates Method for GetClientTemplates

Get client scopes belonging to the realm Returns a list of client scopes belonging to the realm

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@return ApiGetClientTemplatesRequest
*/
func (a *ClientScopesAPIService) GetClientTemplates(ctx context.Context, realm string) ApiGetClientTemplatesRequest {
	return ApiGetClientTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
	}
}

// Execute executes the request
//
//	@return []ClientScopeRepresentation
func (a *ClientScopesAPIService) GetClientTemplatesExecute(r ApiGetClientTemplatesRequest) ([]ClientScopeRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ClientScopeRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.GetClientTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostClientScopesRequest struct {
	ctx                       context.Context
	ApiService                *ClientScopesAPIService
	realm                     string
	clientScopeRepresentation *ClientScopeRepresentation
}

// ClientScopeRepresentation
func (r ApiPostClientScopesRequest) ClientScopeRepresentation(clientScopeRepresentation ClientScopeRepresentation) ApiPostClientScopesRequest {
	r.clientScopeRepresentation = &clientScopeRepresentation
	return r
}

func (r ApiPostClientScopesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostClientScopesExecute(r)
}

/*
PostClientScopes Method for PostClientScopes

Create a new client scope Client Scope’s name must be unique!

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@return ApiPostClientScopesRequest
*/
func (a *ClientScopesAPIService) PostClientScopes(ctx context.Context, realm string) ApiPostClientScopesRequest {
	return ApiPostClientScopesRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
	}
}

// Execute executes the request
func (a *ClientScopesAPIService) PostClientScopesExecute(r ApiPostClientScopesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.PostClientScopes")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clientScopeRepresentation
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostClientTemplatesRequest struct {
	ctx                       context.Context
	ApiService                *ClientScopesAPIService
	realm                     string
	clientScopeRepresentation *ClientScopeRepresentation
}

// ClientScopeRepresentation
func (r ApiPostClientTemplatesRequest) ClientScopeRepresentation(clientScopeRepresentation ClientScopeRepresentation) ApiPostClientTemplatesRequest {
	r.clientScopeRepresentation = &clientScopeRepresentation
	return r
}

func (r ApiPostClientTemplatesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostClientTemplatesExecute(r)
}

/*
PostClientTemplates Method for PostClientTemplates

Create a new client scope Client Scope’s name must be unique!

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@return ApiPostClientTemplatesRequest
*/
func (a *ClientScopesAPIService) PostClientTemplates(ctx context.Context, realm string) ApiPostClientTemplatesRequest {
	return ApiPostClientTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
	}
}

// Execute executes the request
func (a *ClientScopesAPIService) PostClientTemplatesExecute(r ApiPostClientTemplatesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.PostClientTemplates")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clientScopeRepresentation
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPutClientScopeRequest struct {
	ctx                       context.Context
	ApiService                *ClientScopesAPIService
	realm                     string
	id                        string
	clientScopeRepresentation *ClientScopeRepresentation
}

// ClientScopeRepresentation
func (r ApiPutClientScopeRequest) ClientScopeRepresentation(clientScopeRepresentation ClientScopeRepresentation) ApiPutClientScopeRequest {
	r.clientScopeRepresentation = &clientScopeRepresentation
	return r
}

func (r ApiPutClientScopeRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutClientScopeExecute(r)
}

/*
PutClientScope Method for PutClientScope

Update the client scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param id
	@return ApiPutClientScopeRequest
*/
func (a *ClientScopesAPIService) PutClientScope(ctx context.Context, realm string, id string) ApiPutClientScopeRequest {
	return ApiPutClientScopeRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		id:         id,
	}
}

// Execute executes the request
func (a *ClientScopesAPIService) PutClientScopeExecute(r ApiPutClientScopeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.PutClientScope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-scopes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clientScopeRepresentation
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPutClientTemplateRequest struct {
	ctx                       context.Context
	ApiService                *ClientScopesAPIService
	realm                     string
	id                        string
	clientScopeRepresentation *ClientScopeRepresentation
}

// ClientScopeRepresentation
func (r ApiPutClientTemplateRequest) ClientScopeRepresentation(clientScopeRepresentation ClientScopeRepresentation) ApiPutClientTemplateRequest {
	r.clientScopeRepresentation = &clientScopeRepresentation
	return r
}

func (r ApiPutClientTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutClientTemplateExecute(r)
}

/*
PutClientTemplate Method for PutClientTemplate

Update the client scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param id
	@return ApiPutClientTemplateRequest
*/
func (a *ClientScopesAPIService) PutClientTemplate(ctx context.Context, realm string, id string) ApiPutClientTemplateRequest {
	return ApiPutClientTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		id:         id,
	}
}

// Execute executes the request
func (a *ClientScopesAPIService) PutClientTemplateExecute(r ApiPutClientTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientScopesAPIService.PutClientTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/client-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clientScopeRepresentation
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
