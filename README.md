# duckdb-go-bindings

This repository wraps DuckDB's C API calls in Go native types and functions.

The main module (`github.com/duckdb/duckdb-go-bindings`) does not link any pre-built static library.

TODO: example on static linking
TODO: example on dynamic linking

There are also a few pre-built static libraries for different OS + architecture combinations.
Here's a list:
- `github.com/duckdb/duckdb-go-bindings/`...
  - `darwin_amd64`
  - `darwin_arm64`
  - `linux_amd64`
  - `linux_arm64`
  - `windows_amd64`

The table below shows which DuckDB version corresponds to which module version.

| duckdb version | duckdb-go-bindings | darwin amd64 | darwin arm64 | linux amd64 | linux arm64 | windows amd64 |
| ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- |
| `v1.1.3` | `v0.1.2` | `v0.1.0` | `v0.1.0` | `v0.1.1` | `v0.1.1` | `v0.1.1` |
