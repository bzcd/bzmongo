package bzmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Collection *mongo.Collection
}

// Get one by id
func (c *Collection) Get(ctx context.Context, out interface{}, filter interface{}) (interface{}, error) {
	err := c.Collection.FindOne(ctx, filter).Decode(out)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return out, nil
}

// Gets more by filter
func (c *Collection) Gets(ctx context.Context, out interface{}, filter interface{}) error {
	cur, err := c.Collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	return cur.All(ctx, out)
}

// Insert document
func (c *Collection) Insert(ctx context.Context, document interface{}) error {
	_, err := c.Collection.InsertOne(ctx, document)
	return err
}

// Update one
func (c *Collection) Update(ctx context.Context, filter interface{}, update interface{}) error {
	_, err := c.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (c *Collection) InsertOrUpdate(ctx context.Context, filter interface{}, update interface{}) {
	c.Collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true))
}
