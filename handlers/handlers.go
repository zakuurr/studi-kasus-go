package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"studi-kasus/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandlers struct {
	user models.Service
}

func NewUserHandler(user models.Service) *userHandlers {
	return &userHandlers{user}
}

func (h *userHandlers) UserHandlers(c *gin.Context) {
	user, err := h.user.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *userHandlers) UserByIdHandlers(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	user, err := h.user.FindById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": http.StatusNotFound,
			"result":  "Data Tidak Ada",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *userHandlers) DeleteUserHandlers(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	user, err := h.user.DeleteUser(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *userHandlers) UserByLimitOffset(c *gin.Context) {
	aString := c.Param("id")
	actionstring := c.Param("action")

	a, err := strconv.Atoi(aString)
	action, err := strconv.Atoi(actionstring)

	user, err := h.user.FindOffsetLimit(int(a), int(action))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *userHandlers) AddUserHandlers(c *gin.Context) {
	var userRequest models.User

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	userRequest.Id = rand.Intn(100)
	user, err := h.user.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func (h *userHandlers) UpdateUserHandler(c *gin.Context) {
	var userRequest models.User

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	user, err := h.user.UpdateUser(id, userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
