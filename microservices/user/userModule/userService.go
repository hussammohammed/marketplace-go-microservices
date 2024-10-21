package user

import (
	"fmt"

	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/db"
	protos "github.com/hussammohammed/marketplace-go-microservices/microservices/user/protos"
	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/repository"
	"gopkg.in/mgo.v2/bson"
)

type IuserService interface {
	CreateUser(rr *protos.CreateUserRequest) error
	IsUserExist(loginVM db.LoginViewModel) bool
}
type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(iUserRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository: iUserRepository}
}
func (u *UserService) CreateUser(rr *protos.CreateUserRequest) error {
	user := rr.GetUser()
	buyer := rr.GetBuyer()
	seller := rr.GetSeller()
	signupModel := &db.SignupViewModel{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.PhoneNumber,
	}

	userType := "user"
	if buyer != nil {
		userType = "buyer"
		signupModel.Buyer = &db.Buyer{
			ShippingAddress:        buyer.ShippingAddress,
			PaymentMethods:         buyer.PaymentMethods,
			BirthDate:              buyer.DateOfBirth,
			ProfilePicture:         buyer.ProfilePicture,
			PreferredPaymentMethod: buyer.PreferredPaymentMethod,
		}
	}
	if seller != nil {
		userType = "seller"
		signupModel.Seller = &db.Seller{
			CompanyName:  seller.CompanyName,
			BusinessType: seller.BusinessType,
		}
	}
	err := u.userRepository.Insert(signupModel, userType)
	return err
}

func (u *UserService) IsUserExist(loginVM db.LoginViewModel) bool {
	_, err := u.userRepository.FindOne(&bson.M{
		"email":    loginVM.Email,
		"password": loginVM.Password})

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
