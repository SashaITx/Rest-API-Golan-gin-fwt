package handler

import (
	rest "Rest_API_Golan-gin-fwt"
	"Rest_API_Golan-gin-fwt/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	services *service.AuthService
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type error struct {
	Message string `json:"message"`
}

func NewHandler(services *service.AuthService) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")

	auth.POST("/sign-up", func(c *gin.Context) {
		var input rest.User

		if err := c.BindJSON(&input); err != nil {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		id, err := h.services.CreateUser(input)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	})

	auth.POST("/sign-in", func(c *gin.Context) {
		var input signInInput

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
