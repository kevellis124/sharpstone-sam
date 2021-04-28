package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("Returns a stubbed list of nba odds", func(t *testing.T) {
		result, err := GetTodaysOdds("nba")

		if err != nil {
			t.Fatal("Function returned an error when it shouldn't")
		}
		assert.Equal(t, "test", result)
	})
}
