package duckdb_go_bindings

/*
#include <duckdb.h>
*/
import "C"
import "unsafe"

// ------------------------------------------------------------------ //
// Enums
// ------------------------------------------------------------------ //

// Type wraps duckdb_type.
type Type C.duckdb_type

func (t *Type) data() C.duckdb_type {
	return C.duckdb_type(*t)
}

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
	Success State = C.DuckDBSuccess
	Error   State = C.DuckDBError
)

// PendingState wraps duckdb_pending_state.
type PendingState C.duckdb_pending_state

const (
	PendingResultReady      PendingState = C.DUCKDB_PENDING_RESULT_READY
	PendingResultNotReady   PendingState = C.DUCKDB_PENDING_RESULT_NOT_READY
	PendingError            PendingState = C.DUCKDB_PENDING_ERROR
	PendingNoTasksAvailable PendingState = C.DUCKDB_PENDING_NO_TASKS_AVAILABLE
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
	ErrorInvalid              ErrorType = C.DUCKDB_ERROR_INVALID
	ErrorOutOfRange           ErrorType = C.DUCKDB_ERROR_OUT_OF_RANGE
	ErrorConversion           ErrorType = C.DUCKDB_ERROR_CONVERSION
	ErrorUnknownType          ErrorType = C.DUCKDB_ERROR_UNKNOWN_TYPE
	ErrorDecimal              ErrorType = C.DUCKDB_ERROR_DECIMAL
	ErrorMismatchType         ErrorType = C.DUCKDB_ERROR_MISMATCH_TYPE
	ErrorDivideByZero         ErrorType = C.DUCKDB_ERROR_DIVIDE_BY_ZERO
	ErrorObjectSize           ErrorType = C.DUCKDB_ERROR_OBJECT_SIZE
	ErrorInvalidType          ErrorType = C.DUCKDB_ERROR_INVALID_TYPE
	ErrorSerialization        ErrorType = C.DUCKDB_ERROR_SERIALIZATION
	ErrorTransaction          ErrorType = C.DUCKDB_ERROR_TRANSACTION
	ErrorNotImplemented       ErrorType = C.DUCKDB_ERROR_NOT_IMPLEMENTED
	ErrorExpression           ErrorType = C.DUCKDB_ERROR_EXPRESSION
	ErrorCatalog              ErrorType = C.DUCKDB_ERROR_CATALOG
	ErrorParser               ErrorType = C.DUCKDB_ERROR_PARSER
	ErrorPlanner              ErrorType = C.DUCKDB_ERROR_PLANNER
	ErrorScheduler            ErrorType = C.DUCKDB_ERROR_SCHEDULER
	ErrorExecutor             ErrorType = C.DUCKDB_ERROR_EXECUTOR
	ErrorConstraint           ErrorType = C.DUCKDB_ERROR_CONSTRAINT
	ErrorIndex                ErrorType = C.DUCKDB_ERROR_INDEX
	ErrorStat                 ErrorType = C.DUCKDB_ERROR_STAT
	ErrorConnection           ErrorType = C.DUCKDB_ERROR_CONNECTION
	ErrorSyntax               ErrorType = C.DUCKDB_ERROR_SYNTAX
	ErrorSettings             ErrorType = C.DUCKDB_ERROR_SETTINGS
	ErrorBinder               ErrorType = C.DUCKDB_ERROR_BINDER
	ErrorNetwork              ErrorType = C.DUCKDB_ERROR_NETWORK
	ErrorOptimizer            ErrorType = C.DUCKDB_ERROR_OPTIMIZER
	ErrorNullPointer          ErrorType = C.DUCKDB_ERROR_NULL_POINTER
	ErrorErrorIO              ErrorType = C.DUCKDB_ERROR_IO
	ErrorInterrupt            ErrorType = C.DUCKDB_ERROR_INTERRUPT
	ErrorFatal                ErrorType = C.DUCKDB_ERROR_FATAL
	ErrorInternal             ErrorType = C.DUCKDB_ERROR_INTERNAL
	ErrorInvalidInput         ErrorType = C.DUCKDB_ERROR_INVALID_INPUT
	ErrorOutOfMemory          ErrorType = C.DUCKDB_ERROR_OUT_OF_MEMORY
	ErrorPermission           ErrorType = C.DUCKDB_ERROR_PERMISSION
	ErrorParameterNotResolved ErrorType = C.DUCKDB_ERROR_PARAMETER_NOT_RESOLVED
	ErrorParameterNotAllowed  ErrorType = C.DUCKDB_ERROR_PARAMETER_NOT_ALLOWED
	ErrorDependency           ErrorType = C.DUCKDB_ERROR_DEPENDENCY
	ErrorHTTP                 ErrorType = C.DUCKDB_ERROR_HTTP
	ErrorMissingExtension     ErrorType = C.DUCKDB_ERROR_MISSING_EXTENSION
	ErrorAutoload             ErrorType = C.DUCKDB_ERROR_AUTOLOAD
	ErrorSequence             ErrorType = C.DUCKDB_ERROR_SEQUENCE
	ErrorInvalidConfiguration ErrorType = C.DUCKDB_INVALID_CONFIGURATION
)

// CastMode wraps duckdb_cast_mode.
type CastMode C.duckdb_cast_mode

const (
	CastNormal CastMode = C.DUCKDB_CAST_NORMAL
	CastTry    CastMode = C.DUCKDB_CAST_TRY
)

// ------------------------------------------------------------------ //
// Types
// ------------------------------------------------------------------ //

// NOTE: No wrapping for C.idx_t.

// *duckdb_delete_callback_t
// *duckdb_task_state

// Date wraps duckdb_date.
type Date struct {
	Days int32
}

func (date *Date) data() C.duckdb_date {
	var result C.duckdb_date
	result.days = C.int32_t(date.Days)
	return result
}

// DateStruct wraps duckdb_date_struct.
type DateStruct struct {
	Year  int32
	Month int8
	Day   int8
}

// Time wraps duckdb_time.
type Time struct {
	Micros int64
}

func (t *Time) data() C.duckdb_time {
	var result C.duckdb_time
	result.micros = C.int64_t(t.Micros)
	return result
}

// TimeStruct wraps duckdb_time_struct.
type TimeStruct struct {
	Hour   int8
	Min    int8
	Sec    int8
	Micros int32
}

// TimeTZ wraps duckdb_time_tz.
type TimeTZ struct {
	Bits uint64
}

func (t *TimeTZ) data() C.duckdb_time_tz {
	var result C.duckdb_time_tz
	result.bits = C.uint64_t(t.Bits)
	return result
}

// TimeTZStruct wraps duckdb_time_tz_struct.
type TimeTZStruct struct {
	Time   TimeStruct
	Offset int32
}

// Timestamp wraps duckdb_timestamp.
type Timestamp struct {
	Micros int64
}

func (ts *Timestamp) data() C.duckdb_timestamp {
	var result C.duckdb_timestamp
	result.micros = C.int64_t(ts.Micros)
	return result
}

// TimestampStruct wraps duckdb_timestamp_struct.
type TimestampStruct struct {
	Date DateStruct
	Time TimeStruct
}

// Interval wraps duckdb_interval.
type Interval struct {
	Months int32
	Days   int32
	Micros int64
}

func (interval *Interval) data() C.duckdb_interval {
	var result C.duckdb_interval
	result.months = C.int32_t(interval.Months)
	result.days = C.int32_t(interval.Days)
	result.micros = C.int64_t(interval.Micros)
	return result
}

// HugeInt wraps duckdb_hugeint.
type HugeInt struct {
	Lower uint64
	Upper int64
}

func (hugeInt *HugeInt) data() C.duckdb_hugeint {
	var result C.duckdb_hugeint
	result.lower = C.uint64_t(hugeInt.Lower)
	result.upper = C.int64_t(hugeInt.Upper)
	return result
}

// UHugeInt wraps duckdb_uhugeint.
type UHugeInt struct {
	Lower uint64
	Upper uint64
}

// Decimal wraps duckdb_decimal
type Decimal struct {
	Width uint8
	Scale uint8
	Value HugeInt
}

func (decimal *Decimal) data() C.duckdb_decimal {
	var result C.duckdb_decimal
	result.width = C.uint8_t(decimal.Width)
	result.scale = C.uint8_t(decimal.Scale)
	result.value = decimal.Value.data()
	return result
}

// QueryProgressType wraps duckdb_query_progress_type.
type QueryProgressType struct {
	Percentage         float64
	RowsProcessed      uint64
	TotalRowsToProcess uint64
}

// duckdb_string_t

// ListEntry wraps duckdb_list_entry.
type ListEntry struct {
	Offset uint64
	Length uint64
}

// Column wraps duckdb_column.
type Column struct {
	data C.duckdb_column
}

// ------------------------------------------------------------------ //
// Pointers
// ------------------------------------------------------------------ //

// Vector wraps *duckdb_vector.
type Vector struct {
	Ptr unsafe.Pointer
}

func (vector *Vector) data() C.duckdb_vector {
	return C.duckdb_vector(vector.Ptr)
}

// duckdb_string
// duckdb_blob

// Result wraps duckdb_result.
type Result struct {
	data C.duckdb_result
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

// TODO: Do we need *duckdb_extension_info?

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

// *duckdb_scalar_function_t

// *duckdb_aggregate_function
// *duckdb_aggregate_function_set
// *duckdb_aggregate_state

// *duckdb_aggregate_state_size
// *duckdb_aggregate_init_t
// *duckdb_aggregate_destroy_t
// *duckdb_aggregate_update_t
// *duckdb_aggregate_combine_t
// *duckdb_aggregate_finalize_t

// *duckdb_table_function

// BindInfo wraps *duckdb_bind_info.
type BindInfo struct {
	Ptr unsafe.Pointer
}

func (info *BindInfo) data() C.duckdb_bind_info {
	return C.duckdb_bind_info(info.Ptr)
}

// *duckdb_init_info

// *duckdb_table_function_bind_t
// *duckdb_table_function_init_t
// *duckdb_table_function_t

// *duckdb_cast_function

// *duckdb_cast_function_t

// *duckdb_replacement_scan_info

// *duckdb_replacement_callback_t

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

// TODO: What about duckdb_extension_access?

// ------------------------------------------------------------------ //
// Functions
// ------------------------------------------------------------------ //

// ------------------------------------------------------------------ //
// Open Connect
// ------------------------------------------------------------------ //

// duckdb_open

func OpenExt(path string, outDb *Database, config Config, errMsg *string) State {
	cPath := C.CString(path)
	defer Free(unsafe.Pointer(cPath))

	var err *C.char
	defer Free(unsafe.Pointer(err))

	var db C.duckdb_database
	state := C.duckdb_open_ext(cPath, &db, config.data(), &err)
	outDb.Ptr = unsafe.Pointer(db)
	*errMsg = C.GoString(err)
	return State(state)
}

func Close(db *Database) {
	data := db.data()
	C.duckdb_close(&data)
	db.Ptr = nil
}

func Connect(db Database, outConn *Connection) State {
	var conn C.duckdb_connection
	state := C.duckdb_connect(db.data(), &conn)
	outConn.Ptr = unsafe.Pointer(conn)
	return State(state)
}

// duckdb_interrupt
// duckdb_query_progress

func Disconnect(conn *Connection) {
	data := conn.data()
	C.duckdb_disconnect(&data)
	conn.Ptr = nil
}

// duckdb_library_version

// ------------------------------------------------------------------ //
// Configuration
// ------------------------------------------------------------------ //

func CreateConfig(outConfig *Config) State {
	var config C.duckdb_config
	state := C.duckdb_create_config(&config)
	outConfig.Ptr = unsafe.Pointer(config)
	return State(state)
}

// duckdb_config_count
// duckdb_get_config_flag

func SetConfig(config Config, name string, option string) State {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))

	cOption := C.CString(option)
	defer Free(unsafe.Pointer(cOption))

	state := C.duckdb_set_config(config.data(), cName, cOption)
	return State(state)
}

func DestroyConfig(config *Config) {
	data := config.data()
	C.duckdb_destroy_config(&data)
	config.Ptr = nil
}

// ------------------------------------------------------------------ //
// Query Execution
// ------------------------------------------------------------------ //

// duckdb_query

func DestroyResult(res *Result) {
	C.duckdb_destroy_result(&res.data)
}

// duckdb_column_name
// duckdb_column_type
// duckdb_result_statement_type
// duckdb_column_logical_type
// duckdb_column_count
// duckdb_rows_changed
// duckdb_result_error
// duckdb_result_error_type

// ------------------------------------------------------------------ //
// Result Functions (many are deprecated)
// ------------------------------------------------------------------ //

// duckdb_result_return_type

// ------------------------------------------------------------------ //
// Safe Fetch Functions (all deprecated)
// ------------------------------------------------------------------ //

// ------------------------------------------------------------------ //
// Helpers
// ------------------------------------------------------------------ //

// duckdb_malloc

func Free(ptr unsafe.Pointer) {
	C.duckdb_free(ptr)
}

// duckdb_vector_size
// duckdb_string_is_inlined
// duckdb_string_t_length
// duckdb_string_t_data

// ------------------------------------------------------------------ //
// Date Time Timestamp Helpers
// ------------------------------------------------------------------ //

// duckdb_from_date
// duckdb_to_date
// duckdb_is_finite_date
// duckdb_from_time

func CreateTimeTZ(micros int64, offset int32) TimeTZ {
	timeTZ := C.duckdb_create_time_tz(C.int64_t(micros), C.int32_t(offset))
	return TimeTZ{
		Bits: uint64(timeTZ.bits),
	}
}

// duckdb_from_time_tz
// duckdb_to_time
// duckdb_from_timestamp
// duckdb_to_timestamp
// duckdb_is_finite_timestamp
// duckdb_is_finite_timestamp_s
// duckdb_is_finite_timestamp_ms
// duckdb_is_finite_timestamp_ns

// ------------------------------------------------------------------ //
// Hugeint Helpers
// ------------------------------------------------------------------ //

// duckdb_hugeint_to_double
// duckdb_double_to_hugeint

// ------------------------------------------------------------------ //
// Unsigned Hugeint Helpers
// ------------------------------------------------------------------ //

// duckdb_uhugeint_to_double
// duckdb_double_to_uhugeint

// ------------------------------------------------------------------ //
// Decimal Helpers
// ------------------------------------------------------------------ //

// duckdb_double_to_decimal
// duckdb_decimal_to_double

// ------------------------------------------------------------------ //
// Prepared Statements
// ------------------------------------------------------------------ //

// duckdb_prepare

func DestroyPrepare(preparedStmt *PreparedStatement) {
	data := preparedStmt.data()
	C.duckdb_destroy_prepare(&data)
	preparedStmt.Ptr = nil
}

func PrepareError(preparedStmt PreparedStatement) string {
	err := C.duckdb_prepare_error(preparedStmt.data())
	return C.GoString(err)
}

func NParams(preparedStmt PreparedStatement) uint64 {
	count := C.duckdb_nparams(preparedStmt.data())
	return uint64(count)
}

func ParameterName(preparedStmt PreparedStatement, index uint64) string {
	cName := C.duckdb_parameter_name(preparedStmt.data(), C.idx_t(index))
	defer Free(unsafe.Pointer(cName))
	return C.GoString(cName)
}

func ParamType(preparedStmt PreparedStatement, index uint64) Type {
	t := C.duckdb_param_type(preparedStmt.data(), C.idx_t(index))
	return Type(t)
}

// duckdb_param_logical_type
// duckdb_clear_bindings

func PreparedStatementType(preparedStmt PreparedStatement) StatementType {
	t := C.duckdb_prepared_statement_type(preparedStmt.data())
	return StatementType(t)
}

// ------------------------------------------------------------------ //
// Bind Values To Prepared Statements
// ------------------------------------------------------------------ //

func BindValue(preparedStmt PreparedStatement, index uint64, v Value) State {
	state := C.duckdb_bind_value(preparedStmt.data(), C.idx_t(index), v.data())
	return State(state)
}

// duckdb_bind_parameter_index

func BindBoolean(preparedStmt PreparedStatement, index uint64, v bool) State {
	state := C.duckdb_bind_boolean(preparedStmt.data(), C.idx_t(index), C.bool(v))
	return State(state)
}

func BindInt8(preparedStmt PreparedStatement, index uint64, v int8) State {
	state := C.duckdb_bind_int8(preparedStmt.data(), C.idx_t(index), C.int8_t(v))
	return State(state)
}

func BindInt16(preparedStmt PreparedStatement, index uint64, v int16) State {
	state := C.duckdb_bind_int16(preparedStmt.data(), C.idx_t(index), C.int16_t(v))
	return State(state)
}

func BindInt32(preparedStmt PreparedStatement, index uint64, v int32) State {
	state := C.duckdb_bind_int32(preparedStmt.data(), C.idx_t(index), C.int32_t(v))
	return State(state)
}

func BindInt64(preparedStmt PreparedStatement, index uint64, v int64) State {
	state := C.duckdb_bind_int64(preparedStmt.data(), C.idx_t(index), C.int64_t(v))
	return State(state)
}

func BindHugeInt(preparedStmt PreparedStatement, index uint64, v HugeInt) State {
	state := C.duckdb_bind_hugeint(preparedStmt.data(), C.idx_t(index), v.data())
	return State(state)
}

// duckdb_bind_uhugeint

func BindDecimal(preparedStmt PreparedStatement, index uint64, v Decimal) State {
	state := C.duckdb_bind_decimal(preparedStmt.data(), C.idx_t(index), v.data())
	return State(state)
}

func BindUInt8(preparedStmt PreparedStatement, index uint64, v uint8) State {
	state := C.duckdb_bind_uint8(preparedStmt.data(), C.idx_t(index), C.uint8_t(v))
	return State(state)
}

func BindUInt16(preparedStmt PreparedStatement, index uint64, v uint16) State {
	state := C.duckdb_bind_uint16(preparedStmt.data(), C.idx_t(index), C.uint16_t(v))
	return State(state)
}

func BindUInt32(preparedStmt PreparedStatement, index uint64, v uint32) State {
	state := C.duckdb_bind_uint32(preparedStmt.data(), C.idx_t(index), C.uint32_t(v))
	return State(state)
}

func BindUInt64(preparedStmt PreparedStatement, index uint64, v uint64) State {
	state := C.duckdb_bind_uint64(preparedStmt.data(), C.idx_t(index), C.uint64_t(v))
	return State(state)
}

func BindFloat(preparedStmt PreparedStatement, index uint64, v float32) State {
	state := C.duckdb_bind_float(preparedStmt.data(), C.idx_t(index), C.float(v))
	return State(state)
}

func BindDouble(preparedStmt PreparedStatement, index uint64, v float64) State {
	state := C.duckdb_bind_double(preparedStmt.data(), C.idx_t(index), C.double(v))
	return State(state)
}

func BindDate(preparedStmt PreparedStatement, index uint64, v Date) State {
	state := C.duckdb_bind_date(preparedStmt.data(), C.idx_t(index), v.data())
	return State(state)
}

func BindTime(preparedStmt PreparedStatement, index uint64, v Time) State {
	state := C.duckdb_bind_time(preparedStmt.data(), C.idx_t(index), v.data())
	return State(state)
}

func BindTimestamp(preparedStmt PreparedStatement, index uint64, v Timestamp) State {
	state := C.duckdb_bind_timestamp(preparedStmt.data(), C.idx_t(index), v.data())
	return State(state)
}

// duckdb_bind_timestamp_tz

func BindInterval(preparedStmt PreparedStatement, index uint64, v Interval) State {
	state := C.duckdb_bind_interval(preparedStmt.data(), C.idx_t(index), v.data())
	return State(state)
}

func BindVarchar(preparedStmt PreparedStatement, index uint64, v string) State {
	cStr := C.CString(v)
	defer Free(unsafe.Pointer(cStr))

	state := C.duckdb_bind_varchar(preparedStmt.data(), C.idx_t(index), cStr)
	return State(state)
}

// duckdb_bind_varchar_length

func BindBlob(preparedStmt PreparedStatement, index uint64, v []byte) State {
	cBytes := C.CBytes(v)
	defer Free(unsafe.Pointer(cBytes))

	state := C.duckdb_bind_blob(preparedStmt.data(), C.idx_t(index), cBytes, C.idx_t(len(v)))
	return State(state)
}

func BindNull(preparedStmt PreparedStatement, index uint64) State {
	state := C.duckdb_bind_null(preparedStmt.data(), C.idx_t(index))
	return State(state)
}

// ------------------------------------------------------------------ //
// Execute Prepared Statements (many are deprecated)
// ------------------------------------------------------------------ //

// duckdb_execute_prepared

// ------------------------------------------------------------------ //
// Extract Statements
// ------------------------------------------------------------------ //

func ExtractStatements(conn Connection, query string, outExtractedStmts *ExtractedStatements) uint64 {
	cQuery := C.CString(query)
	defer Free(unsafe.Pointer(cQuery))

	var extractedStmts C.duckdb_extracted_statements
	count := C.duckdb_extract_statements(conn.data(), cQuery, &extractedStmts)
	outExtractedStmts.Ptr = unsafe.Pointer(extractedStmts)
	return uint64(count)
}

func PrepareExtractedStatement(conn Connection, extractedStmts ExtractedStatements, index uint64, outPreparedStmt *PreparedStatement) State {
	var preparedStmt C.duckdb_prepared_statement
	state := C.duckdb_prepare_extracted_statement(conn.data(), extractedStmts.data(), C.idx_t(index), &preparedStmt)
	outPreparedStmt.Ptr = unsafe.Pointer(preparedStmt)
	return State(state)
}

func ExtractStatementsError(extractedStmts ExtractedStatements) string {
	err := C.duckdb_extract_statements_error(extractedStmts.data())
	return C.GoString(err)
}

func DestroyExtracted(extractedStmts *ExtractedStatements) {
	data := extractedStmts.data()
	C.duckdb_destroy_extracted(&data)
	extractedStmts.Ptr = nil
}

// ------------------------------------------------------------------ //
// Pending Result Interface
// ------------------------------------------------------------------ //

// duckdb_pending_prepared
// duckdb_destroy_pending
// duckdb_pending_error
// duckdb_pending_execute_task
// duckdb_pending_execute_check_state
// duckdb_execute_pending
// duckdb_pending_execution_is_finished

// ------------------------------------------------------------------ //
// Value Interface
// ------------------------------------------------------------------ //

func DestroyValue(v *Value) {
	data := v.data()
	C.duckdb_destroy_value(&data)
	v.Ptr = nil
}

// duckdb_create_varchar
// duckdb_create_varchar_length
// duckdb_create_bool
// duckdb_create_int8
// duckdb_create_uint8
// duckdb_create_int16
// duckdb_create_uint16
// duckdb_create_int32
// duckdb_create_uint32
// duckdb_create_uint64
// duckdb_create_int64
// duckdb_create_hugeint
// duckdb_create_uhugeint
// duckdb_create_varint
// duckdb_create_decimal
// duckdb_create_float
// duckdb_create_double
// duckdb_create_date
// duckdb_create_time

func CreateTimeTZValue(timeTZ TimeTZ) Value {
	v := C.duckdb_create_time_tz_value(timeTZ.data())
	return Value{
		Ptr: unsafe.Pointer(v),
	}
}

// duckdb_create_timestamp
// duckdb_create_timestamp_tz
// duckdb_create_timestamp_s
// duckdb_create_timestamp_ms
// duckdb_create_timestamp_ns
// duckdb_create_interval
// duckdb_create_blob
// duckdb_create_bit
// duckdb_create_uuid
// duckdb_get_bool
// duckdb_get_int8
// duckdb_get_uint8
// duckdb_get_int16
// duckdb_get_uint16
// duckdb_get_int32
// duckdb_get_uint32
// duckdb_get_int64
// duckdb_get_uint64
// duckdb_get_hugeint
// duckdb_get_uhugeint
// duckdb_get_varint
// duckdb_get_decimal
// duckdb_get_float
// duckdb_get_double
// duckdb_get_date
// duckdb_get_time
// duckdb_get_time_tz
// duckdb_get_timestamp
// duckdb_get_timestamp_tz
// duckdb_get_timestamp_s
// duckdb_get_timestamp_ms
// duckdb_get_timestamp_ns
// duckdb_get_interval
// duckdb_get_value_type
// duckdb_get_blob
// duckdb_get_bit
// duckdb_get_uuid

func GetVarchar(v Value) string {
	cStr := C.duckdb_get_varchar(v.data())
	defer Free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

// duckdb_create_struct_value
// duckdb_create_list_value
// duckdb_create_array_value

func GetMapSize(v Value) uint64 {
	size := C.duckdb_get_map_size(v.data())
	return uint64(size)
}

func GetMapKey(v Value, index uint64) Value {
	value := C.duckdb_get_map_key(v.data(), C.idx_t(index))
	return Value{
		Ptr: unsafe.Pointer(value),
	}
}

func GetMapValue(v Value, index uint64) Value {
	value := C.duckdb_get_map_value(v.data(), C.idx_t(index))
	return Value{
		Ptr: unsafe.Pointer(value),
	}
}

// duckdb_is_null_value
// duckdb_create_null_value
// duckdb_get_list_size
// duckdb_get_list_child
// duckdb_create_enum_value
// duckdb_get_enum_value
// duckdb_get_struct_child

// ------------------------------------------------------------------ //
// Logical Type Interface
// ------------------------------------------------------------------ //

func CreateLogicalType(t Type) LogicalType {
	logicalType := C.duckdb_create_logical_type(t.data())
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

// duckdb_logical_type_get_alias
// duckdb_logical_type_set_alias

func CreateListType(child LogicalType) LogicalType {
	logicalType := C.duckdb_create_list_type(child.data())
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func CreateArrayType(child LogicalType, size uint64) LogicalType {
	logicalType := C.duckdb_create_array_type(child.data(), C.idx_t(size))
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func CreateMapType(key LogicalType, value LogicalType) LogicalType {
	logicalType := C.duckdb_create_map_type(key.data(), value.data())
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

// duckdb_create_union_type

func CreateStructType(types []LogicalType, names []string) LogicalType {
	// FIXME: Unify code with other slice allocation functions.

	count := len(types)
	size := C.size_t(unsafe.Sizeof(C.duckdb_logical_type(nil)))
	typesSlice := (*[1 << 31]C.duckdb_logical_type)(C.malloc(C.size_t(count) * size))

	size = C.size_t(unsafe.Sizeof((*C.char)(nil)))
	namesSlice := (*[1 << 31]*C.char)(C.malloc(C.size_t(count) * size))

	for i, t := range types {
		(*typesSlice)[i] = t.data()
	}
	for i, name := range names {
		(*namesSlice)[i] = C.CString(name)
	}

	cTypes := (*C.duckdb_logical_type)(unsafe.Pointer(typesSlice))
	cNames := (**C.char)(unsafe.Pointer(namesSlice))
	logicalType := C.duckdb_create_struct_type(cTypes, cNames, C.idx_t(count))

	for i := 0; i < count; i++ {
		C.duckdb_destroy_logical_type(&typesSlice[i])
		Free(unsafe.Pointer((*namesSlice)[i]))
	}
	Free(unsafe.Pointer(typesSlice))
	Free(unsafe.Pointer(namesSlice))

	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func CreateEnumType(names []string) LogicalType {
	// FIXME: Unify code with other slice allocation functions.

	count := len(names)
	size := C.size_t(unsafe.Sizeof((*C.char)(nil)))
	namesSlice := (*[1 << 31]*C.char)(C.malloc(C.size_t(count) * size))

	for i, name := range names {
		(*namesSlice)[i] = C.CString(name)
	}
	cNames := (**C.char)(unsafe.Pointer(namesSlice))
	logicalType := C.duckdb_create_enum_type(cNames, C.idx_t(count))

	for i := 0; i < count; i++ {
		Free(unsafe.Pointer((*namesSlice)[i]))
	}
	Free(unsafe.Pointer(namesSlice))

	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func CreateDecimalType(width uint8, scale uint8) LogicalType {
	logicalType := C.duckdb_create_decimal_type(C.uint8_t(width), C.uint8_t(scale))
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func GetTypeId(logicalType LogicalType) Type {
	t := C.duckdb_get_type_id(logicalType.data())
	return Type(t)
}

// duckdb_decimal_width
// duckdb_decimal_scale
// duckdb_decimal_internal_type
// duckdb_enum_internal_type
// duckdb_enum_dictionary_size
// duckdb_enum_dictionary_value
// duckdb_list_type_child_type
// duckdb_array_type_child_type
// duckdb_array_type_array_size
// duckdb_map_type_key_type
// duckdb_map_type_value_type
// duckdb_struct_type_child_count
// duckdb_struct_type_child_name
// duckdb_struct_type_child_type
// duckdb_union_type_member_count
// duckdb_union_type_member_name
// duckdb_union_type_member_type

func DestroyLogicalType(logicalType *LogicalType) {
	data := logicalType.data()
	C.duckdb_destroy_logical_type(&data)
	logicalType.Ptr = nil
}

// duckdb_register_logical_type

// ------------------------------------------------------------------ //
// Data Chunk Interface
// ------------------------------------------------------------------ //

// duckdb_create_data_chunk
// duckdb_destroy_data_chunk
// duckdb_data_chunk_reset
// duckdb_data_chunk_get_column_count
// duckdb_data_chunk_get_vector
// duckdb_data_chunk_get_size
// duckdb_data_chunk_set_size

// ------------------------------------------------------------------ //
// Vector Interface
// ------------------------------------------------------------------ //

// duckdb_vector_get_column_type
// duckdb_vector_get_data
// duckdb_vector_get_validity
// duckdb_vector_ensure_validity_writable
// duckdb_vector_assign_string_element
// duckdb_vector_assign_string_element_len
// duckdb_list_vector_get_child
// duckdb_list_vector_get_size
// duckdb_list_vector_set_size
// duckdb_list_vector_reserve
// duckdb_struct_vector_get_child
// duckdb_array_vector_get_child

// ------------------------------------------------------------------ //
// Validity Mask Functions
// ------------------------------------------------------------------ //

// duckdb_validity_row_is_valid
// duckdb_validity_set_row_validity
// duckdb_validity_set_row_invalid
// duckdb_validity_set_row_valid

// ------------------------------------------------------------------ //
// Scalar Functions
// ------------------------------------------------------------------ //

func CreateScalarFunction() ScalarFunction {
	f := C.duckdb_create_scalar_function()
	return ScalarFunction{
		Ptr: unsafe.Pointer(f),
	}
}

func DestroyScalarFunction(f *ScalarFunction) {
	data := f.data()
	C.duckdb_destroy_scalar_function(&data)
	f.Ptr = nil
}

func ScalarFunctionSetName(f ScalarFunction, name string) {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))
	C.duckdb_scalar_function_set_name(f.data(), cName)
}

func ScalarFunctionSetVarargs(f ScalarFunction, logicalType LogicalType) {
	C.duckdb_scalar_function_set_varargs(f.data(), logicalType.data())
}

func ScalarFunctionSetSpecialHandling(f ScalarFunction) {
	C.duckdb_scalar_function_set_special_handling(f.data())
}

func ScalarFunctionSetVolatile(f ScalarFunction) {
	C.duckdb_scalar_function_set_volatile(f.data())
}

func ScalarFunctionAddParameter(f ScalarFunction, logicalType LogicalType) {
	C.duckdb_scalar_function_add_parameter(f.data(), logicalType.data())
}

func ScalarFunctionSetReturnType(f ScalarFunction, logicalType LogicalType) {
	C.duckdb_scalar_function_set_return_type(f.data(), logicalType.data())
}

func ScalarFunctionSetExtraInfo(f ScalarFunction, extraInfoPtr unsafe.Pointer, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_delete_callback_t(callbackPtr)
	C.duckdb_scalar_function_set_extra_info(f.data(), extraInfoPtr, callback)
}

func ScalarFunctionSetFunction(f ScalarFunction, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_scalar_function_t(callbackPtr)
	C.duckdb_scalar_function_set_function(f.data(), callback)
}

func RegisterScalarFunction(conn Connection, f ScalarFunction) State {
	state := C.duckdb_register_scalar_function(conn.data(), f.data())
	return State(state)
}

func ScalarFunctionGetExtraInfo(info FunctionInfo) unsafe.Pointer {
	ptr := C.duckdb_scalar_function_get_extra_info(info.data())
	return unsafe.Pointer(ptr)
}

func ScalarFunctionSetError(info FunctionInfo, err string) {
	cErr := C.CString(err)
	defer Free(unsafe.Pointer(cErr))
	C.duckdb_scalar_function_set_error(info.data(), cErr)
}

func CreateScalarFunctionSet(name string) ScalarFunctionSet {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))

	set := C.duckdb_create_scalar_function_set(cName)
	return ScalarFunctionSet{
		Ptr: unsafe.Pointer(set),
	}
}

func DestroyScalarFunctionSet(set *ScalarFunctionSet) {
	data := set.data()
	C.duckdb_destroy_scalar_function_set(&data)
	set.Ptr = nil
}

func AddScalarFunctionToSet(set ScalarFunctionSet, f ScalarFunction) State {
	state := C.duckdb_add_scalar_function_to_set(set.data(), f.data())
	return State(state)
}

func RegisterScalarFunctionSet(conn Connection, f ScalarFunctionSet) State {
	state := C.duckdb_register_scalar_function_set(conn.data(), f.data())
	return State(state)
}

// ------------------------------------------------------------------ //
// Aggregate Functions
// ------------------------------------------------------------------ //

// duckdb_create_aggregate_function
// duckdb_destroy_aggregate_function
// duckdb_aggregate_function_set_name
// duckdb_aggregate_function_add_parameter
// duckdb_aggregate_function_set_return_type
// duckdb_aggregate_function_set_functions
// duckdb_aggregate_function_set_destructor
// duckdb_register_aggregate_function
// duckdb_aggregate_function_set_special_handling
// duckdb_aggregate_function_set_extra_info
// duckdb_aggregate_function_get_extra_info
// duckdb_aggregate_function_set_error
// duckdb_create_aggregate_function_set
// duckdb_destroy_aggregate_function_set
// duckdb_add_aggregate_function_to_set
// duckdb_register_aggregate_function_set

// ------------------------------------------------------------------ //
// Table Functions
// ------------------------------------------------------------------ //

// duckdb_create_table_function
// duckdb_destroy_table_function
// duckdb_table_function_set_name
// duckdb_table_function_add_parameter
// duckdb_table_function_add_named_parameter
// duckdb_table_function_set_extra_info
// duckdb_table_function_set_bind
// duckdb_table_function_set_init
// duckdb_table_function_set_local_init
// duckdb_table_function_set_function
// duckdb_table_function_supports_projection_pushdown
// duckdb_register_table_function

// ------------------------------------------------------------------ //
// Table Function Bind
// ------------------------------------------------------------------ //

// duckdb_bind_get_extra_info
// duckdb_bind_add_result_column
// duckdb_bind_get_parameter_count
// duckdb_bind_get_parameter
// duckdb_bind_get_named_parameter
// duckdb_bind_set_bind_data
// duckdb_bind_set_cardinality

func BindSetError(info BindInfo, err string) {
	cErr := C.CString(err)
	defer Free(unsafe.Pointer(cErr))
	C.duckdb_bind_set_error(info.data(), cErr)
}

// ------------------------------------------------------------------ //
// Table Function Init
// ------------------------------------------------------------------ //

// duckdb_init_get_extra_info
// duckdb_init_get_bind_data
// duckdb_init_set_init_data
// duckdb_init_get_column_count
// duckdb_init_get_column_index
// duckdb_init_set_max_threads
// duckdb_init_set_error

// ------------------------------------------------------------------ //
// Table Function
// ------------------------------------------------------------------ //

// duckdb_function_get_extra_info
// duckdb_function_get_bind_data
// duckdb_function_get_init_data
// duckdb_function_get_local_init_data
// duckdb_function_set_error

// ------------------------------------------------------------------ //
// Replacement Scans
// ------------------------------------------------------------------ //

// duckdb_add_replacement_scan
// duckdb_replacement_scan_set_function_name
// duckdb_replacement_scan_add_parameter
// duckdb_replacement_scan_set_error

// ------------------------------------------------------------------ //
// Profiling Info
// ------------------------------------------------------------------ //

func GetProfilingInfo(conn Connection) ProfilingInfo {
	info := C.duckdb_get_profiling_info(conn.data())
	return ProfilingInfo{
		Ptr: unsafe.Pointer(info),
	}
}

// duckdb_profiling_info_get_value

func ProfilingInfoGetMetrics(info ProfilingInfo) Value {
	value := C.duckdb_profiling_info_get_metrics(info.data())
	return Value{
		Ptr: unsafe.Pointer(value),
	}
}

func ProfilingInfoGetChildCount(info ProfilingInfo) uint64 {
	count := C.duckdb_profiling_info_get_child_count(info.data())
	return uint64(count)
}

func ProfilingInfoGetChild(info ProfilingInfo, index uint64) ProfilingInfo {
	child := C.duckdb_profiling_info_get_child(info.data(), C.idx_t(index))
	return ProfilingInfo{
		Ptr: unsafe.Pointer(child),
	}
}

// ------------------------------------------------------------------ //
// Appender
// ------------------------------------------------------------------ //

func AppenderCreate(conn Connection, schema string, table string, outAppender *Appender) State {
	cSchema := C.CString(schema)
	defer Free(unsafe.Pointer(cSchema))

	cTable := C.CString(table)
	defer Free(unsafe.Pointer(cTable))

	var appender C.duckdb_appender
	state := C.duckdb_appender_create(conn.data(), cSchema, cTable, &appender)
	outAppender.Ptr = unsafe.Pointer(appender)
	return State(state)
}

// duckdb_appender_create_ext

func AppenderColumnCount(appender Appender) uint64 {
	count := C.duckdb_appender_column_count(appender.data())
	return uint64(count)
}

func AppenderColumnType(appender Appender, index uint64) LogicalType {
	logicalType := C.duckdb_appender_column_type(appender.data(), C.idx_t(index))
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func AppenderError(appender Appender) string {
	err := C.duckdb_appender_error(appender.data())
	return C.GoString(err)
}

func AppenderFlush(appender Appender) State {
	state := C.duckdb_appender_flush(appender.data())
	return State(state)
}

func AppenderClose(appender Appender) State {
	state := C.duckdb_appender_close(appender.data())
	return State(state)
}

func AppenderDestroy(appender *Appender) State {
	data := appender.data()
	state := C.duckdb_appender_destroy(&data)
	appender.Ptr = nil
	return State(state)
}

// duckdb_appender_add_column
// duckdb_appender_clear_columns
// duckdb_appender_begin_row
// duckdb_appender_end_row
// duckdb_append_default
// duckdb_append_bool
// duckdb_append_int8
// duckdb_append_int16
// duckdb_append_int32
// duckdb_append_int64
// duckdb_append_hugeint
// duckdb_append_uint8
// duckdb_append_uint16
// duckdb_append_uint32
// duckdb_append_uint64
// duckdb_append_uhugeint
// duckdb_append_float
// duckdb_append_double
// duckdb_append_date
// duckdb_append_time
// duckdb_append_timestamp
// duckdb_append_interval
// duckdb_append_varchar
// duckdb_append_varchar_length
// duckdb_append_blob
// duckdb_append_null
// duckdb_append_value

func AppendDataChunk(appender Appender, chunk DataChunk) State {
	state := C.duckdb_append_data_chunk(appender.data(), chunk.data())
	return State(state)
}

// ------------------------------------------------------------------ //
// Table Description
// ------------------------------------------------------------------ //

// duckdb_table_description_create
// duckdb_table_description_create_ext
// duckdb_table_description_destroy
// duckdb_table_description_error
// duckdb_column_has_default
// duckdb_table_description_get_column_name
// duckdb_execute_tasks
// duckdb_create_task_state
// duckdb_execute_tasks_state
// duckdb_execute_n_tasks_state
// duckdb_finish_execution
// duckdb_task_state_is_finished
// duckdb_destroy_task_state
// duckdb_execution_is_finished
// duckdb_fetch_chunk
// duckdb_create_cast_function
// duckdb_cast_function_set_source_type
// duckdb_cast_function_set_target_type
// duckdb_cast_function_set_implicit_cast_cost
// duckdb_cast_function_set_function
// duckdb_cast_function_set_extra_info
// duckdb_cast_function_get_extra_info
// duckdb_cast_function_get_cast_mode
// duckdb_cast_function_set_error
// duckdb_cast_function_set_row_error
// duckdb_register_cast_function
// duckdb_destroy_cast_function

// duckdb_row_count
// duckdb_column_data
// duckdb_nullmask_data
// duckdb_result_get_chunk
// duckdb_result_is_streaming
// duckdb_result_chunk_count
// duckdb_value_boolean
// duckdb_value_int8
// duckdb_value_int16
// duckdb_value_int32

func ValueInt64(res *Result, col uint64, row uint64) int64 {
	v := C.duckdb_value_int64(&res.data, C.idx_t(col), C.idx_t(row))
	return int64(v)
}

// duckdb_value_hugeint
// duckdb_value_uhugeint
// duckdb_value_decimal
// duckdb_value_uint8
// duckdb_value_uint16
// duckdb_value_uint32
// duckdb_value_uint64
// duckdb_value_float
// duckdb_value_double
// duckdb_value_date
// duckdb_value_time
// duckdb_value_timestamp
// duckdb_value_interval
// duckdb_value_varchar
// duckdb_value_string
// duckdb_value_varchar_internal
// duckdb_value_string_internal
// duckdb_value_blob
// duckdb_value_is_null
// duckdb_execute_prepared_streaming
// duckdb_pending_prepared_streaming

// ------------------------------------------------------------------ //
// Arrow Interface
// ------------------------------------------------------------------ //

// duckdb_query_arrow

func QueryArrowSchema(arrow Arrow, outSchema *ArrowSchema) State {
	var arrowSchema C.duckdb_arrow_schema
	state := C.duckdb_query_arrow_schema(arrow.data(), &arrowSchema)
	outSchema.Ptr = unsafe.Pointer(arrowSchema)
	return State(state)
}

// duckdb_prepared_arrow_schema
// duckdb_result_arrow_array

func QueryArrowArray(arrow Arrow, outArray *ArrowArray) State {
	var arrowArray C.duckdb_arrow_array
	state := C.duckdb_query_arrow_array(arrow.data(), &arrowArray)
	outArray.Ptr = unsafe.Pointer(arrowArray)
	return State(state)
}

// duckdb_arrow_column_count

func ArrowRowCount(arrow Arrow) uint64 {
	count := C.duckdb_arrow_row_count(arrow.data())
	return uint64(count)
}

// duckdb_arrow_rows_changed

func QueryArrowError(arrow Arrow) string {
	err := C.duckdb_query_arrow_error(arrow.data())
	return C.GoString(err)
}

func DestroyArrow(arrow *Arrow) {
	data := arrow.data()
	C.duckdb_destroy_arrow(&data)
	arrow.Ptr = nil
}

// duckdb_destroy_arrow
// duckdb_destroy_arrow_stream

func ExecutePreparedArrow(preparedStmt PreparedStatement, outArrow *Arrow) State {
	var arrow C.duckdb_arrow
	state := C.duckdb_execute_prepared_arrow(preparedStmt.data(), &arrow)
	outArrow.Ptr = unsafe.Pointer(arrow)
	return State(state)
}

func ArrowScan(conn Connection, table string, stream ArrowStream) State {
	cTable := C.CString(table)
	defer Free(unsafe.Pointer(cTable))

	state := C.duckdb_arrow_scan(conn.data(), cTable, stream.data())
	return State(state)
}

// duckdb_arrow_array_scan
// duckdb_stream_fetch_chunk

// duckdb_create_instance_cache
// duckdb_get_or_create_from_cache
// duckdb_destroy_instance_cache

// duckdb_append_default_to_chunk

// ------------------------------------------------------------------ //
// Go Bindings Helper
// ------------------------------------------------------------------ //

func MallocLogicalTypeSlice(count uint64) (unsafe.Pointer, []LogicalType) {
	// FIXME: Unify code with other slice allocation functions.

	var dummy C.duckdb_logical_type
	size := C.size_t(unsafe.Sizeof(dummy))

	ptr := unsafe.Pointer(C.malloc(C.size_t(count) * size))
	slice := (*[1 << 30]C.duckdb_logical_type)(ptr)[:count:count]

	var outSlice []LogicalType
	for i := 0; i < int(count); i++ {
		outSlice = append(outSlice, LogicalType{Ptr: unsafe.Pointer(slice[i])})
	}
	return ptr, outSlice
}
