package bzmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	*mongo.Collection
}

func (c *Collection) FindMany(ctx context.Context, out interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cur, err := c.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return cur.All(ctx, out)
}
