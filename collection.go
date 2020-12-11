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
	coll, _ := c.Collection.Clone()
	err := coll.FindOne(ctx, filter).Decode(out)
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
	coll, _ := c.Collection.Clone()
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}

	return cur.All(ctx, out)
}

// Insert document
func (c *Collection) Insert(ctx context.Context, document interface{}) error {
	coll, _ := c.Collection.Clone()
	_, err := coll.InsertOne(ctx, document)
	return err
}

// Update one
func (c *Collection) Update(ctx context.Context, filter interface{}, update interface{}) error {
	coll, _ := c.Collection.Clone()
	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}

func (c *Collection) InsertOrUpdate(ctx context.Context, filter interface{}, update interface{}) {
	coll, _ := c.Collection.Clone()
	coll.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true))
}

func (c *Collection) InsertOrUpdate2(ctx context.Context, out interface{}, filter interface{}, update interface{}) (interface{}, error) {
	coll, _ := c.Collection.Clone()
	err := coll.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true)).Decode(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Collection) Count(ctx context.Context, filter interface{}) (int64, error) {
	coll, _ := c.Collection.Clone()
	return coll.CountDocuments(ctx, filter)
}
