package services

import (
	"log"
	"profesimu/dto"
	"profesimu/entity"
	"profesimu/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	// UpdateUser(user dto.UserUpdateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (s *authService) VerifyCredential(email string, password string) interface{} {
	res := s.userRepository.VerifyCredential(email, password)
	if value, ok := res.(entity.User); ok {
		comparedPassword := comparedPassword(value.Password, []byte(password))
		if value.Email == email && comparedPassword == true {
			return res
		}
		return false
	}
	return false
}

func (s *authService) CreateUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := s.userRepository.InsertUser(userToCreate)
	return res
}

func (s *authService) FindByEmail(email string) entity.User {
	return s.userRepository.FindByEmail(email)
}

func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparedPassword(hashedPass string, plainPass []byte) bool {
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
