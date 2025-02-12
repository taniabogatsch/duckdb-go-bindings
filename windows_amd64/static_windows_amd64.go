package duckdb_go_bindings

/*
#cgo CPPFLAGS: -DDUCKDB_STATIC_BUILD
#cgo LDFLAGS: -lduckdb
#cgo LDFLAGS: -lws2_32 -lwsock32 -lrstrtmgr -lstdc++ -lm --static -L${SRCDIR}/libs
#include <duckdb.h>
*/
import "C"
