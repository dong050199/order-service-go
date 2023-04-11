package handler

import (
	"net/http"
	"order-service/pkg/wrapper"
	"order-service/service/model/request"
	"order-service/service/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useUsecase usecase.IuserUsecase
}

func NewUserhandler(
	useUsecase usecase.IuserUsecase,
) UserHandler {
	return UserHandler{
		useUsecase: useUsecase,
	}
}

// User godoc
// @Summary Get my record page for non login usersas
// @Description Get my record page for non login usersas
// @Tags user
// @Accept json
// @Produce json
// @Param model body request.UserRegister true "model"
// @Success 200 {object} wrapper.Response{data=string} "success"
// @Failure 400 {object} wrapper.Response
// @Failure 500 {object} wrapper.Response
// @Router  /user/register [post]
func (h *UserHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.UserRegister
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		token, err := h.useUsecase.UserRegister(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		wrapper.JSONOk(c, token)
	}
}

// User godoc
// @Summary API for get token from user name email and password
// @Description API for get token from user name email and password
// @Tags user
// @Accept json
// @Produce json
// @Param model body request.UserLogin true "model"
// @Success 200 {object} wrapper.Response{data=string} "success"
// @Failure 400 {object} wrapper.Response
// @Failure 500 {object} wrapper.Response
// @Router /user/login [post]
func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.UserLogin
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		token, err := h.useUsecase.UserLogin(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		wrapper.JSONOk(c, token)
	}
}
