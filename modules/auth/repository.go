package auth

import (
	"boilerplate/models"
	"boilerplate/utils"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Repository interface {
	Register(input *models.User) error
	Login(input *models.User) (*models.User, error)
	SaveToken(key string, value interface{}, expiration time.Duration) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db: db, redis: redis}
}

func (r *repository) Register(user *models.User) error {
	db := r.db.Model(&user).Session(&gorm.Session{SkipDefaultTransaction: true})

	if db.First(&user, "email = ?", user.Email).RowsAffected > 0 {
		return errors.New("email registered")
	}

	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) Login(input *models.User) (*models.User, error) {
	var user models.User
	db := r.db.Model(&user)

	user.Email = input.Email
	user.Password = input.Password

	if db.First(&user, "email = ?", input.Email).RowsAffected == 0 {
		return nil, errors.New("User account is not registered")
	}

	if !user.Active {
		return nil, errors.New("User account is not active")
	}

	if err := utils.CheckPassword(input.Password, user.Password); err != nil {
		return nil, errors.New("Username or password is wrong")
	}

	return &user, nil
}

func (r *repository) SaveToken(key string, value interface{}, expiration time.Duration) error {
	return r.redis.Set(r.redis.Context(), key, value, expiration).Err()
}
