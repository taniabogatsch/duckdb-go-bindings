DUCKDB_VERSION=v1.3.0

fetch.static.libs:
	cd ${DIRECTORY} && \
	curl -OL https://github.com/duckdb/duckdb/releases/download/${DUCKDB_VERSION}/${FILENAME}.zip && \
	rm *.a && \
	rm -f duckdb.h && \
	unzip ${FILENAME}.zip && \
	rm -f ${FILENAME}.zip

fetch.custom.duckdb:
	cd .. && \
	git clone ${DIRECTORY} && \
	cd duckdb && \
	git checkout ${BRANCH}

extract.custom.libs:
	cd .. && \
	rm -rf duckdb-go-bindings/custom-duckdb-linux-amd64/*.a && \
	rm -f duckdb-go-bindings/custom-duckdb-linux-amd64/duckdb.h && \
	mv duckdb/src/include/duckdb.h duckdb-go-bindings/custom-duckdb-linux-amd64/duckdb.h && \
	cp duckdb/build/release/libs/* duckdb-go-bindings/custom-duckdb-linux-amd64/. && \
	cd duckdb-go-bindings

update.binding:
	rm -f ${DIRECTORY}/bindings.go && \
	cp bindings.go ${DIRECTORY}/bindings.go

test.dynamic.lib:
	mkdir dynamic-dir && \
	cd dynamic-dir && \
	curl -OL https://github.com/duckdb/duckdb/releases/download/${DUCKDB_VERSION}/${FILENAME}.zip && \
	unzip ${FILENAME}.zip
