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
	"time"
)


type RateLimitSettingsAPI interface {

	/*
	GetRateLimitSettingsAdminNotifications Retrieve the Rate Limit Admin Notification Settings

	Retrieves the currently configured Rate Limit Admin Notification Settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRateLimitSettingsAdminNotificationsRequest
	*/
	GetRateLimitSettingsAdminNotifications(ctx context.Context) ApiGetRateLimitSettingsAdminNotificationsRequest

	// GetRateLimitSettingsAdminNotificationsExecute executes the request
	//  @return RateLimitAdminNotifications
	// TODU
	GetRateLimitSettingsAdminNotificationsExecute(r ApiGetRateLimitSettingsAdminNotificationsRequest) (*APIResponse, error)

	/*
	GetRateLimitSettingsPerClient Retrieve the Per-Client Rate Limit Settings

	Retrieves the currently configured Per-Client Rate Limit Settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRateLimitSettingsPerClientRequest
	*/
	GetRateLimitSettingsPerClient(ctx context.Context) ApiGetRateLimitSettingsPerClientRequest

	// GetRateLimitSettingsPerClientExecute executes the request
	//  @return PerClientRateLimitSettings
	// TODU
	GetRateLimitSettingsPerClientExecute(r ApiGetRateLimitSettingsPerClientRequest) (*APIResponse, error)

	/*
	GetRateLimitSettingsWarningThreshold Retrieve the Rate Limit Warning Threshold Percentage

	Retrieves the currently configured threshold for warning notifications when the API's rate limit is exceeded

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRateLimitSettingsWarningThresholdRequest
	*/
	GetRateLimitSettingsWarningThreshold(ctx context.Context) ApiGetRateLimitSettingsWarningThresholdRequest

	// GetRateLimitSettingsWarningThresholdExecute executes the request
	//  @return RateLimitWarningThresholdResponse
	// TODU
	GetRateLimitSettingsWarningThresholdExecute(r ApiGetRateLimitSettingsWarningThresholdRequest) (*APIResponse, error)

	/*
	ReplaceRateLimitSettingsAdminNotifications Replace the Rate Limit Admin Notification Settings

	Replaces the Rate Limit Admin Notification Settings and returns the configured properties

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceRateLimitSettingsAdminNotificationsRequest
	*/
	ReplaceRateLimitSettingsAdminNotifications(ctx context.Context) ApiReplaceRateLimitSettingsAdminNotificationsRequest

	// ReplaceRateLimitSettingsAdminNotificationsExecute executes the request
	//  @return RateLimitAdminNotifications
	// TODU
	ReplaceRateLimitSettingsAdminNotificationsExecute(r ApiReplaceRateLimitSettingsAdminNotificationsRequest) (*APIResponse, error)

	/*
	ReplaceRateLimitSettingsPerClient Replace the Per-Client Rate Limit Settings

	Replaces the Per-Client Rate Limit Settings and returns the configured properties

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceRateLimitSettingsPerClientRequest
	*/
	ReplaceRateLimitSettingsPerClient(ctx context.Context) ApiReplaceRateLimitSettingsPerClientRequest

	// ReplaceRateLimitSettingsPerClientExecute executes the request
	//  @return PerClientRateLimitSettings
	// TODU
	ReplaceRateLimitSettingsPerClientExecute(r ApiReplaceRateLimitSettingsPerClientRequest) (*APIResponse, error)

	/*
	ReplaceRateLimitSettingsWarningThreshold Replace the Rate Limit Warning Threshold Percentage

	Replaces the Rate Limit Warning Threshold Percentage and returns the configured property

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceRateLimitSettingsWarningThresholdRequest
	*/
	ReplaceRateLimitSettingsWarningThreshold(ctx context.Context) ApiReplaceRateLimitSettingsWarningThresholdRequest

	// ReplaceRateLimitSettingsWarningThresholdExecute executes the request
	//  @return RateLimitWarningThresholdResponse
	// TODU
	ReplaceRateLimitSettingsWarningThresholdExecute(r ApiReplaceRateLimitSettingsWarningThresholdRequest) (*APIResponse, error)
}

// RateLimitSettingsAPIService RateLimitSettingsAPI service
type RateLimitSettingsAPIService service

type ApiGetRateLimitSettingsAdminNotificationsRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	// TODU
	data       interface{}
	retryCount int32
}


// TODU
func (r ApiGetRateLimitSettingsAdminNotificationsRequest) Data (data interface{}) ApiGetRateLimitSettingsAdminNotificationsRequest {
	r.data = data
	return r
}

// TODU
func (r ApiGetRateLimitSettingsAdminNotificationsRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetRateLimitSettingsAdminNotificationsExecute(r)
}

/*
GetRateLimitSettingsAdminNotifications Retrieve the Rate Limit Admin Notification Settings

Retrieves the currently configured Rate Limit Admin Notification Settings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRateLimitSettingsAdminNotificationsRequest
*/
// TODU

func (a *RateLimitSettingsAPIService) GetRateLimitSettingsAdminNotifications(ctx context.Context) ApiGetRateLimitSettingsAdminNotificationsRequest {
	return ApiGetRateLimitSettingsAdminNotificationsRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitAdminNotifications

func (a *RateLimitSettingsAPIService) GetRateLimitSettingsAdminNotificationsExecute(r ApiGetRateLimitSettingsAdminNotificationsRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.GetRateLimitSettingsAdminNotifications")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/admin-notifications"

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
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

type ApiGetRateLimitSettingsPerClientRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	// TODU
	data       interface{}
	retryCount int32
}


// TODU
func (r ApiGetRateLimitSettingsPerClientRequest) Data (data interface{}) ApiGetRateLimitSettingsPerClientRequest {
	r.data = data
	return r
}

// TODU
func (r ApiGetRateLimitSettingsPerClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetRateLimitSettingsPerClientExecute(r)
}

/*
GetRateLimitSettingsPerClient Retrieve the Per-Client Rate Limit Settings

Retrieves the currently configured Per-Client Rate Limit Settings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRateLimitSettingsPerClientRequest
*/
// TODU

func (a *RateLimitSettingsAPIService) GetRateLimitSettingsPerClient(ctx context.Context) ApiGetRateLimitSettingsPerClientRequest {
	return ApiGetRateLimitSettingsPerClientRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return PerClientRateLimitSettings

func (a *RateLimitSettingsAPIService) GetRateLimitSettingsPerClientExecute(r ApiGetRateLimitSettingsPerClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.GetRateLimitSettingsPerClient")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/per-client"

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
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

type ApiGetRateLimitSettingsWarningThresholdRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	// TODU
	data       interface{}
	retryCount int32
}


// TODU
func (r ApiGetRateLimitSettingsWarningThresholdRequest) Data (data interface{}) ApiGetRateLimitSettingsWarningThresholdRequest {
	r.data = data
	return r
}

// TODU
func (r ApiGetRateLimitSettingsWarningThresholdRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetRateLimitSettingsWarningThresholdExecute(r)
}

/*
GetRateLimitSettingsWarningThreshold Retrieve the Rate Limit Warning Threshold Percentage

Retrieves the currently configured threshold for warning notifications when the API's rate limit is exceeded

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRateLimitSettingsWarningThresholdRequest
*/
// TODU

func (a *RateLimitSettingsAPIService) GetRateLimitSettingsWarningThreshold(ctx context.Context) ApiGetRateLimitSettingsWarningThresholdRequest {
	return ApiGetRateLimitSettingsWarningThresholdRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitWarningThresholdResponse

func (a *RateLimitSettingsAPIService) GetRateLimitSettingsWarningThresholdExecute(r ApiGetRateLimitSettingsWarningThresholdRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.GetRateLimitSettingsWarningThreshold")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/warning-threshold"

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
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

type ApiReplaceRateLimitSettingsAdminNotificationsRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	rateLimitAdminNotifications *RateLimitAdminNotifications
	// TODU
	data       interface{}
	retryCount int32
}

func (r ApiReplaceRateLimitSettingsAdminNotificationsRequest) RateLimitAdminNotifications(rateLimitAdminNotifications RateLimitAdminNotifications) ApiReplaceRateLimitSettingsAdminNotificationsRequest {
	r.rateLimitAdminNotifications = &rateLimitAdminNotifications
	return r
}


// TODU
func (r ApiReplaceRateLimitSettingsAdminNotificationsRequest) Data (data interface{}) ApiReplaceRateLimitSettingsAdminNotificationsRequest {
	r.data = data
	return r
}

// TODU
func (r ApiReplaceRateLimitSettingsAdminNotificationsRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ReplaceRateLimitSettingsAdminNotificationsExecute(r)
}

/*
ReplaceRateLimitSettingsAdminNotifications Replace the Rate Limit Admin Notification Settings

Replaces the Rate Limit Admin Notification Settings and returns the configured properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReplaceRateLimitSettingsAdminNotificationsRequest
*/
// TODU

func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsAdminNotifications(ctx context.Context) ApiReplaceRateLimitSettingsAdminNotificationsRequest {
	return ApiReplaceRateLimitSettingsAdminNotificationsRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitAdminNotifications

func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsAdminNotificationsExecute(r ApiReplaceRateLimitSettingsAdminNotificationsRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.ReplaceRateLimitSettingsAdminNotifications")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/admin-notifications"

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
	// localVarPostBody = r.rateLimitAdminNotifications
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
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

type ApiReplaceRateLimitSettingsPerClientRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	perClientRateLimitSettings *PerClientRateLimitSettings
	// TODU
	data       interface{}
	retryCount int32
}

func (r ApiReplaceRateLimitSettingsPerClientRequest) PerClientRateLimitSettings(perClientRateLimitSettings PerClientRateLimitSettings) ApiReplaceRateLimitSettingsPerClientRequest {
	r.perClientRateLimitSettings = &perClientRateLimitSettings
	return r
}


// TODU
func (r ApiReplaceRateLimitSettingsPerClientRequest) Data (data interface{}) ApiReplaceRateLimitSettingsPerClientRequest {
	r.data = data
	return r
}

// TODU
func (r ApiReplaceRateLimitSettingsPerClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ReplaceRateLimitSettingsPerClientExecute(r)
}

/*
ReplaceRateLimitSettingsPerClient Replace the Per-Client Rate Limit Settings

Replaces the Per-Client Rate Limit Settings and returns the configured properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReplaceRateLimitSettingsPerClientRequest
*/
// TODU

func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsPerClient(ctx context.Context) ApiReplaceRateLimitSettingsPerClientRequest {
	return ApiReplaceRateLimitSettingsPerClientRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return PerClientRateLimitSettings

func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsPerClientExecute(r ApiReplaceRateLimitSettingsPerClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.ReplaceRateLimitSettingsPerClient")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/per-client"

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
	// localVarPostBody = r.perClientRateLimitSettings
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
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

type ApiReplaceRateLimitSettingsWarningThresholdRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	rateLimitWarningThreshold *RateLimitWarningThresholdRequest
	// TODU
	data       interface{}
	retryCount int32
}

func (r ApiReplaceRateLimitSettingsWarningThresholdRequest) RateLimitWarningThreshold(rateLimitWarningThreshold RateLimitWarningThresholdRequest) ApiReplaceRateLimitSettingsWarningThresholdRequest {
	r.rateLimitWarningThreshold = &rateLimitWarningThreshold
	return r
}


// TODU
func (r ApiReplaceRateLimitSettingsWarningThresholdRequest) Data (data interface{}) ApiReplaceRateLimitSettingsWarningThresholdRequest {
	r.data = data
	return r
}

// TODU
func (r ApiReplaceRateLimitSettingsWarningThresholdRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ReplaceRateLimitSettingsWarningThresholdExecute(r)
}

/*
ReplaceRateLimitSettingsWarningThreshold Replace the Rate Limit Warning Threshold Percentage

Replaces the Rate Limit Warning Threshold Percentage and returns the configured property

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReplaceRateLimitSettingsWarningThresholdRequest
*/
// TODU

func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsWarningThreshold(ctx context.Context) ApiReplaceRateLimitSettingsWarningThresholdRequest {
	return ApiReplaceRateLimitSettingsWarningThresholdRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitWarningThresholdResponse

func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsWarningThresholdExecute(r ApiReplaceRateLimitSettingsWarningThresholdRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		// TODU
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.ReplaceRateLimitSettingsWarningThreshold")
	if err != nil {
		// TODU
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/warning-threshold"

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
	// localVarPostBody = r.rateLimitWarningThreshold
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		// TODU
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
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
