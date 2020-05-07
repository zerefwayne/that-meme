package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

// UploadImageHandler ...
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	// Setting maxImageSize as 1MB
	var maxImageSize int64 = 1 << 20

	// Parsing the multipart Form data
	r.ParseMultipartForm(maxImageSize)

	file, handler, err := r.FormFile("newMeme")

	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("name: %+v\nsize: %+v\nheader: %+v\n", handler.Filename, handler.Size, handler.Header)

	extension := extractExtension(handler.Filename)

	newFileName := fmt.Sprintf("meme*.%s", extension)

	tempfile, err := ioutil.TempFile("data/memes", newFileName)

	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer tempfile.Close()

	imageBytes, err := ioutil.ReadAll(file)

	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	tempfile.Write(imageBytes)

	respondWithSuccess(w, "successfully uploaded the file")

}
