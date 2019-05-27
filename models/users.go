package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate/validators"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"

	"golang.org/x/crypto/bcrypt"
)

// User is a struct of user records
type User struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	CreatedAt            time.Time `form:"-" json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `form:"-" json:"updated_at" db:"updated_at"`
	Email                string    `form:"email" json:"email" db:"email"`
	EncPassword          string    `db:"enc_password"`
	Password             string    `form:"password" db:"-"`
	PasswordConfirmation string    `form:"password_confirmation" db:"-"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
// example https://github.com/gobuffalo/authrecipe/blob/master/models/user.go
// validators https://github.com/gobuffalo/validate/tree/master/validators
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: u.Email},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error

	return validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: u.Email},
		&validators.StringIsPresent{Name: "EncPassword", Field: u.EncPassword},
		&validators.StringLengthInRange{Name: "Password", Field: u.Password, Min: 6, Max: 60},
		&validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirmation},
		&validators.FuncValidator{
			Field:   u.Email,
			Name:    "Email",
			Message: "%s is already taken",
			Fn: func() bool {
				var res bool
				q := tx.Where("email = ?", u.Email)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				res, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !res
			},
		},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// GeneratePassword generates hashed pass
func (u *User) GeneratePassword(pass string) (string, error) {
	saltedBytes := []byte(pass)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hash := string(hashedBytes[:])
	return hash, nil
}

// CheckPassword checks if password match
func (u *User) CheckPassword(pass string) bool {
	hash := []byte(u.EncPassword)
	p := []byte(pass)

	if err := bcrypt.CompareHashAndPassword(hash, p); err != nil {
		return false
	}
	return true
}
