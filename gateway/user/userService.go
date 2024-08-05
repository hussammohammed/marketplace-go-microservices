package user

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hussammohammed/marketplace-go-microservices/gateway/helpers"
	userMicroService "github.com/hussammohammed/marketplace-go-microservices/gateway/server/grpcClients/protos/user"
)

var (
	issuer        = "micro marketplace gateway"
	audience      = "micro marketplace  gateway and microservices"
	secretKey     = []byte(os.Getenv("JWT_SECRET_KEY"))
	revokedTokens = make(map[string]bool)
)

type IUserService interface {
	Login(loginDto LoginDto) (string, error)
	LogOut(token string) (bool, error)
	ValidateAuthToken(tokenStr string) (jwt.MapClaims, error)
}
type UserService struct {
	userMicroService userMicroService.UserClient
	cryptHelper      helpers.ICryptHelper
}

func NewUserService(userMicroSvc userMicroService.UserClient, iCryptHelper helpers.ICryptHelper) *UserService {
	return &UserService{userMicroService: userMicroSvc, cryptHelper: iCryptHelper}
}
func (u *UserService) Login(loginDto LoginDto) (string, error) {
	if loginDto.Email == "" || loginDto.Password == "" {
		return "", nil
	}
	isUserExist, err := u.isUserExist(loginDto)
	if err != nil {
		return "", err
	}
	if isUserExist {
		return generateJwtToken()
	}
	return "", nil
}

func (u *UserService) isUserExist(loginDto LoginDto) (bool, error) {
	//u.userMicroService.CheckHealth()
	dbPass, _ := u.cryptHelper.HashPassword("123456")
	comparErr := u.cryptHelper.ComparePasswords(dbPass, loginDto.Password)
	if comparErr != nil {
		return false, comparErr
	}
	if loginDto.Email == "xyz@gmail.com" {
		return true, nil
	}
	return false, nil
}

func (u *UserService) LogOut(token string) (bool, error) {
	revokeToken(token)
	return true, nil
}

func generateJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		"nbf": time.Now().Unix(),                    // Token is not valid before current time
		"iss": issuer,                               // Issuer claim
		"aud": audience,                             // Audience claim
		"jti": "unique_token_id",                    // Unique token identifier
	})
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (u *UserService) ValidateAuthToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method and reject "none"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check token expiration
		exp := claims["exp"].(float64)
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, fmt.Errorf("token has expired")
		}
		// Check not before claim
		nbf, nbfOk := claims["nbf"].(float64)
		if nbfOk && time.Unix(int64(nbf), 0).After(time.Now()) {
			return nil, fmt.Errorf("token not valid yet")
		}

		// Check issuer claim
		iss := claims["iss"].(string)
		if iss != issuer {
			return nil, fmt.Errorf("invalid issuer")
		}

		// Check audience claim
		aud := claims["aud"].(string)
		if aud != audience {
			return nil, fmt.Errorf("invalid audience")
		}

		// Check if the token is revoked
		tokenID := claims["jti"].(string) // Assuming a unique identifier for each token
		if revokedTokens[tokenID] {
			return nil, fmt.Errorf("token has been revoked")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func revokeToken(tokenID string) {
	revokedTokens[tokenID] = true
}
