package config

var (
	SECRET     string
	SERVERPORT int16
)

// const code feature
const (
	FEAT_USER_CODE         = "001"
	FEAT_AUTH_CODE         = "002"
	FEAT_QUESTIONAIRE_CODE = "003"
	FEAT_PATIENT_CODE      = "004"
)

// const code layer
const (
	LAYER_DATA_CODE       = "DATA"
	LAYER_SERVICE_CODE    = "SERVICE"
	LAYER_HANDLER_CODE    = "HANDLER"
	LAYER_DEFAULT_CODE    = "DEFAULT"
	RESPONSE_SUCCESS_CODE = "OK"
)

// const db error
const (
	// ErrRecordNotFound record not found error
	DB_ERR_RECORD_NOT_FOUND = "error data not found"
	// ErrMissingWhereClause missing where clause
	DB_ERR_MISSING_WHERE_CLAUSE = "error WHERE conditions required"
	// ErrUnsupportedRelation unsupported relations
	DB_ERR_UNSUPPORTED_RELATION = "error unsupported relations"
	// ErrInvalidData unsupported data
	DB_ERR_INVALID_DATA = "error unsupported data"
	// ErrInvalidField invalid field
	DB_ERR_INVALID_FIELD = "error invalid field"
	// ErrPreloadNotAllowed preload is not allowed when count is used
	DB_ERR_PRELOAD_NOT_ALLOWED = "error preload is not allowed when count is used"
	// ErrInvalidDB
	DB_ERR_INVALID_DB = "error invalid database"
	// ErrPrimaryKeyRequired
	DB_ERR_PRIMARY_KEY_REQUIRED = "error primary key required"
	// ErrDuplicatedKey
	DB_ERR_DUPLICATE_KEY = "duplicated key not allowed"
)

// Failed JWT Response
const (
	JWT_InvalidJwtToken       string = "jwt token missing or invalid"
	JWT_FailedCastingJwtToken string = "failed to cast claims as jwt.MapClaims"
	JWT_FailedCreateToken     string = "failed generate token"
)

// auth
const (
	ERR_AuthWrongCredentials = "wrong email/password"
)

// input request body
const (
	REQ_InvalidParam      string = "invalid param"
	REQ_InvalidIdParam    string = "invaild id param"
	REQ_InvalidPageParam  string = "invalid page param"
	REQ_InvalidLimitParam string = "invalid limit param"
	REQ_ErrorBindData     string = "error bind data"
)

// validation input file
const (
	VAL_InvalidImageFileType string = "invalid image file type"
	VAL_InvalidFileSize      string = "invalid file size"
)

// validation input
const (
	VAL_InvalidValidation string = "validation error"
)
