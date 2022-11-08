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

	dq.Where("invited = ?", true)
	dq.Glue(&query)
	dq.Attr(&query, "ORDER BY id")
	dq.Limit(&query, 10)
	dq.Offset(&query, 10)

	args := dq.Args()

	require.Equal(t, "SELECT * FROM user WHERE invited = ? ORDER BY id LIMIT ? OFFSET ?", query)
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
				dq.Where("role = ?", "admin")
			},
			expectedQuery: "SELECT id FROM user WHERE role = ?",
		},
		{
			query: "SELECT id FROM user",
			handler: func(dq *Dynamic) {
				dq.Where("role = ?", "admin")
				dq.Where("age = ?", 18)
			},
			expectedQuery: "SELECT id FROM user WHERE role = ? AND age = ?",
		},
		{
			query: "SELECT id FROM user",
			handler: func(dq *Dynamic) {
				dq.Where("role = ? OR invited = ?", "admin", true)
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

func TestDynamic_Limit(t *testing.T) {
	table := []struct {
		query         string
		limit         int64
		expectedQuery string
	}{
		{
			query:         "SELECT id FROM user",
			limit:         0,
			expectedQuery: "SELECT id FROM user",
		},
		{
			query:         "SELECT id FROM user",
			limit:         15,
			expectedQuery: "SELECT id FROM user LIMIT ?",
		},
	}

	for _, item := range table {
		var dq Dynamic

		dq.Limit(&item.query, item.limit)

		require.Equal(t, item.query, item.expectedQuery)
	}
}

func TestDynamic_Offset(t *testing.T) {
	table := []struct {
		query         string
		offset        int64
		expectedQuery string
	}{
		{
			query:         "SELECT id FROM user",
			offset:        0,
			expectedQuery: "SELECT id FROM user",
		},
		{
			query:         "SELECT id FROM user",
			offset:        15,
			expectedQuery: "SELECT id FROM user OFFSET ?",
		},
	}

	for _, item := range table {
		var dq Dynamic

		dq.Offset(&item.query, item.offset)

		require.Equal(t, item.query, item.expectedQuery)
	}
}
