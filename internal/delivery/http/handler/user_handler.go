package handler

import (
	"net/http"
	domain "splunk_soar_clone/internal/domain/entity"
	"splunk_soar_clone/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUseCase *user.UserUseCase
}
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   string `json:"role_id" binding:"required"`
}

func NewAuthHandler(userUseCase *user.UserUseCase) *AuthHandler {
	return &AuthHandler{
		userUseCase: userUseCase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := h.userUseCase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": gin.H{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
			"expires_at":    token.Expiry,
		},
		"user": gin.H{
			"user_id":  user.UserID,
			"role_id":  user.RoleID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
func (h *AuthHandler) CreateUser(c *gin.Context) {
	roleID, _ := c.Get("role_id")
    if roleID != "1" {
        c.JSON(http.StatusForbidden, gin.H{"error": "only admin can create users"})
        return
    }

    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create user
    user, err := h.userUseCase.CreateUser(&domain.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
		RoleID:       req.RoleID,
	})
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User created successfully",
        "user": gin.H{
            "user_id":  user.UserID,
            "username": user.Username,
            "email":    user.Email,
            "role_id":  user.RoleID,
        },
    })
}
