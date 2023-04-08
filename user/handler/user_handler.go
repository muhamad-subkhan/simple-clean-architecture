package handler

import (
	"clean/models"
	"clean/user"
	"clean/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	user user.UserUsecase
}

func Routes(r *gin.Engine, user user.UserUsecase) {
	handler := userHandler{user: user}

	r.POST("/add-user", handler.createUser)
	r.GET("/users", handler.ReadAll)
}


func (h *userHandler) createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.Bind(&user)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Oops, something went wrong")
		return
	}

	if user.Id == 0 {
		response.Error(ctx, http.StatusBadRequest, "You must provide an id")
		return
	}

	if user.Name == "" {
		response.Error(ctx, http.StatusBadRequest, "You must provide a name")
		return
	}

	if user.Age == 0 || user.Age <= 0 {
		response.Error(ctx, http.StatusBadRequest, "You must provide an age")
		return
	}

	data, err := h.user.Create(&user)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Oops, something went wrong")
		return
	}

	response.Succes(ctx, data)
	
}


func (h *userHandler) ReadAll(ctx *gin.Context) {
	user, err := h.user.ReadAll()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Oops, something went wrong")
		return
	}

	if len(user) <= 0 {
		response.Error(ctx, http.StatusNotFound, "User not found")
		return
	}

	response.Succes(ctx, user)
}

