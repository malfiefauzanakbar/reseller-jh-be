package common

import (
	"strconv"
	"strings"

	"reseller-jh-be/base"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleReqPagination(c *gin.Context) base.ReqPagination {
	keyword := strings.ToLower(c.Query("keyword"))
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	reqPagination := base.ReqPagination{
		Keyword:  keyword,
		Page:     page,
		PageSize: pageSize,
	}

	return reqPagination
}

func HandlePagination(totalItems int, reqPagination base.ReqPagination) (pagination base.Pagination) {
	totalPages := (totalItems + reqPagination.PageSize - 1) / reqPagination.PageSize

	pagination.Page = reqPagination.Page
	pagination.PageSize = reqPagination.PageSize
	pagination.TotalPages = totalPages
	pagination.TotalItems = totalItems

	return pagination
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
