package data

import (
	"github.com/hashicorp/go-hclog"
	"github.com/thetkpark/golang-todo/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(username string, password string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}

type GormUserRepository struct {
	db  *gorm.DB
	log hclog.Logger
}

func NewGormUserRepository(db *gorm.DB, log hclog.Logger) *GormUserRepository {
	return &GormUserRepository{db: db, log: log}
}

func (repo *GormUserRepository) Create(username string, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: password,
	}

	tx := repo.db.Create(user)
	if tx.Error != nil {
		repo.log.Error("Error saving user to db", tx.Error.Error())
		return nil, tx.Error
	}
	return user, nil
}

func (repo *GormUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	tx := repo.db.Where(&models.User{Username: username}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
