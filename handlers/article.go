package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/afifialaa/blog/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	article := models.Article{
		ID:      primitive.NewObjectID(),
		User:    r.FormValue("user"),
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	// Validate article

	var result bool = models.CreateArticle(article)
	if result == true {
		res := map[string]string{"msg": "Article was created"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"err": "Failed to create article"}
	json.NewEncoder(w).Encode(res)
	return
}

//
func FetchArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	var user string = r.FormValue("user")
	result := models.FetchArticles(user)

	json.NewEncoder(w).Encode(result)
	return
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	var id string = r.FormValue("id")
	result := models.DeleteArticle(id)

	if result == false {
		res := map[string]string{"err": "Failed to delete article"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"msg": "Article was deleted"}
	json.NewEncoder(w).Encode(res)
	return

}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	idPrimitive, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
	article := models.Article{
		ID:      idPrimitive,
		User:    r.FormValue("user"),
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	result := models.UpdateArticle(article)

	if result == false {
		res := map[string]string{"err": "Failed to update article"}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]string{"msg": "Article was updated"}
	json.NewEncoder(w).Encode(res)
	return
}
