package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/models"
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

	// Reading the file

	file, header, err := r.FormFile("newMeme")

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer file.Close()

	// File Upload

	fileURL, err := uploadToS3(file, header)

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Reading the rest form data

	meme := new(models.Meme)

	meme.FileURL = fileURL
	meme.Name = r.FormValue("name")
	meme.Tags = strings.Split(r.FormValue("tags"), ";")
	meme.Description = r.FormValue("description")
	meme.Text = r.FormValue("text")
	meme.Origin = r.FormValue("origin")
	meme.CreatedAt = time.Now()
	meme.UpdatedAt = time.Now()

	err = models.InsertMeme(meme)

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*meme)

}
