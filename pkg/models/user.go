package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"_id,omitempty" bson:"_id,omitempty"`
}
