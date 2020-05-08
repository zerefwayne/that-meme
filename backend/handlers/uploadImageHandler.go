package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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

// AddToIndex ...
func AddToIndex(m *models.Meme) error {

	body, err := json.Marshal(m)

	if err != nil {
		return err
	}

	bodyString := string(body)

	req := esapi.IndexRequest{
		Index:   "memes",
		Body:    strings.NewReader(bodyString),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), config.Config.ES)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	return nil

}

// UploadImageHandler ...
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	// REQUEST PARSE ====================================================

	var maxImageSize int64 = 1 << 20

	r.ParseMultipartForm(maxImageSize)

	file, header, err := r.FormFile("newMeme")

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer file.Close()

	// REQUEST PARSE ====================================================

	// S3 START ==========================================================

	fileURL, err := uploadToS3(file, header)

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// S3 END ============================================================

	// MONGODB START =====================================================

	meme := new(models.Meme)

	meme.FileURL = fileURL
	meme.Name = r.FormValue("name")
	meme.Tags = r.FormValue("tags")
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

	// MONGODB END =====================================================

	// ADD TO ES =======================================================

	err = AddToIndex(meme)

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// ADD TO ES =======================================================

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*meme)

}
