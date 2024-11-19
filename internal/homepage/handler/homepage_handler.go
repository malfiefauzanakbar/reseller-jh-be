package handler

import (
	"net/http"

	"reseller-jh-be/base"
	"reseller-jh-be/constant"
	"reseller-jh-be/internal/homepage/request"
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

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": homepage,
	}).Info("GetHomepage")

	base.RespondSuccess(c, homepage, nil)
}

func (h *HomepageHandler) UpdateHomepage(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - UpdateHomepage =====")

	var reqHomepage request.ReqHomepage
	if err := c.ShouldBind(&reqHomepage); err != nil {
		common.Log.Error("Func ShouldBind: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	file, err := c.FormFile("banner_image")
	if err != nil {
		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	homepage, err := h.Service.UpdateHomepage(c, &reqHomepage, file)
	if err != nil {
		common.Log.Error("Func UpdateHomepage: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": homepage,
	}).Info("UpdateHomepage")

	base.RespondSuccess(c, homepage, nil)
}
