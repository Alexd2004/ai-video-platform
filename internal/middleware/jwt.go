package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	appauth "video-platform/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const bearerPrefix = "Bearer "

// GinUserIDKey is the key used with gin.Context.Set / Get for the authenticated user's id (JWT sub).
const GinUserIDKey = "userID"

// userIDRequestCtxKey stores the same id on *http.Request.Context() for code that does not use Gin getters.
type userIDRequestCtxKey struct{}

// RequireJWT returns middleware that requires a valid Authorization: Bearer <JWT>.
// It verifies the signature and expiry with the same secret used at login, then stores the user id in:
//   - gin context: c.GetString(GinUserIDKey)
//   - request context: UserIDFromRequest(c.Request)
func RequireJWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.GetHeader("Authorization")
		if raw == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		if !strings.HasPrefix(raw, bearerPrefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization must be Bearer <token>"})
			return
		}
		tokenString := strings.TrimSpace(strings.TrimPrefix(raw, bearerPrefix))
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is empty"})
			return
		}

		claims, err := appauth.ParseAccessToken(secret, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": jwtErrorMessage(err)})
			return
		}
		if claims.Subject == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is missing subject (user id)"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), userIDRequestCtxKey{}, claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Set(GinUserIDKey, claims.Subject)
		c.Next()
	}
}

// UserIDFromRequest returns the user id set by RequireJWT, if present.
func UserIDFromRequest(r *http.Request) (string, bool) {
	v, ok := r.Context().Value(userIDRequestCtxKey{}).(string)
	return v, ok
}

func jwtErrorMessage(err error) string {
	switch {
	case errors.Is(err, jwt.ErrTokenExpired):
		return "token has expired"
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return "invalid token signature"
	case errors.Is(err, jwt.ErrTokenMalformed):
		return "malformed token"
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		return "token is not valid yet"
	case errors.Is(err, jwt.ErrTokenInvalidClaims):
		return "invalid token claims"
	case errors.Is(err, jwt.ErrTokenInvalidAudience):
		return "invalid token audience"
	case errors.Is(err, jwt.ErrTokenInvalidIssuer):
		return "invalid token issuer"
	case errors.Is(err, jwt.ErrTokenUnverifiable):
		return "token could not be verified"
	case errors.Is(err, jwt.ErrTokenRequiredClaimMissing):
		return "token is missing a required claim"
	default:
		return "invalid or unreadable token"
	}
}
