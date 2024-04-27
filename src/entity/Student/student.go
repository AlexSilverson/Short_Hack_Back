package entityStudent

type Student struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"secondname"`
	Telegram  string `json:"telegram"`
	Password  string `json:"password"`
	StudentID uint   `json:"studentid"` //студенческий
	MentorId  uint   `json:"mentorid"`
}
