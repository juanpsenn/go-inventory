package user

import (
	"net/http"

	"example.com/users/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h Handler) CreateUser(c *gin.Context) {
	var newUser domain.User

	// Translate to domain
	// Bind the JSON payload to newUser struct.
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ====== Consume service
	userID, err := h.UserService.Create(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ======

	// Translate to response
	// Respond with the new user object.
	c.JSON(http.StatusCreated, gin.H{"user_id": userID})
}
