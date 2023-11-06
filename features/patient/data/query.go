package data

import (
	"log"

	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type patientQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) patient.PatientDataInterface {
	return &patientQuery{
		db: db,
	}
}

// for dashboard
// CountAllPatient implements patient.PatientDataInterface.
func (repo *patientQuery) CountAllPatient() (int, error) {
	var countAttemp int64
	tx := repo.db.Model(&Patient{}).Count(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// CountPartner implements patient.PatientDataInterface.
func (repo *patientQuery) CountPartner(partnerId string) (int, error) {
	var count int64
	tx := repo.db.Model(&Patient{}).Where("partner_id = ?", partnerId).Count(&count)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}

	return int(count), nil
}

// CheckByEmailAndPhone implements patient.PatientDataInterface.
func (repo *patientQuery) CheckByEmailAndPhone(email string, phone string) (*patient.Core, error) {
	var data Patient
	tx := repo.db.Where("email = ? AND phone = ?", email, phone).Find(&data)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}

	dataCore := data.ModelToCore()
	return &dataCore, nil
}

// SelectByEmailOrPhone implements patient.PatientDataInterface.
func (repo *patientQuery) SelectByEmailOrPhone(str string) (*patient.Core, error) {
	var data Patient
	tx := repo.db.Where("email = ? OR phone = ?", str, str).Find(&data)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}

	dataCore := data.ModelToCore()
	return &dataCore, nil
}

// Delete implements patient.PatientDataInterface.
func (repo *patientQuery) Delete(id string) error {
	tx := repo.db.Where("id = ?", id).Delete(&Patient{})
	if tx.Error != nil {
		return helpers.CheckQueryErrorMessage(tx.Error)
	}
	return nil
}

// Insert implements patient.PatientDataInterface.
func (repo *patientQuery) Insert(data patient.Core) (*patient.Core, error) {
	var input = PatientCoreToModel(data)
	input.ID = uuid.New().String()

	if data.PartnerID != nil {
		input.PartnerID = data.PartnerID
	}

	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}
	dataPatient, _ := repo.SelectById(input.ID)
	return dataPatient, nil

}

// Select implements patient.PatientDataInterface.
func (repo *patientQuery) Select(search string, offset int, limit int) ([]patient.Core, error) {
	var patient []Patient
	txSelect := repo.db.Preload("Partner")
	if search != "" {
		txSelect.Where("name like ? OR email LIKE ?", "%"+search+"%", "%"+search+"%").Session(&gorm.Session{})
	}

	tx := txSelect.Offset(offset).Limit(limit).Find(&patient)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}
	dataCore := ListModelToCore(patient)
	return dataCore, nil
}

// SelectById implements patient.PatientDataInterface.
func (repo *patientQuery) SelectById(id string) (*patient.Core, error) {
	var patient Patient
	txSelect := repo.db.First(&patient, "id = ?", id)
	if txSelect.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(txSelect.Error)
	}
	dataCore := patient.ModelToCore()
	return &dataCore, nil
}

// Update implements patient.PatientDataInterface.
func (repo *patientQuery) Update(id string, data patient.Core) (*patient.Core, error) {
	var input = PatientCoreToModel(data)
	tx := repo.db.Model(&Patient{}).Where("id = ?", id).Updates(input)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}

	var patient Patient
	txSelect := repo.db.First(&patient, "id = ?", id)
	if txSelect.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(txSelect.Error)
	}
	dataCore := patient.ModelToCore()
	return &dataCore, nil
}

// SelectAllNIK implements patient.PatientDataInterface.
func (repo *patientQuery) SelectAllNIK() ([]string, error) {
	var niks []string
	txSelect := repo.db.Raw("SELECT nik FROM patients where nik != '' AND deleted_at IS NULL").Scan(&niks)
	if txSelect.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(txSelect.Error)
	}
	log.Println("niks", niks)
	return niks, nil
}
