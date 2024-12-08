package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u User) GetID() string {
	return u.ID
}

func (u User) GetCollectionName() string {
	return "users"
}
