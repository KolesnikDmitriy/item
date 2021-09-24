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

	t.Run("valid id", func(t *testing.T) {
		t.Parallel()

		req := item.GetItemRequest{Id: 1}

		res, err := api.Item(&req)

		require.NoError(t, err)
		assert.EqualValues(t, "Very Great Book", res.Description)
		assert.EqualValues(t, "Book", res.Title)
	})

	t.Run("invalid request", func(t *testing.T) {
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
				_, err := api.Item(&tc.req)

				require.Error(t, err)
				require.Contains(t, err.Error(), "wrong id")
			})
		}

	})
}
