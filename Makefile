DUCKDB_VERSION=v1.2.0
DUCKDB_DOWNLOAD_PATH=https://github.com/duckdb/duckdb/releases/download/

#https://github.com/duckdb/duckdb/releases/download/v1.2.0/static-lib-osx-arm64.zip

fetch.artefacts:
	rm -f darwin_amd64/duckdb.h
	rm -f darwin_arm64/duckdb.h
	rm -f linux_amd64/duckdb.h
	rm -f linux_arm64/duckdb.h
	rm -f windows_amd64/duckdb.h
