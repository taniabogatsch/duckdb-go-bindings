package duckdb_go_bindings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	require.Equal(t, IdxT(2048), VectorSize())
}
