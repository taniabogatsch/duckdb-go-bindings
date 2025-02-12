package duckdb_go_bindings_linux_amd64

/*
#cgo CPPFLAGS: -DDUCKDB_STATIC_BUILD
#cgo LDFLAGS: -lstdc++ -lm -ldl -lduckdb_part1 -lduckdb_part2 -L${SRCDIR}/libs
#include <duckdb.h>
*/
import "C"
