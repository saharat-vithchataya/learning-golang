package services

import (
	"banking/logs"
	"banking/repository"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte("super-shy-secret-and-change-me")

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{repo: repo}
}

func (srv userService) SignUp(request SignUpRequest) (*SignUpResponse, error) {
	if request.Username == "" || request.Password == "" {
		logs.Error("Invalid credentials")
		return nil, errors.New("invalid credentials")
	}
	// hashedPassword := bc
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	newUser, err := srv.repo.Create(repository.User{
		Username: request.Username,
		Password: string(hashedPassword),
	})
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &SignUpResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
	}, nil
}

func (srv userService) Login(request LoginRequest) (*LoginResponse, error) {
	user, err := srv.repo.GetByUsername(request.Username)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("invalid credentials")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("invalid credentials")
	}

	// Create a new token using HMAC method
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["issuer"] = strconv.Itoa(user.ID)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &LoginResponse{
		AccessToken: tokenString,
	}, nil
}
