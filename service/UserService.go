package service

type UserService interface {
	GetUserByID(userID int) (*User, error)
	CreateUser(user *User) error
	// Add more methods as needed
}
