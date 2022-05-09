package user

import (
	"fmt"
	"github.com/saaitt/user_authentication/model"
	"strconv"
	"time"
)

type SQLUserRepo interface {
	Create(user *model.User) error
	ListAll(page int, cancelled bool) ([]model.User, error)
	FindUserByNationalID(nationalID string) (*model.User, error)
	DeleteUserByID(userID string) error
	UpdateUser(user *model.User) error
}

type Service struct {
	Repo SQLUserRepo
}

func (s Service) CreateUser(req CreateUserRequest) (*model.User, error) {

	user := model.User{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		NationalID: req.NationalID,
		Cancelled:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	existingUser, err := s.Repo.FindUserByNationalID(user.NationalID)
	if err != nil {
		return nil, err
	}
	if existingUser.ID == "" {
		err := s.Repo.Create(&user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return existingUser, nil

}

func (s Service) ListUsers(pageStr string, cancelledSubStr string) ([]model.User, error) {
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Printf("there was an error converting page param string to int in List user data %+v", err)
		page = 0
	}
	if cancelledSubStr == "" {
		cancelledSubStr = "false"
	}
	cancelledSubscription, err := strconv.ParseBool(cancelledSubStr)
	if err != nil {
		fmt.Printf("there was an error converting cancelled_subscription string to boolean in List user data %+v", err)
		cancelledSubscription = false
	}
	users, err := s.Repo.ListAll(page, cancelledSubscription)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s Service) DeleteUser(userID string) error {
	if userID == "" {
		return fmt.Errorf("no user id was passed")
	}
	err := s.Repo.DeleteUserByID(userID)
	if err != nil {
		return err
	}
	return nil
}

