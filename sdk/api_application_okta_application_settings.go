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

type ApplicationOktaApplicationSettingsAPI interface {

	/*
		GetFirstPartyAppSettings Retrieve the Okta app settings

		Retrieves the settings for the first party Okta app

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appName `appName` of the application
		@return ApiGetFirstPartyAppSettingsRequest
	*/
	GetFirstPartyAppSettings(ctx context.Context, appName string) ApiGetFirstPartyAppSettingsRequest

	// GetFirstPartyAppSettingsExecute executes the request
	//  @return AdminConsoleSettings
	GetFirstPartyAppSettingsExecute(r ApiGetFirstPartyAppSettingsRequest) (*APIResponse, error)

	/*
		ReplaceFirstPartyAppSettings Replace the Okta app settings

		Replaces the settings for the first party Okta app

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appName `appName` of the application
		@return ApiReplaceFirstPartyAppSettingsRequest
	*/
	ReplaceFirstPartyAppSettings(ctx context.Context, appName string) ApiReplaceFirstPartyAppSettingsRequest

	// ReplaceFirstPartyAppSettingsExecute executes the request
	//  @return AdminConsoleSettings
	ReplaceFirstPartyAppSettingsExecute(r ApiReplaceFirstPartyAppSettingsRequest) (*APIResponse, error)
}

// ApplicationOktaApplicationSettingsAPIService ApplicationOktaApplicationSettingsAPI service
type ApplicationOktaApplicationSettingsAPIService service

type ApiGetFirstPartyAppSettingsRequest struct {
	ctx        context.Context
	ApiService ApplicationOktaApplicationSettingsAPI
	appName    string
	data       interface{}
	retryCount int32
}

func (r ApiGetFirstPartyAppSettingsRequest) Data(data interface{}) ApiGetFirstPartyAppSettingsRequest {
	r.data = data
	return r
}

func (r ApiGetFirstPartyAppSettingsRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetFirstPartyAppSettingsExecute(r)
}

/*
GetFirstPartyAppSettings Retrieve the Okta app settings

Retrieves the settings for the first party Okta app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appName `appName` of the application
 @return ApiGetFirstPartyAppSettingsRequest
*/

func (a *ApplicationOktaApplicationSettingsAPIService) GetFirstPartyAppSettings(ctx context.Context, appName string) ApiGetFirstPartyAppSettingsRequest {
	return ApiGetFirstPartyAppSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		appName:    appName,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AdminConsoleSettings

func (a *ApplicationOktaApplicationSettingsAPIService) GetFirstPartyAppSettingsExecute(r ApiGetFirstPartyAppSettingsRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationOktaApplicationSettingsAPIService.GetFirstPartyAppSettings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/first-party-app-settings/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)

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
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
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
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiReplaceFirstPartyAppSettingsRequest struct {
	ctx                  context.Context
	ApiService           ApplicationOktaApplicationSettingsAPI
	appName              string
	adminConsoleSettings *AdminConsoleSettings
	data                 interface{}
	retryCount           int32
}

func (r ApiReplaceFirstPartyAppSettingsRequest) AdminConsoleSettings(adminConsoleSettings AdminConsoleSettings) ApiReplaceFirstPartyAppSettingsRequest {
	r.adminConsoleSettings = &adminConsoleSettings
	return r
}

func (r ApiReplaceFirstPartyAppSettingsRequest) Data(data interface{}) ApiReplaceFirstPartyAppSettingsRequest {
	r.data = data
	return r
}

func (r ApiReplaceFirstPartyAppSettingsRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ReplaceFirstPartyAppSettingsExecute(r)
}

/*
ReplaceFirstPartyAppSettings Replace the Okta app settings

Replaces the settings for the first party Okta app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appName `appName` of the application
 @return ApiReplaceFirstPartyAppSettingsRequest
*/

func (a *ApplicationOktaApplicationSettingsAPIService) ReplaceFirstPartyAppSettings(ctx context.Context, appName string) ApiReplaceFirstPartyAppSettingsRequest {
	return ApiReplaceFirstPartyAppSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		appName:    appName,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AdminConsoleSettings

func (a *ApplicationOktaApplicationSettingsAPIService) ReplaceFirstPartyAppSettingsExecute(r ApiReplaceFirstPartyAppSettingsRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationOktaApplicationSettingsAPIService.ReplaceFirstPartyAppSettings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/first-party-app-settings/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)

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
	// localVarPostBody = r.adminConsoleSettings
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
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
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
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}
