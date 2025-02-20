package duckdb_go_bindings

/*
#include <duckdb.h>
*/
import "C"
import (
	"log"
	"sync/atomic"
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

// ------------------------------------------------------------------ //
// Functions
// ------------------------------------------------------------------ //

// ------------------------------------------------------------------ //
// Open Connect
// ------------------------------------------------------------------ //

// duckdb_open

// OpenExt wraps duckdb_open_ext.
// outDb must be closed with Close.
func OpenExt(path string, outDb *Database, config Config, errMsg *string) State {
	cPath := C.CString(path)
	defer Free(unsafe.Pointer(cPath))

	var err *C.char
	defer Free(unsafe.Pointer(err))

	var db C.duckdb_database
	state := State(C.duckdb_open_ext(cPath, &db, config.data(), &err))
	outDb.Ptr = unsafe.Pointer(db)
	*errMsg = C.GoString(err)

	if debugMode {
		allocCounters.db.Add(1)
	}
	return state
}

// Close wraps duckdb_close.
func Close(db *Database) {
	if debugMode {
		allocCounters.db.Add(-1)
	}
	if db.Ptr == nil {
		return
	}
	data := db.data()
	C.duckdb_close(&data)
	db.Ptr = nil
}

// Connect wraps duckdb_connect.
// outConn must be disconnected with Disconnect.
func Connect(db Database, outConn *Connection) State {
	var conn C.duckdb_connection
	state := State(C.duckdb_connect(db.data(), &conn))
	outConn.Ptr = unsafe.Pointer(conn)

	if debugMode {
		allocCounters.conn.Add(1)
	}
	return state
}

func Interrupt(conn Connection) {
	C.duckdb_interrupt(conn.data())
}

// duckdb_query_progress

// Disconnect wraps duckdb_disconnect.
func Disconnect(conn *Connection) {
	if debugMode {
		allocCounters.conn.Add(-1)
	}
	if conn.Ptr == nil {
		return
	}
	data := conn.data()
	C.duckdb_disconnect(&data)
	conn.Ptr = nil
}

// duckdb_library_version

// ------------------------------------------------------------------ //
// Configuration
// ------------------------------------------------------------------ //

// CreateConfig wraps duckdb_create_config.
// outConfig must be destroyed with DestroyConfig.
func CreateConfig(outConfig *Config) State {
	var config C.duckdb_config
	state := State(C.duckdb_create_config(&config))
	outConfig.Ptr = unsafe.Pointer(config)

	if debugMode {
		allocCounters.config.Add(1)
	}
	return state
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

// DestroyConfig wraps duckdb_destroy_config.
func DestroyConfig(config *Config) {
	if debugMode {
		allocCounters.config.Add(-1)
	}
	if config.Ptr == nil {
		return
	}
	data := config.data()
	C.duckdb_destroy_config(&data)
	config.Ptr = nil
}

// ------------------------------------------------------------------ //
// Query Execution
// ------------------------------------------------------------------ //

// duckdb_query

// DestroyResult wraps duckdb_destroy_result.
func DestroyResult(res *Result) {
	if debugMode {
		allocCounters.res.Add(-1)
	}
	if res == nil {
		return
	}
	C.duckdb_destroy_result(&res.data)
	res = nil
}

func ColumnName(res *Result, col uint64) string {
	name := C.duckdb_column_name(&res.data, C.idx_t(col))
	return C.GoString(name)
}

func ColumnType(res *Result, col uint64) Type {
	t := C.duckdb_column_type(&res.data, C.idx_t(col))
	return Type(t)
}

// duckdb_result_statement_type

// ColumnLogicalType wraps duckdb_column_logical_type.
// The return value must be destroyed with DestroyLogicalType.
func ColumnLogicalType(res *Result, col uint64) LogicalType {
	logicalType := C.duckdb_column_logical_type(&res.data, C.idx_t(col))

	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func ColumnCount(res *Result) uint64 {
	count := C.duckdb_column_count(&res.data)
	return uint64(count)
}

// duckdb_rows_changed

func ResultError(res *Result) string {
	err := C.duckdb_result_error(&res.data)
	return C.GoString(err)
}

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

func VectorSize() uint64 {
	size := C.duckdb_vector_size()
	return uint64(size)
}

func StringIsInlined(strT StringT) bool {
	isInlined := C.duckdb_string_is_inlined(C.duckdb_string_t(strT))
	return bool(isInlined)
}

func StringTLength(strT StringT) uint32 {
	length := C.duckdb_string_t_length(C.duckdb_string_t(strT))
	return uint32(length)
}

func StringTData(strT *StringT) string {
	cStr := C.duckdb_string_t(*strT)
	length := StringTLength(*strT)
	if StringIsInlined(*strT) {
		return string(C.GoBytes(unsafe.Pointer(C.duckdb_string_t_data(&cStr)), C.int(length)))
	}
	return string(C.GoBytes(unsafe.Pointer(C.duckdb_string_t_data(&cStr)), C.int(length)))
}

// ------------------------------------------------------------------ //
// Date Time Timestamp Helpers
// ------------------------------------------------------------------ //

func FromDate(date Date) DateStruct {
	dateStruct := C.duckdb_from_date(C.duckdb_date(date))
	return DateStruct(dateStruct)
}

// duckdb_to_date
// duckdb_is_finite_date
// duckdb_from_time

func CreateTimeTZ(micros int64, offset int32) TimeTZ {
	timeTZ := C.duckdb_create_time_tz(C.int64_t(micros), C.int32_t(offset))
	return TimeTZ(timeTZ)
}

func FromTimeTZ(ti TimeTZ) TimeTZStruct {
	timeTZStruct := C.duckdb_from_time_tz(C.duckdb_time_tz(ti))
	return TimeTZStruct(timeTZStruct)
}

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

// DestroyPrepare wraps duckdb_destroy_prepare.
func DestroyPrepare(preparedStmt *PreparedStatement) {
	if debugMode {
		allocCounters.preparedStmt.Add(-1)
	}
	if preparedStmt.Ptr == nil {
		return
	}
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
	state := C.duckdb_bind_hugeint(preparedStmt.data(), C.idx_t(index), C.duckdb_hugeint(v))
	return State(state)
}

// duckdb_bind_uhugeint

func BindDecimal(preparedStmt PreparedStatement, index uint64, v Decimal) State {
	state := C.duckdb_bind_decimal(preparedStmt.data(), C.idx_t(index), C.duckdb_decimal(v))
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
	state := C.duckdb_bind_date(preparedStmt.data(), C.idx_t(index), C.duckdb_date(v))
	return State(state)
}

func BindTime(preparedStmt PreparedStatement, index uint64, v Time) State {
	state := C.duckdb_bind_time(preparedStmt.data(), C.idx_t(index), C.duckdb_time(v))
	return State(state)
}

func BindTimestamp(preparedStmt PreparedStatement, index uint64, v Timestamp) State {
	state := C.duckdb_bind_timestamp(preparedStmt.data(), C.idx_t(index), C.duckdb_timestamp(v))
	return State(state)
}

// duckdb_bind_timestamp_tz

func BindInterval(preparedStmt PreparedStatement, index uint64, v Interval) State {
	state := C.duckdb_bind_interval(preparedStmt.data(), C.idx_t(index), C.duckdb_interval(v))
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

// ExtractStatements wraps duckdb_extract_statements.
// outExtractedStmts must be destroyed with DestroyExtracted.
func ExtractStatements(conn Connection, query string, outExtractedStmts *ExtractedStatements) uint64 {
	cQuery := C.CString(query)
	defer Free(unsafe.Pointer(cQuery))

	var extractedStmts C.duckdb_extracted_statements
	count := C.duckdb_extract_statements(conn.data(), cQuery, &extractedStmts)
	outExtractedStmts.Ptr = unsafe.Pointer(extractedStmts)

	if debugMode {
		allocCounters.extractedStmts.Add(1)
	}
	return uint64(count)
}

// PrepareExtractedStatement wraps duckdb_prepare_extracted_statement.
// outPreparedStmt must be destroyed with DestroyPrepare.
func PrepareExtractedStatement(conn Connection, extractedStmts ExtractedStatements, index uint64, outPreparedStmt *PreparedStatement) State {
	var preparedStmt C.duckdb_prepared_statement
	state := C.duckdb_prepare_extracted_statement(conn.data(), extractedStmts.data(), C.idx_t(index), &preparedStmt)
	outPreparedStmt.Ptr = unsafe.Pointer(preparedStmt)

	if debugMode {
		allocCounters.preparedStmt.Add(1)
	}
	return State(state)
}

func ExtractStatementsError(extractedStmts ExtractedStatements) string {
	err := C.duckdb_extract_statements_error(extractedStmts.data())
	return C.GoString(err)
}

// DestroyExtracted wraps duckdb_destroy_extracted.
func DestroyExtracted(extractedStmts *ExtractedStatements) {
	if debugMode {
		allocCounters.extractedStmts.Add(-1)
	}
	if extractedStmts.Ptr == nil {
		return
	}
	data := extractedStmts.data()
	C.duckdb_destroy_extracted(&data)
	extractedStmts.Ptr = nil
}

// ------------------------------------------------------------------ //
// Pending Result Interface
// ------------------------------------------------------------------ //

// PendingPrepared wraps duckdb_pending_prepared.
// outPendingRes must be destroyed with DestroyPending.
func PendingPrepared(preparedStmt PreparedStatement, outPendingRes *PendingResult) State {
	var pendingRes C.duckdb_pending_result
	state := C.duckdb_pending_prepared(preparedStmt.data(), &pendingRes)
	outPendingRes.Ptr = unsafe.Pointer(pendingRes)

	if debugMode {
		allocCounters.pendingRes.Add(1)
	}
	return State(state)
}

// DestroyPending wraps duckdb_destroy_pending.
func DestroyPending(pendingRes *PendingResult) {
	if debugMode {
		allocCounters.pendingRes.Add(-1)
	}
	if pendingRes.Ptr == nil {
		return
	}
	data := pendingRes.data()
	C.duckdb_destroy_pending(&data)
	pendingRes.Ptr = nil
}

func PendingError(pendingRes PendingResult) string {
	err := C.duckdb_pending_error(pendingRes.data())
	return C.GoString(err)
}

// duckdb_pending_execute_task
// duckdb_pending_execute_check_state

// ExecutePending wraps duckdb_execute_pending.
// outRes must be destroyed with DestroyResult.
func ExecutePending(res PendingResult, outRes *Result) State {
	state := State(C.duckdb_execute_pending(res.data(), &outRes.data))
	if debugMode {
		allocCounters.res.Add(1)
	}
	return state
}

// duckdb_pending_execution_is_finished

// ------------------------------------------------------------------ //
// Value Interface
// ------------------------------------------------------------------ //

// DestroyValue wraps duckdb_destroy_value.
func DestroyValue(v *Value) {
	if debugMode {
		allocCounters.v.Add(-1)
	}
	if v.Ptr == nil {
		return
	}
	data := v.data()
	C.duckdb_destroy_value(&data)
	v.Ptr = nil
}

// CreateVarchar wraps duckdb_create_varchar.
// The return value must be destroyed with DestroyValue.
func CreateVarchar(str string) Value {
	cStr := C.CString(str)
	defer Free(unsafe.Pointer(cStr))
	v := C.duckdb_create_varchar(cStr)

	if debugMode {
		allocCounters.v.Add(1)
	}
	return Value{
		Ptr: unsafe.Pointer(v),
	}
}

// duckdb_create_varchar_length
// duckdb_create_bool
// duckdb_create_int8
// duckdb_create_uint8
// duckdb_create_int16
// duckdb_create_uint16
// duckdb_create_int32
// duckdb_create_uint32
// duckdb_create_uint64

// CreateInt64 wraps duckdb_create_int64.
// The return value must be destroyed with DestroyValue.
func CreateInt64(val int64) Value {
	v := C.duckdb_create_int64(C.int64_t(val))
	if debugMode {
		allocCounters.v.Add(1)
	}
	return Value{
		Ptr: unsafe.Pointer(v),
	}
}

// duckdb_create_hugeint
// duckdb_create_uhugeint
// duckdb_create_varint
// duckdb_create_decimal
// duckdb_create_float
// duckdb_create_double
// duckdb_create_date
// duckdb_create_time

// CreateTimeTZValue wraps duckdb_create_time_tz_value.
// The return value must be destroyed with DestroyValue.
func CreateTimeTZValue(timeTZ TimeTZ) Value {
	v := C.duckdb_create_time_tz_value(C.duckdb_time_tz(timeTZ))
	if debugMode {
		allocCounters.v.Add(1)
	}
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

func GetBool(v Value) bool {
	val := C.duckdb_get_bool(v.data())
	return bool(val)
}

func GetInt8(v Value) int8 {
	val := C.duckdb_get_int8(v.data())
	return int8(val)
}

func GetUInt8(v Value) uint8 {
	val := C.duckdb_get_uint8(v.data())
	return uint8(val)
}

func GetInt16(v Value) int16 {
	val := C.duckdb_get_int16(v.data())
	return int16(val)
}

func GetUInt16(v Value) uint16 {
	val := C.duckdb_get_uint16(v.data())
	return uint16(val)
}

func GetInt32(v Value) int32 {
	val := C.duckdb_get_int32(v.data())
	return int32(val)
}

func GetUInt32(v Value) uint32 {
	val := C.duckdb_get_uint32(v.data())
	return uint32(val)
}

func GetInt64(v Value) int64 {
	val := C.duckdb_get_int64(v.data())
	return int64(val)
}

func GetUInt64(v Value) uint64 {
	val := C.duckdb_get_uint64(v.data())
	return uint64(val)
}

func GetHugeInt(v Value) HugeInt {
	val := C.duckdb_get_hugeint(v.data())
	return HugeInt(val)
}

// duckdb_get_uhugeint
// duckdb_get_varint
// duckdb_get_decimal

func GetFloat(v Value) float32 {
	val := C.duckdb_get_float(v.data())
	return float32(val)
}

func GetDouble(v Value) float64 {
	val := C.duckdb_get_double(v.data())
	return float64(val)
}

func GetDate(v Value) Date {
	val := C.duckdb_get_date(v.data())
	return Date(val)
}

func GetTime(v Value) Time {
	val := C.duckdb_get_time(v.data())
	return Time(val)
}

func GetTimeTZ(v Value) TimeTZ {
	val := C.duckdb_get_time_tz(v.data())
	return TimeTZ(val)
}

func GetTimestamp(v Value) Timestamp {
	val := C.duckdb_get_timestamp(v.data())
	return Timestamp(val)
}

// duckdb_get_timestamp_tz
// duckdb_get_timestamp_s
// duckdb_get_timestamp_ms
// duckdb_get_timestamp_ns

func GetInterval(v Value) Interval {
	val := C.duckdb_get_interval(v.data())
	return Interval(val)
}

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

// GetMapKey wraps duckdb_get_map_key.
// The return value must be destroyed with DestroyValue.
func GetMapKey(v Value, index uint64) Value {
	value := C.duckdb_get_map_key(v.data(), C.idx_t(index))
	if debugMode {
		allocCounters.v.Add(1)
	}
	return Value{
		Ptr: unsafe.Pointer(value),
	}
}

// GetMapValue wraps duckdb_get_map_value.
// The return value must be destroyed with DestroyValue.
func GetMapValue(v Value, index uint64) Value {
	value := C.duckdb_get_map_value(v.data(), C.idx_t(index))
	if debugMode {
		allocCounters.v.Add(1)
	}
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

// CreateLogicalType wraps duckdb_create_logical_type.
// The return value must be destroyed with DestroyLogicalType.
func CreateLogicalType(t Type) LogicalType {
	logicalType := C.duckdb_create_logical_type(C.duckdb_type(t))
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func LogicalTypeGetAlias(logicalType LogicalType) string {
	alias := C.duckdb_logical_type_get_alias(logicalType.data())
	defer Free(unsafe.Pointer(alias))
	return C.GoString(alias)
}

// duckdb_logical_type_set_alias

// CreateListType wraps duckdb_create_list_type.
// The return value must be destroyed with DestroyLogicalType.
func CreateListType(child LogicalType) LogicalType {
	logicalType := C.duckdb_create_list_type(child.data())
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

// CreateArrayType wraps duckdb_create_array_type.
// The return value must be destroyed with DestroyLogicalType.
func CreateArrayType(child LogicalType, size uint64) LogicalType {
	logicalType := C.duckdb_create_array_type(child.data(), C.idx_t(size))
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

// CreateMapType wraps duckdb_create_map_type.
// The return value must be destroyed with DestroyLogicalType.
func CreateMapType(key LogicalType, value LogicalType) LogicalType {
	logicalType := C.duckdb_create_map_type(key.data(), value.data())
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

// duckdb_create_union_type

// CreateStructType wraps duckdb_create_struct_type.
// The return value must be destroyed with DestroyLogicalType.
func CreateStructType(types []LogicalType, names []string) LogicalType {
	count := len(types)

	typesSlice := allocLogicalTypeSlice(types)
	defer Free(typesSlice)
	typesPtr := (*C.duckdb_logical_type)(typesSlice)

	namesSlice := (*[1 << 31]*C.char)(C.malloc(C.size_t(count) * charSize))
	defer Free(unsafe.Pointer(namesSlice))

	for i, name := range names {
		(*namesSlice)[i] = C.CString(name)
	}
	namesPtr := (**C.char)(unsafe.Pointer(namesSlice))

	// Create the STRUCT type.
	logicalType := C.duckdb_create_struct_type(typesPtr, namesPtr, C.idx_t(count))

	for i := 0; i < count; i++ {
		Free(unsafe.Pointer((*namesSlice)[i]))
	}

	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

// CreateEnumType wraps duckdb_create_enum_type.
// The return value must be destroyed with DestroyLogicalType.
func CreateEnumType(names []string) LogicalType {
	count := len(names)

	namesSlice := (*[1 << 31]*C.char)(C.malloc(C.size_t(count) * charSize))
	defer Free(unsafe.Pointer(namesSlice))

	for i, name := range names {
		(*namesSlice)[i] = C.CString(name)
	}
	namesPtr := (**C.char)(unsafe.Pointer(namesSlice))

	// Create the ENUM type.
	logicalType := C.duckdb_create_enum_type(namesPtr, C.idx_t(count))

	for i := 0; i < count; i++ {
		Free(unsafe.Pointer((*namesSlice)[i]))
	}

	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

// CreateDecimalType wraps duckdb_create_decimal_type.
// The return value must be destroyed with DestroyLogicalType.
func CreateDecimalType(width uint8, scale uint8) LogicalType {
	logicalType := C.duckdb_create_decimal_type(C.uint8_t(width), C.uint8_t(scale))
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func GetTypeId(logicalType LogicalType) Type {
	t := C.duckdb_get_type_id(logicalType.data())
	return Type(t)
}

func DecimalWidth(logicalType LogicalType) uint8 {
	width := C.duckdb_decimal_width(logicalType.data())
	return uint8(width)
}

func DecimalScale(logicalType LogicalType) uint8 {
	scale := C.duckdb_decimal_scale(logicalType.data())
	return uint8(scale)
}

func DecimalInternalType(logicalType LogicalType) Type {
	t := C.duckdb_decimal_internal_type(logicalType.data())
	return Type(t)
}

func EnumInternalType(logicalType LogicalType) Type {
	t := C.duckdb_enum_internal_type(logicalType.data())
	return Type(t)
}

func EnumDictionarySize(logicalType LogicalType) uint32 {
	size := C.duckdb_enum_dictionary_size(logicalType.data())
	return uint32(size)
}

func EnumDictionaryValue(logicalType LogicalType, index uint64) string {
	str := C.duckdb_enum_dictionary_value(logicalType.data(), C.idx_t(index))
	defer Free(unsafe.Pointer(str))
	return C.GoString(str)
}

// ListTypeChildType wraps duckdb_list_type_child_type.
// The return value must be destroyed with DestroyLogicalType.
func ListTypeChildType(logicalType LogicalType) LogicalType {
	child := C.duckdb_list_type_child_type(logicalType.data())
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(child),
	}
}

// ArrayTypeChildType wraps duckdb_array_type_child_type.
// The return value must be destroyed with DestroyLogicalType.
func ArrayTypeChildType(logicalType LogicalType) LogicalType {
	child := C.duckdb_array_type_child_type(logicalType.data())
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(child),
	}
}

func ArrayTypeArraySize(logicalType LogicalType) uint64 {
	size := C.duckdb_array_type_array_size(logicalType.data())
	return uint64(size)
}

// MapTypeKeyType wraps duckdb_map_type_key_type.
// The return value must be destroyed with DestroyLogicalType.
func MapTypeKeyType(logicalType LogicalType) LogicalType {
	key := C.duckdb_map_type_key_type(logicalType.data())
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(key),
	}
}

// MapTypeValueType wraps duckdb_map_type_value_type.
// The return value must be destroyed with DestroyLogicalType.
func MapTypeValueType(logicalType LogicalType) LogicalType {
	value := C.duckdb_map_type_value_type(logicalType.data())
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(value),
	}
}

func StructTypeChildCount(logicalType LogicalType) uint64 {
	count := C.duckdb_struct_type_child_count(logicalType.data())
	return uint64(count)
}

func StructTypeChildName(logicalType LogicalType, index uint64) string {
	cName := C.duckdb_struct_type_child_name(logicalType.data(), C.idx_t(index))
	defer Free(unsafe.Pointer(cName))
	return C.GoString(cName)
}

// StructTypeChildType wraps duckdb_struct_type_child_type.
// The return value must be destroyed with DestroyLogicalType.
func StructTypeChildType(logicalType LogicalType, index uint64) LogicalType {
	child := C.duckdb_struct_type_child_type(logicalType.data(), C.idx_t(index))
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(child),
	}
}

// duckdb_union_type_member_count
// duckdb_union_type_member_name
// duckdb_union_type_member_type

// DestroyLogicalType wraps duckdb_destroy_logical_type.
func DestroyLogicalType(logicalType *LogicalType) {
	if debugMode {
		allocCounters.logicalType.Add(-1)
	}
	if logicalType.Ptr == nil {
		return
	}
	data := logicalType.data()
	C.duckdb_destroy_logical_type(&data)
	logicalType.Ptr = nil
}

// duckdb_register_logical_type

// ------------------------------------------------------------------ //
// Data Chunk Interface
// ------------------------------------------------------------------ //

// CreateDataChunk wraps duckdb_create_data_chunk.
// The return value must be destroyed with DestroyDataChunk.
func CreateDataChunk(types []LogicalType) DataChunk {
	count := len(types)

	typesSlice := allocLogicalTypeSlice(types)
	typesPtr := (*C.duckdb_logical_type)(typesSlice)
	defer Free(unsafe.Pointer(typesPtr))

	chunk := C.duckdb_create_data_chunk(typesPtr, C.idx_t(count))
	if debugMode {
		allocCounters.chunk.Add(1)
	}
	return DataChunk{
		Ptr: unsafe.Pointer(chunk),
	}
}

// DestroyDataChunk wraps duckdb_destroy_data_chunk.
func DestroyDataChunk(chunk *DataChunk) {
	if debugMode {
		allocCounters.chunk.Add(-1)
	}
	if chunk.Ptr == nil {
		return
	}
	data := chunk.data()
	C.duckdb_destroy_data_chunk(&data)
	chunk.Ptr = nil
}

// duckdb_data_chunk_reset

func DataChunkGetColumnCount(chunk DataChunk) uint64 {
	count := C.duckdb_data_chunk_get_column_count(chunk.data())
	return uint64(count)
}

func DataChunkGetVector(chunk DataChunk, index uint64) Vector {
	vec := C.duckdb_data_chunk_get_vector(chunk.data(), C.idx_t(index))
	return Vector{
		Ptr: unsafe.Pointer(vec),
	}
}

func DataChunkGetSize(chunk DataChunk) uint64 {
	size := C.duckdb_data_chunk_get_size(chunk.data())
	return uint64(size)
}

func DataChunkSetSize(chunk DataChunk, size uint64) {
	C.duckdb_data_chunk_set_size(chunk.data(), C.idx_t(size))
}

// ------------------------------------------------------------------ //
// Vector Interface
// ------------------------------------------------------------------ //

// VectorGetColumnType wraps duckdb_vector_get_column_type.
// The return value must be destroyed with DestroyLogicalType.
func VectorGetColumnType(vec Vector) LogicalType {
	logicalType := C.duckdb_vector_get_column_type(vec.data())
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
	return LogicalType{
		Ptr: unsafe.Pointer(logicalType),
	}
}

func VectorGetData(vec Vector) unsafe.Pointer {
	ptr := C.duckdb_vector_get_data(vec.data())
	return unsafe.Pointer(ptr)
}

func VectorGetValidity(vec Vector) unsafe.Pointer {
	mask := C.duckdb_vector_get_validity(vec.data())
	return unsafe.Pointer(mask)
}

func VectorEnsureValidityWritable(vec Vector) {
	C.duckdb_vector_ensure_validity_writable(vec.data())
}

func VectorAssignStringElement(vec Vector, index uint64, str string) {
	cStr := C.CString(str)
	defer Free(unsafe.Pointer(cStr))
	C.duckdb_vector_assign_string_element(vec.data(), C.idx_t(index), cStr)
}

func VectorAssignStringElementLen(vec Vector, index uint64, blob []byte, len uint64) {
	cBytes := (*C.char)(C.CBytes(blob))
	defer Free(unsafe.Pointer(cBytes))
	C.duckdb_vector_assign_string_element_len(vec.data(), C.idx_t(index), cBytes, C.idx_t(len))
}

func ListVectorGetChild(vec Vector) Vector {
	child := C.duckdb_list_vector_get_child(vec.data())
	return Vector{
		Ptr: unsafe.Pointer(child),
	}
}

func ListVectorGetSize(vec Vector) uint64 {
	size := C.duckdb_list_vector_get_size(vec.data())
	return uint64(size)
}

func ListVectorSetSize(vec Vector, size uint64) State {
	state := C.duckdb_list_vector_set_size(vec.data(), C.idx_t(size))
	return State(state)
}

func ListVectorReserve(vec Vector, capacity uint64) State {
	state := C.duckdb_list_vector_reserve(vec.data(), C.idx_t(capacity))
	return State(state)
}

func StructVectorGetChild(vec Vector, index uint64) Vector {
	child := C.duckdb_struct_vector_get_child(vec.data(), C.idx_t(index))
	return Vector{
		Ptr: unsafe.Pointer(child),
	}
}

func ArrayVectorGetChild(vec Vector) Vector {
	child := C.duckdb_array_vector_get_child(vec.data())
	return Vector{
		Ptr: unsafe.Pointer(child),
	}
}

// ------------------------------------------------------------------ //
// Validity Mask Functions
// ------------------------------------------------------------------ //

// duckdb_validity_row_is_valid
// duckdb_validity_set_row_validity

func ValiditySetRowInvalid(maskPtr unsafe.Pointer, row uint64) {
	mask := (*C.uint64_t)(unsafe.Pointer(maskPtr))
	C.duckdb_validity_set_row_invalid(mask, C.idx_t(row))
}

// duckdb_validity_set_row_valid

// ------------------------------------------------------------------ //
// Scalar Functions
// ------------------------------------------------------------------ //

// CreateScalarFunction wraps duckdb_create_scalar_function.
// The return value must be destroyed with DestroyScalarFunction.
func CreateScalarFunction() ScalarFunction {
	f := C.duckdb_create_scalar_function()
	if debugMode {
		allocCounters.scalarFunc.Add(1)
	}
	return ScalarFunction{
		Ptr: unsafe.Pointer(f),
	}
}

// DestroyScalarFunction wraps duckdb_destroy_scalar_function.
func DestroyScalarFunction(f *ScalarFunction) {
	if debugMode {
		allocCounters.scalarFunc.Add(-1)
	}
	if f.Ptr == nil {
		return
	}
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

// CreateScalarFunctionSet wraps duckdb_create_scalar_function_set.
// The return value must be destroyed with DestroyScalarFunctionSet.
func CreateScalarFunctionSet(name string) ScalarFunctionSet {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))

	set := C.duckdb_create_scalar_function_set(cName)
	if debugMode {
		allocCounters.scalarFuncSet.Add(1)
	}
	return ScalarFunctionSet{
		Ptr: unsafe.Pointer(set),
	}
}

// DestroyScalarFunctionSet wraps duckdb_destroy_scalar_function_set.
func DestroyScalarFunctionSet(set *ScalarFunctionSet) {
	if debugMode {
		allocCounters.scalarFuncSet.Add(-1)
	}
	if set.Ptr == nil {
		return
	}
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

// CreateTableFunction wraps duckdb_create_table_function.
// The return value must be destroyed with DestroyTableFunction.
func CreateTableFunction() TableFunction {
	f := C.duckdb_create_table_function()
	if debugMode {
		allocCounters.tableFunc.Add(1)
	}
	return TableFunction{
		Ptr: unsafe.Pointer(f),
	}
}

// DestroyTableFunction wraps duckdb_destroy_table_function.
func DestroyTableFunction(f *TableFunction) {
	if debugMode {
		allocCounters.tableFunc.Add(-1)
	}
	if f.Ptr == nil {
		return
	}
	data := f.data()
	C.duckdb_destroy_table_function(&data)
	f.Ptr = nil
}

func TableFunctionSetName(f TableFunction, name string) {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))
	C.duckdb_table_function_set_name(f.data(), cName)
}

func TableFunctionAddParameter(f TableFunction, logicalType LogicalType) {
	C.duckdb_table_function_add_parameter(f.data(), logicalType.data())
}

func TableFunctionAddNamedParameter(f TableFunction, name string, logicalType LogicalType) {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))
	C.duckdb_table_function_add_named_parameter(f.data(), cName, logicalType.data())
}

func TableFunctionSetExtraInfo(f TableFunction, extraInfoPtr unsafe.Pointer, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_delete_callback_t(callbackPtr)
	C.duckdb_table_function_set_extra_info(f.data(), extraInfoPtr, callback)
}

func TableFunctionSetBind(f TableFunction, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_table_function_bind_t(callbackPtr)
	C.duckdb_table_function_set_bind(f.data(), callback)
}

func TableFunctionSetInit(f TableFunction, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_table_function_init_t(callbackPtr)
	C.duckdb_table_function_set_init(f.data(), callback)
}

func TableFunctionSetLocalInit(f TableFunction, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_table_function_init_t(callbackPtr)
	C.duckdb_table_function_set_local_init(f.data(), callback)
}

func TableFunctionSetFunction(f TableFunction, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_table_function_t(callbackPtr)
	C.duckdb_table_function_set_function(f.data(), callback)
}

func TableFunctionSupportsProjectionPushdown(f TableFunction, pushdown bool) {
	C.duckdb_table_function_supports_projection_pushdown(f.data(), C.bool(pushdown))
}

func RegisterTableFunction(conn Connection, f TableFunction) State {
	state := C.duckdb_register_table_function(conn.data(), f.data())
	return State(state)
}

// ------------------------------------------------------------------ //
// Table Function Bind
// ------------------------------------------------------------------ //

func BindGetExtraInfo(info BindInfo) unsafe.Pointer {
	ptr := C.duckdb_bind_get_extra_info(info.data())
	return unsafe.Pointer(ptr)
}

func BindAddResultColumn(info BindInfo, name string, logicalType LogicalType) {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))

	C.duckdb_bind_add_result_column(info.data(), cName, logicalType.data())
}

// duckdb_bind_get_parameter_count

// BindGetParameter wraps duckdb_bind_get_parameter.
// The return value must be destroyed with DestroyValue.
func BindGetParameter(info BindInfo, index uint64) Value {
	v := C.duckdb_bind_get_parameter(info.data(), C.idx_t(index))
	if debugMode {
		allocCounters.v.Add(1)
	}
	return Value{
		Ptr: unsafe.Pointer(v),
	}
}

// BindGetNamedParameter wraps duckdb_bind_get_named_parameter.
// The return value must be destroyed with DestroyValue.
func BindGetNamedParameter(info BindInfo, name string) Value {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))

	v := C.duckdb_bind_get_named_parameter(info.data(), cName)
	if debugMode {
		allocCounters.v.Add(1)
	}
	return Value{
		Ptr: unsafe.Pointer(v),
	}
}

func BindSetBindData(info BindInfo, bindDataPtr unsafe.Pointer, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_delete_callback_t(callbackPtr)
	C.duckdb_bind_set_bind_data(info.data(), bindDataPtr, callback)
}

func BindSetCardinality(info BindInfo, cardinality uint64, exact bool) {
	C.duckdb_bind_set_cardinality(info.data(), C.idx_t(cardinality), C.bool(exact))
}

func BindSetError(info BindInfo, err string) {
	cErr := C.CString(err)
	defer Free(unsafe.Pointer(cErr))
	C.duckdb_bind_set_error(info.data(), cErr)
}

// ------------------------------------------------------------------ //
// Table Function Init
// ------------------------------------------------------------------ //

// duckdb_init_get_extra_info

func InitGetBindData(info InitInfo) unsafe.Pointer {
	ptr := C.duckdb_init_get_bind_data(info.data())
	return unsafe.Pointer(ptr)
}

func InitSetInitData(info InitInfo, initDataPtr unsafe.Pointer, callbackPtr unsafe.Pointer) {
	callback := C.duckdb_delete_callback_t(callbackPtr)
	C.duckdb_init_set_init_data(info.data(), initDataPtr, callback)
}

func InitGetColumnCount(info InitInfo) uint64 {
	count := C.duckdb_init_get_column_count(info.data())
	return uint64(count)
}

func InitGetColumnIndex(info InitInfo, index uint64) uint64 {
	colIndex := C.duckdb_init_get_column_index(info.data(), C.idx_t(index))
	return uint64(colIndex)
}

func InitSetMaxThreads(info InitInfo, max uint64) {
	C.duckdb_init_set_max_threads(info.data(), C.idx_t(max))
}

// duckdb_init_set_error

// ------------------------------------------------------------------ //
// Table Function
// ------------------------------------------------------------------ //

// duckdb_function_get_extra_info

func FunctionGetBindData(info FunctionInfo) unsafe.Pointer {
	ptr := C.duckdb_function_get_bind_data(info.data())
	return unsafe.Pointer(ptr)
}

// duckdb_function_get_init_data

func FunctionGetLocalInitData(info FunctionInfo) unsafe.Pointer {
	ptr := C.duckdb_function_get_local_init_data(info.data())
	return unsafe.Pointer(ptr)
}

func FunctionSetError(info FunctionInfo, err string) {
	cErr := C.CString(err)
	defer Free(unsafe.Pointer(cErr))
	C.duckdb_function_set_error(info.data(), cErr)
}

// ------------------------------------------------------------------ //
// Replacement Scans
// ------------------------------------------------------------------ //

func AddReplacementScan(db Database, callbackPtr unsafe.Pointer, extraData unsafe.Pointer, deleteCallbackPtr unsafe.Pointer) {
	callback := C.duckdb_replacement_callback_t(callbackPtr)
	deleteCallback := C.duckdb_delete_callback_t(deleteCallbackPtr)
	C.duckdb_add_replacement_scan(db.data(), callback, extraData, deleteCallback)
}

func ReplacementScanSetFunctionName(info ReplacementScanInfo, name string) {
	cName := C.CString(name)
	defer Free(unsafe.Pointer(cName))
	C.duckdb_replacement_scan_set_function_name(info.data(), cName)
}

func ReplacementScanAddParameter(info ReplacementScanInfo, v Value) {
	C.duckdb_replacement_scan_add_parameter(info.data(), v.data())
}

func ReplacementScanSetError(info ReplacementScanInfo, err string) {
	cErr := C.CString(err)
	defer Free(unsafe.Pointer(cErr))
	C.duckdb_replacement_scan_set_error(info.data(), cErr)
}

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

// ProfilingInfoGetMetrics wraps duckdb_profiling_info_get_metrics.
// The return value must be destroyed with DestroyValue.
func ProfilingInfoGetMetrics(info ProfilingInfo) Value {
	v := C.duckdb_profiling_info_get_metrics(info.data())
	if debugMode {
		allocCounters.v.Add(1)
	}
	return Value{
		Ptr: unsafe.Pointer(v),
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

// AppenderCreate wraps duckdb_appender_create.
// outAppender must be destroyed with AppenderDestroy.
func AppenderCreate(conn Connection, schema string, table string, outAppender *Appender) State {
	cSchema := C.CString(schema)
	defer Free(unsafe.Pointer(cSchema))

	cTable := C.CString(table)
	defer Free(unsafe.Pointer(cTable))

	var appender C.duckdb_appender
	state := State(C.duckdb_appender_create(conn.data(), cSchema, cTable, &appender))
	outAppender.Ptr = unsafe.Pointer(appender)

	if debugMode {
		allocCounters.appender.Add(1)
	}
	return state
}

// duckdb_appender_create_ext

func AppenderColumnCount(appender Appender) uint64 {
	count := C.duckdb_appender_column_count(appender.data())
	return uint64(count)
}

// AppenderColumnType wraps duckdb_appender_column_type.
// The return value must be destroyed with DestroyLogicalType.
func AppenderColumnType(appender Appender, index uint64) LogicalType {
	logicalType := C.duckdb_appender_column_type(appender.data(), C.idx_t(index))
	if debugMode {
		allocCounters.logicalType.Add(1)
	}
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

// AppenderDestroy wraps duckdb_appender_destroy.
func AppenderDestroy(appender *Appender) State {
	if debugMode {
		allocCounters.appender.Add(-1)
	}
	if appender.Ptr == nil {
		return StateSuccess
	}
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

// ResultGetChunk wraps duckdb_result_get_chunk.
// The return value must be destroyed with DestroyDataChunk.
func ResultGetChunk(res Result, index uint64) DataChunk {
	chunk := C.duckdb_result_get_chunk(res.data, C.idx_t(index))
	if debugMode {
		allocCounters.chunk.Add(1)
	}
	return DataChunk{
		Ptr: unsafe.Pointer(chunk),
	}
}

// duckdb_result_is_streaming

func ResultChunkCount(res Result) uint64 {
	count := C.duckdb_result_chunk_count(res.data)
	return uint64(count)
}

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
	state := C.duckdb_query_arrow_schema(arrow.data(), (*C.duckdb_arrow_schema)(outSchema.Ptr))
	return State(state)
}

// duckdb_prepared_arrow_schema
// duckdb_result_arrow_array

func QueryArrowArray(arrow Arrow, outArray *ArrowArray) State {
	state := C.duckdb_query_arrow_array(arrow.data(), (*C.duckdb_arrow_array)(outArray.Ptr))
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

// DestroyArrow wraps duckdb_destroy_arrow.
func DestroyArrow(arrow *Arrow) {
	if debugMode {
		allocCounters.arrow.Add(-1)
	}
	if arrow.Ptr == nil {
		return
	}
	data := arrow.data()
	C.duckdb_destroy_arrow(&data)
	arrow.Ptr = nil
}

// duckdb_destroy_arrow_stream

// ExecutePreparedArrow wraps duckdb_execute_prepared_arrow.
// outArrow must be destroyed with DestroyArrow.
func ExecutePreparedArrow(preparedStmt PreparedStatement, outArrow *Arrow) State {
	var arrow C.duckdb_arrow
	state := State(C.duckdb_execute_prepared_arrow(preparedStmt.data(), &arrow))
	outArrow.Ptr = unsafe.Pointer(arrow)
	if debugMode {
		allocCounters.arrow.Add(1)
	}
	return state
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

func ValidityMaskValueIsValid(maskPtr unsafe.Pointer, index uint64) bool {
	entryIdx := index / 64
	idxInEntry := index % 64
	slice := (*[1 << 31]C.uint64_t)(maskPtr)
	isValid := slice[entryIdx] & (C.uint64_t(1) << idxInEntry)
	return uint64(isValid) != 0
}

const (
	logicalTypeSize = C.size_t(unsafe.Sizeof((C.duckdb_logical_type)(nil)))
	charSize        = C.size_t(unsafe.Sizeof((*C.char)(nil)))
)

func allocLogicalTypeSlice(types []LogicalType) unsafe.Pointer {
	count := len(types)

	// Initialize the memory of the logical types.
	typesSlice := (*[1 << 31]C.duckdb_logical_type)(C.malloc(C.size_t(count) * logicalTypeSize))
	for i, t := range types {
		// We only copy the pointers.
		// The actual types live in types.
		(*typesSlice)[i] = t.data()
	}
	return unsafe.Pointer(typesSlice)
}

// ------------------------------------------------------------------ //
// Memory Safety
// ------------------------------------------------------------------ //

type allocationCounters struct {
	db             atomic.Int64
	conn           atomic.Int64
	config         atomic.Int64
	logicalType    atomic.Int64
	preparedStmt   atomic.Int64
	extractedStmts atomic.Int64
	pendingRes     atomic.Int64
	res            atomic.Int64
	v              atomic.Int64
	chunk          atomic.Int64
	scalarFunc     atomic.Int64
	scalarFuncSet  atomic.Int64
	tableFunc      atomic.Int64
	appender       atomic.Int64
	arrow          atomic.Int64
}

var allocCounters = allocationCounters{}

func VerifyAllocationCounters() {
	dbCount := allocCounters.db.Load()
	if dbCount != 0 {
		log.Fatalf("db count is %d", dbCount)
	}
	connCount := allocCounters.conn.Load()
	if connCount != 0 {
		log.Fatalf("conn count is %d", connCount)
	}
	configCount := allocCounters.config.Load()
	if configCount != 0 {
		log.Fatalf("config count is %d", configCount)
	}
	logicalTypeCount := allocCounters.logicalType.Load()
	if logicalTypeCount != 0 {
		log.Fatalf("logical type count is %d", logicalTypeCount)
	}
	preparedStmtCount := allocCounters.preparedStmt.Load()
	if preparedStmtCount != 0 {
		log.Fatalf("preparesd statement count is %d", preparedStmtCount)
	}
	extractedStmtsCount := allocCounters.extractedStmts.Load()
	if extractedStmtsCount != 0 {
		log.Fatalf("extracted statements count is %d", extractedStmtsCount)
	}
	pendingResCount := allocCounters.pendingRes.Load()
	if pendingResCount != 0 {
		log.Fatalf("pending res count is %d", pendingResCount)
	}
	resCount := allocCounters.res.Load()
	if resCount != 0 {
		log.Fatalf("res count is %d", resCount)
	}
	vCount := allocCounters.v.Load()
	if vCount != 0 {
		log.Fatalf("v count is %d", vCount)
	}
	chunkCount := allocCounters.chunk.Load()
	if chunkCount != 0 {
		log.Fatalf("chunk count is %d", chunkCount)
	}
	scalarFuncCount := allocCounters.scalarFunc.Load()
	if scalarFuncCount != 0 {
		log.Fatalf("scalar function count is %d", scalarFuncCount)
	}
	scalarFuncSetCount := allocCounters.scalarFuncSet.Load()
	if scalarFuncSetCount != 0 {
		log.Fatalf("scalar function set count is %d", scalarFuncSetCount)
	}
	tableFuncCount := allocCounters.tableFunc.Load()
	if tableFuncCount != 0 {
		log.Fatalf("table function count is %d", tableFuncCount)
	}
	appenderCount := allocCounters.appender.Load()
	if appenderCount != 0 {
		log.Fatalf("appender count is %d", appenderCount)
	}
	arrowCount := allocCounters.arrow.Load()
	if arrowCount != 0 {
		log.Fatalf("arrow count is %d", arrowCount)
	}
}
