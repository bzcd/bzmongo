package bzmongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type MapStringOptions map[string]*Options

var mgr = map[string]*Mongo{}

func Init(mo MapStringOptions) {
	for key, opt := range mo {
		cli, err := New(opt)
		if err != nil {
			continue
		}

		mgr[key] = cli
	}
}

func InitAndConnect(ctx context.Context, mo MapStringOptions) {
	for key, opt := range mo {
		cli, err := Connect(ctx, opt)
		if err != nil {
			continue
		}

		mgr[key] = cli
	}
}

func GetMongo(name string) (*Mongo, error) {
	m, ok := mgr[name]
	if !ok {
		return nil, errors.New("No mongo named: " + name)
	}

	return m, nil
}

func GetDatabase(name, db string) *mongo.Database {
	m, _ := GetMongo(name)
	if m == nil {
		return nil
	}

	return m.Database(db)
}

func GetCollection(name, db, coll string) *Collection {
	m, _ := GetMongo(name)
	if m == nil {
		return nil
	}

	return m.Collection(db, coll)
}
