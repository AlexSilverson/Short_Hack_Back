package entityMentor

type Mentor struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Password     string `json:"password"`
	Skills       []uint
	Telegram     string   `json:"telegram"`
	Diploma      string   `json:"diploma"`
	Photo        string   `json:"photo"`
	FieldOfStudy string   `json:"fieldofstudy"`
	Feedback     []string `json:"feedback"`
}
