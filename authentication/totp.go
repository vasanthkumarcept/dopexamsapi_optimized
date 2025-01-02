package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func topgenerate1(c *gin.Context) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "gateway.cept.gov.in",
		AccountName: "Vams",
		Secret:      []byte("aea41edccefffad065fc8a9ddf6661853f29776085be6e00f235cfdc820ff39a"),
		Period:      90,
		Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"key": key.Secret()})

}
