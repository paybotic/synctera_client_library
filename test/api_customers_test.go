/*
Synctera API

Testing CustomersAPIService

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

func Test_synctera_client_CustomersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomersAPIService CreateCustomer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService GetCustomer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.GetCustomer(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService ListCustomers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomersAPI.ListCustomers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService PatchCustomer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.PatchCustomer(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService UpdateCustomer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.UpdateCustomer(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
