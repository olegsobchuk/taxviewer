package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// TaxCurrencies is the map of allowed currencies
var TaxCurrencies = map[int]string{
	0: "usd",
	1: "eur",
	2: "hrn",
}

// Tax table for taxes records
type Tax struct {
	ID        uuid.UUID     `json:"id" db:"id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
	Date      time.Time     `json:"date" db:"date"`
	Amount    float64       `json:"amount" db:"amount"`
	Currency  int           `json:"currency" db:"currency"`
	Exchange  float64       `json:"exchange" db:"exchange"`
	Exchanged nulls.Float64 `form:"-" db:"exchanged"`
}

// String is not required by pop and may be deleted
func (t Tax) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Taxes is not required by pop and may be deleted
type Taxes []Tax

// String is not required by pop and may be deleted
func (t Taxes) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Tax) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: t.Currency, Name: "Currency"},
		&validators.FuncValidator{
			Fn: func() bool {
				_, ok := TaxCurrencies[t.Currency]
				return ok
			},
			Name:    "Currency",
			Message: "Currency must be choosen",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				return t.Amount == float64(t.Amount)
			},
			Name: "Amount",
		},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Tax) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Tax) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
