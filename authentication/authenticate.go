package authentication

import (
	"context"
	"fmt"
	"hash/crc32"
	"recruit/ent"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
)

type JWTClaims struct {
	Username   string `json:"username"`
	Officename string `json:"ofcname"`
	jwt.StandardClaims
}

const (
	JWTSecretKey = "6baa3304-8cab-43bb-91cc-62cca2c0a7e9" // Change this to your preferred secret key
)

func CreateToken(userid string, client *ent.Client) string {
	// User exists, generate JWT token
	claims := JWTClaims{
		Username: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		// http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return "failed to generate token"
	}

	crc := crc32.ChecksumIEEE([]byte(signedToken))

	// Take the last 10 digits of the CRC32 checksum and pad with leading zeros
	crcDigits := fmt.Sprintf("%010d", crc)

	fmt.Println("10-digit number:", crcDigits)

	client.UserMaster.Update().SetExamCodePS(int32(crc)).Save(context.Background())

	return signedToken

}

func ValidateToken(tokenString string, client *ent.Client) (bool, error) {

	// Remove the "Bearer " prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	fmt.Println("inside token validation function..!")
	fmt.Println("token is - \n", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})
	fmt.Println("inside token validation function..1")

	if err != nil {
		return false, fmt.Errorf("token validation failed: %v", err)
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	fmt.Println("inside token validation function..2")

	// otp, err := client.UserMaster.Query().
	// 	Where(usermaster.User(userID)).
	// 	Select(usermaster.FieldOTP).
	// 	First(ctx)

	return true, nil

}

// var tokenStore = struct {
// 	m map[string]time.Time
// 	sync.RWMutex
// }{m: make(map[string]time.Time)}

// func CreateToken(userid string) (string, error) {
// 	// Lock token store for writing
// 	tokenStore.Lock()
// 	defer tokenStore.Unlock()

// 	// Invalidate all existing tokens for the user
// 	for user, _ := range tokenStore.m {
// 		if user == userid {
// 			delete(tokenStore.m, user)
// 		}
// 	}

// 	// User exists, generate JWT token
// 	claims := JWTClaims{
// 		Username: userid,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString([]byte(JWTSecretKey))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate token: %v", err)
// 	}

// 	// Update the token store with the new token
// 	tokenStore.m[userid] = time.Now()

// 	return signedToken, nil
// }

// func ValidateToken(tokenString string) (bool, error) {
// 	// Remove the "Bearer " prefix if present
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	// Parse the token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Check the signing method
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		// Return the secret key for validation
// 		return []byte(JWTSecretKey), nil
// 	})
// 	if err != nil {
// 		return false, fmt.Errorf("token parsing failed: %v", err)
// 	}

// 	// Check if the token is valid
// 	if token.Valid {
// 		return true, nil
// 	} else if ve, ok := err.(*jwt.ValidationError); ok {
// 		// Check the validation error type
// 		if ve.Errors&jwt.ValidationErrorExpired != 0 {
// 			return false, fmt.Errorf("token is expired")
// 		}
// 		return false, fmt.Errorf("invalid token: %v", err)
// 	} else {
// 		return false, fmt.Errorf("invalid token: %v", err)
// 	}
// }

// func ValidateToken(tokenString string) (bool, error) {
// 	// Remove the "Bearer " prefix if present
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	fmt.Println("inside token validation function..!")
// 	fmt.Println("token is - \n", tokenString)

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(JWTSecretKey), nil
// 	})
// 	fmt.Println("inside token validation function..1")

// 	if err != nil {
// 		return false, fmt.Errorf("token validation failed: %v", err)
// 	}

// 	if !token.Valid {
// 		return false, fmt.Errorf("invalid token")
// 	}

// 	fmt.Println("inside token validation function..2")

// 	return true, nil
// }
