package helpers

import (
	"errors"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"gorm.io/gorm"
)

func CheckQueryErrorMessage(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(config.DB_ERR_RECORD_NOT_FOUND)
	} else if errors.Is(err, gorm.ErrMissingWhereClause) {
		return errors.New(config.DB_ERR_MISSING_WHERE_CLAUSE)
	} else if errors.Is(err, gorm.ErrUnsupportedRelation) {
		return errors.New(config.DB_ERR_UNSUPPORTED_RELATION)
	} else if errors.Is(err, gorm.ErrInvalidData) {
		return errors.New(config.DB_ERR_INVALID_DATA)
	} else if errors.Is(err, gorm.ErrInvalidField) {
		return errors.New(config.DB_ERR_INVALID_FIELD)
	} else if errors.Is(err, gorm.ErrPreloadNotAllowed) {
		return errors.New(config.DB_ERR_PRELOAD_NOT_ALLOWED)
	} else if errors.Is(err, gorm.ErrInvalidDB) {
		return errors.New(config.DB_ERR_INVALID_DB)
	} else if errors.Is(err, gorm.ErrPrimaryKeyRequired) {
		return errors.New(config.DB_ERR_PRIMARY_KEY_REQUIRED)
	} else {
		return err
	}
}
