package middleware

import (
	"fmt"
	"net/http"
	"order-service/pkg/config"
	"order-service/pkg/constants"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

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

func TokenValid(c *gin.Context) error {
	err := ExtractTokenID(c)
	if err != nil {
		return fmt.Errorf("Can't extract token")
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	tokenString := strings.Split(bearerToken, " ")
	if len(tokenString) == 2 {
		if tokenString[0] != "Bearer" {
			return ""
		}
		return tokenString[1]
	}
	return ""
}

// return 0 if can't extract userID from token
func ExtractUserIDFromContext(c *gin.Context) int {
	return c.GetInt(constants.UserID)
}

func ExtractTokenID(c *gin.Context) error {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JwtConfig().APISecret), nil
	})
	if err != nil {
		return fmt.Errorf("Can't parse token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := fmt.Sprint(claims[constants.UserID])
		if len(userID) == 0 {
			return nil
		}
		// set and get from conetx cross middlerware
		c.Set(constants.UserID, userID)
		return nil
	}
	return nil
}
