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

// RolesByIDAPIService RolesByIDAPI service
type RolesByIDAPIService service

type ApiDeleteRolesByIdRequest struct {
	ctx        context.Context
	ApiService *RolesByIDAPIService
	realm      string
	roleId     string
}

func (r ApiDeleteRolesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRolesByIdExecute(r)
}

/*
DeleteRolesById Method for DeleteRolesById

Delete the role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId id of role
	@return ApiDeleteRolesByIdRequest
*/
func (a *RolesByIDAPIService) DeleteRolesById(ctx context.Context, realm string, roleId string) ApiDeleteRolesByIdRequest {
	return ApiDeleteRolesByIdRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
func (a *RolesByIDAPIService) DeleteRolesByIdExecute(r ApiDeleteRolesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.DeleteRolesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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

type ApiDeleteRolesByIdCompositesRequest struct {
	ctx                context.Context
	ApiService         *RolesByIDAPIService
	realm              string
	roleId             string
	roleRepresentation *RoleRepresentation
}

// RoleRepresentation
func (r ApiDeleteRolesByIdCompositesRequest) RoleRepresentation(roleRepresentation RoleRepresentation) ApiDeleteRolesByIdCompositesRequest {
	r.roleRepresentation = &roleRepresentation
	return r
}

func (r ApiDeleteRolesByIdCompositesRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRolesByIdCompositesExecute(r)
}

/*
DeleteRolesByIdComposites Method for DeleteRolesByIdComposites

Remove a set of roles from the role’s composite

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId
	@return ApiDeleteRolesByIdCompositesRequest
*/
func (a *RolesByIDAPIService) DeleteRolesByIdComposites(ctx context.Context, realm string, roleId string) ApiDeleteRolesByIdCompositesRequest {
	return ApiDeleteRolesByIdCompositesRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
func (a *RolesByIDAPIService) DeleteRolesByIdCompositesExecute(r ApiDeleteRolesByIdCompositesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.DeleteRolesByIdComposites")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}/composites"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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
	localVarPostBody = r.roleRepresentation
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

type ApiGetRolesByIdRequest struct {
	ctx        context.Context
	ApiService *RolesByIDAPIService
	realm      string
	roleId     string
}

func (r ApiGetRolesByIdRequest) Execute() (*RoleRepresentation, *http.Response, error) {
	return r.ApiService.GetRolesByIdExecute(r)
}

/*
GetRolesById Method for GetRolesById

Get a specific role’s representation

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId id of role
	@return ApiGetRolesByIdRequest
*/
func (a *RolesByIDAPIService) GetRolesById(ctx context.Context, realm string, roleId string) ApiGetRolesByIdRequest {
	return ApiGetRolesByIdRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
//
//	@return RoleRepresentation
func (a *RolesByIDAPIService) GetRolesByIdExecute(r ApiGetRolesByIdRequest) (*RoleRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RoleRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.GetRolesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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

type ApiGetRolesByIdCompositesRequest struct {
	ctx        context.Context
	ApiService *RolesByIDAPIService
	realm      string
	roleId     string
	first      *string
	max        *string
	search     *string
}

func (r ApiGetRolesByIdCompositesRequest) First(first string) ApiGetRolesByIdCompositesRequest {
	r.first = &first
	return r
}

func (r ApiGetRolesByIdCompositesRequest) Max(max string) ApiGetRolesByIdCompositesRequest {
	r.max = &max
	return r
}

func (r ApiGetRolesByIdCompositesRequest) Search(search string) ApiGetRolesByIdCompositesRequest {
	r.search = &search
	return r
}

func (r ApiGetRolesByIdCompositesRequest) Execute() ([]RoleRepresentation, *http.Response, error) {
	return r.ApiService.GetRolesByIdCompositesExecute(r)
}

/*
GetRolesByIdComposites Method for GetRolesByIdComposites

Get role’s children Returns a set of role’s children provided the role is a composite.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId
	@return ApiGetRolesByIdCompositesRequest
*/
func (a *RolesByIDAPIService) GetRolesByIdComposites(ctx context.Context, realm string, roleId string) ApiGetRolesByIdCompositesRequest {
	return ApiGetRolesByIdCompositesRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
//
//	@return []RoleRepresentation
func (a *RolesByIDAPIService) GetRolesByIdCompositesExecute(r ApiGetRolesByIdCompositesRequest) ([]RoleRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []RoleRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.GetRolesByIdComposites")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}/composites"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.first != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "first", r.first, "form", "")
	}
	if r.max != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max", r.max, "form", "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
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

type ApiGetRolesByIdCompositesClientRequest struct {
	ctx        context.Context
	ApiService *RolesByIDAPIService
	realm      string
	roleId     string
	clientUuid string
}

func (r ApiGetRolesByIdCompositesClientRequest) Execute() ([]RoleRepresentation, *http.Response, error) {
	return r.ApiService.GetRolesByIdCompositesClientExecute(r)
}

/*
GetRolesByIdCompositesClient Method for GetRolesByIdCompositesClient

Get client-level roles for the client that are in the role’s composite

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId
	@param clientUuid
	@return ApiGetRolesByIdCompositesClientRequest
*/
func (a *RolesByIDAPIService) GetRolesByIdCompositesClient(ctx context.Context, realm string, roleId string, clientUuid string) ApiGetRolesByIdCompositesClientRequest {
	return ApiGetRolesByIdCompositesClientRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
		clientUuid: clientUuid,
	}
}

// Execute executes the request
//
//	@return []RoleRepresentation
func (a *RolesByIDAPIService) GetRolesByIdCompositesClientExecute(r ApiGetRolesByIdCompositesClientRequest) ([]RoleRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []RoleRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.GetRolesByIdCompositesClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}/composites/clients/{clientUuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientUuid"+"}", url.PathEscape(parameterValueToString(r.clientUuid, "clientUuid")), -1)

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

type ApiGetRolesByIdCompositesRealmRequest struct {
	ctx        context.Context
	ApiService *RolesByIDAPIService
	realm      string
	roleId     string
}

func (r ApiGetRolesByIdCompositesRealmRequest) Execute() ([]RoleRepresentation, *http.Response, error) {
	return r.ApiService.GetRolesByIdCompositesRealmExecute(r)
}

/*
GetRolesByIdCompositesRealm Method for GetRolesByIdCompositesRealm

Get realm-level roles that are in the role’s composite

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId
	@return ApiGetRolesByIdCompositesRealmRequest
*/
func (a *RolesByIDAPIService) GetRolesByIdCompositesRealm(ctx context.Context, realm string, roleId string) ApiGetRolesByIdCompositesRealmRequest {
	return ApiGetRolesByIdCompositesRealmRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
//
//	@return []RoleRepresentation
func (a *RolesByIDAPIService) GetRolesByIdCompositesRealmExecute(r ApiGetRolesByIdCompositesRealmRequest) ([]RoleRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []RoleRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.GetRolesByIdCompositesRealm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}/composites/realm"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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

type ApiGetRolesByIdManagementPermissionsRequest struct {
	ctx        context.Context
	ApiService *RolesByIDAPIService
	realm      string
	roleId     string
}

func (r ApiGetRolesByIdManagementPermissionsRequest) Execute() (*ManagementPermissionReference, *http.Response, error) {
	return r.ApiService.GetRolesByIdManagementPermissionsExecute(r)
}

/*
GetRolesByIdManagementPermissions Method for GetRolesByIdManagementPermissions

Return object stating whether role Authorization permissions have been initialized or not and a reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId
	@return ApiGetRolesByIdManagementPermissionsRequest
*/
func (a *RolesByIDAPIService) GetRolesByIdManagementPermissions(ctx context.Context, realm string, roleId string) ApiGetRolesByIdManagementPermissionsRequest {
	return ApiGetRolesByIdManagementPermissionsRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
//
//	@return ManagementPermissionReference
func (a *RolesByIDAPIService) GetRolesByIdManagementPermissionsExecute(r ApiGetRolesByIdManagementPermissionsRequest) (*ManagementPermissionReference, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ManagementPermissionReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.GetRolesByIdManagementPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}/management/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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

type ApiPostRolesByIdCompositesRequest struct {
	ctx                context.Context
	ApiService         *RolesByIDAPIService
	realm              string
	roleId             string
	roleRepresentation *RoleRepresentation
}

// RoleRepresentation
func (r ApiPostRolesByIdCompositesRequest) RoleRepresentation(roleRepresentation RoleRepresentation) ApiPostRolesByIdCompositesRequest {
	r.roleRepresentation = &roleRepresentation
	return r
}

func (r ApiPostRolesByIdCompositesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostRolesByIdCompositesExecute(r)
}

/*
PostRolesByIdComposites Method for PostRolesByIdComposites

Make the role a composite role by associating some child roles

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId
	@return ApiPostRolesByIdCompositesRequest
*/
func (a *RolesByIDAPIService) PostRolesByIdComposites(ctx context.Context, realm string, roleId string) ApiPostRolesByIdCompositesRequest {
	return ApiPostRolesByIdCompositesRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
func (a *RolesByIDAPIService) PostRolesByIdCompositesExecute(r ApiPostRolesByIdCompositesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.PostRolesByIdComposites")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}/composites"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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
	localVarPostBody = r.roleRepresentation
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

type ApiPutRolesByIdRequest struct {
	ctx                context.Context
	ApiService         *RolesByIDAPIService
	realm              string
	roleId             string
	roleRepresentation *RoleRepresentation
}

// RoleRepresentation
func (r ApiPutRolesByIdRequest) RoleRepresentation(roleRepresentation RoleRepresentation) ApiPutRolesByIdRequest {
	r.roleRepresentation = &roleRepresentation
	return r
}

func (r ApiPutRolesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutRolesByIdExecute(r)
}

/*
PutRolesById Method for PutRolesById

Update the role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId id of role
	@return ApiPutRolesByIdRequest
*/
func (a *RolesByIDAPIService) PutRolesById(ctx context.Context, realm string, roleId string) ApiPutRolesByIdRequest {
	return ApiPutRolesByIdRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
func (a *RolesByIDAPIService) PutRolesByIdExecute(r ApiPutRolesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.PutRolesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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
	localVarPostBody = r.roleRepresentation
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

type ApiPutRolesByIdManagementPermissionsRequest struct {
	ctx                           context.Context
	ApiService                    *RolesByIDAPIService
	realm                         string
	roleId                        string
	managementPermissionReference *ManagementPermissionReference
}

// ManagementPermissionReference
func (r ApiPutRolesByIdManagementPermissionsRequest) ManagementPermissionReference(managementPermissionReference ManagementPermissionReference) ApiPutRolesByIdManagementPermissionsRequest {
	r.managementPermissionReference = &managementPermissionReference
	return r
}

func (r ApiPutRolesByIdManagementPermissionsRequest) Execute() (*ManagementPermissionReference, *http.Response, error) {
	return r.ApiService.PutRolesByIdManagementPermissionsExecute(r)
}

/*
PutRolesByIdManagementPermissions Method for PutRolesByIdManagementPermissions

Return object stating whether role Authorization permissions have been initialized or not and a reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realm realm name (not id!)
	@param roleId
	@return ApiPutRolesByIdManagementPermissionsRequest
*/
func (a *RolesByIDAPIService) PutRolesByIdManagementPermissions(ctx context.Context, realm string, roleId string) ApiPutRolesByIdManagementPermissionsRequest {
	return ApiPutRolesByIdManagementPermissionsRequest{
		ApiService: a,
		ctx:        ctx,
		realm:      realm,
		roleId:     roleId,
	}
}

// Execute executes the request
//
//	@return ManagementPermissionReference
func (a *RolesByIDAPIService) PutRolesByIdManagementPermissionsExecute(r ApiPutRolesByIdManagementPermissionsRequest) (*ManagementPermissionReference, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ManagementPermissionReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesByIDAPIService.PutRolesByIdManagementPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realm}/roles-by-id/{role-id}/management/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"role-id"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.managementPermissionReference
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
