/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// InternalAccountsApiService InternalAccountsApi service
type InternalAccountsApiService service

type InternalAccountsApiAddInternalAccountsRequest struct {
	ctx context.Context
	ApiService *InternalAccountsApiService
	internalAccount *InternalAccount
	idempotencyKey *string
}

func (r InternalAccountsApiAddInternalAccountsRequest) InternalAccount(internalAccount InternalAccount) InternalAccountsApiAddInternalAccountsRequest {
	r.internalAccount = &internalAccount
	return r
}

// An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key
func (r InternalAccountsApiAddInternalAccountsRequest) IdempotencyKey(idempotencyKey string) InternalAccountsApiAddInternalAccountsRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r InternalAccountsApiAddInternalAccountsRequest) Execute() (*InternalAccount, *http.Response, error) {
	return r.ApiService.AddInternalAccountsExecute(r)
}

/*
AddInternalAccounts Add internal accounts

Add an internal account


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InternalAccountsApiAddInternalAccountsRequest
*/
func (a *InternalAccountsApiService) AddInternalAccounts(ctx context.Context) InternalAccountsApiAddInternalAccountsRequest {
	return InternalAccountsApiAddInternalAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InternalAccount
func (a *InternalAccountsApiService) AddInternalAccountsExecute(r InternalAccountsApiAddInternalAccountsRequest) (*InternalAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InternalAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InternalAccountsApiService.AddInternalAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/internal_accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.internalAccount == nil {
		return localVarReturnValue, nil, reportError("internalAccount is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.idempotencyKey != nil {
		localVarHeaderParams["Idempotency-Key"] = parameterToString(*r.idempotencyKey, "")
	}
	// body params
	localVarPostBody = r.internalAccount
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type InternalAccountsApiListInternalAccountsRequest struct {
	ctx context.Context
	ApiService *InternalAccountsApiService
	limit *int32
	pageToken *string
	id *[]string
}

func (r InternalAccountsApiListInternalAccountsRequest) Limit(limit int32) InternalAccountsApiListInternalAccountsRequest {
	r.limit = &limit
	return r
}

func (r InternalAccountsApiListInternalAccountsRequest) PageToken(pageToken string) InternalAccountsApiListInternalAccountsRequest {
	r.pageToken = &pageToken
	return r
}

// Internal account ID(s). Multiple IDs can be provided as a comma-separated list. 
func (r InternalAccountsApiListInternalAccountsRequest) Id(id []string) InternalAccountsApiListInternalAccountsRequest {
	r.id = &id
	return r
}

func (r InternalAccountsApiListInternalAccountsRequest) Execute() (*InternalAccountsList, *http.Response, error) {
	return r.ApiService.ListInternalAccountsExecute(r)
}

/*
ListInternalAccounts List internal accounts

Returns a list of all internal accounts.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InternalAccountsApiListInternalAccountsRequest
*/
func (a *InternalAccountsApiService) ListInternalAccounts(ctx context.Context) InternalAccountsApiListInternalAccountsRequest {
	return InternalAccountsApiListInternalAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InternalAccountsList
func (a *InternalAccountsApiService) ListInternalAccountsExecute(r InternalAccountsApiListInternalAccountsRequest) (*InternalAccountsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InternalAccountsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InternalAccountsApiService.ListInternalAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/internal_accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type InternalAccountsApiPatchInternalAccountRequest struct {
	ctx context.Context
	ApiService *InternalAccountsApiService
	internalAccountId string
	internalAccountPatch *InternalAccountPatch
}

// Internal Account fields to be patched
func (r InternalAccountsApiPatchInternalAccountRequest) InternalAccountPatch(internalAccountPatch InternalAccountPatch) InternalAccountsApiPatchInternalAccountRequest {
	r.internalAccountPatch = &internalAccountPatch
	return r
}

func (r InternalAccountsApiPatchInternalAccountRequest) Execute() (*InternalAccount, *http.Response, error) {
	return r.ApiService.PatchInternalAccountExecute(r)
}

/*
PatchInternalAccount Patch internal account

Patch internal account.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param internalAccountId Unique identifier for the internal account.
 @return InternalAccountsApiPatchInternalAccountRequest
*/
func (a *InternalAccountsApiService) PatchInternalAccount(ctx context.Context, internalAccountId string) InternalAccountsApiPatchInternalAccountRequest {
	return InternalAccountsApiPatchInternalAccountRequest{
		ApiService: a,
		ctx: ctx,
		internalAccountId: internalAccountId,
	}
}

// Execute executes the request
//  @return InternalAccount
func (a *InternalAccountsApiService) PatchInternalAccountExecute(r InternalAccountsApiPatchInternalAccountRequest) (*InternalAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InternalAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InternalAccountsApiService.PatchInternalAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/internal_accounts/{internal_account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"internal_account_id"+"}", url.PathEscape(parameterToString(r.internalAccountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.internalAccountPatch == nil {
		return localVarReturnValue, nil, reportError("internalAccountPatch is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.internalAccountPatch
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 461 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
