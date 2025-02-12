package duckdb_go_bindings

/*
#cgo CPPFLAGS: -DDUCKDB_STATIC_BUILD
#cgo LDFLAGS:  -lduckdb_part1 -lduckdb_part2  -lstdc++ -lm -ldl -L${SRCDIR}/libs
#include <duckdb.h>
*/
import "C"
