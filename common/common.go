package common

import(

	"fmt"

	"time"

	"crypto/md5"

	"encoding/base64"

	"github.com/dgrijalva/jwt-go"

)

type Claims struct {

	AccountId 				string `json:"acount_id"`

	jwt.StandardClaims

}

func Base64Encode(s []byte ) []byte {

	return []byte(base64.StdEncoding.EncodeToString(s))

}

func ToMD5(encod string ) (decode string ) {

	md5Ctx := md5.New()

	md5Ctx.Write([]byte(encod))

	cipher := md5Ctx.Sum(nil)

	return string(Base64Encode(cipher))

}

//创建token
func CreateToken(account string,secret string ) (string,int64) {

	expireTime := time.Now().Add(time.Hour*1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	nclaims := make(jwt.MapClaims)

	nclaims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	
	nclaims["iat"] = time.Now().Unix()

	nclaims["account_id"] = account
	
	token.Claims = nclaims
	
	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString,expireTime

}

//校验
func TokenAuth(signedToken,secret string ) string {

	fmt.Println("signedToken :"+signedToken )

	// token, err := jwt.ParseWithClaims(signedToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

	// 	return []byte(secret), nil

	// })

	token, err := jwt.Parse(signedToken,func(token *jwt.Token) (interface{}, error) {

		return []byte(secret), nil

	})

	if err != nil {

		fmt.Println(err)

		return ""

	}

	if !token.Valid {

		return ""

	}

	beego.Debug("Token token:", token)

	claims, ok := token.Claims.(jwt.MapClaims)

	beego.Debug("Token:", claims)

	if !ok {

		return ""

	}
	
	accountId := claims["account_id"].(string)

	return accountId

}

func GetSecretStr() string {

	return "secret"

}
