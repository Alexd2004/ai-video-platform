package auth

import (
	"database/sql"
	"errors"
	"net/http"

	// Handlers live in package auth; JWT signing lives in internal/auth — import alias avoids a name clash.
	appauth "video-platform/internal/auth"

	"github.com/gin-gonic/gin"
)

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const selectUserByEmail = `SELECT id, password_hash FROM users WHERE email = $1`

// PostLogin returns a Gin handler for POST /auth/login.

// The flow here essentialyl is we bind the json, load the user by tier email verify the hash, then issue then a JWT
// Wrong email and wrong password both return the same 401 message so callers cannot tell which field failed. TODO: Rewwrok this
func PostLogin(db *sql.DB, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body loginBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		// Validation
		if body.Email == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
			return
		}

		var userID, passwordHash string

		//
		err := db.QueryRowContext(c.Request.Context(), selectUserByEmail, body.Email).Scan(&userID, &passwordHash)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// We keep the error code like this cause apparently its good to not tell a hacker if the emai  l exists or not
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
				return
			}

			// Fail
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not look up user"})
			return
		}

		// Verify the password
		if err := CheckPasswordHash(body.Password, passwordHash); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}

		// token for user
		token, err := appauth.SignAccessToken(jwtSecret, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not issue token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
