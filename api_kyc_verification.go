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


// KYCVerificationApiService KYCVerificationApi service
type KYCVerificationApiService service

type KYCVerificationApiCreateCustomerVerificationResultRequest struct {
	ctx context.Context
	ApiService *KYCVerificationApiService
	customerId string
	customerVerificationResult *CustomerVerificationResult
	idempotencyKey *string
}

// Customer verification result to create.
func (r KYCVerificationApiCreateCustomerVerificationResultRequest) CustomerVerificationResult(customerVerificationResult CustomerVerificationResult) KYCVerificationApiCreateCustomerVerificationResultRequest {
	r.customerVerificationResult = &customerVerificationResult
	return r
}

// An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key
func (r KYCVerificationApiCreateCustomerVerificationResultRequest) IdempotencyKey(idempotencyKey string) KYCVerificationApiCreateCustomerVerificationResultRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r KYCVerificationApiCreateCustomerVerificationResultRequest) Execute() (*CustomerVerificationResult, *http.Response, error) {
	return r.ApiService.CreateCustomerVerificationResultExecute(r)
}

/*
CreateCustomerVerificationResult Create a customer verification result

Create a verification result for a customer.
This endpoint will be removed in API v1. Use `POST /v0/verifications` instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The customer's unique identifier
 @return KYCVerificationApiCreateCustomerVerificationResultRequest

Deprecated
*/
func (a *KYCVerificationApiService) CreateCustomerVerificationResult(ctx context.Context, customerId string) KYCVerificationApiCreateCustomerVerificationResultRequest {
	return KYCVerificationApiCreateCustomerVerificationResultRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return CustomerVerificationResult
// Deprecated
func (a *KYCVerificationApiService) CreateCustomerVerificationResultExecute(r KYCVerificationApiCreateCustomerVerificationResultRequest) (*CustomerVerificationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerVerificationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCVerificationApiService.CreateCustomerVerificationResult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}/verifications"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerVerificationResult == nil {
		return localVarReturnValue, nil, reportError("customerVerificationResult is required and must be specified")
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
	localVarPostBody = r.customerVerificationResult
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

type KYCVerificationApiGetVerificationRequest struct {
	ctx context.Context
	ApiService *KYCVerificationApiService
	customerId string
	verificationId string
}

func (r KYCVerificationApiGetVerificationRequest) Execute() (*CustomerVerificationResult, *http.Response, error) {
	return r.ApiService.GetVerificationExecute(r)
}

/*
GetVerification Get verification result

Get verification result by ID.

This endpoint will be removed in API v1. Use `GET /v0/verifications/{verification_id}` instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The customer's unique identifier
 @param verificationId Unique identifier for the verification.
 @return KYCVerificationApiGetVerificationRequest

Deprecated
*/
func (a *KYCVerificationApiService) GetVerification(ctx context.Context, customerId string, verificationId string) KYCVerificationApiGetVerificationRequest {
	return KYCVerificationApiGetVerificationRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
		verificationId: verificationId,
	}
}

// Execute executes the request
//  @return CustomerVerificationResult
// Deprecated
func (a *KYCVerificationApiService) GetVerificationExecute(r KYCVerificationApiGetVerificationRequest) (*CustomerVerificationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerVerificationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCVerificationApiService.GetVerification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}/verifications/{verification_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"verification_id"+"}", url.PathEscape(parameterToString(r.verificationId, "")), -1)

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

type KYCVerificationApiListVerificationsRequest struct {
	ctx context.Context
	ApiService *KYCVerificationApiService
	customerId string
	includeHistory *bool
	limit *int32
	pageToken *string
}

// If true, include old (inactive) records as well.
func (r KYCVerificationApiListVerificationsRequest) IncludeHistory(includeHistory bool) KYCVerificationApiListVerificationsRequest {
	r.includeHistory = &includeHistory
	return r
}

func (r KYCVerificationApiListVerificationsRequest) Limit(limit int32) KYCVerificationApiListVerificationsRequest {
	r.limit = &limit
	return r
}

func (r KYCVerificationApiListVerificationsRequest) PageToken(pageToken string) KYCVerificationApiListVerificationsRequest {
	r.pageToken = &pageToken
	return r
}

func (r KYCVerificationApiListVerificationsRequest) Execute() (*CustomerVerificationResultList, *http.Response, error) {
	return r.ApiService.ListVerificationsExecute(r)
}

/*
ListVerifications List verification results

List verification results.

This endpoint will be removed in API v1. Use `GET /v0/verifications?customer_id={customer_id}` instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The customer's unique identifier
 @return KYCVerificationApiListVerificationsRequest

Deprecated
*/
func (a *KYCVerificationApiService) ListVerifications(ctx context.Context, customerId string) KYCVerificationApiListVerificationsRequest {
	return KYCVerificationApiListVerificationsRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return CustomerVerificationResultList
// Deprecated
func (a *KYCVerificationApiService) ListVerificationsExecute(r KYCVerificationApiListVerificationsRequest) (*CustomerVerificationResultList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerVerificationResultList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCVerificationApiService.ListVerifications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}/verifications"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeHistory != nil {
		localVarQueryParams.Add("include_history", parameterToString(*r.includeHistory, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
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

type KYCVerificationApiVerifyCustomerRequest struct {
	ctx context.Context
	ApiService *KYCVerificationApiService
	customerId string
	customerVerification *CustomerVerification
	idempotencyKey *string
}

// Customer verification request.
func (r KYCVerificationApiVerifyCustomerRequest) CustomerVerification(customerVerification CustomerVerification) KYCVerificationApiVerifyCustomerRequest {
	r.customerVerification = &customerVerification
	return r
}

// An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key
func (r KYCVerificationApiVerifyCustomerRequest) IdempotencyKey(idempotencyKey string) KYCVerificationApiVerifyCustomerRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r KYCVerificationApiVerifyCustomerRequest) Execute() (*CustomerVerifyResponse, *http.Response, error) {
	return r.ApiService.VerifyCustomerExecute(r)
}

/*
VerifyCustomer Verify a customer's identity

Initiate identity verification for a customer and run the specified identity checks.

Verifying a personal customer requires that the following fields already be set:
* `first_name`
* `last_name`
* `dob`
* `email`
* `phone_number`
* `legal_address`
* `ssn`

This endpoint will be removed in API v1. Use `POST /v0/verifications/verify` instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The customer's unique identifier
 @return KYCVerificationApiVerifyCustomerRequest

Deprecated
*/
func (a *KYCVerificationApiService) VerifyCustomer(ctx context.Context, customerId string) KYCVerificationApiVerifyCustomerRequest {
	return KYCVerificationApiVerifyCustomerRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return CustomerVerifyResponse
// Deprecated
func (a *KYCVerificationApiService) VerifyCustomerExecute(r KYCVerificationApiVerifyCustomerRequest) (*CustomerVerifyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerVerifyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCVerificationApiService.VerifyCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}/verify"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerVerification == nil {
		return localVarReturnValue, nil, reportError("customerVerification is required and must be specified")
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
	localVarPostBody = r.customerVerification
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
