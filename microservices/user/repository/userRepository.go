package repository

import (
	"time"

	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	UserCollection = "users"
)

type IUserRepository interface {
	Insert(user *db.SignupViewModel, userType string) error
	FindOne(query *bson.M) (*db.UserModel, error)
}

type UserRepository struct {
	db *db.DB
}

func NewUserRepository(database *db.DB) *UserRepository {
	return &UserRepository{db: database}
}
func (r *UserRepository) getSession() *mgo.Session {
	session := r.db.Clone()      // Clone session for concurrent operations on the database.
	session.SetSafe(&mgo.Safe{}) // set the safety mode for the session
	return session
}

func (r *UserRepository) Insert(user *db.SignupViewModel, userType string) error {
	session := r.getSession()
	defer session.Close()
	collection := session.DB("").C(UserCollection)
	err := collection.Insert(&db.UserModel{
		Id:        bson.NewObjectId(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		Type:      userType,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Buyer:     user.Buyer,
		Seller:    user.Seller,
	})
	return err
}

func (r *UserRepository) FindOne(query *bson.M) (*db.UserModel, error) {
	session := r.getSession()
	defer session.Close()
	collection := session.DB("").C(UserCollection)
	user := db.UserModel{}
	err := collection.Find(query).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
