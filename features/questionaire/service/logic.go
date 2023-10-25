package service

import (
	"errors"

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

// for dashboard
// CountQuestionerAttempt implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) CountQuestionerAttempt() (int, error) {
	questAttemptCount, errQuestAttempt := service.questionaireData.CountQuestionerAttempt()
	return questAttemptCount, errQuestAttempt
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
		dataAttempt, countAttempt, errCountAttempt := service.questionaireData.CountTestAttempt(patientFound.ID)
		if errCountAttempt != nil {
			return "", countAttempt, errCountAttempt
		}
		// if patient / partner already take test attempt
		if countAttempt != 0 {
			countAttemptAnswer, errCountAttemptAnswer := service.questionaireData.CheckCountAttemptAnswer(patientFound.ID)
			if errCountAttemptAnswer != nil {
				return "", countAttempt, errCountAttemptAnswer
			}

			// if patient / partner have take test attempt and they have already answer questioner
			if countAttemptAnswer != 0 {
				return "", countAttempt, errors.New(config.VAL_InvalidValidation)
			}

			//else if patient / partner already take test attempt, but they haven't filled the answer yet
			//send email invitation link again
			go helpers.SendMailQuestionerLink(patientFound.Email, dataAttempt.CodeAttempt, service.cfg.GMAIL_APP_PASSWORD)
			return dataAttempt.CodeAttempt, countAttempt, nil
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

	sumAllQuestions, errSumAllQuestions := service.questionaireData.CountAllQuestion()
	if errSumAllQuestions != nil {
		return errSumAllQuestions
	}

	//validate apakah jumlah jawaban yang dikirim client sudah sama dengan banyaknya pertanyaan
	if len(data) != sumAllQuestions {
		return errors.New(config.VAL_IncompleteAnswer)
	}
	// decrypt codeAttempt to idAttempt
	idAttempt := encrypt.DecryptText(codeAttempt, service.cfg.AES_GCM_SECRET)

	err := service.questionaireData.InsertAnswer(idAttempt, data)
	return err
}

// FindTestAttempt implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) GetTestAttempt(status string, page int, perPage int) (dataAttempt []questionaire.CoreAttempt, err error) {
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
	return service.questionaireData.FindTestAttempt(status, offset, perPage)
}

// GetAllAnswerByAttempt implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) GetAllAnswerByAttempt(idAttempt string, page int, perPage int) (dataAnswer []questionaire.CoreAnswer, err error) {
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
	return service.questionaireData.FindAllAnswerByAttempt(idAttempt, offset, perPage)
}

// GetAll implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) GetAll() ([]questionaire.Core, error) {
	return service.questionaireData.SelectAll()
}
