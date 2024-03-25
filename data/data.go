package data

import (
	"cbam_api/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Client   *mongo.Client
	dataName string
}

type ReportRepository interface {
	GetQReport(filter interface{}) (*model.QReport, error)
}

func DBClient(uri, dataName string) (*MongoDBClient, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}
	return &MongoDBClient{
		Client:   client,
		dataName: dataName,
	}, nil
}

func (m *MongoDBClient) GetQReport(filter bson.M) (model.QReport, error) {
	var qReport model.QReport
	collection := m.Client.Database(m.dataName).Collection("test_report")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&qReport)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.QReport{}, fmt.Errorf("no document found with the provided filter")
		}
		log.Println(err)
		return model.QReport{}, fmt.Errorf("error fetching document: %w", err)
	}
	return qReport, nil
}

func SaveQuarterlyReportToDatabase(report *model.QReport, m *MongoDBClient, collectionName string) error {
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": report}
	_, err := m.Client.Database(m.dataName).Collection(collectionName).UpdateOne(context.Background(), bson.M{"ReportId": report.DraftReportId}, update, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert quarterly report: %w", err)
	}
	return nil
}

func (m *MongoDBClient) DeleteAllQReports(collectionName string) error {
	_, err := m.Client.Database(m.dataName).Collection(collectionName).DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to delete all quarterly reports: %w", err)
	}
	return nil
}

func SetupDatabase(MongoUri, dataName string) *MongoDBClient {
	client, err := DBClient(MongoUri, dataName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	return client
}
