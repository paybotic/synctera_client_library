/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
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


// DigitalWalletTokensAPIService DigitalWalletTokensAPI service
type DigitalWalletTokensAPIService service

type ApiCreateDigitalWalletAppleRequest struct {
	ctx context.Context
	ApiService *DigitalWalletTokensAPIService
	cardId string
	appleDigitalWalletProvisionRequest *AppleDigitalWalletProvisionRequest
	idempotencyKey *string
}

// Request to provision digital wallet card data to pass to Apple Pay digital wallet.  &lt;br&gt;Cannot be used outside of production.&lt;/br&gt; 
func (r ApiCreateDigitalWalletAppleRequest) AppleDigitalWalletProvisionRequest(appleDigitalWalletProvisionRequest AppleDigitalWalletProvisionRequest) ApiCreateDigitalWalletAppleRequest {
	r.appleDigitalWalletProvisionRequest = &appleDigitalWalletProvisionRequest
	return r
}

// An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry.
func (r ApiCreateDigitalWalletAppleRequest) IdempotencyKey(idempotencyKey string) ApiCreateDigitalWalletAppleRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateDigitalWalletAppleRequest) Execute() (*AppleDigitalWalletProvisionResponse, *http.Response, error) {
	return r.ApiService.CreateDigitalWalletAppleExecute(r)
}

/*
CreateDigitalWalletApple Create digital wallet token provision request for Apple Pay

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cardId
 @return ApiCreateDigitalWalletAppleRequest
*/
func (a *DigitalWalletTokensAPIService) CreateDigitalWalletApple(ctx context.Context, cardId string) ApiCreateDigitalWalletAppleRequest {
	return ApiCreateDigitalWalletAppleRequest{
		ApiService: a,
		ctx: ctx,
		cardId: cardId,
	}
}

// Execute executes the request
//  @return AppleDigitalWalletProvisionResponse
func (a *DigitalWalletTokensAPIService) CreateDigitalWalletAppleExecute(r ApiCreateDigitalWalletAppleRequest) (*AppleDigitalWalletProvisionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppleDigitalWalletProvisionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DigitalWalletTokensAPIService.CreateDigitalWalletApple")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cards/{card_id}/digital_wallet_tokens/applepay"
	localVarPath = strings.Replace(localVarPath, "{"+"card_id"+"}", url.PathEscape(parameterValueToString(r.cardId, "cardId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appleDigitalWalletProvisionRequest == nil {
		return localVarReturnValue, nil, reportError("appleDigitalWalletProvisionRequest is required and must be specified")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	// body params
	localVarPostBody = r.appleDigitalWalletProvisionRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiCreateDigitalWalletGoogleRequest struct {
	ctx context.Context
	ApiService *DigitalWalletTokensAPIService
	cardId string
	googleDigitalWalletProvisionRequest *GoogleDigitalWalletProvisionRequest
	idempotencyKey *string
}

// Request to provision digital wallet card data to pass to Google Pay digital wallet. &lt;br&gt;Cannot be used outside of production.&lt;/br&gt; 
func (r ApiCreateDigitalWalletGoogleRequest) GoogleDigitalWalletProvisionRequest(googleDigitalWalletProvisionRequest GoogleDigitalWalletProvisionRequest) ApiCreateDigitalWalletGoogleRequest {
	r.googleDigitalWalletProvisionRequest = &googleDigitalWalletProvisionRequest
	return r
}

// An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry.
func (r ApiCreateDigitalWalletGoogleRequest) IdempotencyKey(idempotencyKey string) ApiCreateDigitalWalletGoogleRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateDigitalWalletGoogleRequest) Execute() (*GoogleDigitalWalletProvisionResponse, *http.Response, error) {
	return r.ApiService.CreateDigitalWalletGoogleExecute(r)
}

/*
CreateDigitalWalletGoogle Create digital wallet token provision request for Google Pay

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cardId
 @return ApiCreateDigitalWalletGoogleRequest
*/
func (a *DigitalWalletTokensAPIService) CreateDigitalWalletGoogle(ctx context.Context, cardId string) ApiCreateDigitalWalletGoogleRequest {
	return ApiCreateDigitalWalletGoogleRequest{
		ApiService: a,
		ctx: ctx,
		cardId: cardId,
	}
}

// Execute executes the request
//  @return GoogleDigitalWalletProvisionResponse
func (a *DigitalWalletTokensAPIService) CreateDigitalWalletGoogleExecute(r ApiCreateDigitalWalletGoogleRequest) (*GoogleDigitalWalletProvisionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GoogleDigitalWalletProvisionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DigitalWalletTokensAPIService.CreateDigitalWalletGoogle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cards/{card_id}/digital_wallet_tokens/googlepay"
	localVarPath = strings.Replace(localVarPath, "{"+"card_id"+"}", url.PathEscape(parameterValueToString(r.cardId, "cardId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.googleDigitalWalletProvisionRequest == nil {
		return localVarReturnValue, nil, reportError("googleDigitalWalletProvisionRequest is required and must be specified")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	// body params
	localVarPostBody = r.googleDigitalWalletProvisionRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiGetDigitalWalletTokenRequest struct {
	ctx context.Context
	ApiService *DigitalWalletTokensAPIService
	digitalWalletTokenId string
}

func (r ApiGetDigitalWalletTokenRequest) Execute() (*DigitalWalletTokenResponse, *http.Response, error) {
	return r.ApiService.GetDigitalWalletTokenExecute(r)
}

/*
GetDigitalWalletToken Get Digital Wallet Token

Get the details about the digital wallet token of a card <br>NB: Digital wallet tokens cannot be created outside of production</br>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param digitalWalletTokenId
 @return ApiGetDigitalWalletTokenRequest
*/
func (a *DigitalWalletTokensAPIService) GetDigitalWalletToken(ctx context.Context, digitalWalletTokenId string) ApiGetDigitalWalletTokenRequest {
	return ApiGetDigitalWalletTokenRequest{
		ApiService: a,
		ctx: ctx,
		digitalWalletTokenId: digitalWalletTokenId,
	}
}

// Execute executes the request
//  @return DigitalWalletTokenResponse
func (a *DigitalWalletTokensAPIService) GetDigitalWalletTokenExecute(r ApiGetDigitalWalletTokenRequest) (*DigitalWalletTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DigitalWalletTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DigitalWalletTokensAPIService.GetDigitalWalletToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cards/digital_wallet_tokens/{digital_wallet_token_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"digital_wallet_token_id"+"}", url.PathEscape(parameterValueToString(r.digitalWalletTokenId, "digitalWalletTokenId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiListDigitalWalletTokensRequest struct {
	ctx context.Context
	ApiService *DigitalWalletTokensAPIService
	pageToken *string
	limit *int32
	tokenState *DigitalWalletTokenState
	cardId *string
}

func (r ApiListDigitalWalletTokensRequest) PageToken(pageToken string) ApiListDigitalWalletTokensRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListDigitalWalletTokensRequest) Limit(limit int32) ApiListDigitalWalletTokensRequest {
	r.limit = &limit
	return r
}

// The status of the Digital Wallet Token
func (r ApiListDigitalWalletTokensRequest) TokenState(tokenState DigitalWalletTokenState) ApiListDigitalWalletTokensRequest {
	r.tokenState = &tokenState
	return r
}

func (r ApiListDigitalWalletTokensRequest) CardId(cardId string) ApiListDigitalWalletTokensRequest {
	r.cardId = &cardId
	return r
}

func (r ApiListDigitalWalletTokensRequest) Execute() (*TokenListResponse, *http.Response, error) {
	return r.ApiService.ListDigitalWalletTokensExecute(r)
}

/*
ListDigitalWalletTokens List Digital Wallet Tokens

List Digital Wallet Tokens <br>NB: Digital wallet tokens cannot be created outside of production</br>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDigitalWalletTokensRequest
*/
func (a *DigitalWalletTokensAPIService) ListDigitalWalletTokens(ctx context.Context) ApiListDigitalWalletTokensRequest {
	return ApiListDigitalWalletTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TokenListResponse
func (a *DigitalWalletTokensAPIService) ListDigitalWalletTokensExecute(r ApiListDigitalWalletTokensRequest) (*TokenListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TokenListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DigitalWalletTokensAPIService.ListDigitalWalletTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cards/digital_wallet_tokens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.tokenState != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token_state", r.tokenState, "")
	}
	if r.cardId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "card_id", r.cardId, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiUpdateDigitalWalletTokenStatusRequest struct {
	ctx context.Context
	ApiService *DigitalWalletTokensAPIService
	digitalWalletTokenId string
	digitalWalletTokenEditRequest *DigitalWalletTokenEditRequest
	idempotencyKey *string
}

// Update Digital wallet token status
func (r ApiUpdateDigitalWalletTokenStatusRequest) DigitalWalletTokenEditRequest(digitalWalletTokenEditRequest DigitalWalletTokenEditRequest) ApiUpdateDigitalWalletTokenStatusRequest {
	r.digitalWalletTokenEditRequest = &digitalWalletTokenEditRequest
	return r
}

// An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry.
func (r ApiUpdateDigitalWalletTokenStatusRequest) IdempotencyKey(idempotencyKey string) ApiUpdateDigitalWalletTokenStatusRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiUpdateDigitalWalletTokenStatusRequest) Execute() (*DigitalWalletTokenResponse, *http.Response, error) {
	return r.ApiService.UpdateDigitalWalletTokenStatusExecute(r)
}

/*
UpdateDigitalWalletTokenStatus Update Digital Wallet Token's life cycle status

The status of a digital wallet token can be updated as, ACTIVE to SUSPENDED, SUSPENDED to ACTIVE, ACTIVE to TERMINATED or SUSPENDED to TERMINATED.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param digitalWalletTokenId
 @return ApiUpdateDigitalWalletTokenStatusRequest
*/
func (a *DigitalWalletTokensAPIService) UpdateDigitalWalletTokenStatus(ctx context.Context, digitalWalletTokenId string) ApiUpdateDigitalWalletTokenStatusRequest {
	return ApiUpdateDigitalWalletTokenStatusRequest{
		ApiService: a,
		ctx: ctx,
		digitalWalletTokenId: digitalWalletTokenId,
	}
}

// Execute executes the request
//  @return DigitalWalletTokenResponse
func (a *DigitalWalletTokensAPIService) UpdateDigitalWalletTokenStatusExecute(r ApiUpdateDigitalWalletTokenStatusRequest) (*DigitalWalletTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DigitalWalletTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DigitalWalletTokensAPIService.UpdateDigitalWalletTokenStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cards/digital_wallet_tokens/{digital_wallet_token_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"digital_wallet_token_id"+"}", url.PathEscape(parameterValueToString(r.digitalWalletTokenId, "digitalWalletTokenId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.digitalWalletTokenEditRequest == nil {
		return localVarReturnValue, nil, reportError("digitalWalletTokenEditRequest is required and must be specified")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	// body params
	localVarPostBody = r.digitalWalletTokenEditRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
