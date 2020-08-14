package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/poker/lib"
	"github.com/uzimaru0000/poker/usecase"
	"github.com/uzimaru0000/poker/usecase/impl"
	"net/http"
)

const (
	ClaimKey = "claim"
)

func AuthMiddleware(tokenizer usecase.Tokenizer) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := c.GetHeader("Authorization")
		if !lib.TokenValidation(rawToken) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := lib.ExtractingToken(rawToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format token"})
			c.Abort()
			return
		}

		rawClaim, err := tokenizer.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		claim, ok := rawClaim.(*impl.Claim)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
			c.Abort()
			return
		}

		c.Set(ClaimKey, claim)
		c.Next()
	}
}
