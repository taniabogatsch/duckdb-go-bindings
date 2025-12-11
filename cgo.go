//go:build !duckdb_use_lib && !duckdb_use_static_lib

package duckdb_go_bindings

/*
#cgo CPPFLAGS: -I${SRCDIR}/include -DDUCKDB_STATIC_BUILD

#cgo darwin,arm64 LDFLAGS: -lduckdb_static -lautocomplete_extension -lcore_functions_extension -licu_extension -ljson_extension -lparquet_extension -ltpcds_extension -ltpch_extension -lduckdb_fastpforlib -lduckdb_fmt -lduckdb_fsst -lduckdb_hyperloglog -lduckdb_mbedtls -lduckdb_miniz -lduckdb_pg_query -lduckdb_re2 -lduckdb_skiplistlib -lduckdb_utf8proc -lduckdb_yyjson -lduckdb_zstd -lc++ -L${SRCDIR}/lib/darwin-arm64
#cgo darwin,amd64 LDFLAGS: -lduckdb_static -lautocomplete_extension -lcore_functions_extension -licu_extension -ljson_extension -lparquet_extension -ltpcds_extension -ltpch_extension -lduckdb_fastpforlib -lduckdb_fmt -lduckdb_fsst -lduckdb_hyperloglog -lduckdb_mbedtls -lduckdb_miniz -lduckdb_pg_query -lduckdb_re2 -lduckdb_skiplistlib -lduckdb_utf8proc -lduckdb_yyjson -lduckdb_zstd -lc++ -L${SRCDIR}/lib/darwin-amd64

#cgo linux,arm64 LDFLAGS: -rdynamic -lduckdb_static -lautocomplete_extension -lcore_functions_extension -licu_extension -ljson_extension -lparquet_extension -ljemalloc_extension -ltpcds_extension -ltpch_extension -lduckdb_fastpforlib -lduckdb_fmt -lduckdb_fsst -lduckdb_hyperloglog -lduckdb_mbedtls -lduckdb_miniz -lduckdb_pg_query -lduckdb_re2 -lduckdb_skiplistlib -lduckdb_utf8proc -lduckdb_yyjson -lduckdb_zstd -lstdc++ -lm -ldl -L${SRCDIR}/lib/linux-arm64
#cgo linux,amd64 LDFLAGS: -rdynamic -lduckdb_static -lautocomplete_extension -lcore_functions_extension -licu_extension -ljson_extension -lparquet_extension -ljemalloc_extension -ltpcds_extension -ltpch_extension -lduckdb_fastpforlib -lduckdb_fmt -lduckdb_fsst -lduckdb_hyperloglog -lduckdb_mbedtls -lduckdb_miniz -lduckdb_pg_query -lduckdb_re2 -lduckdb_skiplistlib -lduckdb_utf8proc -lduckdb_yyjson -lduckdb_zstd -lstdc++ -lm -ldl -L${SRCDIR}/lib/linux-amd64

#cgo windows,amd64 LDFLAGS: -lduckdb_static -lautocomplete_extension -lcore_functions_extension -licu_extension -ljson_extension -lparquet_extension -ltpcds_extension -ltpch_extension -lduckdb_fastpforlib -lduckdb_fmt -lduckdb_fsst -lduckdb_hyperloglog -lduckdb_mbedtls -lduckdb_miniz -lduckdb_pg_query -lduckdb_re2 -lduckdb_skiplistlib -lduckdb_utf8proc -lduckdb_yyjson -lduckdb_zstd -lws2_32 -lwsock32 -lrstrtmgr -lstdc++ -lm --static -L${SRCDIR}/lib/windows-amd64

#include <duckdb.h>
*/
import "C"
