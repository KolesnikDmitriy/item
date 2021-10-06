//go:build e2e
// +build e2e

package tests

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	item "github.com/KolesnikDmitriy/item/pkg/api"
)

func TestItemGRPC(t *testing.T) {
	t.Parallel()

	t.Run("should return valid description and title by valid id", func(t *testing.T) {
		t.Parallel()

		req := item.GetItemRequest{Id: 1}

		res, err := api.Item(&req)

		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, "Book", res.Title)
		assert.Equal(t, "Very Great Book", res.Description)
	})

	t.Run("should retrun error by invalid request", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name string
			req  item.GetItemRequest
		}{
			{name: "zero value", req: item.GetItemRequest{}},
			{name: "negative value", req: item.GetItemRequest{Id: -1}},
			{name: "max negative value", req: item.GetItemRequest{Id: math.MinInt64}},
		}
		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				_, err := api.Item(&tc.req)

				require.Error(t, err)
				require.Contains(t, err.Error(), "wrong id")
			})
		}

	})
}
