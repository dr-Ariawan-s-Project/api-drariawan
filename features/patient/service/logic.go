package service

import "github.com/dr-ariawan-s-project/api-drariawan/features/patient"

type patientService struct {
	patientData patient.PatientDataInterface
}

func New(repo patient.PatientDataInterface) patient.PatientServiceInterface {
	return &patientService{
		patientData: repo,
	}
}

// CheckByEmailOrPhone implements patient.PatientServiceInterface.
func (service *patientService) CheckByEmailOrPhone(email string, phone string) (*patient.Core, error) {
	panic("unimplemented")
}

// Delete implements patient.PatientServiceInterface.
func (service *patientService) Delete(id string) error {
	panic("unimplemented")
}

// FindAll implements patient.PatientServiceInterface.
func (service *patientService) FindAll(search string, page int, perPage int) ([]patient.Core, error) {
	panic("unimplemented")
}

// FindById implements patient.PatientServiceInterface.
func (service *patientService) FindById(id string) (*patient.Core, error) {
	panic("unimplemented")
}

// Insert implements patient.PatientServiceInterface.
func (service *patientService) Insert(data patient.Core) error {
	panic("unimplemented")
}

// Update implements patient.PatientServiceInterface.
func (service *patientService) Update(data patient.Core) (*patient.Core, error) {
	panic("unimplemented")
}
