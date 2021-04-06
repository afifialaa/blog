package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/afifialaa/blog/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	articlePrimID, _ := primitive.ObjectIDFromHex(r.FormValue("article_id"))
	comment := models.Comment{
		ID:         primitive.NewObjectID(),
		User:       r.FormValue("user"),
		Body:       r.FormValue("body"),
		Posted_At:  time.Now(),
		Article_ID: articlePrimID,
	}

	// Validate article

	var result bool = models.CreateComment(comment)
	if result == true {
		models.PostCommentES(comment)
		res := map[string]string{"msg": "Comment was created"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"err": "Failed to create comment"}
	json.NewEncoder(w).Encode(res)
	return
}

func FetchComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	var article string = r.FormValue("article_id")
	result := models.FetchComments(article)

	json.NewEncoder(w).Encode(result)
	return
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	var id string = r.FormValue("id")
	result := models.DeleteComment(id)

	if result == false {
		res := map[string]string{"err": "Failed to delete comment"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"msg": "Comment was deleted"}
	json.NewEncoder(w).Encode(res)
	return
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	commentPrimID, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
	comment := models.Comment{
		ID:   commentPrimID,
		User: r.FormValue("user"),
		Body: r.FormValue("body"),
	}

	result := models.UpdateComment(comment)

	if result == false {
		res := map[string]string{"err": "Failed to update comment"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"msg": "Comment was updated"}
	json.NewEncoder(w).Encode(res)
	return
}
