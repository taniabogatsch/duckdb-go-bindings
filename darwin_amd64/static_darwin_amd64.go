package duckdb_go_bindings

/*
#cgo CPPFLAGS: -DDUCKDB_STATIC_BUILD
#cgo LDFLAGS: -lduckdb
#cgo LDFLAGS: -lc++ -L${SRCDIR}/libs
#include <duckdb.h>
*/
import "C"
