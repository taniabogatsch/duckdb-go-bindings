split.artefacts:
	mkdir ${DIRECTORY}/libs/core && \
	mkdir ${DIRECTORY}/libs/corefunctions && \
	mv ${DIRECTORY}/libs/libduckdb.a ${DIRECTORY}/libs/core/libduckdb.a && \
	cd ${DIRECTORY}/libs/core && \
	${AR} -x libduckdb.a && \
	ls && \
	rm -f __.SYMDEF && \
	rm -f libduckdb.a && \
	mv ub_duckdb_core_functions_algebraic.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_distributive.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_nested.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_regression.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_array.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_bit.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_blob.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_date.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_debug.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_enum.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_generic.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_list.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_map.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_math.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_operators.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_random.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_string.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_struct.cpp.${FILETYPE} ../corefunctions/ && \
    mv ub_duckdb_core_functions_union.cpp.${FILETYPE} ../corefunctions/ && \
    mv lambda_functions.cpp.${FILETYPE} ../corefunctions/ && \
    cd .. && \
	${AR} r libcore.a core/* && \
	${AR} r libcorefunctions.a corefunctions/* && \
	rm -rf core && \
	rm -rf corefunctions

fetch.static.lib:
	cd ${DIRECTORY} && \
	curl -OL https://github.com/duckdb/duckdb/releases/download/${VERSION}/${FILENAME}.zip && \
	rm -rf libs && \
	mkdir libs && \
	rm -f duckdb.h && \
	unzip ${FILENAME}.zip && \
	mv libduckdb_bundle.a libs/libduckdb.a && \
	rm -f ${FILENAME}.zip
