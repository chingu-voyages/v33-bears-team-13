package record

import (
	"context"
	// "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type RemoteRecord interface {
	Close(ctx context.Context) error
	AddSummary(ctx context.Context, summary string) error
	Read(ctx context.Context) ([]string, error)
}

type mongoDbRecord struct {
	client *mongo.Client
	coll *mongo.Collection
}

type summary struct {
	Text       string             `bson:"text"`
}

// type summary string `bson:"text"`

func Open(ctx context.Context) (RemoteRecord, error) {

	uri := "" // REPLACE WITH MONGODB URL

	// os.Getenv("MONGODB_URI") // TODO Either pass URI in to Open or have Open read the URI from
  // uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	recordCollection := client.Database("weather").Collection("record")

	return mongoDbRecord{client, recordCollection}, err
} // TODO Rather than Open, which is to do with databases, it'd be nice if the caller could instantiate a Record like a regular record (though there's always the possibility of returning an error, so that has to be accounted for in the interface)

func (r mongoDbRecord) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}

func (r mongoDbRecord) AddSummary(ctx context.Context, summary string) error {
	// summary := Summary { Text: "Current temperature in București..." }
	//
	// _, err2 := recordCollection.InsertOne(ctx, summary)
	// if err2 != nil {
	//     log.Fatal(err2)
	// }

	_, err := r.coll.InsertOne(ctx, bson.D{{Key: "text", Value: summary}}) // TODO without the key, just the value

	return err
}

func (r mongoDbRecord) Read(ctx context.Context) ([]string, error) {
	cursor, err := r.coll.Find(ctx, bson.D{})
	if err != nil {
			return nil, err
	}

	var summaries []summary
	if err = cursor.All(ctx, &summaries); err != nil {
			return nil, err
	}

	// fmt.Println(summaries) // [{Current temperature in București...} {Current temperature in București...}] // For now, map this to []string for the sake of the external interface. TODO figure out how to use bson so that you're inserting strings into the collection (so you'll get []string back).

	var summaryStrings []string
	for _, summary := range summaries {
		summaryStrings = append(summaryStrings, summary.Text)
	}

	return summaryStrings, nil
} // TODO See if using the record type, eg fmt.Println(record), can automatically call Read so you don't have to explicitly do record.Read()