//go:build e2e
// +build e2e

package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestItemREST(t *testing.T) {
	t.Parallel()

	t.Run("should return valid description and title by valid id", func(t *testing.T) {
		t.Parallel()

		res, err := http.Get(fmt.Sprintf("http://%v/v1/item/1", cfg.RestUrl))

		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		require.NoError(t, err)
		var item struct {
			Title       string
			Description string
		}
		err = json.Unmarshal(body, &item)
		require.NoError(t, err)
		assert.Equal(t, "Book", item.Title)
		assert.Equal(t, "Very Great Book", item.Description)
	})

	t.Run("should retrun error by invalid request", func(t *testing.T) {
		t.Parallel()

		res, err := http.Get(fmt.Sprintf("http://%v/v1/item/0", cfg.RestUrl))

		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}
