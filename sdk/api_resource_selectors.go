/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ResourceSelectorsAPI interface {

	/*
		CreateResourceSelector Create a Resource Selector

		Creates a Resource Selector

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateResourceSelectorRequest
	*/
	CreateResourceSelector(ctx context.Context) ApiCreateResourceSelectorRequest

	// CreateResourceSelectorExecute executes the request
	//  @return ResourceSelectorResponseSchema
	// TODU
	CreateResourceSelectorExecute(r ApiCreateResourceSelectorRequest) (*APIResponse, error)

	/*
		DeleteResourceSelector Delete a Resource Selector

		Deletes a Resource Selector by `resourceSelectorId`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param resourceSelectorId `id` of a Resource Selector
		@return ApiDeleteResourceSelectorRequest
	*/
	DeleteResourceSelector(ctx context.Context, resourceSelectorId string) ApiDeleteResourceSelectorRequest

	// DeleteResourceSelectorExecute executes the request
	// TODU
	DeleteResourceSelectorExecute(r ApiDeleteResourceSelectorRequest) (*APIResponse, error)

	/*
		GetResourceSelector Retrieve a Resource Selector

		Retrieves a Resource Selector by `resourceSelectorId`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param resourceSelectorId `id` of a Resource Selector
		@return ApiGetResourceSelectorRequest
	*/
	GetResourceSelector(ctx context.Context, resourceSelectorId string) ApiGetResourceSelectorRequest

	// GetResourceSelectorExecute executes the request
	//  @return ResourceSelectorResponseSchema
	// TODU
	GetResourceSelectorExecute(r ApiGetResourceSelectorRequest) (*APIResponse, error)

	/*
		ListResourceSelectors List all Resource Selectors

		Lists all Resource Selectors

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListResourceSelectorsRequest
	*/
	ListResourceSelectors(ctx context.Context) ApiListResourceSelectorsRequest

	// ListResourceSelectorsExecute executes the request
	//  @return ResourceSelectorsSchema
	// TODU
	ListResourceSelectorsExecute(r ApiListResourceSelectorsRequest) (*APIResponse, error)

	/*
		UpdateResourceSelector Update a Resource Selector

		Updates a Resource Selector by `resourceSelectorId`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param resourceSelectorId `id` of a Resource Selector
		@return ApiUpdateResourceSelectorRequest
	*/
	UpdateResourceSelector(ctx context.Context, resourceSelectorId string) ApiUpdateResourceSelectorRequest

	// UpdateResourceSelectorExecute executes the request
	//  @return ResourceSelectorResponseSchema
	// TODU
	UpdateResourceSelectorExecute(r ApiUpdateResourceSelectorRequest) (*APIResponse, error)
}

// ResourceSelectorsAPIService ResourceSelectorsAPI service
type ResourceSelectorsAPIService service

type ApiCreateResourceSelectorRequest struct {
	ctx        context.Context
	ApiService ResourceSelectorsAPI
	instance   *ResourceSelectorCreateRequestSchema
	// TODU
	data       interface{}
	retryCount int32
}

func (r ApiCreateResourceSelectorRequest) Instance(instance ResourceSelectorCreateRequestSchema) ApiCreateResourceSelectorRequest {
	r.instance = &instance
	return r
}

// TODU
func (r ApiCreateResourceSelectorRequest) Data(data interface{}) ApiCreateResourceSelectorRequest {
	r.data = data
	return r
}

// TODU
func (r ApiCreateResourceSelectorRequest) Execute() (*APIResponse, error) {
	return r.ApiService.CreateResourceSelectorExecute(r)
}

/*
CreateResourceSelector Create a Resource Selector

Creates a Resource Selector

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateResourceSelectorRequest
*/
// TODU

func (a *ResourceSelectorsAPIService) CreateResourceSelector(ctx context.Context) ApiCreateResourceSelectorRequest {
	return ApiCreateResourceSelectorRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ResourceSelectorResponseSchema

func (a *ResourceSelectorsAPIService) CreateResourceSelectorExecute(r ApiCreateResourceSelectorRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSelectorsAPIService.CreateResourceSelector")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/resource-selectors"

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
	// TODU
	// body params
	// localVarPostBody = r.instance
	localVarPostBody = r.data
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiDeleteResourceSelectorRequest struct {
	ctx                context.Context
	ApiService         ResourceSelectorsAPI
	resourceSelectorId string
	// TODU
	data       interface{}
	retryCount int32
}

// TODU
func (r ApiDeleteResourceSelectorRequest) Data(data interface{}) ApiDeleteResourceSelectorRequest {
	r.data = data
	return r
}

// TODU
func (r ApiDeleteResourceSelectorRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteResourceSelectorExecute(r)
}

/*
DeleteResourceSelector Delete a Resource Selector

Deletes a Resource Selector by `resourceSelectorId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSelectorId `id` of a Resource Selector
 @return ApiDeleteResourceSelectorRequest
*/
// TODU

func (a *ResourceSelectorsAPIService) DeleteResourceSelector(ctx context.Context, resourceSelectorId string) ApiDeleteResourceSelectorRequest {
	return ApiDeleteResourceSelectorRequest{
		ApiService:         a,
		ctx:                ctx,
		resourceSelectorId: resourceSelectorId,
		retryCount:         0,
	}
}

// Execute executes the request

func (a *ResourceSelectorsAPIService) DeleteResourceSelectorExecute(r ApiDeleteResourceSelectorRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSelectorsAPIService.DeleteResourceSelector")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/resource-selectors/{resourceSelectorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSelectorId"+"}", url.PathEscape(parameterToString(r.resourceSelectorId, "")), -1)

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
	// TODU
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiGetResourceSelectorRequest struct {
	ctx                context.Context
	ApiService         ResourceSelectorsAPI
	resourceSelectorId string
	// TODU
	data       interface{}
	retryCount int32
}

// TODU
func (r ApiGetResourceSelectorRequest) Data(data interface{}) ApiGetResourceSelectorRequest {
	r.data = data
	return r
}

// TODU
func (r ApiGetResourceSelectorRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetResourceSelectorExecute(r)
}

/*
GetResourceSelector Retrieve a Resource Selector

Retrieves a Resource Selector by `resourceSelectorId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSelectorId `id` of a Resource Selector
 @return ApiGetResourceSelectorRequest
*/
// TODU

func (a *ResourceSelectorsAPIService) GetResourceSelector(ctx context.Context, resourceSelectorId string) ApiGetResourceSelectorRequest {
	return ApiGetResourceSelectorRequest{
		ApiService:         a,
		ctx:                ctx,
		resourceSelectorId: resourceSelectorId,
		retryCount:         0,
	}
}

// Execute executes the request
//  @return ResourceSelectorResponseSchema

func (a *ResourceSelectorsAPIService) GetResourceSelectorExecute(r ApiGetResourceSelectorRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSelectorsAPIService.GetResourceSelector")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/resource-selectors/{resourceSelectorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSelectorId"+"}", url.PathEscape(parameterToString(r.resourceSelectorId, "")), -1)

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
	// TODU
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiListResourceSelectorsRequest struct {
	ctx        context.Context
	ApiService ResourceSelectorsAPI
	after      *string
	limit      *int32
	// TODU
	data       interface{}
	retryCount int32
}

// Specifies the pagination cursor for the next page of Resource Selectors
func (r ApiListResourceSelectorsRequest) After(after string) ApiListResourceSelectorsRequest {
	r.after = &after
	return r
}

// Specifies the number of results returned. Defaults to &#x60;10&#x60;
func (r ApiListResourceSelectorsRequest) Limit(limit int32) ApiListResourceSelectorsRequest {
	r.limit = &limit
	return r
}

// TODU
func (r ApiListResourceSelectorsRequest) Data(data interface{}) ApiListResourceSelectorsRequest {
	r.data = data
	return r
}

// TODU
func (r ApiListResourceSelectorsRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ListResourceSelectorsExecute(r)
}

/*
ListResourceSelectors List all Resource Selectors

Lists all Resource Selectors

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListResourceSelectorsRequest
*/
// TODU

func (a *ResourceSelectorsAPIService) ListResourceSelectors(ctx context.Context) ApiListResourceSelectorsRequest {
	return ApiListResourceSelectorsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ResourceSelectorsSchema

func (a *ResourceSelectorsAPIService) ListResourceSelectorsExecute(r ApiListResourceSelectorsRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSelectorsAPIService.ListResourceSelectors")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/resource-selectors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
	// TODU
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiUpdateResourceSelectorRequest struct {
	ctx                context.Context
	ApiService         ResourceSelectorsAPI
	resourceSelectorId string
	instance           *ResourceSelectorPatchRequestSchema
	// TODU
	data       interface{}
	retryCount int32
}

func (r ApiUpdateResourceSelectorRequest) Instance(instance ResourceSelectorPatchRequestSchema) ApiUpdateResourceSelectorRequest {
	r.instance = &instance
	return r
}

// TODU
func (r ApiUpdateResourceSelectorRequest) Data(data interface{}) ApiUpdateResourceSelectorRequest {
	r.data = data
	return r
}

// TODU
func (r ApiUpdateResourceSelectorRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UpdateResourceSelectorExecute(r)
}

/*
UpdateResourceSelector Update a Resource Selector

Updates a Resource Selector by `resourceSelectorId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSelectorId `id` of a Resource Selector
 @return ApiUpdateResourceSelectorRequest
*/
// TODU

func (a *ResourceSelectorsAPIService) UpdateResourceSelector(ctx context.Context, resourceSelectorId string) ApiUpdateResourceSelectorRequest {
	return ApiUpdateResourceSelectorRequest{
		ApiService:         a,
		ctx:                ctx,
		resourceSelectorId: resourceSelectorId,
		retryCount:         0,
	}
}

// Execute executes the request
//  @return ResourceSelectorResponseSchema

func (a *ResourceSelectorsAPIService) UpdateResourceSelectorExecute(r ApiUpdateResourceSelectorRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSelectorsAPIService.UpdateResourceSelector")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/resource-selectors/{resourceSelectorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSelectorId"+"}", url.PathEscape(parameterToString(r.resourceSelectorId, "")), -1)

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
	// TODU
	// body params
	// localVarPostBody = r.instance
	localVarPostBody = r.data
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			// TODU
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				// TODU
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		// TODU
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}
