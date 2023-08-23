package data

import (
	"errors"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/auth"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type authQuery struct {
	db  *gorm.DB
	cfg *config.AppConfig
}

func New(db *gorm.DB, cfg *config.AppConfig) auth.AuthDataInterface {
	return &authQuery{
		db:  db,
		cfg: cfg,
	}
}

// CreateToken implements auth.AuthDataInterface.
func (repo *authQuery) CreateToken(id int, role string) (token string, err error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenJWT.SignedString([]byte(repo.cfg.JWT_SECRET))
}

// GetUserByEmail implements auth.AuthDataInterface.
func (repo *authQuery) GetUserByEmail(email string) (auth.UserCore, error) {
	var record auth.UserCore
	tx := repo.db.Table("users").Select("id", "email", "password", "name", "role").Where("email = ?", email).Scan(&record)
	if tx.Error != nil {
		return auth.UserCore{}, helpers.CheckQueryErrorMessage(tx.Error)
	}

	if tx.RowsAffected == 0 {
		return auth.UserCore{}, errors.New(config.DB_ERR_RECORD_NOT_FOUND)
	}
	return record, nil
}
