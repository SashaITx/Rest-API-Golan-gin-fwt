package handler

import (
	"Rest_API_Golan-gin-fwt/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignUp struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Handler struct {
	services *service.AuthService
}

func NewHandler(services *service.AuthService) *Handler {
	return &Handler{services: services}
}

type error struct {
	Message string `json:"message"`
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")

	auth.POST("/sign-up", func(c *gin.Context) {
		var input UserSignUp

		if err := c.BindJSON(&input); err != nil {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		id, err := h.services.CreateUser(service.User{Name: input.Name, Username: input.Username, PasswordHash: input.Password})
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	})

	auth.POST("/sign-in", func(c *gin.Context) {
		var input UserSignIn

		if err := c.BindJSON(&input); err != nil {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		token, err := h.services.GenerateToken(input.Username, input.Password)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
		})
	})

	return router
}

func NewErrorResponse(c *gin.Context, statusCode int, massage string) {
	logrus.Error(massage)
	c.AbortWithStatusJSON(statusCode, error{massage})
}
