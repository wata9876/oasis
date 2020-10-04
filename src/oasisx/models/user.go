package models

import (
	"regexp"
	"time"

	"github.com/wcl48/valval"
)

// User is struct of user
type User struct {
	ID        int64     `gorm:"primary_key"`
	Name      string    `sql:"size:255"`
	CreatedAt time.Time `sql:"not null;type:date"`
	UpdatedAT time.Time `sql:"not null;type:date"`
	DeletedAt time.Time `sql:"not null;type:date"`
}

//UserValidate バリデーション
func UserValidate(user User) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(user)
}
