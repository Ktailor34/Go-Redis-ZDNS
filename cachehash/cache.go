package cachehash

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func getclient() *redis.Client {
	//Start Client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func Exists(key interface{}) bool {

	//start clients
	rdb := getclient()

	//Exists
	val := rdb.Exists(ctx, key.(string))
	result, err := val.Result()
	if err != nil {
		panic(err)
	}
	if result != 0 {
		return true
	}
	return false
}

func SetCacheValue(key interface{}, value interface{}) bool {

	//start clients
	rdb := getclient()

	//Set
	err := rdb.Set(ctx, key.(string), value.(string), 0).Err()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func GetCacheValue(key interface{}) interface{} {

	//start clients
	rdb := getclient()

	//Get
	val, err := rdb.Get(ctx, key.string).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
		return val
	}
	return val
}

// func ExampleClient(key string, value string) {
// 	SetCacheValue(key, value)
// 	GetCacheValue(key)
// }
