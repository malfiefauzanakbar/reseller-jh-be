package repository

import (	
	"reseller-jh-be/base"
	"reseller-jh-be/internal/reseller/model"
	"reseller-jh-be/internal/reseller/request"
	"reseller-jh-be/internal/reseller/response"
	"time"

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
		Joins("inner join public.status s on r.status_id = s.id").
		Order("r.created_at desc").Limit(reqPagination.PageSize).Offset((reqPagination.Page - 1) * reqPagination.PageSize)

	if reqPagination.Keyword != "" {
		query = query.Where("(lower(r.fullname) like ? OR lower(r.whatsapp_no) like ? OR lower(r.email) like ? OR r.nik like ? OR lower(r.address) like ?)", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%", "%"+reqPagination.Keyword+"%")
	}

	if statusID > 0 {
		query = query.Where("r.status_id = ? ", statusID)
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

func (r *ResellerRepository) CountResellers(reqReseller request.ReqReseller) (reseller response.RespResellerDashboard, err error) {
	reseller.Total, err = r.FilterCountResellers(reqReseller, "total")
	if err != nil {
		return reseller, err
	}

	reseller.Unread, err = r.FilterCountResellers(reqReseller, "unread")
	if err != nil {
		return reseller, err
	}

	reseller.Read, err = r.FilterCountResellers(reqReseller, "read")
	if err != nil {
		return reseller, err
	}

	return reseller, nil
}

func (r *ResellerRepository) FilterCountResellers(reqReseller request.ReqReseller, countType string) (result int64, err error) {
	var reseller response.RespResellerDashboard
	query := r.DB.Model(&model.Reseller{})
	if reqReseller.StartDate != "" && reqReseller.EndDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", reqReseller.StartDate, reqReseller.EndDate)
	}
	if countType == "unread" {
		if err = query.Where("status_id = 1").Count(&reseller.Unread).Error; err != nil {
			return result, err
		}
		result = reseller.Unread
	} else if countType == "read" {
		if err = query.Where("status_id = 2").Count(&reseller.Read).Error; err != nil {
			return result, err
		}
		result = reseller.Read
	} else {
		if err = query.Count(&reseller.Total).Error; err != nil {
			return result, err
		}
		result = reseller.Total
	}

	return result, nil
}

func (r *ResellerRepository) ResellersChart(reqReseller request.ReqReseller) (resp response.RespResellerChart, err error) {
	var results []response.ResellerChart
	now := time.Now()
	today := now.Format("2006-01-02 15:04:05")

	query := `
		SELECT 
			TO_CHAR(created_at, 'DD Mon') AS date,
			COUNT(*) AS count
		FROM resellers
		WHERE DATE(created_at) BETWEEN $1 AND $2
		GROUP BY TO_CHAR(created_at, 'DD Mon')
		ORDER BY MIN(created_at)
	`

	if reqReseller.StartDate == "" || reqReseller.EndDate == "" {
		reqReseller.StartDate = today
		reqReseller.EndDate = today
	}
	if err = r.DB.Raw(query, reqReseller.StartDate, reqReseller.EndDate).Scan(&results).Error; err != nil {
		return resp, err
	}

	var categories []string
	var data []int

	for _, result := range results {
		categories = append(categories, result.Date)
		data = append(data, result.Count)
	}

	resp.Categories = categories
	resp.Data = data

	return resp, nil
}

func (r *ResellerRepository)ExportExcelResellers(statusID int64, reqReseller request.ReqReseller) (resellers []response.RespReseller, err error) {		
	query := r.DB.Table("resellers AS r").Select("r.*, s.name AS status_name").
		Joins("inner join public.status s on r.status_id = s.id").
		Order("r.created_at desc")	

	if statusID > 0 {
		query = query.Where("r.status_id = ? ", statusID)
	}

	if reqReseller.StartDate != "" && reqReseller.EndDate != "" {
		query = query.Where("DATE(r.created_at) BETWEEN ? AND ? ", reqReseller.StartDate, reqReseller.EndDate)
	}

	err = query.Find(&resellers).Error
	if err != nil {
		return resellers, err
	}

	return resellers, err
}