package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xenakil/cheddartrack/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type userService struct {
	db        *gorm.DB
	jwtSecret string
}

func NewUserService(db *gorm.DB, jwtSecret string) UserService {
	return &userService{db, jwtSecret}
}

func (s *userService) Register(email, password string) error {
	var user model.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err == nil {
		return errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user = model.User{
		Email:    email,
		Password: string(hash),
	}
	return s.db.Create(&user).Error
}

func (s *userService) Login(email, password string) (string, error) {
	var user model.User

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s *userService) GetById(id uint) (*model.User, error) {

	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
