package duckdb_go_bindings

/*
#include <duckdb.h>
*/
import "C"
import (
	"unsafe"
)

// ------------------------------------------------------------------ //
// Enums
// ------------------------------------------------------------------ //

// Type wraps duckdb_type.
type Type C.duckdb_type

const (
	TypeInvalid     Type = C.DUCKDB_TYPE_INVALID
	TypeBoolean     Type = C.DUCKDB_TYPE_BOOLEAN
	TypeTinyInt     Type = C.DUCKDB_TYPE_TINYINT
	TypeSmallInt    Type = C.DUCKDB_TYPE_SMALLINT
	TypeInteger     Type = C.DUCKDB_TYPE_INTEGER
	TypeBigInt      Type = C.DUCKDB_TYPE_BIGINT
	TypeUTinyInt    Type = C.DUCKDB_TYPE_UTINYINT
	TypeUSmallInt   Type = C.DUCKDB_TYPE_USMALLINT
	TypeUInteger    Type = C.DUCKDB_TYPE_UINTEGER
	TypeUBigInt     Type = C.DUCKDB_TYPE_UBIGINT
	TypeFloat       Type = C.DUCKDB_TYPE_FLOAT
	TypeDouble      Type = C.DUCKDB_TYPE_DOUBLE
	TypeTimestamp   Type = C.DUCKDB_TYPE_TIMESTAMP
	TypeDate        Type = C.DUCKDB_TYPE_DATE
	TypeTime        Type = C.DUCKDB_TYPE_TIME
	TypeInterval    Type = C.DUCKDB_TYPE_INTERVAL
	TypeHugeInt     Type = C.DUCKDB_TYPE_HUGEINT
	TypeUHugeInt    Type = C.DUCKDB_TYPE_UHUGEINT
	TypeVarchar     Type = C.DUCKDB_TYPE_VARCHAR
	TypeBlob        Type = C.DUCKDB_TYPE_BLOB
	TypeDecimal     Type = C.DUCKDB_TYPE_DECIMAL
	TypeTimestampS  Type = C.DUCKDB_TYPE_TIMESTAMP_S
	TypeTimestampMS Type = C.DUCKDB_TYPE_TIMESTAMP_MS
	TypeTimestampNS Type = C.DUCKDB_TYPE_TIMESTAMP_NS
	TypeEnum        Type = C.DUCKDB_TYPE_ENUM
	TypeList        Type = C.DUCKDB_TYPE_LIST
	TypeStruct      Type = C.DUCKDB_TYPE_STRUCT
	TypeMap         Type = C.DUCKDB_TYPE_MAP
	TypeArray       Type = C.DUCKDB_TYPE_ARRAY
	TypeUUID        Type = C.DUCKDB_TYPE_UUID
	TypeUnion       Type = C.DUCKDB_TYPE_UNION
	TypeBit         Type = C.DUCKDB_TYPE_BIT
	TypeTimeTZ      Type = C.DUCKDB_TYPE_TIME_TZ
	TypeTimestampTZ Type = C.DUCKDB_TYPE_TIMESTAMP_TZ
	TypeAny         Type = C.DUCKDB_TYPE_ANY
	TypeVarInt      Type = C.DUCKDB_TYPE_VARINT
	TypeSQLNull     Type = C.DUCKDB_TYPE_SQLNULL
)

// State wraps duckdb_state.
type State C.duckdb_state

const (
	StateSuccess State = C.DuckDBSuccess
	StateError   State = C.DuckDBError
)

// PendingState wraps duckdb_pending_state.
type PendingState C.duckdb_pending_state

const (
	PendingStateResultReady      PendingState = C.DUCKDB_PENDING_RESULT_READY
	PendingStateResultNotReady   PendingState = C.DUCKDB_PENDING_RESULT_NOT_READY
	PendingStateError            PendingState = C.DUCKDB_PENDING_ERROR
	PendingStateNoTasksAvailable PendingState = C.DUCKDB_PENDING_NO_TASKS_AVAILABLE
)

// ResultType wraps duckdb_result_type.
type ResultType C.duckdb_result_type

const (
	ResultTypeInvalid     ResultType = C.DUCKDB_RESULT_TYPE_INVALID
	ResultTypeChangedRows ResultType = C.DUCKDB_RESULT_TYPE_CHANGED_ROWS
	ResultTypeNothing     ResultType = C.DUCKDB_RESULT_TYPE_NOTHING
	ResultTypeQueryResult ResultType = C.DUCKDB_RESULT_TYPE_QUERY_RESULT
)

// StatementType wraps duckdb_statement_type.
type StatementType C.duckdb_statement_type

const (
	StatementTypeInvalid     StatementType = C.DUCKDB_STATEMENT_TYPE_INVALID
	StatementTypeSelect      StatementType = C.DUCKDB_STATEMENT_TYPE_SELECT
	StatementTypeInsert      StatementType = C.DUCKDB_STATEMENT_TYPE_INSERT
	StatementTypeUpdate      StatementType = C.DUCKDB_STATEMENT_TYPE_UPDATE
	StatementTypeExplain     StatementType = C.DUCKDB_STATEMENT_TYPE_EXPLAIN
	StatementTypeDelete      StatementType = C.DUCKDB_STATEMENT_TYPE_DELETE
	StatementTypePrepare     StatementType = C.DUCKDB_STATEMENT_TYPE_PREPARE
	StatementTypeCreate      StatementType = C.DUCKDB_STATEMENT_TYPE_CREATE
	StatementTypeExecute     StatementType = C.DUCKDB_STATEMENT_TYPE_EXECUTE
	StatementTypeAlter       StatementType = C.DUCKDB_STATEMENT_TYPE_ALTER
	StatementTypeTransaction StatementType = C.DUCKDB_STATEMENT_TYPE_TRANSACTION
	StatementTypeCopy        StatementType = C.DUCKDB_STATEMENT_TYPE_COPY
	StatementTypeAnalyze     StatementType = C.DUCKDB_STATEMENT_TYPE_ANALYZE
	StatementTypeVariableSet StatementType = C.DUCKDB_STATEMENT_TYPE_VARIABLE_SET
	StatementTypeCreateFunc  StatementType = C.DUCKDB_STATEMENT_TYPE_CREATE_FUNC
	StatementTypeDrop        StatementType = C.DUCKDB_STATEMENT_TYPE_DROP
	StatementTypeExport      StatementType = C.DUCKDB_STATEMENT_TYPE_EXPORT
	StatementTypePragma      StatementType = C.DUCKDB_STATEMENT_TYPE_PRAGMA
	StatementTypeVacuum      StatementType = C.DUCKDB_STATEMENT_TYPE_VACUUM
	StatementTypeCall        StatementType = C.DUCKDB_STATEMENT_TYPE_CALL
	StatementTypeSet         StatementType = C.DUCKDB_STATEMENT_TYPE_SET
	StatementTypeLoad        StatementType = C.DUCKDB_STATEMENT_TYPE_LOAD
	StatementTypeRelation    StatementType = C.DUCKDB_STATEMENT_TYPE_RELATION
	StatementTypeExtension   StatementType = C.DUCKDB_STATEMENT_TYPE_EXTENSION
	StatementTypeLogicalPlan StatementType = C.DUCKDB_STATEMENT_TYPE_LOGICAL_PLAN
	StatementTypeAttach      StatementType = C.DUCKDB_STATEMENT_TYPE_ATTACH
	StatementTypeDetach      StatementType = C.DUCKDB_STATEMENT_TYPE_DETACH
	StatementTypeMulti       StatementType = C.DUCKDB_STATEMENT_TYPE_MULTI
)

// ErrorType wraps duckdb_error_type.
type ErrorType C.duckdb_error_type

const (
	ErrorTypeInvalid              ErrorType = C.DUCKDB_ERROR_INVALID
	ErrorTypeOutOfRange           ErrorType = C.DUCKDB_ERROR_OUT_OF_RANGE
	ErrorTypeConversion           ErrorType = C.DUCKDB_ERROR_CONVERSION
	ErrorTypeUnknownType          ErrorType = C.DUCKDB_ERROR_UNKNOWN_TYPE
	ErrorTypeDecimal              ErrorType = C.DUCKDB_ERROR_DECIMAL
	ErrorTypeMismatchType         ErrorType = C.DUCKDB_ERROR_MISMATCH_TYPE
	ErrorTypeDivideByZero         ErrorType = C.DUCKDB_ERROR_DIVIDE_BY_ZERO
	ErrorTypeObjectSize           ErrorType = C.DUCKDB_ERROR_OBJECT_SIZE
	ErrorTypeInvalidType          ErrorType = C.DUCKDB_ERROR_INVALID_TYPE
	ErrorTypeSerialization        ErrorType = C.DUCKDB_ERROR_SERIALIZATION
	ErrorTypeTransaction          ErrorType = C.DUCKDB_ERROR_TRANSACTION
	ErrorTypeNotImplemented       ErrorType = C.DUCKDB_ERROR_NOT_IMPLEMENTED
	ErrorTypeExpression           ErrorType = C.DUCKDB_ERROR_EXPRESSION
	ErrorTypeCatalog              ErrorType = C.DUCKDB_ERROR_CATALOG
	ErrorTypeParser               ErrorType = C.DUCKDB_ERROR_PARSER
	ErrorTypePlanner              ErrorType = C.DUCKDB_ERROR_PLANNER
	ErrorTypeScheduler            ErrorType = C.DUCKDB_ERROR_SCHEDULER
	ErrorTypeExecutor             ErrorType = C.DUCKDB_ERROR_EXECUTOR
	ErrorTypeConstraint           ErrorType = C.DUCKDB_ERROR_CONSTRAINT
	ErrorTypeIndex                ErrorType = C.DUCKDB_ERROR_INDEX
	ErrorTypeStat                 ErrorType = C.DUCKDB_ERROR_STAT
	ErrorTypeConnection           ErrorType = C.DUCKDB_ERROR_CONNECTION
	ErrorTypeSyntax               ErrorType = C.DUCKDB_ERROR_SYNTAX
	ErrorTypeSettings             ErrorType = C.DUCKDB_ERROR_SETTINGS
	ErrorTypeBinder               ErrorType = C.DUCKDB_ERROR_BINDER
	ErrorTypeNetwork              ErrorType = C.DUCKDB_ERROR_NETWORK
	ErrorTypeOptimizer            ErrorType = C.DUCKDB_ERROR_OPTIMIZER
	ErrorTypeNullPointer          ErrorType = C.DUCKDB_ERROR_NULL_POINTER
	ErrorTypeErrorIO              ErrorType = C.DUCKDB_ERROR_IO
	ErrorTypeInterrupt            ErrorType = C.DUCKDB_ERROR_INTERRUPT
	ErrorTypeFatal                ErrorType = C.DUCKDB_ERROR_FATAL
	ErrorTypeInternal             ErrorType = C.DUCKDB_ERROR_INTERNAL
	ErrorTypeInvalidInput         ErrorType = C.DUCKDB_ERROR_INVALID_INPUT
	ErrorTypeOutOfMemory          ErrorType = C.DUCKDB_ERROR_OUT_OF_MEMORY
	ErrorTypePermission           ErrorType = C.DUCKDB_ERROR_PERMISSION
	ErrorTypeParameterNotResolved ErrorType = C.DUCKDB_ERROR_PARAMETER_NOT_RESOLVED
	ErrorTypeParameterNotAllowed  ErrorType = C.DUCKDB_ERROR_PARAMETER_NOT_ALLOWED
	ErrorTypeDependency           ErrorType = C.DUCKDB_ERROR_DEPENDENCY
	ErrorTypeHTTP                 ErrorType = C.DUCKDB_ERROR_HTTP
	ErrorTypeMissingExtension     ErrorType = C.DUCKDB_ERROR_MISSING_EXTENSION
	ErrorTypeAutoload             ErrorType = C.DUCKDB_ERROR_AUTOLOAD
	ErrorTypeSequence             ErrorType = C.DUCKDB_ERROR_SEQUENCE
	ErrorTypeInvalidConfiguration ErrorType = C.DUCKDB_INVALID_CONFIGURATION
)

// CastMode wraps duckdb_cast_mode.
type CastMode C.duckdb_cast_mode

const (
	CastModeNormal CastMode = C.DUCKDB_CAST_NORMAL
	CastModeTry    CastMode = C.DUCKDB_CAST_TRY
)

// ------------------------------------------------------------------ //
// Types
// ------------------------------------------------------------------ //

// NOTE: No wrapping for C.idx_t.

// Types without internal pointers:

type (
	Date              C.duckdb_date
	DateStruct        C.duckdb_date_struct
	Time              C.duckdb_time
	TimeStruct        C.duckdb_time_struct
	TimeTZ            C.duckdb_time_tz
	TimeTZStruct      C.duckdb_time_tz_struct
	Timestamp         C.duckdb_timestamp
	TimestampStruct   C.duckdb_timestamp_struct
	Interval          C.duckdb_interval
	HugeInt           C.duckdb_hugeint
	UHugeInt          C.duckdb_uhugeint
	Decimal           C.duckdb_decimal
	QueryProgressType C.duckdb_query_progress_type
	StringT           C.duckdb_string_t
	ListEntry         C.duckdb_list_entry
)

// duckdb_string
// duckdb_blob
// duckdb_extension_access

// Helper functions for types without internal pointers:

func DateSetDays(date *Date, days int32) {
	date.days = C.int32_t(days)
}

func DateStructGetYear(date *DateStruct) int32 {
	return int32(date.year)
}

func DateStructGetMonth(date *DateStruct) int8 {
	return int8(date.month)
}

func DateStructGetDay(date *DateStruct) int8 {
	return int8(date.day)
}

func TimeGetMicros(ti *Time) int64 {
	return int64(ti.micros)
}

func TimeSetMicros(ti *Time, micros int64) {
	ti.micros = C.int64_t(micros)
}

func TimeStructGetHour(ti *TimeStruct) int8 {
	return int8(ti.hour)
}

func TimeStructGetMinute(ti *TimeStruct) int8 {
	return int8(ti.min)
}

func TimeStructGetSecond(ti *TimeStruct) int8 {
	return int8(ti.sec)
}

func TimeStructGetMicros(ti *TimeStruct) int32 {
	return int32(ti.micros)
}

func TimeTZStructGetTimeStruct(ti *TimeTZStruct) TimeStruct {
	return TimeStruct(ti.time)
}

func TimeTZStructGetOffset(ti *TimeTZStruct) int32 {
	return int32(ti.offset)
}

func TimestampGetMicros(ts *Timestamp) int64 {
	return int64(ts.micros)
}

func TimestampSetMicros(ts *Timestamp, micros int64) {
	ts.micros = C.int64_t(micros)
}

func IntervalGetMonths(i *Interval) int32 {
	return int32(i.months)
}

func IntervalSetMonths(i *Interval, months int32) {
	i.months = C.int32_t(months)
}

func IntervalGetDays(i *Interval) int32 {
	return int32(i.days)
}

func IntervalSetDays(i *Interval, days int32) {
	i.days = C.int32_t(days)
}

func IntervalGetMicros(i *Interval) int64 {
	return int64(i.micros)
}

func IntervalSetMicros(i *Interval, micros int64) {
	i.micros = C.int64_t(micros)
}

func HugeIntGetLower(hugeInt *HugeInt) uint64 {
	return uint64(hugeInt.lower)
}

func HugeIntSetLower(hugeInt *HugeInt, lower uint64) {
	hugeInt.lower = C.uint64_t(lower)
}

func HugeIntGetUpper(hugeInt *HugeInt) int64 {
	return int64(hugeInt.upper)
}

func HugeIntSetUpper(hugeInt *HugeInt, upper int64) {
	hugeInt.upper = C.int64_t(upper)
}

func ListEntryGetOffset(entry *ListEntry) uint64 {
	return uint64(entry.offset)
}

func ListEntrySetOffset(entry *ListEntry, offset uint64) {
	entry.offset = C.uint64_t(offset)
}

func ListEntryGetLength(entry *ListEntry) uint64 {
	return uint64(entry.length)
}

func ListEntrySetLength(entry *ListEntry, length uint64) {
	entry.length = C.uint64_t(length)
}

// Types with internal pointers:

// duckdb_column

// Result wraps duckdb_result.
type Result struct {
	data C.duckdb_result
}

// ------------------------------------------------------------------ //
// Pointer Types
// ------------------------------------------------------------------ //

// NOTE: No wrappings for function pointers.
// *duckdb_delete_callback_t
// *duckdb_task_state
// *duckdb_scalar_function_t
// *duckdb_aggregate_state_size
// *duckdb_aggregate_init_t
// *duckdb_aggregate_destroy_t
// *duckdb_aggregate_update_t
// *duckdb_aggregate_combine_t
// *duckdb_aggregate_finalize_t
// *duckdb_table_function_bind_t
// *duckdb_table_function_init_t
// *duckdb_table_function_t
// *duckdb_cast_function_t
// *duckdb_replacement_callback_t

// NOTE: We export the Ptr of each wrapped type pointer to allow (void *) typedef's of callback functions.
// NOTE: See https://golang.org/issue/19837 and https://golang.org/issue/19835.

// *duckdb_task_state

// Vector wraps *duckdb_vector.
type Vector struct {
	Ptr unsafe.Pointer
}

func (vec *Vector) data() C.duckdb_vector {
	return C.duckdb_vector(vec.Ptr)
}

// Database wraps *duckdb_database.
type Database struct {
	Ptr unsafe.Pointer
}

func (db *Database) data() C.duckdb_database {
	return C.duckdb_database(db.Ptr)
}

// Connection wraps *duckdb_connection.
type Connection struct {
	Ptr unsafe.Pointer
}

func (conn *Connection) data() C.duckdb_connection {
	return C.duckdb_connection(conn.Ptr)
}

// PreparedStatement wraps *duckdb_prepared_statement.
type PreparedStatement struct {
	Ptr unsafe.Pointer
}

func (preparedStmt *PreparedStatement) data() C.duckdb_prepared_statement {
	return C.duckdb_prepared_statement(preparedStmt.Ptr)
}

// ExtractedStatements wraps *duckdb_extracted_statements.
type ExtractedStatements struct {
	Ptr unsafe.Pointer
}

func (extractedStmts *ExtractedStatements) data() C.duckdb_extracted_statements {
	return C.duckdb_extracted_statements(extractedStmts.Ptr)
}

// PendingResult wraps *duckdb_pending_result.
type PendingResult struct {
	Ptr unsafe.Pointer
}

func (pendingRes *PendingResult) data() C.duckdb_pending_result {
	return C.duckdb_pending_result(pendingRes.Ptr)
}

// Appender wraps *duckdb_appender.
type Appender struct {
	Ptr unsafe.Pointer
}

func (appender *Appender) data() C.duckdb_appender {
	return C.duckdb_appender(appender.Ptr)
}

// TableDescription wraps *duckdb_table_description.
type TableDescription struct {
	Ptr unsafe.Pointer
}

func (description *TableDescription) data() C.duckdb_table_description {
	return C.duckdb_table_description(description.Ptr)
}

// Config wraps *duckdb_config.
type Config struct {
	Ptr unsafe.Pointer
}

func (config *Config) data() C.duckdb_config {
	return C.duckdb_config(config.Ptr)
}

// LogicalType wraps *duckdb_logical_type.
type LogicalType struct {
	Ptr unsafe.Pointer
}

func (logicalType *LogicalType) data() C.duckdb_logical_type {
	return C.duckdb_logical_type(logicalType.Ptr)
}

// *duckdb_create_type_info

// DataChunk wraps *duckdb_data_chunk.
type DataChunk struct {
	Ptr unsafe.Pointer
}

func (chunk *DataChunk) data() C.duckdb_data_chunk {
	return C.duckdb_data_chunk(chunk.Ptr)
}

// Value wraps *duckdb_value.
type Value struct {
	Ptr unsafe.Pointer
}

func (v *Value) data() C.duckdb_value {
	return C.duckdb_value(v.Ptr)
}

// ProfilingInfo wraps *duckdb_profiling_info.
type ProfilingInfo struct {
	Ptr unsafe.Pointer
}

func (info *ProfilingInfo) data() C.duckdb_profiling_info {
	return C.duckdb_profiling_info(info.Ptr)
}

// *duckdb_extension_info

// FunctionInfo wraps *duckdb_function_info.
type FunctionInfo struct {
	Ptr unsafe.Pointer
}

func (info *FunctionInfo) data() C.duckdb_function_info {
	return C.duckdb_function_info(info.Ptr)
}

// ScalarFunction wraps *duckdb_scalar_function.
type ScalarFunction struct {
	Ptr unsafe.Pointer
}

func (f *ScalarFunction) data() C.duckdb_scalar_function {
	return C.duckdb_scalar_function(f.Ptr)
}

// ScalarFunctionSet wraps *duckdb_scalar_function_set.
type ScalarFunctionSet struct {
	Ptr unsafe.Pointer
}

func (set *ScalarFunctionSet) data() C.duckdb_scalar_function_set {
	return C.duckdb_scalar_function_set(set.Ptr)
}

// *duckdb_aggregate_function
// *duckdb_aggregate_function_set
// *duckdb_aggregate_state

// TableFunction wraps *duckdb_table_function.
type TableFunction struct {
	Ptr unsafe.Pointer
}

func (f *TableFunction) data() C.duckdb_table_function {
	return C.duckdb_table_function(f.Ptr)
}

// BindInfo wraps *duckdb_bind_info.
type BindInfo struct {
	Ptr unsafe.Pointer
}

func (info *BindInfo) data() C.duckdb_bind_info {
	return C.duckdb_bind_info(info.Ptr)
}

// InitInfo wraps *C.duckdb_init_info.
type InitInfo struct {
	Ptr unsafe.Pointer
}

func (info *InitInfo) data() C.duckdb_init_info {
	return C.duckdb_init_info(info.Ptr)
}

// *duckdb_cast_function

// ReplacementScanInfo wraps *duckdb_replacement_scan.
type ReplacementScanInfo struct {
	Ptr unsafe.Pointer
}

func (info *ReplacementScanInfo) data() C.duckdb_replacement_scan_info {
	return C.duckdb_replacement_scan_info(info.Ptr)
}

// Arrow wraps *duckdb_arrow.
type Arrow struct {
	Ptr unsafe.Pointer
}

func (arrow *Arrow) data() C.duckdb_arrow {
	return C.duckdb_arrow(arrow.Ptr)
}

// ArrowStream wraps *duckdb_arrow_stream.
type ArrowStream struct {
	Ptr unsafe.Pointer
}

func (stream *ArrowStream) data() C.duckdb_arrow_stream {
	return C.duckdb_arrow_stream(stream.Ptr)
}

// ArrowSchema wraps *duckdb_arrow_schema.
type ArrowSchema struct {
	Ptr unsafe.Pointer
}

func (schema *ArrowSchema) data() C.duckdb_arrow_schema {
	return C.duckdb_arrow_schema(schema.Ptr)
}

// ArrowArray wraps *duckdb_arrow_array.
type ArrowArray struct {
	Ptr unsafe.Pointer
}

func (array *ArrowArray) data() C.duckdb_arrow_array {
	return C.duckdb_arrow_array(array.Ptr)
}
