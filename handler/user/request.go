package user

type CreateUserRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	NationalID string `json:"national_id"`
}
