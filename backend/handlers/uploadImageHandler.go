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
	"github.com/zerefwayne/that-meme/config"
)

func respondWithError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "encountered error: %+v\n", err)
}

func respondWithSuccess(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%+v\n", message)
}

func extractExtension(name string) (extension string) {

	fields := strings.Split(name, ".")

	extension = fields[len(fields)-1]

	return

}

func uploadToS3(file multipart.File, header *multipart.FileHeader) error {

	var size int64 = header.Size

	buffer, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	path := "/memes/" + header.Filename

	params := &s3.PutObjectInput{
		Bucket:        aws.String("thatmemedev"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}

	resp, err := config.Config.S3.PutObject(params)

	if err != nil {
		return err
	}

	fmt.Println(resp.String())

	return nil
}

// UploadImageHandler ...
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	// Setting maxImageSize as 1MB
	var maxImageSize int64 = 1 << 20

	// Parsing the multipart Form data
	r.ParseMultipartForm(maxImageSize)

	file, header, err := r.FormFile("newMeme")

	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer file.Close()

	err = uploadToS3(file, header)

	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	respondWithSuccess(w, "successfully uploaded the file")

}
