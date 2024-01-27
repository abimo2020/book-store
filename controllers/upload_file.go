package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/abimo2020/book-store/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/labstack/echo/v4"
)

func UploadFileController(c echo.Context) (err error) {
	utils.LoadEnv()

	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	f, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to open file")
	}
	defer f.Close()

	endpoint := os.Getenv("BUCKET_ENDPOINT")
	region := "nyc3"
	session := session.Must(session.NewSession(&aws.Config{
		Endpoint:    &endpoint,
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(os.Getenv("BUCKET_ACCESS_KEY"), os.Getenv("BUCKET_SECRET_KEY"), ""),
	}))
	uploader := s3manager.NewUploader(session)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(file.Filename),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("error: %v", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to upload file")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success to upload file",
		"image_path": result.Location,
	})
}
