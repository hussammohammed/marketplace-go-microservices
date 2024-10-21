package db

type SignupViewModel struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Password string  `json:"password"`
	Phone    string  `json:"phone"`
	Buyer    *Buyer  `json:"buyer"`
	Seller   *Seller `json:"seller"`
}
