package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ethereum-explorer/logger"

	log "github.com/sirupsen/logrus"
)

type DB struct {
	Client *mongo.Client
	Collections map[string]*mongo.Collection
}

type Document interface{}
type Documents []interface{}

func NewDB(ctx context.Context, mongoUri string, dbName string, colNames []string) (*DB, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)

	logger.Logger.WithFields(log.Fields{
		"uri": mongoUri,
	}).Info("Connecting Mongo DB")

	// MongoDB 연결
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logger.Logger.WithError(err).Error("Mongo DB Connect Error")
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logger.Logger.WithError(err).Error("Mongo DB Ping Error")
		return nil, err
	}

	var collections map[string]*mongo.Collection

	for _, colName := range colNames {
		collections[colName] = client.Database(dbName).Collection(colName)
	}

	return &DB{client, collections}, nil
}

func (db *DB) InsertOneDocument(colName string, document Document) error {
	_, err := db.Collections[colName].InsertOne(context.TODO(), document)
	if err != nil {}
	return nil
}

func (db *DB) InsertManyDocument(colName string, documents Documents) error {
	_, err := db.Collections[colName].InsertMany(context.TODO(), documents)
	if err != nil {}
	return nil
}

func (db *DB) ReadDocument(colName string, keyName string, key string) Documents {
	filter := bson.M{keyName:key}

	cur, err := db.Collections[colName].Find(context.Background(), filter)

	var documents Documents

	if err != nil {}

	for cur.Next(context.Background()) {
		var document Document
		err := cur.Decode(&document)
		if err != nil {}
		documents = append(documents, document)
	}

	return documents
}
