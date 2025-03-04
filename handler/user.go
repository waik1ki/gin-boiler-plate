package handler

import (
	"gin-boiler-plate/model"
	"gin-boiler-plate/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func newUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// CreateUser: POST /users - 사용자 생성
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ErrorResponse(c, http.StatusBadRequest)
		return
	}
	if err := h.userService.CreateUser(&user); err != nil {
		ErrorResponse(c, http.StatusInternalServerError)
		return
	}
	SuccessResponse(c, http.StatusCreated, user)
}

// GetUsers: GET /users - 전체 사용자 조회
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError)
		return
	}
	SuccessResponse(c, http.StatusOK, users)
}

// GetUser: GET /users/:id - 특정 사용자 조회
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest)
		return
	}
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError)
		return
	}
	SuccessResponse(c, http.StatusOK, user)
}

// UpdateUser: PUT /users/:id - 사용자 수정
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest)
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ErrorResponse(c, http.StatusBadRequest)
		return
	}
	user.ID = uint(id)
	if err := h.userService.UpdateUser(&user); err != nil {
		ErrorResponse(c, http.StatusInternalServerError)
		return
	}
	SuccessResponse(c, http.StatusOK, user)
}

// DeleteUser: DELETE /users/:id - 사용자 삭제
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest)
		return
	}
	if err := h.userService.DeleteUser(uint(id)); err != nil {
		ErrorResponse(c, http.StatusInternalServerError)
		return
	}
	SuccessResponse(c, http.StatusOK, nil)
}
