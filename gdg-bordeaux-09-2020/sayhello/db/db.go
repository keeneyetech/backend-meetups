package db

import "github.com/go-redis/redis"

type DB struct {
	*redis.Client
}

func New(rdb *redis.Client) *DB {
	return &DB{Client: rdb}
}

func (db *DB) Hit(IP string) error {
	return db.Incr("hits:" + IP).Err()
}
