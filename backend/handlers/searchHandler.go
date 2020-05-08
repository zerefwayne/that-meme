package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/utils"
)

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
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	resBody := make(map[string]interface{})

	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Println(resBody)

	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(resBody["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(resBody["took"].(float64)),
	)

	for _, hit := range resBody["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successful search")

}
