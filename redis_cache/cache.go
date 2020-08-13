package main

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

func GetCacheValue(key interface{}) (interface{}, bool) {

	// start clients
	rdb := getclient()

	// Get
	val, err := rdb.Get(ctx, key.(string)).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
		return val, false
	}
	return nil, false
}

func DeleteCacheValue(key interface{}) (interface{}, bool) {

	rdb := getclient()
	val, err := rdb.Del(ctx, key.(string)).Result()
	if err != nil {
		panic(err)
	}

	return val, true
}

type CacheHash struct {
	// sync.Mutex
	// h       map[interface{}]*list.Element
	// l       *list.List
	// len     int
	// maxLen  int
	// ejectCB func(interface{}, interface{})
}

type keyValue struct {
	Key   interface{}
	Value interface{}
}

func (c *CacheHash) Init(maxLen int) {

}

// 4,294,967,295 Max amount of values
// 512Mb max for each entry
// Call this if over max values?
func (c *CacheHash) Eject() {

}

// check to see if cache full?
func (c *CacheHash) Add(k interface{}, v interface{}) bool {

	//REDIS CACHE SET
	return SetCacheValue(k, v)
}

// First does not exist in Dictionary
func (c *CacheHash) First() (interface{}, interface{}) {
	return "", ""
}

// Last does not exist in Dictionary
func (c *CacheHash) Last() (interface{}, interface{}) {
	return "", ""
}

// Get returns value and err if true
// On error case val = nil
func (c *CacheHash) Get(k interface{}) (interface{}, bool) {
	val, err := GetCacheValue(k)
	return val, err
}

//Mimiced function behavior
func (c *CacheHash) GetNoMove(k interface{}) (interface{}, bool) {
	// e, ok := c.h[k]
	// if ok {
	// 	return e.Value.(keyValue).Value, ok
	// }
	// return nil, ok

	if Exists(k) {
		val, myerror := GetCacheValue(k)
		return val, myerror
	}
	return nil, false
}

// Checks for existance
func (c *CacheHash) Has(k interface{}) bool {
	return Exists(k)
}

//Deletes key returns value and bool
func (c *CacheHash) Delete(k interface{}) (interface{}, bool) {
	return DeleteCacheValue(k)
}

// Obsolete
func (c *CacheHash) Len() int {
	return 0
}

// Obsolete
func (c *CacheHash) RegisterCB(newCB func(interface{}, interface{})) {

}

func ExampleClient(key string, value string) {
	SetCacheValue(key, value)
	GetCacheValue(key)
	fmt.Println(Exists(key))
	fmt.Println(Exists("Hello world"))
	DeleteCacheValue(key)
	GetCacheValue(key)
}

func main() {
	ExampleClient("Hello", "World")
}
