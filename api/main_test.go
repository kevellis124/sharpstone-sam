package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("Test invokes get odds", func(t *testing.T) {
		result, _ := handler(events.APIGatewayProxyRequest{})
		assert.Equal(t, "test", result.Body)
		assert.Equal(t, 200, result.StatusCode)
	})
}
