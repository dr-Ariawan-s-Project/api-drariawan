package questionaire

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

type QuestionaireDataInterface interface {
	SelectAll() ([]Core, error)
}

type QuestionaireServiceInterface interface {
	GetAll() ([]Core, error)
}
