package duckdb_go_bindings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOpen(t *testing.T) {
	require.Equal(t, uint64(2048), VectorSize())
}
