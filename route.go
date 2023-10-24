package main

import (
	"mongodb/handler"
	"mongodb/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func routeInit(conn *mongo.Database) *mux.Router {
	articleRepository := model.NewArticleRepository(conn)
	articleHandler := handler.NewArticleHandler(articleRepository)

	r := mux.NewRouter()

	r.HandleFunc("/articles", CORSMiddleware(articleHandler.GetAllArticles)).Methods(http.MethodGet)
	r.HandleFunc("/articles", CORSMiddleware(articleHandler.InsertNewArticle)).Methods(http.MethodPost)
	r.HandleFunc("/articles/{id}", CORSMiddleware(articleHandler.UpdateArticle)).Methods(http.MethodPut)
	r.HandleFunc("/articles/{id}", CORSMiddleware(articleHandler.DeleteArticle)).Methods(http.MethodDelete)
	r.HandleFunc("/articles/{id}", CORSMiddleware(articleHandler.GetArticleById)).Methods(http.MethodGet)

	return r
}

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}
