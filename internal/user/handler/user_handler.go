package handler

import (
	"net/http"

	"reseller-jh-be/base"
	"reseller-jh-be/constant"
	"reseller-jh-be/internal/user/request"
	"reseller-jh-be/internal/user/service"
	"reseller-jh-be/pkg/common"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: *service,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - Login =====")

	var reqUser request.ReqLogin
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		common.Log.Error("Func ShouldBindJSON: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	user, err := h.Service.Login(c, reqUser)
	if err != nil {
		common.Log.Error("Func Login: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": user,
	}).Info("Login")

	base.RespondSuccess(c, constant.Success, user, nil)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - CreateUser =====")

	var reqUser request.ReqRegister
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		common.Log.Error("Func ShouldBindJSON: ", err)

		// resp := base.BaseResp{
		// 	Status:  constant.Error,
		// 	Message: err.Error(),
		// }
		// c.JSON(http.StatusBadRequest, resp)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	user, err := h.Service.CreateUser(reqUser)
	if err != nil {
		common.Log.Error("Func CreateUser: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": user,
	}).Info("CreateUser")

	// resp := base.BaseResp{
	// 	Status: constant.Success,
	// 	Data:   user,
	// }
	// c.JSON(http.StatusCreated, resp)
	base.RespondSuccess(c, constant.Success, user, nil)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - GetUser =====")

	id := c.Param("id")
	user, err := h.Service.GetUser(id)
	if err != nil {
		common.Log.Error("Func GetUser: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": user,
	}).Info("GetUser")

	base.RespondSuccess(c, constant.Success, user, nil)
}
