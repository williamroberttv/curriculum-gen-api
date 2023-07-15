package repositories

import (
	"fmt"
	"time"

	"github.com/williamroberttv/curriculum-gen-api/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *models.User) (*models.User, error)
	Find(id string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	Update(id string, user *models.User) (*models.User, error)
	Delete(id string) error
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func NewUserRepositoryDb(db *gorm.DB) *UserRepositoryDb {
	return &UserRepositoryDb{Db: db}
}

func (repo *UserRepositoryDb) Init(db *gorm.DB) {
	repo.Db = db
}

func (repo UserRepositoryDb) Insert(user *models.User) (*models.User, error) {
	user.BeforeCreate(repo.Db)

	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDb) Find(id string) (*models.User, error) {
	var user models.User

	repo.Db.First(&user, "id = ?", id)

	if user.ID == "" {
		return nil, fmt.Errorf("User not found")
	}
	return &user, nil
}

func (repo UserRepositoryDb) Update(id string, user *models.User) (*models.User, error) {
	user.UpdatedAt = time.Now()
	err := repo.Db.Save(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDb) Delete(id string) error {
	var user models.User

	repo.Db.First(&user, "id = ?", id)

	if user.ID == "" {
		return fmt.Errorf("User not found")
	}

	err := repo.Db.Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositoryDb) FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	repo.Db.First(&user, "email = ?", email)

	if user.ID == "" {
		return nil, fmt.Errorf("User not found")
	}
	return &user, nil
}