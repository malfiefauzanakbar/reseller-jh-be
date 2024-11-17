package service

import (
	"strconv"

	"reseller-jh-be/base"
	"reseller-jh-be/internal/reseller/model"
	"reseller-jh-be/internal/reseller/repository"
	"reseller-jh-be/internal/reseller/request"
	"reseller-jh-be/internal/reseller/response"
	"reseller-jh-be/pkg/common"
)

type ResellerService struct {
	Repo repository.ResellerRepository
}

type ResellerServiceInterface interface {
}

func NewResellerService(repo *repository.ResellerRepository) *ResellerService {
	return &ResellerService{
		Repo: *repo,
	}
}

func (s *ResellerService) CreateReseller(reqReseller *model.Reseller) (reseller *model.Reseller, err error) {
	reqReseller.StatusID = 1
	reseller, err = s.Repo.CreateReseller(reqReseller)
	if err != nil {
		return nil, err
	}

	return reseller, nil
}

func (s *ResellerService) GetAllReseller(reqReseller request.ReqReseller, reqPagination base.ReqPagination) (resellers []response.RespReseller, pagination base.Pagination, err error) {
	statusID, _ := strconv.ParseInt(reqReseller.Type, 10, 64)
	resellers, count, err := s.Repo.GetAllReseller(statusID, reqReseller, reqPagination)
	if err != nil {
		return nil, pagination, err
	}

	pagination = common.HandlePagination(int(count), reqPagination)

	return resellers, pagination, nil
}

func (s *ResellerService) GetReseller(id string) (reseller *model.Reseller, err error) {
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	reseller, err = s.Repo.GetReseller(ID)
	if err != nil {
		return nil, err
	}

	return reseller, nil
}

func (s *ResellerService) UpdateReseller(id string, reqReseller *model.Reseller) (reseller *model.Reseller, err error) {
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	reseller, err = s.Repo.GetReseller(ID)
	if err != nil {
		return nil, err
	}

	reseller.Fullname = reqReseller.Fullname
	reseller.WhatsappNo = reqReseller.WhatsappNo
	reseller.WhatsappLink = reqReseller.WhatsappLink
	reseller.Email = reqReseller.Email
	reseller.NIK = reqReseller.NIK
	reseller.Address = reqReseller.Address
	reseller.StatusID = reqReseller.StatusID
	err = s.Repo.UpdateReseller(reseller)
	if err != nil {
		return nil, err
	}

	return reseller, nil
}

func (s *ResellerService) ReadReseller(id string) (reseller *model.Reseller, err error) {
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	reseller, err = s.Repo.GetReseller(ID)
	if err != nil {
		return nil, err
	}

	reseller.StatusID = 2
	err = s.Repo.UpdateReseller(reseller)
	if err != nil {
		return nil, err
	}

	return reseller, nil
}

func (s *ResellerService) DeleteReseller(id string) (err error) {
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	_, err = s.Repo.GetReseller(ID)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteReseller(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ResellerService) CountResellers(reqReseller request.ReqReseller) (reseller response.RespResellerDashboard, err error) {
	reseller, err = s.Repo.CountResellers(reqReseller)
	if err != nil {
		return reseller, err
	}

	return reseller, nil
}

func (s *ResellerService) ResellersChart(reqReseller request.ReqReseller) (resp response.RespResellerChart, err error) {
	resp, err = s.Repo.ResellersChart(reqReseller)
	if err != nil {
		return resp, err
	}

	return resp, nil
}