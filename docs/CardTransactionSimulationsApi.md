# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulateAuthorization**](CardTransactionSimulationsApi.md#SimulateAuthorization) | **Post** /cards/transaction_simulations/authorization | Simulate authorization
[**SimulateAuthorizationAdvice**](CardTransactionSimulationsApi.md#SimulateAuthorizationAdvice) | **Post** /cards/transaction_simulations/authorization/advice | Simulate authorization advice
[**SimulateBalanceInquiry**](CardTransactionSimulationsApi.md#SimulateBalanceInquiry) | **Post** /cards/transaction_simulations/financial/balance_inquiry | Simulate balance inquiry
[**SimulateClearing**](CardTransactionSimulationsApi.md#SimulateClearing) | **Post** /cards/transaction_simulations/clearing | Simulate clearing or refund
[**SimulateFinancial**](CardTransactionSimulationsApi.md#SimulateFinancial) | **Post** /cards/transaction_simulations/financial | Simulate financial
[**SimulateFinancialAdvice**](CardTransactionSimulationsApi.md#SimulateFinancialAdvice) | **Post** /cards/transaction_simulations/financial/advice | Simulate financial advice
[**SimulateOriginalCredit**](CardTransactionSimulationsApi.md#SimulateOriginalCredit) | **Post** /cards/transaction_simulations/financial/original_credit | Simulate OCT
[**SimulateReversal**](CardTransactionSimulationsApi.md#SimulateReversal) | **Post** /cards/transaction_simulations/reversal | Simulate reversal
[**SimulateWithdrawal**](CardTransactionSimulationsApi.md#SimulateWithdrawal) | **Post** /cards/transaction_simulations/financial/withdrawal | Simulate ATM withdrawal

# **SimulateAuthorization**
> SimulationResponseModel SimulateAuthorization(ctx, optional)
Simulate authorization

Simulate an `authorization` type transaction by including the `card_token` and other authorization details in your request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardTransactionSimulationsApiSimulateAuthorizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardTransactionSimulationsApiSimulateAuthorizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthRequestModel**](AuthRequestModel.md)| Authorization details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateAuthorizationAdvice**
> SimulationResponseModel SimulateAuthorizationAdvice(ctx, optional)
Simulate authorization advice

An authorization advice allows an amount to be decreased after the authorization. This endpoint allows you to simulate post-swipe adjustments.  Simulate an `authorization.advice` type transaction by including the `original_transaction_token` and other authorization details in your request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardTransactionSimulationsApiSimulateAuthorizationAdviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardTransactionSimulationsApiSimulateAuthorizationAdviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthorizationAdviceModel**](AuthorizationAdviceModel.md)| Authorization advice details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateBalanceInquiry**
> SimulationResponseModel SimulateBalanceInquiry(ctx, body)
Simulate balance inquiry

Simulate a `pindebit.balanceinquiry` type transaction by sending a POST request to the `/simulate/financial/balanceinquiry` endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BalanceInquiryRequestModel**](BalanceInquiryRequestModel.md)| Balance inquiry details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateClearing**
> SimulationResponseModel SimulateClearing(ctx, optional)
Simulate clearing or refund

Simulate an `authorization.clearing` type transaction by including the `original_transaction_token` and `amount` in your request. To simulate a refund type transaction, set the `is_refund` field to true. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardTransactionSimulationsApiSimulateClearingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardTransactionSimulationsApiSimulateClearingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClearingModel**](ClearingModel.md)| Transaction clearing or refund details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateFinancial**
> SimulationResponseModel SimulateFinancial(ctx, body)
Simulate financial

A \"financial\" is a transaction message class that includes ATM transactions, PIN-debit transactions, and balance inquiries.  Simulate a `pindebit` type transaction by including the `card_token` and `amount` in your request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FinancialRequestModel**](FinancialRequestModel.md)| Financial details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateFinancialAdvice**
> SimulationResponseModel SimulateFinancialAdvice(ctx, body)
Simulate financial advice

Simulate a financial advice by including the `original_transaction_token` and other authorization details in JSON format in the body of the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationAdviceModel**](AuthorizationAdviceModel.md)| Financial advice details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateOriginalCredit**
> SimulationResponseModel SimulateOriginalCredit(ctx, body)
Simulate OCT

This Original Credit Transaction (OCT) enables the cardholder to receive funds on the specified card from an external source via the card network. Use this endpoint to simulate a transaction that is similar to a wire transfer and not linked to any purchase.  Simulate an OCT by including the `card_token`, `amount`, `mid`, and `type` in your request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OriginalCreditRequestModel**](OriginalCreditRequestModel.md)| OCT details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateReversal**
> SimulationResponseModel SimulateReversal(ctx, body)
Simulate reversal

A reversal releases the hold that was placed on account funds by an authorization, thus returning the funds to the account.  Simulate an `authorization.reversal` type transaction by including the `original_transaction_token` and `amount` in your request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReversalModel**](ReversalModel.md)| Reversal details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateWithdrawal**
> SimulationResponseModel SimulateWithdrawal(ctx, body)
Simulate ATM withdrawal

Simulate a `pindebit.atm.withdrawal` type transaction by including the `card_token` and `amount` in your request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WithdrawalRequestModel**](WithdrawalRequestModel.md)| ATM withdrawal details | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

