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

type TemplateAPI interface {

	/*
		CreateSmsTemplate Create an SMS Template

		Creates a new custom SMS template

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateSmsTemplateRequest
	*/
	CreateSmsTemplate(ctx context.Context) ApiCreateSmsTemplateRequest

	// CreateSmsTemplateExecute executes the request
	//  @return SmsTemplate
	CreateSmsTemplateExecute(r ApiCreateSmsTemplateRequest) (*APIResponse, error)

	/*
		DeleteSmsTemplate Delete an SMS Template

		Deletes an SMS template

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param templateId `id` of the Template
		@return ApiDeleteSmsTemplateRequest
	*/
	DeleteSmsTemplate(ctx context.Context, templateId string) ApiDeleteSmsTemplateRequest

	// DeleteSmsTemplateExecute executes the request
	DeleteSmsTemplateExecute(r ApiDeleteSmsTemplateRequest) (*APIResponse, error)

	/*
		GetSmsTemplate Retrieve an SMS Template

		Retrieves a specific template by `id`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param templateId `id` of the Template
		@return ApiGetSmsTemplateRequest
	*/
	GetSmsTemplate(ctx context.Context, templateId string) ApiGetSmsTemplateRequest

	// GetSmsTemplateExecute executes the request
	//  @return SmsTemplate
	GetSmsTemplateExecute(r ApiGetSmsTemplateRequest) (*APIResponse, error)

	/*
		ListSmsTemplates List all SMS Templates

		Lists all custom SMS templates. A subset of templates can be returned that match a template type.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListSmsTemplatesRequest
	*/
	ListSmsTemplates(ctx context.Context) ApiListSmsTemplatesRequest

	// ListSmsTemplatesExecute executes the request
	//  @return []SmsTemplate
	ListSmsTemplatesExecute(r ApiListSmsTemplatesRequest) (*APIResponse, error)

	/*
		ReplaceSmsTemplate Replace an SMS Template

		Replaces the SMS template

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param templateId `id` of the Template
		@return ApiReplaceSmsTemplateRequest
	*/
	ReplaceSmsTemplate(ctx context.Context, templateId string) ApiReplaceSmsTemplateRequest

	// ReplaceSmsTemplateExecute executes the request
	//  @return SmsTemplate
	ReplaceSmsTemplateExecute(r ApiReplaceSmsTemplateRequest) (*APIResponse, error)

	/*
		UpdateSmsTemplate Update an SMS Template

		Updates an SMS template

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param templateId `id` of the Template
		@return ApiUpdateSmsTemplateRequest
	*/
	UpdateSmsTemplate(ctx context.Context, templateId string) ApiUpdateSmsTemplateRequest

	// UpdateSmsTemplateExecute executes the request
	//  @return SmsTemplate
	UpdateSmsTemplateExecute(r ApiUpdateSmsTemplateRequest) (*APIResponse, error)
}

// TemplateAPIService TemplateAPI service
type TemplateAPIService service

type ApiCreateSmsTemplateRequest struct {
	ctx         context.Context
	ApiService  TemplateAPI
	smsTemplate *SmsTemplate
	data        interface{}
	retryCount  int32
}

func (r ApiCreateSmsTemplateRequest) SmsTemplate(smsTemplate SmsTemplate) ApiCreateSmsTemplateRequest {
	r.smsTemplate = &smsTemplate
	return r
}

func (r ApiCreateSmsTemplateRequest) Data(data interface{}) ApiCreateSmsTemplateRequest {
	r.data = data
	return r
}

func (r ApiCreateSmsTemplateRequest) Execute() (*APIResponse, error) {
	return r.ApiService.CreateSmsTemplateExecute(r)
}

/*
CreateSmsTemplate Create an SMS Template

Creates a new custom SMS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateSmsTemplateRequest
*/

func (a *TemplateAPIService) CreateSmsTemplate(ctx context.Context) ApiCreateSmsTemplateRequest {
	return ApiCreateSmsTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return SmsTemplate

func (a *TemplateAPIService) CreateSmsTemplateExecute(r ApiCreateSmsTemplateRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateAPIService.CreateSmsTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/templates/sms"

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
	// localVarPostBody = r.smsTemplate
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
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
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

type ApiDeleteSmsTemplateRequest struct {
	ctx        context.Context
	ApiService TemplateAPI
	templateId string
	data       interface{}
	retryCount int32
}

func (r ApiDeleteSmsTemplateRequest) Data(data interface{}) ApiDeleteSmsTemplateRequest {
	r.data = data
	return r
}

func (r ApiDeleteSmsTemplateRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteSmsTemplateExecute(r)
}

/*
DeleteSmsTemplate Delete an SMS Template

Deletes an SMS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId `id` of the Template
 @return ApiDeleteSmsTemplateRequest
*/

func (a *TemplateAPIService) DeleteSmsTemplate(ctx context.Context, templateId string) ApiDeleteSmsTemplateRequest {
	return ApiDeleteSmsTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		templateId: templateId,
		retryCount: 0,
	}
}

// Execute executes the request

func (a *TemplateAPIService) DeleteSmsTemplateExecute(r ApiDeleteSmsTemplateRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateAPIService.DeleteSmsTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/templates/sms/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetSmsTemplateRequest struct {
	ctx        context.Context
	ApiService TemplateAPI
	templateId string
	data       interface{}
	retryCount int32
}

func (r ApiGetSmsTemplateRequest) Data(data interface{}) ApiGetSmsTemplateRequest {
	r.data = data
	return r
}

func (r ApiGetSmsTemplateRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetSmsTemplateExecute(r)
}

/*
GetSmsTemplate Retrieve an SMS Template

Retrieves a specific template by `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId `id` of the Template
 @return ApiGetSmsTemplateRequest
*/

func (a *TemplateAPIService) GetSmsTemplate(ctx context.Context, templateId string) ApiGetSmsTemplateRequest {
	return ApiGetSmsTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		templateId: templateId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return SmsTemplate

func (a *TemplateAPIService) GetSmsTemplateExecute(r ApiGetSmsTemplateRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateAPIService.GetSmsTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/templates/sms/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiListSmsTemplatesRequest struct {
	ctx          context.Context
	ApiService   TemplateAPI
	templateType *string
	data         interface{}
	retryCount   int32
}

func (r ApiListSmsTemplatesRequest) TemplateType(templateType string) ApiListSmsTemplatesRequest {
	r.templateType = &templateType
	return r
}

func (r ApiListSmsTemplatesRequest) Data(data interface{}) ApiListSmsTemplatesRequest {
	r.data = data
	return r
}

func (r ApiListSmsTemplatesRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ListSmsTemplatesExecute(r)
}

/*
ListSmsTemplates List all SMS Templates

Lists all custom SMS templates. A subset of templates can be returned that match a template type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListSmsTemplatesRequest
*/

func (a *TemplateAPIService) ListSmsTemplates(ctx context.Context) ApiListSmsTemplatesRequest {
	return ApiListSmsTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []SmsTemplate

func (a *TemplateAPIService) ListSmsTemplatesExecute(r ApiListSmsTemplatesRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateAPIService.ListSmsTemplates")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/templates/sms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.templateType != nil {
		localVarQueryParams.Add("templateType", parameterToString(*r.templateType, ""))
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
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
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

type ApiReplaceSmsTemplateRequest struct {
	ctx         context.Context
	ApiService  TemplateAPI
	templateId  string
	smsTemplate *SmsTemplate
	data        interface{}
	retryCount  int32
}

func (r ApiReplaceSmsTemplateRequest) SmsTemplate(smsTemplate SmsTemplate) ApiReplaceSmsTemplateRequest {
	r.smsTemplate = &smsTemplate
	return r
}

func (r ApiReplaceSmsTemplateRequest) Data(data interface{}) ApiReplaceSmsTemplateRequest {
	r.data = data
	return r
}

func (r ApiReplaceSmsTemplateRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ReplaceSmsTemplateExecute(r)
}

/*
ReplaceSmsTemplate Replace an SMS Template

Replaces the SMS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId `id` of the Template
 @return ApiReplaceSmsTemplateRequest
*/

func (a *TemplateAPIService) ReplaceSmsTemplate(ctx context.Context, templateId string) ApiReplaceSmsTemplateRequest {
	return ApiReplaceSmsTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		templateId: templateId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return SmsTemplate

func (a *TemplateAPIService) ReplaceSmsTemplateExecute(r ApiReplaceSmsTemplateRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateAPIService.ReplaceSmsTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/templates/sms/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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
	// localVarPostBody = r.smsTemplate
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
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiUpdateSmsTemplateRequest struct {
	ctx         context.Context
	ApiService  TemplateAPI
	templateId  string
	smsTemplate *SmsTemplate
	data        interface{}
	retryCount  int32
}

func (r ApiUpdateSmsTemplateRequest) SmsTemplate(smsTemplate SmsTemplate) ApiUpdateSmsTemplateRequest {
	r.smsTemplate = &smsTemplate
	return r
}

func (r ApiUpdateSmsTemplateRequest) Data(data interface{}) ApiUpdateSmsTemplateRequest {
	r.data = data
	return r
}

func (r ApiUpdateSmsTemplateRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UpdateSmsTemplateExecute(r)
}

/*
UpdateSmsTemplate Update an SMS Template

Updates an SMS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId `id` of the Template
 @return ApiUpdateSmsTemplateRequest
*/

func (a *TemplateAPIService) UpdateSmsTemplate(ctx context.Context, templateId string) ApiUpdateSmsTemplateRequest {
	return ApiUpdateSmsTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		templateId: templateId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return SmsTemplate

func (a *TemplateAPIService) UpdateSmsTemplateExecute(r ApiUpdateSmsTemplateRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateAPIService.UpdateSmsTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/templates/sms/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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
	// localVarPostBody = r.smsTemplate
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
		return nil, err
	}
	localVarHTTPResponse, err = a.client.Do(r.ctx, req)
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
		if localVarHTTPResponse.StatusCode == 404 {
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
