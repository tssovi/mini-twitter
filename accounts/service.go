package accounts

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/robbert229/jwt"

	"../common"
	"../db/models"
)

var db = common.Dependencies.DB

// IAccountsService Service exposing accounts related operations
type IAccountsService struct {
}

func (as *IAccountsService) registerUser(data io.ReadCloser) error {
	user := models.User{}
	err := json.NewDecoder(data).Decode(&user)
	if err != nil {
		return errors.New("failed to decode form data")
	}
	user.Password = getHash(user.Password)
	errs := db.Create(&user).GetErrors()
	if len(errs) > 0 {
		return errors.New("failed to save user in the database")
	}
	return nil
}

func (as *IAccountsService) getUser(userID int) []int {
	followers := []int{5, 4, 3, 2, 1}
	return followers
}

func (as *IAccountsService) login(data io.ReadCloser) (*models.User, string, error) {
	serializer := LoginSerializer{}
	err := json.NewDecoder(data).Decode(&serializer)
	if err != nil {
		return nil, "", errors.New("failed to decode form data")
	}
	serializer.Password = getHash(serializer.Password)
	user := models.User{}
	errs := db.Where(&models.User{Username: serializer.Username, Password: serializer.Password}).First(&user).GetErrors()
	if len(errs) > 0 {
		return nil, "", errors.New("invalid credentials")
	}
	token := getToken(&user)
	return &user, token, nil
}

func getHash(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))
	str := fmt.Sprintf("%x", hasher.Sum(nil))
	return str
}

func getToken(user *models.User) string {
	algorithm := jwt.HmacSha256(common.JwtSecret)
	claims := jwt.NewClaim()
	exp := time.Now().Add(time.Duration(time.Hour * common.JwtExpiry))
	claims.SetTime("exp", exp)
	claims.Set("userID", user.ID)
	token, err := algorithm.Encode(claims)
	if err != nil {
		log.Fatalln("failed to sign JWT token")
	}
	return token
}

// AccountsService Singleton instance of AccountsService
var AccountsService = IAccountsService{}
