package service

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"reseller-jh-be/base"
	"reseller-jh-be/internal/reseller/model"
	"reseller-jh-be/internal/reseller/repository"
	"reseller-jh-be/internal/reseller/request"
	"reseller-jh-be/internal/reseller/response"
	"reseller-jh-be/pkg/common"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

func (s *ResellerService) CreateReseller(c *gin.Context, reqReseller *model.Reseller, file *multipart.FileHeader) (reseller *model.Reseller, err error) {
	filePath, err := common.UploadFile(c, file, "")
	if err != nil {
		return nil, err
	}

	reqReseller.StatusID = 1
	reqReseller.KTP = filePath
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

func (s *ResellerService) ExportExcelResellers(reqReseller request.ReqReseller) (resp response.RespExportReseller, err error) {
	statusID, _ := strconv.ParseInt(reqReseller.Type, 10, 64)
	resellers, err := s.Repo.ExportExcelResellers(statusID, reqReseller)
	if err != nil {
		return resp, err
	}
	localLocation := time.Now().Location()
	parseStartDate, err := time.ParseInLocation("2006-01-02", reqReseller.StartDate, localLocation)
	if err != nil {
		return resp, err
	}
	startDate := parseStartDate.Format("02 Jan 2006")
	parseEndDate, err := time.ParseInLocation("2006-01-02", reqReseller.EndDate, localLocation)
	if err != nil {
		return resp, err
	}
	endDate := parseEndDate.Format("02 Jan 2006")
	f := excelize.NewFile()

	sheetName := "Sheet1"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Periode")
	f.SetCellValue(sheetName, "B1", startDate+" - "+endDate)

	f.SetCellValue(sheetName, "A3", "Nama Lengkap")
	f.SetCellValue(sheetName, "B3", "Whatsapp")
	f.SetCellValue(sheetName, "C3", "NIK")
	f.SetCellValue(sheetName, "D3", "Alamat")
	f.SetCellValue(sheetName, "E3", "Status")
	f.SetCellValue(sheetName, "F3", "Email")
	f.SetCellValue(sheetName, "G3", "Tanggal Daftar")

	for i, reseller := range resellers {
		row := i + 4
		formattedCreatedAt := reseller.CreatedAt.Format("02 Jan 2006")
		f.SetCellValue(sheetName, "A"+strconv.Itoa(row), reseller.Fullname)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(row), reseller.WhatsappNo)
		f.SetCellValue(sheetName, "C"+strconv.Itoa(row), reseller.NIK)
		f.SetCellValue(sheetName, "D"+strconv.Itoa(row), reseller.Address)
		f.SetCellValue(sheetName, "E"+strconv.Itoa(row), reseller.StatusName)
		f.SetCellValue(sheetName, "F"+strconv.Itoa(row), reseller.Email)
		f.SetCellValue(sheetName, "G"+strconv.Itoa(row), formattedCreatedAt)
	}

	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		os.Mkdir("./uploads", os.ModePerm)
	}

	folderPath := "./uploads/"
	fileName := "reseller.xlsx"
	filePath := filepath.Join(folderPath, fileName)
	if err := f.SaveAs(filePath); err != nil {
		return resp, err
	}

	resp.Filename = os.Getenv("DIR_FILE") + "uploads/" + fileName

	return resp, nil
}
