package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"

	"golang.org/x/crypto/bcrypt"
)

// User is a struct of user records
type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `form:"-" json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `form:"-" json:"updated_at" db:"updated_at"`
	Email       string    `form:"email" json:"email" db:"email"`
	EncPassword string    `db:"enc_password"`
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
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
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
