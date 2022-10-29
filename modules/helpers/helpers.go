package helpers

import (
	"fmt"
	"net/http"
	"time"
	"vessl/modules/users"

	"github.com/gin-gonic/gin"
)

func ValidateApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {

		ApiKey := c.Request.Header.Get("Authorization")

		for _, dbApiKeyData := range users.GetApiKeys() {
			if ApiKey == dbApiKeyData.ApiKey {
				if time.Now().UnixMilli()-dbApiKeyData.Ts > int64(time.Duration(8*time.Hour).Milliseconds()) {
					fmt.Printf("API Key: %s expired", ApiKey)
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication token expired"})
					return
				}
				return
			}
		}

		fmt.Printf("Found 0 results for API Key: %s", ApiKey)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
		return
	}
}
