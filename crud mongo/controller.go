package controllers

import (
	"context"
	"crudproject/model"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb+srv://kadambalaaditya99_db_user:vrETiXTvprdAdtzQ@cluster0.z09ikv3.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	dbName           = "movies"
	collectionName   = "trendlist"
)

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("Error connecting to MongoDB: " + err.Error())
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic("MongoDB ping failed: " + err.Error())
	}
	collection = client.Database(dbName).Collection(collectionName)
}

// -------------------- CREATE --------------------
func Insertdata(m model.Movie) (primitive.ObjectID, error) {
	m.Id = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertedRec, err := collection.InsertOne(ctx, m)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return insertedRec.InsertedID.(primitive.ObjectID), nil
}

func CreateHandler(c *gin.Context) {
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := Insertdata(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie created",
		"id":      id,
		"movie":   movie,
	})
}

// -------------------- READ --------------------
func Getmongo() ([]model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []model.Movie
	for cursor.Next(ctx) {
		var movie model.Movie
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func Readhandler(c *gin.Context) {
	movies, err := Getmongo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": movies})
}

// -------------------- UPDATE --------------------
func UpdateRecord(id string) (*model.Movie, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"watched": true}}

	var updatedMovie model.Movie
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedMovie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &updatedMovie, nil
}

func Updatehandler(c *gin.Context) {
	id := c.Param("id")
	movie, err := UpdateRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if movie == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No movie found with this ID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": movie})
}

// -------------------- DELETE --------------------
func DeleteRecord(id string) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, errors.New("invalid ID format")
	}

	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	count, err := DeleteRecord(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No record found to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted_count": count})
}
