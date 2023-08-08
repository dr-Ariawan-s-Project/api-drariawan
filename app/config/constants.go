package config

var (
	SECRET     string
	SERVERPORT int16
)

// var (
// 	// ErrRecordNotFound record not found error
// 	ErrRecordNotFound = errors.New("error data not found")
// 	// ErrMissingWhereClause missing where clause
// 	ErrMissingWhereClause = errors.New("WHERE conditions required")
// 	// ErrUnsupportedRelation unsupported relations
// 	ErrUnsupportedRelation = errors.New("unsupported relations")
// 	// ErrInvalidData unsupported data
// 	ErrInvalidData = errors.New("unsupported data")
// 	// ErrInvalidField invalid field
// 	ErrInvalidField = errors.New("invalid field")
// 	// ErrPreloadNotAllowed preload is not allowed when count is used
// 	ErrPreloadNotAllowed = errors.New("preload is not allowed when count is used")
// )

// const db error
const (
	// ErrRecordNotFound record not found error
	ERR_RECORD_NOT_FOUND = "error data not found"
	// ErrMissingWhereClause missing where clause
	ERR_MISSING_WHERE_CLAUSE = "error WHERE conditions required"
	// ErrUnsupportedRelation unsupported relations
	ERR_UNSUPPORTED_RELATION = "error unsupported relations"
	// ErrInvalidData unsupported data
	ERR_INVALID_DATA = "error unsupported data"
	// ErrInvalidField invalid field
	ERR_INVALID_FIELD = "error invalid field"
	// ErrPreloadNotAllowed preload is not allowed when count is used
	ERR_PRELOAD_NOT_ALLOWED = "error preload is not allowed when count is used"
	// ErrInvalidDB
	ERR_INVALID_DB = "error invalid database"
	// ErrPrimaryKeyRequired
	ERR_PRIMARY_KEY_REQUIRED = "error primary key required"
)

// Failed Response
const (
	JWT_InvalidJwtToken       string = "jwt token missing or invalid"
	JWT_FailedCastingJwtToken string = "failed to cast claims as jwt.MapClaims"
)

// input request body
const (
	InvalidIdParam    string = "invaild id param"
	InvalidPageParam  string = "invalid page param"
	InvalidLimitParam string = "invalid limit param"
	ErrorBindData     string = "error bind data"
)

// input file
const (
	InvalidImageFileType string = "invalid image file type"
	InvalidFileSize      string = "invalid file size"
)

// Failed
const (
	VALIDATION_InvalidInput string = "invalid input"
)
