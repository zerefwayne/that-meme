package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/utils"
)

type result struct {
	ID     string      `json:"id"`
	Score  float64     `json:"score"`
	Result interface{} `json:"result"`
}

type searchResponse struct {
	Hits    int64    `json:"hits"`
	Time    int64    `json:"time"`
	Results []result `json:"results"`
}

// SearchHandler ...
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	searchQuery := r.URL.Query().Get("q")

	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  searchQuery,
				"fields": []string{"name", "text", "tags^2", "description"},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	es := config.Config.ES

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("memes"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	if res.IsError() {

		var errBody map[string]interface{}

		err := json.NewDecoder(res.Body).Decode(&errBody)

		if err != nil {
			utils.RespondWithError(w, err, http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		errType := errBody["error"].(map[string]interface{})["type"]
		errReason := errBody["error"].(map[string]interface{})["reason"]

		returnError := fmt.Errorf("type: %+v, reason %+v", errType, errReason)

		utils.RespondWithError(w, returnError, http.StatusInternalServerError)
		fmt.Println(err)

		return

	}

	resBody := make(map[string]interface{})

	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	response := new(searchResponse)

	response.Hits = int64(resBody["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	response.Time = int64(resBody["took"].(float64))

	for _, hit := range resBody["hits"].(map[string]interface{})["hits"].([]interface{}) {

		newResult := new(result)

		newResult.ID = hit.(map[string]interface{})["_id"].(string)
		newResult.Score = hit.(map[string]interface{})["_score"].(float64)
		newResult.Result = hit.(map[string]interface{})["_source"]

		response.Results = append(response.Results, *newResult)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

}
