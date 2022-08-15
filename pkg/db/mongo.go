package db

import (
	"context"
	"log"
	"time"

	sf "github.com/sa-/slicefunk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Document[T any] struct {
	value T
}

type Filter[F any] struct {}

// connects to MongoDB
func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dbConfig.ConnectionString).SetDirect(true)
	c, _ := mongo.NewClient(clientOptions)

	err := c.Connect(ctx)

	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	return c
}

func Insert[T any](value Document[T], collectionName string) (any, error) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	collection := c.Database(dbConfig.DBName).Collection(collectionName)

	r, err := collection.InsertOne(ctx, value)

	return r.InsertedID, err
}

func InsertMany[T any](values []Document[T], collectionName string) ([]any, error) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	collection := c.Database(dbConfig.DBName).Collection(collectionName)

	// Convert input to slice of any to satisfy InsertMany
	converted := sf.Map(values, func(item Document[T]) any { return item })
	r, err := collection.InsertMany(ctx, converted)

	return r.InsertedIDs, err
}

func Query[F, T any](filter Filter[F], tableName string) ([]Document[T], error) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	collection := c.Database(dbConfig.ConnectionString).Collection(dbConfig.DBName)
	rs, err := collection.Find(ctx, filter)
	if err != nil {
		return []Document[T]{}, err
	}

	var results []Document[T]
	err = rs.All(ctx, &results)

	return results, err
}

func Update[T any](itemId, collectionName string, newValue Document[T]) error {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	collection := c.Database(dbConfig.DBName).Collection(collectionName)
	oid, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		 return err
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{{Key: "$set", Value: newValue}}
	_, err = collection.UpdateOne(ctx, filter, update)

	return err
}

func UpdateMany[T any](ids []string, newValues []Document[T], collectionName string) error {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	collection := c.Database(dbConfig.DBName).Collection(collectionName)

	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}
	update := bson.D{{Key: "$replaceWith", Value: newValues}}
	_, err := collection.UpdateMany(ctx, filter, update)

	return err
}

func Delete(itemId, collectionName string) error {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	collection := c.Database(dbConfig.ConnectionString).Collection(dbConfig.DBName)
	oid, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	_, err = collection.DeleteOne(ctx, filter)
	return err
}

func DeleteMany(itemIds []string, collectionName string) error {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	collection := c.Database(dbConfig.ConnectionString).Collection(dbConfig.DBName)
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: itemIds}}}}

	_, err := collection.DeleteMany(ctx, filter)
	return err
}
