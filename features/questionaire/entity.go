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

type QuestionaireDataInterface interface {
	SelectAll() ([]Core, error)
	InsertAnswer(idAttempt string, data []CoreAnswer) error
}

type QuestionaireServiceInterface interface {
	GetAll() ([]Core, error)
	InsertAnswer(codeAttempt string, data []CoreAnswer) error
}
