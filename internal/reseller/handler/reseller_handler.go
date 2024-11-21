package handler

import (
	"net/http"

	"reseller-jh-be/base"
	"reseller-jh-be/constant"
	"reseller-jh-be/internal/reseller/model"
	"reseller-jh-be/internal/reseller/request"
	"reseller-jh-be/internal/reseller/service"
	"reseller-jh-be/pkg/common"

	"github.com/gin-gonic/gin"
)

type ResellerHandler struct {
	Service service.ResellerService
}

func NewResellerHandler(service *service.ResellerService) *ResellerHandler {
	return &ResellerHandler{
		Service: *service,
	}
}

func (h *ResellerHandler) CreateReseller(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - CreateReseller =====")

	var reqReseller request.ReqCreateReseller
	if err := c.ShouldBind(&reqReseller); err != nil {
		common.Log.Error("Func ShouldBind: ", err)

		// resp := base.BaseResp{
		// 	Status:  constant.Error,
		// 	Message: err.Error(),
		// }
		// c.JSON(http.StatusBadRequest, resp)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	file, _ := c.FormFile("ktp")	

	reseller, err := h.Service.CreateReseller(c, reqReseller, file)
	if err != nil {
		common.Log.Error("Func CreateReseller: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": reseller,
	}).Info("CreateReseller")

	// resp := base.BaseResp{
	// 	Status: constant.Success,
	// 	Data:   reseller,
	// }
	// c.JSON(http.StatusCreated, resp)
	base.RespondSuccess(c, reseller, nil)
}

func (h *ResellerHandler) GetAllReseller(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - GetAllReseller =====")

	var reqReseller request.ReqReseller
	if err := c.ShouldBindQuery(&reqReseller); err != nil {
		common.Log.Error("Func ShouldBindQuery: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	reqPagination := common.HandleReqPagination(c)
	resellers, pagination, err := h.Service.GetAllReseller(reqReseller, reqPagination)
	if err != nil {
		common.Log.Error("Func GetAllReseller: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": resellers,
	}).Info("GetAllReseller")

	base.RespondSuccess(c, resellers, &pagination)
}

func (h *ResellerHandler) GetReseller(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - GetReseller =====")

	id := c.Param("id")
	reseller, err := h.Service.GetReseller(id)
	if err != nil {
		common.Log.Error("Func GetReseller: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": reseller,
	}).Info("GetReseller")

	base.RespondSuccess(c, reseller, nil)
}

func (h *ResellerHandler) UpdateReseller(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - UpdateReseller =====")

	id := c.Param("id")
	var reqReseller model.Reseller
	if err := c.ShouldBindJSON(&reqReseller); err != nil {
		common.Log.Error("Func ShouldBindJSON: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	reseller, err := h.Service.UpdateReseller(id, &reqReseller)
	if err != nil {
		common.Log.Error("Func UpdateReseller: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": reseller,
	}).Info("UpdateReseller")

	base.RespondSuccess(c, reseller, nil)
}

func (h *ResellerHandler) ReadReseller(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - ReadReseller =====")

	id := c.Param("id")

	reseller, err := h.Service.ReadReseller(id)
	if err != nil {
		common.Log.Error("Func ReadReseller: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": reseller,
	}).Info("ReadReseller")

	base.RespondSuccess(c, reseller, nil)
}

func (h *ResellerHandler) DeleteReseller(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - DeleteReseller =====")

	id := c.Param("id")
	if err := h.Service.DeleteReseller(id); err != nil {
		common.Log.Error("Func UpdateReseller: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	base.RespondSuccess(c, nil, nil)
}

func (h *ResellerHandler) CountResellers(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - CountResellers =====")

	var reqReseller request.ReqReseller
	if err := c.ShouldBindQuery(&reqReseller); err != nil {
		common.Log.Error("Func ShouldBindQuery: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}
	reseller, err := h.Service.CountResellers(reqReseller)
	if err != nil {
		common.Log.Error("Func CountResellers: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": reseller,
	}).Info("ReadReseller")

	base.RespondSuccess(c, reseller, nil)
}

func (h *ResellerHandler) ResellersChart(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - ResellersChart =====")

	var reqReseller request.ReqReseller
	if err := c.ShouldBindQuery(&reqReseller); err != nil {
		common.Log.Error("Func ShouldBindQuery: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}
	resp, err := h.Service.ResellersChart(reqReseller)
	if err != nil {
		common.Log.Error("Func ResellersChart: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": resp,
	}).Info("ReadReseller")

	base.RespondSuccess(c, resp, nil)
}

func (h *ResellerHandler) ExportExcelResellers(c *gin.Context) {
	common.Log.Info("===== HANDLER CALLED - ExportExcelResellers =====")

	var reqReseller request.ReqReseller
	if err := c.ShouldBindQuery(&reqReseller); err != nil {
		common.Log.Error("Func ShouldBindQuery: ", err)

		base.RespondError(c, http.StatusBadRequest, constant.BadRequest, err.Error())
		return
	}

	resp, err := h.Service.ExportExcelResellers(reqReseller)
	if err != nil {
		common.Log.Error("Func ExportExcelResellers: ", err)

		base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err.Error())
		return
	}

	common.Log.WithFields(map[string]interface{}{
		"data": resp,
	}).Info("ExportExcelResellers")

	base.RespondSuccess(c, resp, nil)
}
