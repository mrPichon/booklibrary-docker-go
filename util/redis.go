package util

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitializeRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default db
	})

	InitMockCollection()
	SetBookList(BookCollection)

}

func SetBookList(booklist *Collection) {
	err := rdb.Set(ctx, "booklist", booklist, 0).Err()
	if err != nil {
		log.Println("Couldn't save the booklist, Err: ", err)
	}
}

func GetBookList() *Collection {
	val, err := rdb.Get(ctx, "booklist").Result()
	log.Println("redis get booklist: ", val)
	if err != nil {
		log.Println("Cann't get the collection of books, Err: ", err)
	}

	var books Collection
	err = json.Unmarshal([]byte(val), &books)
	if err != nil {
		panic(err)
	}
	return &books
}
