package usersSearch

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	IPAddress   string `json:"id_address"`
}
