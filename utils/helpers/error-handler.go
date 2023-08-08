package helpers

import (
	"net/http"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
)

func CheckHandlerSuccess(msg string) int {
	switch true {
	case strings.Contains(msg, "insert"):
		return http.StatusCreated
	case strings.Contains(msg, "read"):
		return http.StatusOK
	case strings.Contains(msg, "update"):
		return http.StatusOK
	case strings.Contains(msg, "read"):
		return http.StatusOK
	default:
		return http.StatusOK
	}
}

func CheckHandlerError(err error) (responseCode int, errConst error) {
	switch err.Error() {
	case config.JWT_InvalidJwtToken:
		return http.StatusBadRequest, err

	case config.JWT_FailedCastingJwtToken:
		return http.StatusInternalServerError, err

	case config.ErrorBindData:
		return http.StatusBadRequest, err

	case config.InvalidIdParam:
		return http.StatusBadRequest, err

	case config.InvalidPageParam:
		return http.StatusBadRequest, err

	case config.InvalidLimitParam:
		return http.StatusBadRequest, err

	case config.InvalidImageFileType:
		return http.StatusBadRequest, err

	case config.InvalidFileSize:
		return http.StatusBadRequest, err

	case config.ERR_RECORD_NOT_FOUND:
		return http.StatusBadRequest, err

	default:
		return http.StatusInternalServerError, err
	}
}
