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

	t.Run("valid id", func(t *testing.T) {
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
		assert.EqualValues(t, "Very Great Book", item.Description)
		assert.EqualValues(t, "Book", item.Title)
	})

	t.Run("empty request", func(t *testing.T) {
		t.Parallel()

		res, err := http.Get(fmt.Sprintf("http://%v/v1/item/0", cfg.RestUrl))

		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}
