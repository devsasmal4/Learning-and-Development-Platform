package mongo

import (
	"cb-ldp-backend/config"
	"cb-ldp-backend/constants"
	"context"

	"encoding/base64"
	"os"

	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang/mock/mockgen/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	MongoDbClient *mongo.Client
}
type MongoClient interface {
	RunMigrations() error
	GetCollection(string) *mongo.Collection
	InsertOne(context.Context, string, interface{}) error
	FindOne(context.Context, string, interface{}, *options.FindOneOptions) *mongo.SingleResult
	Find(context.Context, string, interface{}, *options.FindOptions) (*mongo.Cursor, error)
	FindOneAndUpdate(context.Context, string, interface{}, interface{}) *mongo.SingleResult
	FindOneAndDelete(context.Context, string, interface{}) error
	DeleteMany(context.Context, string, interface{}) error
}

var envVar = config.LoadConfig()

func ConnectDB() (MongoClient, error) {
	mongoUrl := envVar["mongo_url"].(string)
	var client *mongo.Client
	var err error
	if envVar["env"] == "development" {
		client, err = mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	} else {
		mongo_pass, err := base64.StdEncoding.DecodeString(os.Getenv("MONGO_INITDB_ROOT_PASSWORD"))
		if err != nil {
			return nil, err
		}
		credentials := options.Credential{
			Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
			Password: string(mongo_pass),
		}
		client, err = mongo.NewClient(options.Client().ApplyURI(mongoUrl).SetAuth(credentials))
	}

	if err != nil {
		log.Println("Error connecting to mongo", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB")
	return &Client{MongoDbClient: client}, nil
}

func (client *Client) RunMigrations() error {
	databaseName := envVar["database_name"].(string)
	migrationConfig := &mongodb.Config{
		DatabaseName: databaseName,
	}
	sourceURL := "file://db/migrations"
	driver, err := mongodb.WithInstance(client.MongoDbClient, migrationConfig)
	migration, err := migrate.NewWithDatabaseInstance(sourceURL, databaseName, driver)
	if err != nil {
		log.Println("Unable to create NewWithDatabaseInstance", err)

		return err
	}
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Println("Failed to run Up migration", err)
		return err
	}

	return nil
}

func (client *Client) GetCollection(collectionName string) *mongo.Collection {
	return client.MongoDbClient.Database(envVar["database_name"].(string)).Collection(collectionName)
}

func (client *Client) InsertOne(ctx context.Context, collectionName string, document interface{}) error {
	collection := client.GetCollection(collectionName)

	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) FindOne(ctx context.Context, collectionName string, filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult {
	collection := client.GetCollection(collectionName)
	result := collection.FindOne(ctx, filter)
	return result
}

func (client *Client) Find(ctx context.Context, collectionName string, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {
	collection := client.GetCollection(collectionName)
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return cursor, err

}

func (client *Client) FindOneAndUpdate(ctx context.Context, collectionName string, filter interface{}, update interface{}) *mongo.SingleResult {
	collection := client.GetCollection(collectionName)
	result := collection.FindOneAndUpdate(ctx, filter, update)
	return result
}

func (client *Client) Paginate(c *gin.Context, field string) *options.FindOptions {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	opts := options.Find().SetSort(bson.D{{field, 1}}).SetLimit(int64(limit)).SetSkip(int64(page))
	return opts
}

func (client *Client) FindOneAndDelete(ctx context.Context, collectionName string, filter interface{}) error {
	collection := client.GetCollection(collectionName)
	var deletedDoc bson.D
	err := collection.FindOneAndDelete(ctx, filter).Decode(&deletedDoc)
	return err
}

func (client *Client) DeleteMany(ctx context.Context, collectionName string, filter interface{}) error {
	collection := client.GetCollection(collectionName)
	_, err := collection.DeleteMany(ctx, filter)
	return err
}
