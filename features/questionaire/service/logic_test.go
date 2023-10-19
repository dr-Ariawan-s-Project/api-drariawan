package service

import (
	"errors"
	"testing"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"github.com/dr-ariawan-s-project/api-drariawan/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestValidate(t *testing.T) {
	patientServ := new(mocks.PatientService)
	questionaireRepo := new(mocks.QuestionaireData)

	inputData := questionaire.Patient{
		Email: "test@mail.com",
		Phone: "081234",
	}

	t.Run("Error search patient email and phone", func(t *testing.T) {
		patientServ.On("CheckByEmailAndPhone", inputData.Email, inputData.Phone).Return(nil, errors.New("error search email and phone")).Once()
		questionaireService := New(questionaireRepo, patientServ, config.InitConfig())
		inputAs := "myself"
		inputPartnerEmail := ""
		respCodeAttempt, respCountAttempt, err := questionaireService.Validate(inputData, inputAs, inputPartnerEmail)
		assert.Error(t, err)
		assert.Equal(t, "", respCodeAttempt)
		assert.Equal(t, 0, respCountAttempt)
	})

	t.Run("if data patient/partner already registered/found. but error when count test attempt based on id patient", func(t *testing.T) {
		responseData := &patient.Core{
			ID:    "0001",
			Email: "test@mail.com",
			Phone: "081234",
		}

		patientServ.On("CheckByEmailAndPhone", inputData.Email, inputData.Phone).Return(responseData, nil).Once()
		questionaireRepo.On("CountTestAttempt", responseData.ID).Return(questionaire.CoreAttempt{}, 0, errors.New("error count test attempt")).Once()
		questionaireService := New(questionaireRepo, patientServ, config.InitConfig())
		inputAs := "myself"
		inputPartnerEmail := ""
		respCodeAttempt, respCountAttempt, err := questionaireService.Validate(inputData, inputAs, inputPartnerEmail)
		assert.Error(t, err)
		assert.Equal(t, "", respCodeAttempt)
		assert.Equal(t, 0, respCountAttempt)
	})
	t.Run("if data patient/partner already registered/found. but already fill test", func(t *testing.T) {
		responseData := &patient.Core{
			ID:    "0001",
			Email: "test@mail.com",
			Phone: "081234",
		}

		resultTestAttempt := questionaire.CoreAttempt{
			Id:           "TEST-0001",
			PatientId:    "0001",
			CodeAttempt:  "CODE-0001",
			NotesAttempt: "notes",
			Score:        0,
			Feedback:     "feedback",
		}
		patientServ.On("CheckByEmailAndPhone", inputData.Email, inputData.Phone).Return(responseData, nil).Once()
		questionaireRepo.On("CountTestAttempt", responseData.ID).Return(resultTestAttempt, 1, nil).Once()
		questionaireRepo.On("CheckCountAttemptAnswer", responseData.ID).Return(1, nil).Once()
		questionaireService := New(questionaireRepo, patientServ, config.InitConfig())
		inputAs := "myself"
		inputPartnerEmail := ""
		respCodeAttempt, respCountAttempt, err := questionaireService.Validate(inputData, inputAs, inputPartnerEmail)
		assert.Error(t, err)
		assert.Equal(t, "", respCodeAttempt)
		assert.Equal(t, 1, respCountAttempt)
	})

	t.Run("data patient not found. but input 'as' not valid", func(t *testing.T) {
		var responseData = new(patient.Core)
		patientServ.On("CheckByEmailAndPhone", inputData.Email, inputData.Phone).Return(responseData, nil).Once()
		questionaireService := New(questionaireRepo, patientServ, config.InitConfig())
		inputAs := "pribadi"
		inputPartnerEmail := ""
		respCodeAttempt, respCountAttempt, err := questionaireService.Validate(inputData, inputAs, inputPartnerEmail)
		assert.Error(t, err)
		assert.Equal(t, "", respCodeAttempt)
		assert.Equal(t, 0, respCountAttempt)
	})
	t.Run("data patient not found. input 'as' valid. but error when insert patient data", func(t *testing.T) {
		var responseDataFound = new(patient.Core)
		var inputDataPatient = patient.Core{
			Email: "test@mail.com",
			Phone: "081234",
		}

		inputAs := "myself"
		inputPartnerEmail := ""
		patientServ.On("CheckByEmailAndPhone", inputData.Email, inputData.Phone).Return(responseDataFound, nil).Once()
		patientServ.On("Insert", inputDataPatient, inputPartnerEmail).Return(nil, errors.New("error insert patient data to database")).Once()
		questionaireService := New(questionaireRepo, patientServ, config.InitConfig())
		respCodeAttempt, respCountAttempt, err := questionaireService.Validate(inputData, inputAs, inputPartnerEmail)
		assert.Error(t, err)
		assert.Equal(t, "", respCodeAttempt)
		assert.Equal(t, 0, respCountAttempt)
	})
	t.Run("data patient not found. input 'as' valid. but failed to insert test attempt", func(t *testing.T) {
		var responseDataFound = new(patient.Core)
		var inputDataPatient = patient.Core{
			Email: "test@mail.com",
			Phone: "081234",
		}
		var responseInsertDataPatient = &patient.Core{
			ID:    "0001",
			Email: "test@mail.com",
			Phone: "081234",
		}

		inputAs := "myself"
		inputPartnerEmail := ""
		patientServ.On("CheckByEmailAndPhone", inputData.Email, inputData.Phone).Return(responseDataFound, nil).Once()
		patientServ.On("Insert", inputDataPatient, inputPartnerEmail).Return(responseInsertDataPatient, nil).Once()
		questionaireRepo.On("InsertTestAttempt", mock.Anything).Return(errors.New("error insert test attempt data")).Once()

		questionaireService := New(questionaireRepo, patientServ, config.InitConfig())
		_, _, err := questionaireService.Validate(inputData, inputAs, inputPartnerEmail)
		assert.Error(t, err)
	})
	t.Run("data patient not found. input 'as' valid. and insert success", func(t *testing.T) {
		var responseDataFound = new(patient.Core)
		var inputDataPatient = patient.Core{
			Email: "test@mail.com",
			Phone: "081234",
		}
		var responseInsertDataPatient = &patient.Core{
			ID:    "0001",
			Email: "test@mail.com",
			Phone: "081234",
		}

		inputAs := "myself"
		inputPartnerEmail := ""
		patientServ.On("CheckByEmailAndPhone", inputData.Email, inputData.Phone).Return(responseDataFound, nil).Once()
		patientServ.On("Insert", inputDataPatient, inputPartnerEmail).Return(responseInsertDataPatient, nil).Once()
		questionaireRepo.On("InsertTestAttempt", mock.Anything).Return(nil).Once()

		questionaireService := New(questionaireRepo, patientServ, config.InitConfig())
		_, _, err := questionaireService.Validate(inputData, inputAs, inputPartnerEmail)
		assert.NoError(t, err)
	})

}

func TestInsertAnswer(t *testing.T) {
	repo := new(mocks.QuestionaireData)
	patientService := new(mocks.PatientService)
	idAttempt := "TEST-0001"
	insertData := []questionaire.CoreAnswer{{
		QuestionId:  1,
		Description: "Ya",
		Score:       0,
	},
	}

	t.Run("Success InsertAnswer", func(t *testing.T) {
		repo.On("InsertAnswer", idAttempt, insertData).Return(nil).Once()
		repo.On("CountAllQuestion").Return(1, nil).Once()
		srv := New(repo, patientService, config.InitConfig())
		codeAttempt := "dFL1UYyVMGuVBJeIuuCoICkqeeanN8NKFT459RojjXCWVVDLyQ=="
		err := srv.InsertAnswer(codeAttempt, insertData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed InsertAnswer. Empty codeAttemp", func(t *testing.T) {
		srv := New(repo, patientService, config.InitConfig())
		codeAttempt := ""
		err := srv.InsertAnswer(codeAttempt, insertData)
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})

}

func TestGetAll(t *testing.T) {
	repo := new(mocks.QuestionaireData)
	patientService := new(mocks.PatientService)
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
		srv := New(repo, patientService, config.InitConfig())
		response, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, responseData[0].Id, response[0].Id)
		assert.Equal(t, responseData[0].Description, response[0].Description)
		repo.AssertExpectations(t)
	})
}
