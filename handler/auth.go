package handler

import (
	"fmt"
	"net/http"
	"project-backend/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type request struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Signin(c *gin.Context) {
	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user, err := h.store.UserRepository().UserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect username or password"})
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect username or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("user", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "success"})
	return
}

func (h *Handler) Signup(c *gin.Context) {
	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user := &model.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}
	if err := h.store.UserRepository().Create(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	session := sessions.Default(c)
	session.Set("user", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "success"})
	return
}

func (h *Handler) WhoAmI(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no such user"})
		return 
	}

	user, err := h.store.UserRepository().UserById(user.(uint))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no such user 2"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return 
}

func (h *Handler) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		fmt.Println(user)

		if user == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := h.store.UserRepository().UserById(user.(uint))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		c.Next()
	}
}
