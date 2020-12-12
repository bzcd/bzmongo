package bzmongo

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MapStringOptions map[string]*Options

type tMgr struct {
	once sync.Once
	m    *Mongo
}

var mgr = map[string]*tMgr{}

func Init(mo MapStringOptions) (err error) {
	for key, opt := range mo {
		cli, e := New(opt)
		if e != nil {
			err = e
			continue
		}

		mgr[key] = &tMgr{m: cli}
	}

	return
}

func InitAndConnect(ctx context.Context, mo MapStringOptions) error {
	for key, opt := range mo {
		cli, err := Connect(ctx, opt)
		if err != nil {
			return err
		}

		v := &tMgr{m: cli}
		v.once.Do(func() {})

		mgr[key] = v
	}

	return nil
}

func GetMongo(name string) (*Mongo, error) {
	m, ok := mgr[name]
	if !ok {
		return nil, errors.New("No mongo named: " + name)
	}

	var err error

	m.once.Do(func() {
		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()
		err = m.m.Connect(ctx)
	})

	return m.m, err
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
