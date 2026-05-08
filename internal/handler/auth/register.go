package auth

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check for more fields and such, the ID should be generated i assume at the DB level
// Along with created and updated at

// Need to figure out where we hash this
type register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Docs i used
// https://go.dev/doc/database/querying
// https://www.practical-go-lessons.com/post/how-to-insert-data-into-a-mysql-database-with-golang-ccbmu7s6qcuc70nnaia0
// https://pkg.go.dev/github.com/gin-gonic/gin#Context.Handler

// I omit ID cause thats made at the DB level itself
const insertUserQuery = `INSERT INTO users (username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`

// PostRegister returns a handler that runs the insert against db.
//
// AI DID this part
// Wire in cmd/api: auth.POST("/register", handlerauth.PostRegister(db)) with group prefix "/auth".
//
// References: https://go.dev/doc/database/querying , Gin JSON binding: https://pkg.go.dev/github.com/gin-gonic/gin#Context.Handler
func PostRegister(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body register
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		if body.Username == "" || body.Email == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username, email, and password are required"})
			return
		}

		// Hash only after policy checks; avoid storing or logging the plaintext password.
		if err := PasswordValidation(body.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hash, err := PasswordHash(body.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not process password"})
			return
		}

		_, err = db.ExecContext(c.Request.Context(), insertUserQuery, body.Username, body.Email, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"username": body.Username, "email": body.Email})
	}
}
