package redisKeys

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

//redis client context
var ctx = context.Background()

//defining the interface so that any type of values can be read in golang environment from redis
var mappedKeys = make(map[string]interface{})

// GoRoutine function which will process each passed key concurrently
func infoGoroutine(key string, rdb *redis.Client, channel chan bool) {
	result, err := rdb.Do(ctx, "INFO", key).Result()
	if err != nil {
		fmt.Println("Error in reading from database.")
	}
	mappedKeys[key] = result
	channel <- true
}

//ProcessKeys function will process redis keys(string slice) for INFO command and return the correponding values in an interface
func ProcessKeys(rdb *redis.Client) map[string]interface{} {

	//processing redis keys
	redisKeys := []string{"CPU", "CLIENTS", "CLUSTER", "SERVER", "MEMORY"}
	// unbuffered boolean value channel which will control GoRoutine
	channel := make(chan bool)

	for _, key := range redisKeys {
		//GoRoutine to process each key concurrently
		go infoGoroutine(key, rdb, channel)
		<-channel
	}
	fmt.Println("Values of the keys recieved succesfully.")
	return mappedKeys
}
