package websocket

import (
	"fmt"
	"log"
	"net/http"

	"../redisKeys"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

//upgrading simple http connection to websocket connection
var upgrader websocket.Upgrader

// WsEndpoint function defines the page with websocket connection
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// "ws" is the upgaraded websocket handle
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")

	//Initialising redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// value of 'n' defines the number of times keys will be fetched from redis and sent to the client
	n := 1
	for i := 1; i <= n; i++ {
		mappedKeys := redisKeys.ProcessKeys(rdb)
		/*
			Iterating over "mappedKeys" which is "map[string]interface{}" type &
			converting each value to "string" type by "fmt.Sprint" function as
			[]interface type cannot be directly converted to []byte and
			[]byte is the only data type which is accepted by "WriteMessage" method
			to print something on client web console.
		*/
		for _, valueOfKey := range mappedKeys {
			stringValueOfKey := fmt.Sprint(valueOfKey)
			err = ws.WriteMessage(1, []byte(stringValueOfKey))
			if err != nil {
				log.Println(err)
			}
		}
	}
}
