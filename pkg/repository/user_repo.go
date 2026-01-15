package repository

import (
	"books/pkg/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	AddToBlacklist(token string, expiresAt time.Time) error
	IsTokenBlacklisted(token string) bool
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) AddToBlacklist(token string, expiresAt time.Time) error {
	blacklist := models.TokenBlacklist{
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return r.db.Create(&blacklist).Error
}

func (r *userRepo) IsTokenBlacklisted(token string) bool {
	var count int64
	r.db.Model(&models.TokenBlacklist{}).Where("token = ?", token).Count(&count)
	return count > 0
}
