package service

type UserServiceImpl struct {
	// You might include any necessary dependencies here
}

func (s *UserServiceImpl) GetUserByID(userID int) (*User, error) {
	return &User{
		ID:       userID,
		Username: "vaibhav",
		Email:    "v.g",
	}, nil
}

func (s *UserServiceImpl) CreateUser(user *User) error {
	userData := &User{
		ID:       1,
		Username: "vaibhav",
		Email:    "v.g",
	}
	print(userData)
	return nil
}
