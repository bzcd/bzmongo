package bzmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	cli *mongo.Client
}

func New(opt ClientOptions) (*Mongo, error) {
	cli, err := mongo.NewClient(opt.ClientOptions())
	if err != nil {
		return nil, err
	}

	return &Mongo{
		cli: cli,
	}, nil
}

func Connect(ctx context.Context, opt ClientOptions) (*Mongo, error) {
	cli, err := mongo.Connect(ctx, opt.ClientOptions())
	if err != nil {
		return nil, err
	}

	return &Mongo{
		cli: cli,
	}, nil
}

func (m *Mongo) Client() *mongo.Client {
	return m.cli
}

func (m *Mongo) Connect(ctx context.Context) error {
	return m.cli.Connect(ctx)
}

func (m *Mongo) Disconnect(ctx context.Context) {
	m.cli.Disconnect(ctx)
}

func (m *Mongo) Database(db string) *mongo.Database {
	return m.cli.Database(db)
}

func (m *Mongo) Collection(db, coll string) *Collection {
	return &Collection{
		Collection: m.cli.Database(db).Collection(coll),
	}
}
