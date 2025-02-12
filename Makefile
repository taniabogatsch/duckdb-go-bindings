#DUCKDB_VERSION=v1.2.0
#DUCKDB_DOWNLOAD_PATH=https://github.com/duckdb/duckdb/releases/download/
#
##https://github.com/duckdb/duckdb/releases/download/v1.2.0/static-lib-osx-arm64.zip

split.linux.amd64.artefacts:
	cd linux_amd64/libs && \
	${AR} -x libduckdb.a && \
	ls && \
	mkdir dir1 && \
	mkdir dir2
