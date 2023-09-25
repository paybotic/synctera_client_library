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


// HoldAPIService HoldAPI service
type HoldAPIService service

type ApiGetHoldRequest struct {
	ctx context.Context
	ApiService *HoldAPIService
	id string
}

func (r ApiGetHoldRequest) Execute() (*PendingTransaction, *http.Response, error) {
	return r.ApiService.GetHoldExecute(r)
}

/*
GetHold Get a pending transaction

Get a pending transaction by its uuid


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Transaction ID
 @return ApiGetHoldRequest
*/
func (a *HoldAPIService) GetHold(ctx context.Context, id string) ApiGetHoldRequest {
	return ApiGetHoldRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PendingTransaction
func (a *HoldAPIService) GetHoldExecute(r ApiGetHoldRequest) (*PendingTransaction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PendingTransaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HoldAPIService.GetHold")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hold/{id}"
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

type ApiListHoldsRequest struct {
	ctx context.Context
	ApiService *HoldAPIService
	includeChildTransactions *bool
	fromDate *string
	toDate *string
	transactionId *string
	status *[]string
	type_ *string
	accountId *string
	idempotencyKey *[]string
	accountNo *string
	excludeJitTransactions *bool
	pageToken *string
	cardId *string
	referenceId *string
	limit *int32
	subtype *string
}

// Include transactions from sub-accounts when listing transactions for a given account
func (r ApiListHoldsRequest) IncludeChildTransactions(includeChildTransactions bool) ApiListHoldsRequest {
	r.includeChildTransactions = &includeChildTransactions
	return r
}

// Only display transactions with a posting date greater than from_date
func (r ApiListHoldsRequest) FromDate(fromDate string) ApiListHoldsRequest {
	r.fromDate = &fromDate
	return r
}

// Only display transactions with a posting date less than or equal to to_date
func (r ApiListHoldsRequest) ToDate(toDate string) ApiListHoldsRequest {
	r.toDate = &toDate
	return r
}

// Only display holds linked to the provided transaction id
func (r ApiListHoldsRequest) TransactionId(transactionId string) ApiListHoldsRequest {
	r.transactionId = &transactionId
	return r
}

// The status of the transaction
func (r ApiListHoldsRequest) Status(status []string) ApiListHoldsRequest {
	r.status = &status
	return r
}

// Only display transactions matching the given type
func (r ApiListHoldsRequest) Type_(type_ string) ApiListHoldsRequest {
	r.type_ = &type_
	return r
}

// Account ID
func (r ApiListHoldsRequest) AccountId(accountId string) ApiListHoldsRequest {
	r.accountId = &accountId
	return r
}

// Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list.
func (r ApiListHoldsRequest) IdempotencyKey(idempotencyKey []string) ApiListHoldsRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// Account number
func (r ApiListHoldsRequest) AccountNo(accountNo string) ApiListHoldsRequest {
	r.accountNo = &accountNo
	return r
}

// Hide \&quot;JIT funding\&quot; transactions from results
func (r ApiListHoldsRequest) ExcludeJitTransactions(excludeJitTransactions bool) ApiListHoldsRequest {
	r.excludeJitTransactions = &excludeJitTransactions
	return r
}

func (r ApiListHoldsRequest) PageToken(pageToken string) ApiListHoldsRequest {
	r.pageToken = &pageToken
	return r
}

// Card ID
func (r ApiListHoldsRequest) CardId(cardId string) ApiListHoldsRequest {
	r.cardId = &cardId
	return r
}

// Reference ID
func (r ApiListHoldsRequest) ReferenceId(referenceId string) ApiListHoldsRequest {
	r.referenceId = &referenceId
	return r
}

func (r ApiListHoldsRequest) Limit(limit int32) ApiListHoldsRequest {
	r.limit = &limit
	return r
}

// Only display transactions matching the given subtype
func (r ApiListHoldsRequest) Subtype(subtype string) ApiListHoldsRequest {
	r.subtype = &subtype
	return r
}

func (r ApiListHoldsRequest) Execute() (*PendingTransactions, *http.Response, error) {
	return r.ApiService.ListHoldsExecute(r)
}

/*
ListHolds List holds

Get paginated list of holds matching the provided filters


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHoldsRequest
*/
func (a *HoldAPIService) ListHolds(ctx context.Context) ApiListHoldsRequest {
	return ApiListHoldsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PendingTransactions
func (a *HoldAPIService) ListHoldsExecute(r ApiListHoldsRequest) (*PendingTransactions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PendingTransactions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HoldAPIService.ListHolds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hold"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeChildTransactions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_child_transactions", r.includeChildTransactions, "")
	}
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_date", r.fromDate, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to_date", r.toDate, "")
	}
	if r.transactionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transaction_id", r.transactionId, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "csv")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "")
	}
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "idempotency_key", r.idempotencyKey, "csv")
	}
	if r.accountNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_no", r.accountNo, "")
	}
	if r.excludeJitTransactions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_jit_transactions", r.excludeJitTransactions, "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
	}
	if r.cardId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "card_id", r.cardId, "")
	}
	if r.referenceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reference_id", r.referenceId, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.subtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subtype", r.subtype, "")
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