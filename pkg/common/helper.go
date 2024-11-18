package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	math "math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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

func Encrypt(plainText string) (string, error) {
	key := os.Getenv("ENCYPTION_KEY")
	keyBytes := []byte(key)
	plainTextBytes := []byte(plainText)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, plainTextBytes, nil)

	result := append(nonce, cipherText...)

	return base64.StdEncoding.EncodeToString(result), nil
}

func Decrypt(cipherText string) (string, error) {
	key := os.Getenv("ENCYPTION_KEY")
	keyBytes := []byte(key)
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherTextBytes) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, cipherTextBytes := cipherTextBytes[:nonceSize], cipherTextBytes[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	math.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[math.Intn(len(letterBytes))]
	}
	return string(b)
}
