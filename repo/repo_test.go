package repo

import (
	"github.com/saaitt/user_authentication/model"
	"testing"
)

func TestDbConnection(t *testing.T) {
	db, err := InitDb()
	if err != nil {
		t.Errorf("test failed, got an error initiating db : %v \n", err )
	}
	var sum int
	db.Raw("SELECT 1+1 AS sum").Scan(&sum)
	if sum != 2 {
		t.Errorf("test failed, expected 2 got : %v \n", sum )
	}
}

func TestUserCreationService(t *testing.T) {
	db, _ := InitDb()
	s := SQLUserRepo{DB: db}
	testUser := model.User{
		FirstName:  "testing fname repo",
		LastName:   "testing lname repo",
		NationalID: "testing nID repo",
		Cancelled: false,
	}

	err := s.Create(&testUser)
	if err != nil {
		t.Errorf("test failed with err: %v  \n", err)
	}
	if testUser.ID == "" {
		t.Errorf("test failed, expecated user to have an id \n" )
	}

}
