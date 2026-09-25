package users_service

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (u UsersService) Test() string {
	return "Hello World!"
}
