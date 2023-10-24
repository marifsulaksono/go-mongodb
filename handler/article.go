package handler

import (
	"encoding/json"
	"mongodb/model"
	"net/http"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	Repo model.ArticleRepository
}

func NewArticleHandler(repo model.ArticleRepository) ArticleHandler {
	return ArticleHandler{
		Repo: repo,
	}
}

func (a *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	result, err := a.Repo.GetAllArticles(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (a *ArticleHandler) GetArticleById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	result, err := a.Repo.GetArticleById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (a *ArticleHandler) InsertNewArticle(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		article model.Article
	)
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.Repo.InsertNewArticle(ctx, &article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success insert article"))
}

func (a *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	var article model.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = a.Repo.UpdateArticle(ctx, id, &article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success update article"))
}

func (a *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := a.Repo.DeleteArticle(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success delete article"))
}
