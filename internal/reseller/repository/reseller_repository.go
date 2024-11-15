package repository

import (
	"reseller-jh-be/base"
	"reseller-jh-be/internal/reseller/model"
	"reseller-jh-be/internal/reseller/request"
	"reseller-jh-be/internal/reseller/response"

	"gorm.io/gorm"
)

type ResellerRepository struct {
	DB *gorm.DB
}

type ResellerRepositoryInterface interface {
}

func NewResellerRepository(DB *gorm.DB) *ResellerRepository {
	return &ResellerRepository{
		DB: DB,
	}
}

func (r *ResellerRepository) CreateReseller(reqReseller *model.Reseller) (*model.Reseller, error) {
	err := r.DB.Create(reqReseller).Error
	if err != nil {
		return nil, err
	}

	return reqReseller, nil
}

func (r *ResellerRepository) GetAllReseller(statusID int64, reqReseller request.ReqReseller, reqPagination base.ReqPagination) (resellers []response.RespReseller, count int64, err error) {
	query := r.DB.Table("resellers AS r").Select("r.*, s.name AS status_name").
		Joins("left join public.status s on r.id = s.id").
		Order("r.created_at desc").Limit(reqPagination.PageSize).Offset((reqPagination.Page - 1) * reqPagination.PageSize)

	if reqPagination.Keyword != "" {
		query = query.Where("(lower(r.fullname) like ? OR lower(r.whatsapp_no) like ? OR lower(r.email) like ? OR r.nik like ? OR lower(r.address) like ?)", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%")
	}

	if statusID > 0 {
		query = query.Where("status_id = ? ", statusID)
	}

	if reqReseller.StartDate != "" && reqReseller.EndDate != "" {
		query = query.Where("DATE(r.created_at) BETWEEN ? AND ? ", reqReseller.StartDate, reqReseller.EndDate)
	}

	err = query.Find(&resellers).Error
	if err != nil {
		return resellers, count, err
	}

	err = query.Count(&count).Error
	if err != nil {
		return resellers, count, err
	}

	return resellers, count, nil
}

func (r *ResellerRepository) GetReseller(ID int64) (*model.Reseller, error) {
	var reseller model.Reseller
	if err := r.DB.Where("id = ? ", ID).First(&reseller).Error; err != nil {
		return nil, err
	}
	return &reseller, nil
}

func (r *ResellerRepository) UpdateReseller(reseller *model.Reseller) error {
	return r.DB.Save(reseller).Error
}

func (r *ResellerRepository) DeleteReseller(ID int64) error {
	return r.DB.Delete(model.Reseller{ID: ID}).Error
}
