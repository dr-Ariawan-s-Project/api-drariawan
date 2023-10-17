package service

import (
	"errors"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/google/uuid"
)

type questionaireService struct {
	questionaireData questionaire.QuestionaireDataInterface
	patientServ      patient.PatientServiceInterface
	cfg              *config.AppConfig
}

func New(repo questionaire.QuestionaireDataInterface, patientServ patient.PatientServiceInterface, cfg *config.AppConfig) questionaire.QuestionaireServiceInterface {
	return &questionaireService{
		questionaireData: repo,
		patientServ:      patientServ,
		cfg:              cfg,
	}
}

// GetDashboard implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) GetDashboard() (questionaire.DashboardCore, error) {
	var dashboardData questionaire.DashboardCore
	questAttemptCount, errQuestAttempt := service.questionaireData.CountQuestionerAttempt()
	// get data from status validated
	questAttemptNeedAssess, errQuestAttemptNeedAssess := service.questionaireData.CountAttemptByStatusAssessment(config.QUESTIONER_ATTEMPT_STATUS_VALIDATED)

	// get data from this month
	t := time.Now()
	questAttemptMonth, errQuestAttemptMonth := service.questionaireData.CountAttemptByMonth(int(t.Month()))

	if errQuestAttempt != nil || errQuestAttemptNeedAssess != nil || errQuestAttemptMonth != nil {
		return dashboardData, errQuestAttempt
	}
	dashboardData.AllQuestioner = questAttemptCount
	dashboardData.NeedAssessQuestioner = questAttemptNeedAssess
	dashboardData.MonthQuestioner = questAttemptMonth

	return dashboardData, nil

}

// Validate implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) Validate(patientData questionaire.Patient, as string, partnerEmail string) (codeAttemp string, countAttempt int, err error) {
	patientFound, errFound := service.patientServ.CheckByEmailAndPhone(patientData.Email, patientData.Phone)
	if errFound != nil {
		return "", 0, errFound
	}

	//data patient / partner already registered or found
	if patientFound.ID != "" {
		// count test attempt
		countAttempt, errCountAttempt := service.questionaireData.CountTestAttempt(patientFound.ID)
		if errCountAttempt != nil {
			return "", countAttempt, errCountAttempt
		}
		if countAttempt != 0 {
			return "", countAttempt, errors.New(config.VAL_InvalidValidation)
		}
	} else {
		data := patient.Core{
			Email: patientData.Email,
			Phone: patientData.Phone,
		}
		if as != config.QUESTIONER_ATTEMPT_SELF && as != config.QUESTIONER_ATTEMPT_PARTNER {
			return "", 0, errors.New(config.VAL_InvalidValidation)
		} else {
			dataNewPatient, errInsert := service.patientServ.Insert(data, partnerEmail)
			if errInsert != nil {
				return "", 0, errInsert
			}
			patientFound.ID = dataNewPatient.ID
		}
	}
	idAttempt := uuid.New().String()
	codeAttemp = encrypt.EncryptText(idAttempt, service.cfg.AES_GCM_SECRET)
	dataTestAttempt := questionaire.CoreAttempt{
		Id:          idAttempt,
		PatientId:   patientFound.ID,
		CodeAttempt: codeAttemp,
	}
	errTestAttempt := service.questionaireData.InsertTestAttempt(dataTestAttempt)
	if errTestAttempt != nil {
		return "", 0, errTestAttempt
	}

	//send email invitation link
	go helpers.SendMailQuestionerLink(patientData.Email, codeAttemp, service.cfg.GMAIL_APP_PASSWORD)

	return codeAttemp, 0, nil
}

// InsertAnswer implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) InsertAnswer(codeAttempt string, data []questionaire.CoreAnswer) error {
	if codeAttempt == "" {
		return errors.New(config.VAL_InvalidValidation)
	}
	// decrypt codeAttempt to idAttempt
	idAttempt := encrypt.DecryptText(codeAttempt, service.cfg.AES_GCM_SECRET)

	err := service.questionaireData.InsertAnswer(idAttempt, data)
	return err
}

// GetAll implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) GetAll() ([]questionaire.Core, error) {
	return service.questionaireData.SelectAll()
}
