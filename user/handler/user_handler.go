package handler

import (
	"clean/models"
	"clean/user"
	"clean/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	user user.UserUsecase
}

func Routes(r *gin.Engine, user user.UserUsecase) {
	handler := userHandler{user: user}

	r.POST("/add-user", handler.createUser)
	r.GET("/users", handler.ReadAll)
	r.GET("/user/:id", handler.GetId)
	r.PATCH("user/:id", handler.Update)
	r.DELETE("/users/:id", handler.Delete)
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

func (h *userHandler) GetId(ctx *gin.Context) {

	id := ctx.Param("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "User not found")
		return
	}

	user, err := h.user.GetId(int64(userId))
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "User not found")
		return
	}

	response.Succes(ctx, user)

}

func (h *userHandler) Update(ctx *gin.Context) {


	
	r := ctx.Request

	ID := ctx.Param("id")
	userId, err := strconv.Atoi(ID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "ID must be an uinteger")
		return
	}

	age := r.FormValue("age")
	userAge, err := strconv.Atoi(age)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "age must be an int")
	}


	req, err := h.user.GetId(int64(userId))
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "User not found")
		return
	}


	request := models.User{
		Id: uint(userId),
		Name: r.FormValue("name"),
		Age: userAge,
		
	}

	if req.Id != 0 {
		req.Id = request.Id
	}

	if req.Name != "" {
		req.Name = request.Name
	}

	if req.Age != 0 {
		req.Age = request.Age
	}
	
	
	data, err := h.user.Update(&request)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Opsss Something went wrong")
		return
	}
	
	response.Succes(ctx, data)
}


func (h *userHandler) Delete(ctx *gin.Context){
	id := ctx.Param("id")
	ID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Invalid ID")
		return
	}

	data, err := h.user.Delete(int64(ID))
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "User not found")
		return
	}

	response.Succes(ctx, data)
}