/*
Synctera API

Testing PersonsAPIService

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

func Test_synctera_client_PersonsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PersonsAPIService CreatePerson", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PersonsAPI.CreatePerson(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonsAPIService CreatePersonalId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PersonsAPI.CreatePersonalId(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonsAPIService DeletePersonalId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var personalIdId string

		resp, httpRes, err := apiClient.PersonsAPI.DeletePersonalId(context.Background(), personalIdId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonsAPIService GetPerson", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonsAPI.GetPerson(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonsAPIService ListPersons", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PersonsAPI.ListPersons(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonsAPIService UpdatePerson", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonsAPI.UpdatePerson(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonsAPIService UpdatePersonalId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var personalIdId string

		resp, httpRes, err := apiClient.PersonsAPI.UpdatePersonalId(context.Background(), personalIdId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}