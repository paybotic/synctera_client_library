/*
Synctera API

Testing MonitoringAPIService

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

func Test_synctera_client_MonitoringAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MonitoringAPIService CreateSubscription", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MonitoringAPI.CreateSubscription(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoringAPIService DeleteSubscription", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.MonitoringAPI.DeleteSubscription(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoringAPIService GetAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var alertId string

		resp, httpRes, err := apiClient.MonitoringAPI.GetAlert(context.Background(), alertId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoringAPIService GetSubscription", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.MonitoringAPI.GetSubscription(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoringAPIService ListAlerts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MonitoringAPI.ListAlerts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoringAPIService ListSubscriptions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MonitoringAPI.ListSubscriptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoringAPIService UpdateAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var alertId string

		httpRes, err := apiClient.MonitoringAPI.UpdateAlert(context.Background(), alertId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
