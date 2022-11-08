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
