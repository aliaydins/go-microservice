package user

import (
	"github.com/aliaydins/microservice/services.user/src/entity"
	jwt_helper "github.com/aliaydins/microservice/services.user/src/pkg/jwt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SignUp(user *entity.User) (*entity.User, error) {

	u, _ := s.repo.FindByEmail(user.Email)
	if u != nil {
		return nil, ErrUserExist
	}

	newUser := entity.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  jwt_helper.GenerateHash(user.Password, "salt"),
	}

	usr, err := s.repo.Create(&newUser)
	if err != nil {
		return user, err
	}

	// create new wallet this user,  publish event here for wallet service

	return usr, nil
}
func (s *Service) ValidateUser(email string, password string, secretKey string) (*entity.User, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, "", ErrUserNotFound
	}
	if user.Password != jwt_helper.GenerateHash(password, "salt") {
		return nil, "", ErrPasswordIncorrect
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"email":     user.Email,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"iat":       time.Now().Unix(),
		"exp": time.Now().Add(12 *
			time.Hour).Unix(),
	})

	accessToken := jwt_helper.GenerateToken(jwtClaims, secretKey)
	return user, accessToken, nil
}
