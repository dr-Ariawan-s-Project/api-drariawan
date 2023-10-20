package service

import (
	"errors"
	"log"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
)

type patientService struct {
	patientData patient.PatientDataInterface
	cfg         *config.AppConfig
}

func New(repo patient.PatientDataInterface, cfg *config.AppConfig) patient.PatientServiceInterface {
	return &patientService{
		patientData: repo,
		cfg:         cfg,
	}
}

// for dashboard
// CountAllPatient implements patient.PatientServiceInterface.
func (service *patientService) CountAllPatient() (int, error) {
	patientCount, errPatientCount := service.patientData.CountAllPatient()
	return patientCount, errPatientCount
}

// CountPartner implements patient.PatientServiceInterface.
func (service *patientService) CountPartner(partnerId string) (int, error) {
	if partnerId == "" {
		return 0, errors.New(config.REQ_InvalidParam)
	}
	count, err := service.patientData.CountPartner(partnerId)
	return count, err
}

// CheckByEmailAndPhone implements patient.PatientServiceInterface.
func (service *patientService) CheckByEmailAndPhone(email string, phone string) (*patient.Core, error) {
	if email == "" || phone == "" {
		return nil, errors.New(config.REQ_InvalidParam)
	}
	response, err := service.patientData.CheckByEmailAndPhone(email, phone)
	return response, err
}

// Delete implements patient.PatientServiceInterface.
func (service *patientService) Delete(id string) error {
	if id == "" {
		return errors.New(config.REQ_InvalidIdParam)
	}
	err := service.patientData.Delete(id)
	return err
}

// FindAll implements patient.PatientServiceInterface.
func (service *patientService) FindAll(search string, page int, perPage int) ([]patient.Core, error) {
	if perPage == 0 {
		perPage = 10
	}
	if page == 0 {
		page = 1
	}
	offset := (page * perPage) - perPage

	if offset < 0 {
		offset = 0
	}
	log.Println("limit", page, "offset", offset)
	response, err := service.patientData.Select(search, offset, perPage)
	return response, err
}

// FindById implements patient.PatientServiceInterface.
func (service *patientService) FindById(id string) (*patient.Core, error) {
	if id == "" {
		return nil, errors.New(config.REQ_InvalidIdParam)
	}
	response, err := service.patientData.SelectById(id)
	return response, err
}

// Insert implements patient.PatientServiceInterface.
func (service *patientService) Insert(data patient.Core, partnerEmail string) (*patient.Core, error) {
	patientDataFound, errCheckPatient := service.patientData.SelectByEmailOrPhone(data.Email)
	if errCheckPatient != nil {
		return nil, errCheckPatient
	}
	// if email already in use
	if patientDataFound.ID != "" {
		return nil, errors.New(config.DB_ERR_DUPLICATE_KEY)
	}

	if data.NIK != "" {
		nikEncrypt := encrypt.EncryptText(data.NIK, service.cfg.AES_GCM_SECRET)
		data.NIK = nikEncrypt
	}

	if data.Password != "" {
		passwordHash, errHash := encrypt.HashPassword(data.Password)
		if errHash != nil {
			return nil, errHash
		}
		data.Password = passwordHash

	}
	if partnerEmail != "" {
		patientDataFound, errCheckPatient := service.patientData.SelectByEmailOrPhone(partnerEmail)
		if errCheckPatient != nil {
			return nil, errCheckPatient
		}
		// if id patient partner not found
		if patientDataFound.ID == "" {
			return nil, errors.New("email partner not found")
		}

		// count how many partner that patient already have
		count, errCount := service.patientData.CountPartner(patientDataFound.ID)
		if errCount != nil {
			return nil, errCheckPatient
		}
		if count != 0 {
			return nil, errors.New("patient already have partner")
		}
		data.PartnerID = &patientDataFound.ID
		log.Println("partner id", &patientDataFound.ID)
	}

	dataPatient, err := service.patientData.Insert(data)
	return dataPatient, err
}

// Update implements patient.PatientServiceInterface.
func (service *patientService) Update(data patient.Core) (*patient.Core, error) {
	if data.ID == "" {
		return nil, errors.New(config.REQ_InvalidIdParam)
	}
	if data.NIK != "" {
		nikEncrypt := encrypt.EncryptText(data.NIK, service.cfg.AES_GCM_SECRET)
		data.NIK = nikEncrypt
	}
	if data.Password != "" {
		passwordHash, errHash := encrypt.HashPassword(data.Password)
		if errHash != nil {
			return nil, errHash
		}
		data.Password = passwordHash
	}
	response, err := service.patientData.Update(data.ID, data)
	return response, err
}
