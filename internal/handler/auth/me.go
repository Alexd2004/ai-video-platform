package auth

import (
	"net/http"

	"video-platform/internal/middleware"

	"github.com/gin-gonic/gin"
)

// GetMe responds with the authenticated user id (set by middleware.RequireJWT).
func GetMe(c *gin.Context) {
	id := c.GetString(middleware.GinUserIDKey)
	c.JSON(http.StatusOK, gin.H{"user_id": id})
}
