package repo

import (
	"fmt"
	"github.com/saaitt/user_authentication/model"
)
import "gorm.io/gorm"

type SQLUserRepo struct {
	DB *gorm.DB
}

func (s SQLUserRepo) Create(user *model.User) error {
	if err := s.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s SQLUserRepo) ListAllWithFilter(page int, cancelled bool) ([]model.User, error) {
	users := []model.User{}
	if err := s.DB.Limit(10).Offset(page*10).Where("cancelled=?", cancelled).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s SQLUserRepo) ListAll(page int) ([]model.User, error) {
	users := []model.User{}
	if err := s.DB.Limit(10).Offset(page*10).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s SQLUserRepo) FindUserByNationalID(nationalID string) (*model.User, error) {
	user := model.User{}
	if err := s.DB.Where("national_id=?", nationalID).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s SQLUserRepo) DeleteUserByID(userID string) error {
	result := s.DB.Where("id = ?", userID).Delete(&model.User{})
	if result.RowsAffected < 1 {
		return fmt.Errorf("userID with id=%v cannot be deleted because it doesn't exist", userID)
	}
	return nil
}

func (s SQLUserRepo) SetUserCancelledToTrue(userID string) error {
	s.DB.Model(&model.User{}).Where("id = ?", userID).Update("cancelled", "true")
	return nil
}
