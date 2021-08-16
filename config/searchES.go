package config

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/afifialaa/blog/database"
)

type SearchResult struct {
    Title string
    Content string
    User string
}

// Returns map[string]interface{}

func SearchArticlesES(word string) []SearchResult {

	ES := database.GetESClient()

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": word,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := ES.Search(
		ES.Search.WithContext(context.Background()),
		ES.Search.WithIndex("article"),
		ES.Search.WithBody(&buf),
		ES.Search.WithTrackTotalHits(true),
		ES.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

    var result []SearchResult

    for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

        // Parse the attributes/fields of the document
        doc := hit.(map[string]interface{})

        // The "_source" data is another map[string]interface{} nested inside of doc
        source := doc["_source"]
        // docID := doc["_id"]
        content := source.(map[string]interface{})["content"].(string)
        title := source.(map[string]interface{})["title"].(string)
        user := source.(map[string]interface{})["user"].(string)

        item := SearchResult{title, content, user}

        result = append(result, item)

        // Get the document's _id and print it out along with _source data
        // fmt.Println("docID:", docID)
    }

	return result
}
