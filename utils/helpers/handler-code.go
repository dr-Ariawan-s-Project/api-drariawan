package helpers

import (
	"net/http"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
)

func CheckHandlerSuccessCode(msg string) int {
	switch true {
	case strings.Contains(msg, "insert") || strings.Contains(msg, "create"):
		return http.StatusCreated
	case strings.Contains(msg, "read") || strings.Contains(msg, "get") || strings.Contains(msg, "update") || strings.Contains(msg, "delete"):
		return http.StatusOK
	default:
		return http.StatusOK
	}
}

func CheckHandlerErrorCode(err error) (responseCode int, layerCode string, errConst error) {
	switch err.Error() {
	case config.ERR_AuthWrongCredentials:
		return http.StatusBadRequest, config.LAYER_SERVICE_CODE, err

	case config.JWT_InvalidJwtToken:
		return http.StatusBadRequest, config.LAYER_HANDLER_CODE, err

	case config.JWT_FailedCastingJwtToken:
		return http.StatusInternalServerError, config.LAYER_HANDLER_CODE, err

	case config.JWT_FailedCreateToken:
		return http.StatusInternalServerError, config.LAYER_SERVICE_CODE, err

	case config.REQ_ErrorBindData:
		return http.StatusBadRequest, config.LAYER_HANDLER_CODE, err

	case config.REQ_InvalidParam:
		return http.StatusBadRequest, config.LAYER_HANDLER_CODE, err

	case config.REQ_InvalidIdParam:
		return http.StatusBadRequest, config.LAYER_HANDLER_CODE, err

	case config.REQ_InvalidPageParam:
		return http.StatusBadRequest, config.LAYER_HANDLER_CODE, err

	case config.REQ_InvalidLimitParam:
		return http.StatusBadRequest, config.LAYER_HANDLER_CODE, err

	case config.VAL_InvalidImageFileType:
		return http.StatusBadRequest, config.LAYER_SERVICE_CODE, err

	case config.VAL_InvalidFileSize:
		return http.StatusBadRequest, config.LAYER_SERVICE_CODE, err

	case config.VAL_InvalidValidation:
		return http.StatusBadRequest, config.LAYER_SERVICE_CODE, err

	case config.DB_ERR_RECORD_NOT_FOUND:
		return http.StatusBadRequest, config.LAYER_DATA_CODE, err

	case config.DB_ERR_MISSING_WHERE_CLAUSE:
		return http.StatusInternalServerError, config.LAYER_DATA_CODE, err

	case config.DB_ERR_UNSUPPORTED_RELATION:
		return http.StatusInternalServerError, config.LAYER_DATA_CODE, err

	case config.DB_ERR_INVALID_DATA:
		return http.StatusInternalServerError, config.LAYER_DATA_CODE, err

	case config.DB_ERR_INVALID_FIELD:
		return http.StatusInternalServerError, config.LAYER_DATA_CODE, err

	case config.DB_ERR_PRELOAD_NOT_ALLOWED:
		return http.StatusInternalServerError, config.LAYER_DATA_CODE, err

	case config.DB_ERR_INVALID_DB:
		return http.StatusInternalServerError, config.LAYER_DATA_CODE, err

	case config.DB_ERR_PRIMARY_KEY_REQUIRED:
		return http.StatusInternalServerError, config.LAYER_DATA_CODE, err

	case config.DB_ERR_DUPLICATE_KEY:
		return http.StatusBadRequest, config.LAYER_DATA_CODE, err

	default:
		return http.StatusInternalServerError, config.LAYER_DEFAULT_CODE, err
	}
}
