package service

import (
	"errors"
	"testing"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"github.com/dr-ariawan-s-project/api-drariawan/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetPagination(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()

	t.Run("Success Get Pagination", func(t *testing.T) {
		totalRows := int64(20)
		page := 1
		perPage := 10
		search := ""
		patientRepo.On("CountByFilter", search).Return(totalRows, nil).Once()
		patientService := New(patientRepo, cfg)
		paginationRes, err := patientService.GetPagination(search, page, perPage)
		assert.NoError(t, err)
		assert.Equal(t, 10, paginationRes.Limit)
		assert.Equal(t, 1, paginationRes.Page)
		assert.Equal(t, 2, paginationRes.TotalPages)
		assert.Equal(t, int64(20), paginationRes.TotalRecords)
		patientRepo.AssertExpectations(t)
	})

	t.Run("Error Count by filter", func(t *testing.T) {
		page := 1
		perPage := 10
		search := ""
		patientRepo.On("CountByFilter", search).Return(int64(0), errors.New("error get count by filter")).Once()
		patientService := New(patientRepo, cfg)
		_, err := patientService.GetPagination(search, page, perPage)
		assert.Error(t, err)
		patientRepo.AssertExpectations(t)
	})

	t.Run("if perPage = 0", func(t *testing.T) {
		totalRows := int64(20)
		page := 1
		perPage := 0
		search := ""
		patientRepo.On("CountByFilter", search).Return(totalRows, nil).Once()
		patientService := New(patientRepo, cfg)
		paginationRes, err := patientService.GetPagination(search, page, perPage)
		assert.NoError(t, err)
		assert.Equal(t, 0, paginationRes.Limit)
		assert.Equal(t, 1, paginationRes.TotalPages)
		patientRepo.AssertExpectations(t)
	})
}

func TestCountAllPatient(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()
	t.Run("Get Count All Patient", func(t *testing.T) {
		patientCount := 20
		patientRepo.On("CountAllPatient").Return(patientCount, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.CountAllPatient()
		assert.NoError(t, err)
		assert.Equal(t, patientCount, response)
		patientRepo.AssertExpectations(t)
	})
}

func TestCountPartner(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()
	t.Run("test Empty partnerid", func(t *testing.T) {

		patientService := New(patientRepo, cfg)
		response, err := patientService.CountPartner("")
		assert.Error(t, err)
		assert.Equal(t, 0, response)
	})
	t.Run("success count partner", func(t *testing.T) {
		patientRepo.On("CountPartner", "PATIENT0001").Return(20, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.CountPartner("PATIENT0001")
		assert.NoError(t, err)
		assert.Equal(t, 20, response)
		patientRepo.AssertExpectations(t)
	})
}

func TestCheckByEmailAndPhone(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()
	t.Run("test Empty email or phone", func(t *testing.T) {

		patientService := New(patientRepo, cfg)
		response, err := patientService.CheckByEmailAndPhone("", "")
		assert.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("success check email and phone", func(t *testing.T) {
		responseData := patient.Core{
			ID:       "PATIENT0001",
			Name:     "Test Patient",
			Email:    "test@alterra.id",
			Password: "qwerty",
			NIK:      "1234567",
			Phone:    "08123456",
		}
		patientRepo.On("CheckByEmailAndPhone", "test@alterra.id", "08123456").Return(&responseData, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.CheckByEmailAndPhone("test@alterra.id", "08123456")
		assert.NoError(t, err)
		assert.Equal(t, responseData.Email, response.Email)
		patientRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()
	t.Run("test Empty id", func(t *testing.T) {

		patientService := New(patientRepo, cfg)
		err := patientService.Delete("")
		assert.Error(t, err)
	})
	t.Run("success delete patient", func(t *testing.T) {
		patientRepo.On("Delete", "PATIENT0001").Return(nil).Once()
		patientService := New(patientRepo, cfg)
		err := patientService.Delete("PATIENT0001")
		assert.NoError(t, err)
		patientRepo.AssertExpectations(t)
	})
}

func TestFindAll(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()
	t.Run("if queryParam page 0 and perPage 0", func(t *testing.T) {
		page := 0
		perPage := 0
		responseData := []patient.Core{
			{
				ID:       "PATIENT0001",
				Name:     "Test Patient",
				Email:    "test@alterra.id",
				Password: "qwerty",
				NIK:      "1234567",
				Phone:    "08123456",
			},
		}
		patientRepo.On("Select", "", page, perPage).Return(responseData, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.FindAll("", 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, responseData[0].Email, response[0].Email)
		patientRepo.AssertExpectations(t)
	})
}

func TestFindById(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()
	t.Run("test Empty id", func(t *testing.T) {
		patientService := New(patientRepo, cfg)
		response, err := patientService.FindById("")
		assert.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("success findbyid patient", func(t *testing.T) {
		responseData := patient.Core{
			ID:       "PATIENT0001",
			Name:     "Test Patient",
			Email:    "test@alterra.id",
			Password: "qwerty",
			NIK:      "1234567",
			Phone:    "08123456",
		}
		patientRepo.On("SelectById", "PATIENT0001").Return(&responseData, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.FindById("PATIENT0001")
		assert.NoError(t, err)
		assert.Equal(t, responseData.ID, response.ID)
		patientRepo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()
	responseData := patient.Core{
		ID:       "PATIENT0001",
		Name:     "Test Patient",
		Email:    "test@alterra.id",
		Password: "qwerty",
		NIK:      "1234567890",
		Phone:    "08123456",
	}

	// untuk data yang not found
	emptyResponseData := patient.Core{}

	inputData := patient.Core{
		ID:       "PATIENT0001",
		Name:     "Test Patient",
		Email:    "test@alterra.id",
		Password: "qwerty",
		NIK:      "1234567890",
		Phone:    "08123456",
	}
	partnerInputData := patient.Core{
		ID:       "PARTNER0001",
		Name:     "Test Patient Partner",
		Email:    "testpartner@alterra.id",
		Password: "qwerty",
		NIK:      "12345678901",
		Phone:    "081234561",
	}

	var partnerIDDummy = "PATIENT0001"
	partnerResponseDataSuccess := patient.Core{
		ID:        "PARTNER0001",
		Name:      "Test Patient Partner",
		Email:     "testpartner@alterra.id",
		Phone:     "081234561",
		PartnerID: &partnerIDDummy,
	}
	t.Run("1. Error when check patient", func(t *testing.T) {
		patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(nil, errors.New("data not found")).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.Insert(inputData, "")
		assert.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("2. Error email already in use", func(t *testing.T) {
		patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(&responseData, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.Insert(inputData, "")
		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("3. error duplicate NIK", func(t *testing.T) {
		responseNIK := []string{
			"r7liBB3MWJRtPxRw+NsyuGCpZjJGf5VaK1FsP2EL+ALlxnBr4HY=", // hasil encrypt dari 1234567890
		}
		patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(&emptyResponseData, nil).Once()
		patientRepo.On("SelectAllNIK").Return(responseNIK, nil).Once()

		patientService := New(patientRepo, cfg)
		response, err := patientService.Insert(inputData, "")
		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("4. NIK not duplicated, but check partner email error", func(t *testing.T) {
		responseNIK := []string{
			"1",
		}
		patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(&emptyResponseData, nil).Once()
		patientRepo.On("SelectAllNIK").Return(responseNIK, nil).Once()
		patientRepo.On("SelectByEmailOrPhone", "testpartner@alterra.id").Return(&emptyResponseData, errors.New("error check/search partner email")).Once()

		patientService := New(patientRepo, cfg)
		response, err := patientService.Insert(inputData, "testpartner@alterra.id")
		assert.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("5. NIK not duplicated, but check patient partner email not found", func(t *testing.T) {
		responseNIK := []string{
			"1",
		}
		patientRepo.On("SelectByEmailOrPhone", "testpartner@alterra.id").Return(&emptyResponseData, nil).Once()
		patientRepo.On("SelectAllNIK").Return(responseNIK, nil).Once()
		patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(&emptyResponseData, nil).Once()
		// patientRepo.On("CountPartner", "PARTNER0001").Return(1, nil).Once()
		// patientRepo.On("Insert", inputData).Return(&responseData, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.Insert(partnerInputData, "test@alterra.id")
		assert.Error(t, err)
		assert.Nil(t, response)
	})
	// t.Run("6. NIK not duplicated, and check patient partner email found, then error when count partner", func(t *testing.T) {
	// 	responseNIK := []string{
	// 		"1",
	// 	}
	// 	patientRepo.On("SelectByEmailOrPhone", "testpartner@alterra.id").Return(&emptyResponseData, nil).Once()
	// 	patientRepo.On("SelectAllNIK").Return(responseNIK, nil).Once()
	// 	patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(&responseData, nil).Once()
	// 	patientRepo.On("CountPartner", responseData.ID).Return(int(0), errors.New("error count partner patient")).Once()

	// 	patientService := New(patientRepo, cfg)
	// 	_, err := patientService.Insert(partnerInputData, "test@alterra.id")
	// 	assert.Error(t, err)
	// })
	t.Run("7. NIK not duplicated, and check patient partner email found, then patient already have partner", func(t *testing.T) {
		responseNIK := []string{
			"1",
		}
		patientRepo.On("SelectByEmailOrPhone", "testpartner@alterra.id").Return(&emptyResponseData, nil).Once()
		patientRepo.On("SelectAllNIK").Return(responseNIK, nil).Once()
		patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(&responseData, nil).Once()
		patientRepo.On("CountPartner", responseData.ID).Return(int(1), nil).Once()
		// patientRepo.On("Insert", inputData).Return(&responseData, nil).Once()
		patientService := New(patientRepo, cfg)
		_, err := patientService.Insert(partnerInputData, "test@alterra.id")
		assert.Error(t, err)
		// assert.Nil(t, response)
	})
	t.Run("8. success", func(t *testing.T) {
		responseNIK := []string{
			"1",
		}
		patientRepo.On("SelectByEmailOrPhone", "testpartner@alterra.id").Return(&emptyResponseData, nil).Once()
		patientRepo.On("SelectAllNIK").Return(responseNIK, nil).Once()
		patientRepo.On("SelectByEmailOrPhone", "test@alterra.id").Return(&responseData, nil).Once()
		patientRepo.On("CountPartner", responseData.ID).Return(int(0), nil).Once()
		patientRepo.On("Insert", partnerResponseDataSuccess).Return(&partnerResponseDataSuccess, nil).Once()
		patientService := New(patientRepo, cfg)
		response, err := patientService.Insert(partnerResponseDataSuccess, "test@alterra.id")
		assert.Nil(t, err)
		assert.Equal(t, partnerResponseDataSuccess.Email, response.Email)
	})

}

func TestUpdate(t *testing.T) {
	patientRepo := new(mocks.PatientData)
	cfg := config.InitConfig()

	t.Run("error id null", func(t *testing.T) {
		inputData := patient.Core{
			Name:     "Test Patient",
			Email:    "test@alterra.id",
			Password: "qwerty",
			NIK:      "1234567890",
			Phone:    "08123456",
		}
		patientService := New(patientRepo, cfg)
		response, err := patientService.Update(inputData)
		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("success update", func(t *testing.T) {
		inputData := patient.Core{
			ID:    "PATIENT0001",
			Name:  "Test Patient",
			Email: "test@alterra.id",
			Phone: "08123456",
		}

		responseData := patient.Core{
			ID:    "PATIENT0001",
			Name:  "Test Patient",
			Email: "test@alterra.id",
			Phone: "08123456",
		}

		patientRepo.On("Update", inputData.ID, inputData).Return(&responseData, nil).Once()

		patientService := New(patientRepo, cfg)
		response, err := patientService.Update(inputData)
		assert.NoError(t, err)
		assert.Equal(t, inputData.Name, response.Name)
	})

	t.Run("error NIK duplicate", func(t *testing.T) {
		inputData := patient.Core{
			ID:       "PATIENT0001",
			Name:     "Test Patient",
			Email:    "test@alterra.id",
			NIK:      "1234567890",
			Password: "qwerty",
			Phone:    "08123456",
		}

		responseData := patient.Core{
			ID:    "PATIENT0001",
			Name:  "Test Patient",
			Email: "test@alterra.id",
			Phone: "08123456",
		}

		responseNIK := []string{
			"r7liBB3MWJRtPxRw+NsyuGCpZjJGf5VaK1FsP2EL+ALlxnBr4HY=", // hasil encrypt dari 1234567890
		}

		patientRepo.On("SelectAllNIK").Return(responseNIK, nil).Once()
		patientRepo.On("Update", inputData.ID, inputData).Return(&responseData, nil).Once()

		patientService := New(patientRepo, cfg)
		response, err := patientService.Update(inputData)
		assert.Error(t, err)
		assert.Nil(t, response)
	})

}
