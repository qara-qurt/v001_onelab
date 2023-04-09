package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository"
)

type IUser interface {
	Create(user model.UserInput) error
	SignIn(user model.SignInInput) (string, error)
	ParseToken(ctx context.Context, token string) (uint, error)
	GetByID(id int) (model.UserResponse, error)
	GetByLogin(login string) (model.User, error)
	GetAll() ([]model.UserResponse, error)
	Delete(id int) error
	Update(user model.UserResponse) error
	ChangePassword(user model.ChangePassword) error
	GetOrderUserBooks() ([]model.UserOrderBooksResponse, error)
}

type User struct {
	repo       repository.IUserRepository
	hmacSecret []byte
}

func NewUser(repo repository.IUserRepository, secret string) *User {
	return &User{
		repo:       repo,
		hmacSecret: []byte(secret),
	}
}

func (u User) Create(user model.UserInput) error {
	hashPass, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass
	return u.repo.Create(user)
}

func (u User) SignIn(user model.SignInInput) (string, error) {
	var res model.User
	res, err := u.repo.GetByLogin(user.Login)
	if err != nil {
		return "", err
	}

	if ok := checkPasswordHash(user.Password, res.Password); !ok {
		return "", model.ErrorPassword
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   strconv.Itoa(int(res.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(u.hmacSecret)
}

func (u User) ParseToken(ctx context.Context, token string) (uint, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpecting signing method: %v", token.Header["alg"])
		}

		return u.hmacSecret, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	return uint(id), nil

}

func (u User) GetByID(id int) (model.UserResponse, error) {
	return u.repo.GetByID(id)
}

func (u User) GetByLogin(login string) (model.User, error) {
	return u.repo.GetByLogin(login)
}

func (u User) GetAll() ([]model.UserResponse, error) {
	return u.repo.GetAll()
}

func (u User) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u User) Update(user model.UserResponse) error {
	return u.repo.Update(user)
}

func (u User) ChangePassword(user model.ChangePassword) error {
	res, err := u.repo.GetByLogin(user.Login)
	if err != nil {
		return err
	}
	ok := checkPasswordHash(user.CurrentPassword, res.Password)
	if !ok {
		return model.ErrorPassword
	}
	hashPass, err := hashPassword(user.NewPassword)
	if err != nil {
		return err
	}
	user.NewPassword = hashPass
	return u.repo.ChangePassword(user)
}

func (u User) GetOrderUserBooks() ([]model.UserOrderBooksResponse, error) {
	userBooks, err := u.repo.GetOrderUserBooks()
	if err != nil {
		return []model.UserOrderBooksResponse{}, err
	}

	res := combineBooks(userBooks)
	return res, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func combineBooks(userBooks []model.UserOrderBooks) []model.UserOrderBooksResponse {
	data := make(map[uint]*model.UserOrderBooksResponse)
	for _, ub := range userBooks {
		if _, ok := data[ub.ID]; !ok {
			data[ub.ID] = &model.UserOrderBooksResponse{
				ID:       ub.ID,
				Login:    ub.Login,
				FullName: ub.FullName,
				Book:     []model.BookWithDate{},
			}
		}
		book := model.BookWithDate{
			ID:          ub.Book.ID,
			Name:        ub.Book.Name,
			Description: ub.Book.Description,
			Author:      ub.Book.Author,
			OrderDate:   ub.Book.OrderDate,
			ReturnDate:  ub.Book.ReturnDate,
		}
		data[ub.ID].AddBook(book)
	}

	result := make([]model.UserOrderBooksResponse, 0, len(data))
	for _, user := range data {
		result = append(result, *user)
	}
	return result
}
