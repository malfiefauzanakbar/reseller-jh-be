package handler

import (	
	"net/http"

	"reseller-jh-be/base"
	"reseller-jh-be/constant"
	"reseller-jh-be/internal/homepage/model"	
	"reseller-jh-be/internal/homepage/service"
	"reseller-jh-be/pkg/common"

	"github.com/gin-gonic/gin"
)

type HomepageHandler struct {
	Service service.HomepageService
}

func NewHomepageHandler(service *service.HomepageService) *HomepageHandler {
	return &HomepageHandler{
		Service: *service,
	}
}

func (h *HomepageHandler) GetHomepage(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - GetHomepage =====")
	
	homepage, err := h.Service.GetHomepage()
	if err != nil {
		common.Log.Error("Func GetHomepage: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.Error, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": homepage,
	}).Info("GetHomepage")

	base.RespondSuccess(c, constant.Success, homepage, nil)
}

func (h *HomepageHandler) UpdateHomepage(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - UpdateHomepage =====")
	
	var reqHomepage model.Homepage
	if err := c.ShouldBindJSON(&reqHomepage); err != nil {
		common.Log.Error("Func ShouldBindJSON: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.Error, err.Error())
		return
	}

	homepage, err := h.Service.UpdateHomepage(&reqHomepage)
	if err != nil {
		common.Log.Error("Func UpdateHomepage: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.Error, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": homepage,
	}).Info("UpdateHomepage")

	base.RespondSuccess(c, constant.Success, homepage, nil)
}
