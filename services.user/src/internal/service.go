package user

import (
	"encoding/json"
	"fmt"
	config2 "github.com/aliaydins/microservice/services.user/src/config"
	"github.com/aliaydins/microservice/services.user/src/entity"
	"github.com/aliaydins/microservice/services.user/src/event"
	jwt_helper "github.com/aliaydins/microservice/shared/jwt"
	"github.com/aliaydins/microservice/shared/rabbitmq"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Service struct {
	repo     *Repository
	rabbitmq *rabbitmq.RabbitMQ
}

func NewService(repo *Repository, rabbitmq *rabbitmq.RabbitMQ) *Service {
	return &Service{
		repo:     repo,
		rabbitmq: rabbitmq,
	}
}

func (s *Service) Register(user *entity.User) (*entity.User, error) {

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
	err = PublishUserCreatedEvent(newUser, s.rabbitmq)
	if err != nil {
		return nil, err
	}

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
		"exp": time.Now().Add(24 *
			time.Hour).Unix(),
	})

	accessToken := jwt_helper.GenerateToken(jwtClaims, secretKey)
	return user, accessToken, nil
}

func PublishUserCreatedEvent(createdUser entity.User, r *rabbitmq.RabbitMQ) error {
	config := config2.LoadConfig(".")
	userCreatedEvent := event.UserCreated{
		UserId: createdUser.ID,
		Email:  createdUser.Email,
	}

	payload, _ := json.Marshal(userCreatedEvent)
	err := r.Publish(payload, config.UserExchange, "UserCreated")
	if err != nil {
		return err
	}
	fmt.Println("UserCreated event published with id ->", userCreatedEvent.UserId)
	return nil

}
