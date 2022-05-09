package user

import (
	"github.com/saaitt/user_authentication/repo"
	"testing"
)

func TestUserCreationService(t *testing.T) {
	db, _ := repo.InitDb()
	s := Service{
		Repo: &repo.SQLUserRepo{
			DB: db,
		},
	}
	testUser := CreateUserRequest{
		FirstName:  "testing fname",
		LastName:   "testing lname",
		NationalID: "testing nID",
	}
	user, err := s.CreateUser(testUser)
	if err != nil {
		t.Errorf("test failed with err: %v  \n", err)
	}
	if user.ID == "" {
		t.Errorf("test failed, expecated user to have an id \n" )
	}

}
