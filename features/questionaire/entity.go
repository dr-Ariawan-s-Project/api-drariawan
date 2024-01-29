package questionaire

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
)

type Core struct {
	Id          uint
	Type        string
	Question    string
	Description string
	UrlVideo    string
	Section     string
	Goto        *uint
	Choices     []Choice
}

type Choice struct {
	Id         uint
	QuestionId uint
	Option     string
	Slugs      string
	Score      int
	Goto       *uint
}

type CoreAnswer struct {
	Id          string
	AttemptId   string
	QuestionId  uint
	Description string
	Score       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Question    Core
}

type CoreAttempt struct {
	Id            string
	PatientId     string
	CodeAttempt   string
	NotesAttempt  string
	Score         int
	AIAccuracy    float64
	AIProbability float64
	AIDiagnosis   string
	Diagnosis     string
	Feedback      string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Patient       Patient
}

type Patient struct {
	ID             string
	Name           string
	Email          string
	Password       string
	NIK            string
	DOB            *time.Time
	Phone          string
	Gender         *string
	MarriageStatus string
	Nationality    string
	PartnerID      *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type QuestionaireDataInterface interface {
	SelectAll() ([]Core, error)
	InsertAnswer(idAttempt string, data []CoreAnswer) error
	CountTestAttempt(patientId string) (dataAttempt CoreAttempt, count int, err error)
	CheckCountAttemptAnswer(patientId string) (int, error)
	InsertTestAttempt(data CoreAttempt) error
	FindTestAttempt(status string, offset int, limit int) (dataAttempt []CoreAttempt, err error)
	FindTestAttemptById(id string) (dataAttempt *CoreAttempt, err error)
	FindAllAnswerByAttempt(idAttemptattempt_id string, offset int, limit int) (dataAnswer []CoreAnswer, err error)
	InsertAssesment(data CoreAttempt) error
	CountAllQuestion() (int, error)
	CountQuestionerAttempt() (int, error)
	CountTestAttemptByFilter(status string) (int64, error)
}

type QuestionaireServiceInterface interface {
	GetAll() ([]Core, error)
	InsertAnswer(codeAttempt string, data []CoreAnswer) error
	Validate(patient Patient, as string, partnerEmail string) (codeAttempt string, countAttempt int, err error)
	GetTestAttempt(status string, page int, perPage int) (dataAttempt []CoreAttempt, err error)
	GetTestAttemptById(id string) (dataAttempt *CoreAttempt, err error)
	GetAllAnswerByAttempt(idAttempt string, page int, perPage int) (dataAnswer []CoreAnswer, err error)
	InsertAssesment(data CoreAttempt) error
	CountQuestionerAttempt() (int, error)
	GetPaginationTestAttempt(status string, page int, perPage int) (helpers.Pagination, error)
}
