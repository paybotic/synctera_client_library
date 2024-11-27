/*
Synctera API

Testing KYCVerificationDeprecatedAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_synctera_client_KYCVerificationDeprecatedAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KYCVerificationDeprecatedAPIService CreateCustomerVerificationResult", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var customerId string
		resp, httpRes, err := apiClient.KYCVerificationDeprecatedAPI.CreateCustomerVerificationResult(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test KYCVerificationDeprecatedAPIService GetVerification", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var verificationId string
		resp, httpRes, err := apiClient.KYCVerificationDeprecatedAPI.GetVerification(context.Background(), verificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test KYCVerificationDeprecatedAPIService ListVerifications", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var customerId string
		resp, httpRes, err := apiClient.KYCVerificationDeprecatedAPI.ListVerifications(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test KYCVerificationDeprecatedAPIService VerifyCustomer", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var customerId string
		verificationRequest := openapiclient.VerificationRequest{
			CustomerID: customerId,
			// Add other required fields based on your API requirements
		}

		resp, httpRes, err := apiClient.KYCVerificationDeprecatedAPI.VerifyCustomer(context.Background(), verificationRequest).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
