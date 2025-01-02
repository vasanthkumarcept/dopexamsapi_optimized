package authentication

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	UID   string `json:"uid"`
	Uname string `json:"uname"` // struct inside a struct is what is to be implemented if you want all values in single struct
	jwt.StandardClaims
}

func ParseToken(tokenString, secret string) (
	claims *JwtCustomClaims,
	err error,
) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func createToken() (token string, err error) {
	//Need to make it dynamic
	var userID uint = 4
	var expireMinutes int = 5
	var secret string = "faskshfdkljadhfkj"

	exp := time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix()
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		ID:    userID,
		UID:   uid,
		Uname: "ram",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))

	return

	/*if err != nil {

		fmt.Println("error in creation of token", err.Error())
	}*/

	//gctx.JSON(http.StatusOK, gin.H{"Token": token})
}

func t1(gctx *gin.Context) {

	token, err := createToken()
	if err != nil {
		fmt.Println("Couldn't create token error")

	}
	refreshtoken, err := refreshToken()

	if err != nil {
		fmt.Println("Couldn't create token error")

	}

	gctx.JSON(http.StatusOK, gin.H{"token": token, "refreshtoken": refreshtoken})

}

func refreshToken() (refreshtoken string, err error) {
	//Need to make these values dynamic
	var userID uint = 4
	var expireMinutes int = 10
	var secret string = "sadf"

	exp := time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix()
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		ID:    userID,
		UID:   uid,
		Uname: "ram",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshtoken, err = jwtToken.SignedString([]byte(secret))

	return

	/*if err != nil {

		fmt.Println("error in creation of token", err.Error())
	}*/

	//gctx.JSON(http.StatusOK, gin.H{"Token": token})
}

func validateRefreshToken(gctx *gin.Context) {

	tokenString := ExtractToken(gctx)

	claims, err := ParseToken(tokenString, "sadf")
	if err != nil {
		gctx.JSON(http.StatusOK, gin.H{"refreshTokenVerified": "False"})

	}

	if claims != nil {
		if claims.Uname == "ram" {
			//gctx.JSON(http.StatusOK, gin.H{"Username": claims.Uname, "Verified": "True"})

			token, err := createToken()
			if err != nil {
				fmt.Println("Couldn't create token error")

			}
			refreshtoken, err := refreshToken()

			if err != nil {
				fmt.Println("Couldn't create token error")

			}

			fmt.Println("Refresh Token validated.. and token generated")

			gctx.JSON(http.StatusOK, gin.H{"token": token, "refreshtoken": refreshtoken})

		}
	}

	//gctx.JSON(http.StatusOK, gin.H{"Username": "", "Verified": "False"})
}

func GenerateToken(user_id uint) (string, error) {

	token_lifespan := 6

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("roshani_here"))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("roshani_here"), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("roshani_here"), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
