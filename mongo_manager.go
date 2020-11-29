package bzmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MapStringOptions map[string]*Options
	mapStringMongo   map[string]*Mongo
)

var mm = mapStringMongo{}

func Init(mo MapStringOptions) {
	for key, opt := range mo {
		cli, err := New(opt)
		if err != nil {
			continue
		}

		mm[key] = cli
	}
}

func InitAndConnect(ctx context.Context, mo MapStringOptions) {
	for key, opt := range mo {
		cli, err := Connect(ctx, opt)
		if err != nil {
			continue
		}

		mm[key] = cli
	}
}

func GetMongo(name string) *Mongo {
	m, ok := mm[name]
	if !ok {
		panic("No mongo named: " + name)
	}

	return m
}

func GetDatabase(name, db string) *mongo.Database {
	return GetMongo(name).Database(db)
}

func GetCollection(name, db, coll string) *Collection {
	return GetMongo(name).Collection(db, coll)
}
