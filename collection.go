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
func (c *Collection) Get(ctx context.Context, out interface{}, filter interface{}, opts ...*options.FindOneOptions) (interface{}, error) {
	coll, _ := c.Collection.Clone()
	err := coll.FindOne(ctx, filter, opts...).Decode(out)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return out, nil
}

// Gets more by filter
func (c *Collection) Gets(ctx context.Context, out interface{}, filter interface{}, opts ...*options.FindOptions) error {
	coll, _ := c.Collection.Clone()
	cur, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return cur.All(ctx, out)
}

// Insert document
func (c *Collection) Insert(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) error {
	coll, _ := c.Collection.Clone()
	_, err := coll.InsertOne(ctx, document, opts...)
	return err
}

// InsertMany document
func (c *Collection) InsertMany(ctx context.Context, document []interface{}, opts ...*options.InsertManyOptions) error {
	coll, _ := c.Collection.Clone()
	_, err := coll.InsertMany(ctx, document, opts...)
	return err
}

// Update one
func (c *Collection) Update(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	coll, _ := c.Collection.Clone()
	_, err := coll.UpdateOne(ctx, filter, update, opts...)
	return err
}

func (c *Collection) InsertOrUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) {
	coll, _ := c.Collection.Clone()
	opts = append(opts, options.FindOneAndUpdate().SetUpsert(true))
	coll.FindOneAndUpdate(ctx, filter, update, opts...)
}

func (c *Collection) InsertOrUpdate2(ctx context.Context, out interface{}, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (interface{}, error) {
	coll, _ := c.Collection.Clone()
	opts = append(opts, options.FindOneAndUpdate().SetUpsert(true))
	err := coll.FindOneAndUpdate(ctx, filter, update, opts...).Decode(out)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Collection) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	coll, _ := c.Collection.Clone()
	return coll.CountDocuments(ctx, filter, opts...)
}
