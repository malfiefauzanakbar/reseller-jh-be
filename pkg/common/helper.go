package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	mathrand "math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	key := os.Getenv("ENCRYPTION_KEY")
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
	key := os.Getenv("ENCRYPTION_KEY")
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
	mathrand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[mathrand.Intn(len(letterBytes))]
	}
	return string(b)
}

func UploadFile(c *gin.Context, file *multipart.FileHeader, fileName string) (resp string, err error) {
	if file != nil {
		if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
			os.Mkdir("./uploads", os.ModePerm)
		}

		if fileName == "" {
			fileName = RandString(10)
		}
		destination := "./uploads/"
		ext := filepath.Ext(file.Filename)
		filePath := filepath.Join(destination, fileName+ext)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			return resp, err
		}

		resp = os.Getenv("DIR_FILE") + filePath
	}

	return resp, nil
}

func VerifyCaptcha(token string) bool {
	secret := os.Getenv("RECAPTCHA_SECRET_KEY")
	verificationURL := os.Getenv("RECAPTCHA_HOST")

	resp, err := http.PostForm(verificationURL, url.Values{
		"secret":   {secret},
		"response": {token},
	})

	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body); err != nil {
		return false
	}

	return true
}
