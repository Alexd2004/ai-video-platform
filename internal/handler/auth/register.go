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

// I omit ID cause thats made at the DB level itself
const insertUserQuery = `INSERT INTO users (username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`

// PostRegister returns a handler that runs the insert against db.
//
//	It connects when you set up routes, after database.Connect: router.POST("/register", auth.PostRegister(db)).
func PostRegister(db *sql.DB) gin.HandlerFunc {

	//https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
	// Good examples above ^ Read through this to learn
	return func(c *gin.Context) {

		// Declare struct type
		var body register

		//This takes in the context and binds the json into that body struct (thats type register)
		if err := c.ShouldBindJSON(&body); err != nil {
			// This is how we return specific errors should there be one back in JSON.
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		// Here is some quick simple validation, ideally there will also be checks where the forum is filled out
		// Prior to being sent, so this edge case should never occur.
		if body.Username == "" || body.Email == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username, email, and password are required"})
			return
		}

		// So apparently hashing it and validating it here is the safe options as
		// In production HTTPS will protect the exposed password in transit
		// On our server itself, if we don't log, return it anywhere. Then it only exists but its not exposed
		// And then the only exposed part should be the hash in the database
		if err := PasswordValidation(body.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// These are my own functions i made in the password.go file
		hash, err := PasswordHash(body.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not process password"})
			return
		}

		/// Here is how we insert the query into the db
		_, err = db.ExecContext(c.Request.Context(), insertUserQuery, body.Username, body.Email, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"username": body.Username, "email": body.Email})
	}
}
