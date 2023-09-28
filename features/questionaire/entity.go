package questionaire

import "time"

type Core struct {
	Id          uint
	Type        string
	Question    string
	Description string
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
}

type CoreAttempt struct {
	Id           string
	PatientId    string
	CodeAttempt  string
	NotesAttempt string
	Score        int
	Feedback     string
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
	CountTestAttempt(patientId string) (int, error)
	InsertTestAttempt(data CoreAttempt) error
}

type QuestionaireServiceInterface interface {
	GetAll() ([]Core, error)
	InsertAnswer(codeAttempt string, data []CoreAnswer) error
	Validate(patient Patient, as string, partnerEmail string) (codeAttempt string, countAttempt int, err error)
}
