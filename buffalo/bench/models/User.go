package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

// User is used by pop to map your users database table to your go code.
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Bio       string    `json:"bio" db:"bio"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
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
// func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
// 	return validate.Validate(
// 		&validators.StringIsPresent{Field: u.Title, Name: "Title"},
// 		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},
// 		&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},
// 		&validators.StringIsPresent{Field: u.Bio, Name: "Bio"},
// 	), nil
// }

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
// func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
// 	return validate.NewErrors(), nil
// }

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
// func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
// 	return validate.NewErrors(), nil
// }
