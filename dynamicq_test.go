package dynamicq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDynamic(t *testing.T) {
	var (
		dq    = Dynamic{}
		query = "SELECT * FROM user"
	)

	dq.Add("invited = ?", true)
	dq.Glue(&query)
	dq.Attr(&query, "SORT BY id")
	dq.Limit(&query, 10)
	dq.Offset(&query, 10)

	args := dq.Args()

	require.Equal(t, "SELECT * FROM user WHERE invited = ? SORT BY id LIMIT ? OFFSET ?", query)
	require.Equal(t, true, args[0])
	require.Equal(t, int64(10), args[1])
	require.Equal(t, int64(10), args[2])
}

func TestDynamic_Glue(t *testing.T) {
	table := []struct {
		query         string
		handler       func(dq *Dynamic)
		expectedQuery string
	}{
		{
			query:         "SELECT id FROM user",
			handler:       func(dq *Dynamic) {},
			expectedQuery: "SELECT id FROM user",
		},
		{
			query: "SELECT id FROM user",
			handler: func(dq *Dynamic) {
				dq.Add("role = ?", "admin")
			},
			expectedQuery: "SELECT id FROM user WHERE role = ?",
		},
		{
			query: "SELECT id FROM user",
			handler: func(dq *Dynamic) {
				dq.Add("role = ?", "admin")
				dq.Add("age = ?", 18)
			},
			expectedQuery: "SELECT id FROM user WHERE role = ? AND age = ?",
		},
		{
			query: "SELECT id FROM user",
			handler: func(dq *Dynamic) {
				dq.Add("role = ? OR invited = ?", "admin", true)
			},
			expectedQuery: "SELECT id FROM user WHERE role = ? OR invited = ?",
		},
	}

	for _, item := range table {
		var dq Dynamic

		item.handler(&dq)
		dq.Glue(&item.query)

		require.Equal(t, item.query, item.expectedQuery)
	}
}
