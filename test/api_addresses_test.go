/*
Synctera API

Testing AddressesAPIService

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

func Test_synctera_client_AddressesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AddressesAPIService CreateAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AddressesAPI.CreateAddress(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressesAPIService GetAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var addressId string

		resp, httpRes, err := apiClient.AddressesAPI.GetAddress(context.Background(), addressId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressesAPIService ListAddresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AddressesAPI.ListAddresses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressesAPIService UpdateAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var addressId string

		resp, httpRes, err := apiClient.AddressesAPI.UpdateAddress(context.Background(), addressId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
