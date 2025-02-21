package duckdb_go_bindings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	require.Equal(t, uint64(2048), VectorSize())
}
