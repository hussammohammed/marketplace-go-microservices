package user

type IUserService interface {
	Login(loginReq LoginDto) (string, error)
}
type UserService struct {
}
