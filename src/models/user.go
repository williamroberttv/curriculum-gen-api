package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/williamroberttv/curriculum-gen-api/src/utils"
)

type User struct {
	ID 	string `valid:"string" json:"id" gorm:"type:uuid;primary_key;"`
	Name string `valid:"type(string),notnull" json:"name" gorm:"type:varchar(255); not null;"`
	Age       int    `valid:"type(int),notnull" json:"age" gorm:"type:integer; not null;"`
	Password	string	`valid:"notnull" json:"password" gorm:"type:varchar(255); not null;"`
	Email			string	`valid:"email" json:"email" gorm:"type:varchar(255);unique_index; not null;"`
	Gender string `valid:"type(string),notnull" json:"gender" gorm:"type:varchar(1); not null;"`
	CreatedAt time.Time `valid:"-" json:"created_at"`
	UpdatedAt time.Time `valid:"-" json:"updated_at"`
	DeletedAt time.Time `valid:"-" json:"deleted_at" gorm:"default:null"`
}

func NewUser() *User {
	return &User{}
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (user *User) Validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	return
}