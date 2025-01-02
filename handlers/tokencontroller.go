package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"recruit/authentication"
	"recruit/ent"
	"recruit/ent/login"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	UID   string `json:"uid"`
	Uname string `json:"uname"` // struct inside a struct is what is to be implemented if you want all values in single struct
	jwt.StandardClaims
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

func ValidateRefreshToken(gctx *gin.Context) {

	tokenString := authentication.ExtractToken(gctx)

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

func ValidateToken(gctx *gin.Context) {
	tokenString := authentication.ExtractToken(gctx)

	claims, err := ParseToken(tokenString, "faskshfdkljadhfkj")

	fmt.Println(claims)
	if err != nil {
		gctx.String(http.StatusUnauthorized, "unauthorised")
		gctx.Abort()
		return
		//gctx.JSON(http.StatusOK, gin.H{"Username": "", "Verified": "False"})

	}
	if claims.Uname != "ram" {
		gctx.String(http.StatusUnauthorized, "unauthorised")
		gctx.Abort()
		return

		//gctx.JSON(http.StatusOK, gin.H{"Username": claims.Uname, "Verified": "True"})
	}

	gctx.Next()

	/*if claims != nil {
		if claims.Uname == "ram" {

			//gctx.JSON(http.StatusOK, gin.H{"Username": claims.Uname, "Verified": "True"})
		}
	 } */

	//gctx.JSON(http.StatusOK, gin.H{"Username": "", "Verified": "False"})
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
}

/* func t(gctx *gin.Context) {

	token, err := createToken()
	if err != nil {
		fmt.Println("Couldn't create token error")

	}
	refreshtoken, err := refreshToken()

	if err != nil {
		fmt.Println("Couldn't create token error")

	}

	gctx.JSON(http.StatusOK, gin.H{"token": token, "refreshtoken": refreshtoken})

} */

//createtokenfrom db /*

func createTokenfromdb(ctx context.Context, client *ent.Client, empid int32) (string, error) {
	// Retrieve dynamic values from the database using the Ent client

	userIDFromDB := uint(0)
	expireMinutesFromDB := 0
	secretFromDB := ""

	// Example of retrieving values from the database using Ent
	// Replace with your own implementation based on your Ent schema
	userEnt, err := client.Login.Query().
		Where(login.EmployeedIDEQ(empid)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			// Handle the case when no entity is found for the given condition
			log.Println("No newUser entry found for the specified employee ID")
			return "", fmt.Errorf("no newUser entry found for the specified employee ID: %w", err)
		}
		// Handle other error cases
		log.Println("Error getting newUser details from the database:", err)
		return "", fmt.Errorf("failed to retrieve newUser details from the database: %w", err)
	}

	// Log the retrieved values
	log.Printf("Retrieved newUser details for employee ID %d:", empid)
	log.Printf("UserID: %d", userEnt.EmployeedID)
	log.Printf("ExpireMinutes: %d", userEnt.ExpireminsToken)
	log.Printf("Secret: %s", userEnt.LoginID)

	userIDFromDB = uint(userEnt.EmployeedID)
	expireMinutesFromDB = int(userEnt.ExpireminsToken)
	//secretFromDB = userEnt.LoginID.String() // Convert UUID to string
	secretFromDB = userEnt.Username

	exp := time.Now().Add(time.Minute * time.Duration(expireMinutesFromDB)).Unix()
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		ID:    userIDFromDB,
		UID:   uid,
		Uname: userEnt.Username,
		//Uname: "ram",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secretFromDB))

	log.Println("Generated token:", token)

	return token, err
}

func Gettoken(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		empid, _ := strconv.ParseInt(id, 10, 32)
		newUsers, err := createTokenfromdb(ctx, client, int32(empid))
		if err != nil {
			log.Println("Invalid empid:", err)
			// Handle the error appropriately
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newUsers})
	}
}
