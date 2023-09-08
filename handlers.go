package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	
)
func(h *Handler)registerUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Passwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Passwd = string(hashedPassword)


	db.Create(&newUser)

	c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) loginUser(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := generateToken(user)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler)getUserInfo(c *gin.Context) {
	// Get the user from the authentication middleware
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) createLead(c *gin.Context){
	userId, _:= c.Get("userID")
	userIDUint, ok := userId.(uint)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format in context"})
        return
    }
	var newLead Lead
	if err := c.BindJSON(&newLead); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()})
		return
		}
	newLead.CreatedByID = userIDUint
	db.Create(&newLead)
	c.JSON(http.StatusCreated, newLead)
}


func (h *Handler) showAllUsersLeads(c *gin.Context,){
	userId, exists := c.Get("userID")
	if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }
	userIDUint, ok := userId.(uint)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format in context"})
        return
    }
	var leads []Lead
	if err := db.Where("created_by_id = ?", userIDUint).Find(&leads).Error;err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, leads)
}

func (h *Handler) updateUserLead(c *gin.Context) {
    id := c.Param("id")

    var lead Lead
    if err := db.First(&lead, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
        return
    }

    // Define a map to hold the updated data.
    updateData := make(map[string]interface{})

    // Attempt to bind the updated lead data from the request body (assuming it's in JSON format).
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

    // Save the updated lead back to the database.
    if err := db.Save(&updateData).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lead"})
        return
    }

    c.JSON(http.StatusOK, updateData) // Respond with the updated lead.
}


func (h *Handler) deleteUserLead(c *gin.Context) {
    id := c.Param("id")
    
    // Implement code to find the user's lead with the given ID in your database.
    var lead Lead
    if err := h.db.First(&lead, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "UserLead not found"})
        return
    }

    // Implement code to delete the user's lead from the database.
    if err := db.Delete(&lead).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete UserLead"})
        return
    }

    c.JSON(http.StatusNoContent, nil) // Respond with a 204 No Content status.
}