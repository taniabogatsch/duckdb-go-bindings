split.artefacts:
	mkdir ${DIRECTORY}/libs/dir1 && \
	mkdir ${DIRECTORY}/libs/dir2 && \
	mv ${DIRECTORY}/libs/libduckdb.a ${DIRECTORY}/libs/dir1/libduckdb.a && \
	cd ${DIRECTORY}/libs/dir1 && \
	${AR} -x libduckdb.a && \
	ls && \
	rm -f __.SYMDEF && \
	rm -f libduckdb.a && \
	cd ../../.. && \
	python3 scripts/move_files.py -src ${DIRECTORY}/libs/dir1/ -dst ${DIRECTORY}/libs/dir2/ && \
	cd ${DIRECTORY}/libs/ && \
	${AR} r libduckdb_part1.a dir1/* && \
	${AR} r libduckdb_part2.a dir2/* && \
	rm -rf dir1 && \
	rm -rf dir2

fetch.static.lib:
	cd ${DIRECTORY} && \
	curl -OL https://github.com/duckdb/duckdb/releases/download/${VERSION}/${FILENAME}.zip && \
	rm -rf libs && \
	mkdir libs && \
	rm -f duckdb.h && \
	unzip ${FILENAME}.zip && \
	mv libduckdb_bundle.a libs/libduckdb.a && \
	rm -f ${FILENAME}.zip
