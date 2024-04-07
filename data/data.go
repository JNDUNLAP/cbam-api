package data

import (
	"cbam_api/model"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func isAllowedCollectionName(collectionName string) bool {
	return collectionName == os.Getenv("REPORTCOLLECTION")
}

type MongoDBClient struct {
	Client   *mongo.Client
	dataName string
}

type ReportRepository interface {
	GetQReport(filter interface{}) (*model.QReport, error)
}

func DBClient(mongoURI string) (*MongoDBClient, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoDBClient{
		Client: client,
	}, nil
}

func UploadReport(ctx context.Context, report *model.QReport, m *MongoDBClient, collectionName string) error {

	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": report}

	_, err := m.Client.Database("Mongo_Test").Collection(collectionName).UpdateOne(ctx, bson.M{"ReportId": report.ReportId}, update, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert quarterly report: %w", err)
	}

	return nil
}

func (m *MongoDBClient) DeleteReports(collectionName string) error {
	if !isAllowedCollectionName(collectionName) {
		return fmt.Errorf("operation not allowed on collection: %s", collectionName)
	}
	_, err := m.Client.Database(m.dataName).Collection(collectionName).DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to delete all quarterly reports: %w", err)
	}
	return nil
}

func SetupDatabase(mongoURI string) (*MongoDBClient, error) {
	client, err := DBClient(mongoURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	return client, nil
}
