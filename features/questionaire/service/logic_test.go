package service

import (
	"testing"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"github.com/dr-ariawan-s-project/api-drariawan/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInsertAnswer(t *testing.T) {
	repo := new(mocks.QuestionaireData)
	idAttempt := "TEST-0001"
	insertData := []questionaire.CoreAnswer{{
		QuestionId:  1,
		Description: "Ya",
		Score:       0,
	},
	}

	t.Run("Success InsertAnswer", func(t *testing.T) {
		repo.On("InsertAnswer", idAttempt, insertData).Return(nil).Once()
		srv := New(repo, config.InitConfig())
		codeAttempt := "dFL1UYyVMGuVBJeIuuCoICkqeeanN8NKFT459RojjXCWVVDLyQ=="
		err := srv.InsertAnswer(codeAttempt, insertData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed InsertAnswer. Empty codeAttemp", func(t *testing.T) {
		srv := New(repo, config.InitConfig())
		codeAttempt := ""
		err := srv.InsertAnswer(codeAttempt, insertData)
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})

}

func TestGetAll(t *testing.T) {
	repo := new(mocks.QuestionaireData)
	responseData := []questionaire.Core{
		{
			Id:          1,
			Type:        "text",
			Question:    "https://linkto.com/video.mp4",
			Description: "Berapa umur anda",
			Goto:        nil,
		},
	}
	t.Run("Success GetAll Question", func(t *testing.T) {
		repo.On("SelectAll").Return(responseData, nil).Once()
		srv := New(repo, config.InitConfig())
		response, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, responseData[0].Id, response[0].Id)
		assert.Equal(t, responseData[0].Description, response[0].Description)
		repo.AssertExpectations(t)
	})
}
