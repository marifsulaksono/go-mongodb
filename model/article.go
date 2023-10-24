package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article struct {
	Id     primitive.ObjectID `bson:"_id" json:"id"`
	Title  string             `bson:"title" json:"title"`
	Body   string             `bson:"body" json:"body"`
	Tags   []string           `bson:"tags" json:"tags"`
	Author string             `bson:"author" json:"author"`
}

type ArticleRepository struct {
	DB *mongo.Database
}

func NewArticleRepository(db *mongo.Database) ArticleRepository {
	return ArticleRepository{
		DB: db,
	}
}

func (a *ArticleRepository) GetAllArticles(ctx context.Context) ([]Article, error) {
	cursor, err := a.DB.Collection("articles").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var articles []Article
	for cursor.Next(ctx) {
		var article Article
		err := cursor.Decode(&article)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return articles, err
}

func (a *ArticleRepository) GetArticleById(ctx context.Context, id string) (Article, error) {
	var article *Article
	articleId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Article{}, err
	}

	filter := bson.M{"_id": articleId}
	err = a.DB.Collection("articles").FindOne(ctx, filter).Decode(&article)
	if err != nil {
		return Article{}, err
	}

	return *article, err
}

func (a *ArticleRepository) InsertNewArticle(ctx context.Context, article *Article) error {
	data := bson.D{
		{Key: "title", Value: article.Title},
		{Key: "body", Value: article.Body},
		{Key: "tags", Value: article.Tags},
		{Key: "author", Value: article.Author},
	}
	_, err := a.DB.Collection("articles").InsertOne(ctx, data)
	return err
}

func (a *ArticleRepository) UpdateArticle(ctx context.Context, id string, article *Article) error {
	articleId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: articleId}}
	dataUpdate := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: article.Title},
		{Key: "body", Value: article.Body},
		{Key: "tags", Value: article.Tags},
		{Key: "author", Value: article.Author},
	}}}

	_, err = a.DB.Collection("articles").UpdateOne(ctx, filter, dataUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleRepository) DeleteArticle(ctx context.Context, id string) error {
	articleId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": articleId}
	_, err = a.DB.Collection("articles").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
