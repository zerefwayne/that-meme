package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/utils"
)

func generateUniqueName(name string) (filename string) {

	fields := strings.Split(name, ".")

	extension := fields[len(fields)-1]

	newFileName := uuid.New().String()

	filename = fmt.Sprintf("m-%s.%s", newFileName, extension)

	return filename

}

func uploadToS3(file multipart.File, header *multipart.FileHeader) (string, error) {

	var size int64 = header.Size

	buffer, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	fileName := generateUniqueName(header.Filename)

	path := "/memes/" + fileName

	params := &s3.PutObjectInput{
		Bucket:        aws.String("thatmemedev"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}

	_, err = config.Config.S3.PutObject(params)

	if err != nil {
		return "", err
	}

	bucketName := "thatmemedev"
	regionName := "ap-south-1"

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com%s", bucketName, regionName, path)

	return fileURL, nil
}

// UploadImageHandler ...
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	// Setting maxImageSize as 1MB
	var maxImageSize int64 = 1 << 20

	// Parsing the multipart Form data
	r.ParseMultipartForm(maxImageSize)

	file, header, err := r.FormFile("newMeme")

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer file.Close()

	fileURL, err := uploadToS3(file, header)

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	utils.RespondWithSuccess(w, fileURL)

}
