package auth

import (
	"crypto/rand"
	"errors"
	"fmt"

	// "github.com/astaxie/beego"
	"encoding/json"
	"io"
	"log"
	"time"
	"tl_mlkit/models"

	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/manucorporat/try"
	"golang.org/x/crypto/scrypt"
	// "crypto/md5"
	// "strconv"
)

// JWT : HEADER PAYLOAD SIGNATURE
const (
	SecretKEY              string = "key"
	DEFAULT_EXPIRE_SECONDS int    = 180 * 10 // default expired 10 minute
	PasswordHashBytes             = 16
)

type MyCustomClaims struct {
	UserID int `json:"UserID"`
	jwt.StandardClaims
}

type JwtPayload struct {
	Username  string `json:"Username"`
	UserID    int    `json:"UserID"`
	IssuedAt  int64  `json:"Iat"`
	ExpiresAt int64  `json:"Exp"`
}

func Check(ctx *context.Context) {

	if ctx.Input.Header("Authorization") == "" && ctx.Input.Method() == "OPTIONS" {

	} else {
		fmt.Print("else block")
		authtoken := ctx.Input.Header("Authorization")
	

		try.This(func() {
			token, err := jwt.ParseWithClaims(authtoken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(SecretKEY), nil
			})

			if err != nil {
				fmt.Print("******************")
			}
			claims, ok := token.Claims.(*MyCustomClaims)
			if ok && token.Valid {
				fmt.Printf("%+v\n", claims)

			} else {

				fmt.Printf("Token fail")
				ctx.Output.SetStatus(403)
				resBody, err := json.Marshal(models.GetCode("token_not_allowed"))
				ctx.Output.Body(resBody)
				if err != nil {
					panic(err)
				}

			}

		}).Finally(func() {

		}).Catch(func(e try.E) {
			if e != nil {
				panic(e)

			}
			fmt.Print("from catch")
			ctx.Output.SetStatus(200)
			resBody, err := json.Marshal(models.GetCode("token_not_allowed"))
			ctx.Output.Body(resBody)
			if err != nil {
				//	panic(err)

			}
		})
	}

}

// validate token
func ValidateToken(tokenString string) (*JwtPayload, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface {
		}, error) {

			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {

		log.Println("ok && token valid")
		logs.Info("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
		logs.Info("Token was issued at ", time.Now().Unix())
		logs.Info("Token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))

		return &JwtPayload{

			Username:  claims.StandardClaims.Issuer,
			UserID:    claims.UserID,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, nil
	} else {
		fmt.Println(err)
		return nil, errors.New("error: failed to validate token")
	}
}

// update token
func RefreshToken(tokenString string) (newTokenString string, err error) {

	// get previous token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface {
		}, error) {

			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {

		return "", err
	}

	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix() //new expired
	newClaims := MyCustomClaims{

		claims.UserID,
		jwt.StandardClaims{

			Issuer:    claims.StandardClaims.Issuer, //name of token issue
			IssuedAt:  time.Now().Unix(),            //time of token issue
			ExpiresAt: expireAt,                     // new expired
		},
	}

	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	// sign the token with a secret
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {

		return "", errors.New("error: failed to generate new fresh json web token")
	}

	return tokenStr, nil
}

// generate salt
func GenerateSalt() (salt string, err error) {

	buf := make([]byte, PasswordHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {

		return "", errors.New("error: failed to generate user's salt")
	}

	return fmt.Sprintf("%x", buf), nil
}

// generate password hash
func GeneratePassHash(password string, salt string) (hash string, err error) {

	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, PasswordHashBytes)
	if err != nil {

		return "", errors.New("error: failed to generate password hash")
	}

	return fmt.Sprintf("%x", h), nil
}

func CheckStatus(tokenString string) (string, int64) {

	jp, err := ValidateToken(tokenString)

	if err != nil {

		// if token has already expired
		fmt.Println("Your token has expired, Please log in again! ")
		return "", -1
	}

	timeDiff := jp.ExpiresAt - time.Now().Unix()
	fmt.Println("timeDiff = ", timeDiff)
	if timeDiff <= 30 {

		// if token is close to expiration, refresh the token
		fmt.Println("Your token would soon be expired")
		newToken, err := RefreshToken(tokenString)
		if err == nil {

			return newToken, timeDiff
		}
	}
	// if token is valid, do nothing
	fmt.Println("Your token is good ")
	return tokenString, timeDiff
}
