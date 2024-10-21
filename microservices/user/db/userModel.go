package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Password  string        `json:"password"`
	Phone     string        `json:"phone"`
	Type      string        `json:"-"` // ["user","buyer","seller"]
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Buyer     *Buyer        `json:"buyer"`
	Seller    *Seller       `json:"seller"`
}

type Buyer struct {
	ShippingAddress        string   `json:"shipping_address" binding:"required"`
	PaymentMethods         []string `json:"payment_methods" binding:"required"`
	BirthDate              string   `json:"birth_date"`
	ProfilePicture         string   `json:"profile_picture"`
	PreferredPaymentMethod string   `json:"preferred_payment_method"`
}

type Seller struct {
	CompanyName  string `json:"company_name" binding:"required"`
	BusinessType string `json:"business_type"`
}

type LoginViewModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
