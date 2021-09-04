package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	item "github.com/KolesnikDmitriy/item/pkg/api"
)

func TestItem(t *testing.T) {
	t.Parallel()

	t.Run("valid id", func(t *testing.T) {
		t.Parallel()

		req := item.GetItemRequest{Id: 1}

		res, err := api.Item(&req)

		require.NoError(t, err)
		assert.EqualValues(t, "Very Great Book", res.Description)
		assert.EqualValues(t, "Book", res.Title)
	})

	t.Run("empty request", func(t *testing.T) {
		t.Parallel()

		req := item.GetItemRequest{}

		_, err := api.Item(&req)

		require.Error(t, err)
		require.Contains(t, err.Error(), "wrong id")
	})
}
