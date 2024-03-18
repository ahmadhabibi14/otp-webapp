package main

import (
	"log"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("failed to connect redis:", err)
		return
	} else {
		log.Println("connected to redis")
	}

	// Set key value pair
	err = client.Set("name", "Habibi", 0).Err()
	if err != nil {
		log.Println("failed to set new value:", err)
	}
	
	// Retrieve a value for a given key
	name, err := client.Get("name").Result()
	if err != nil {
		log.Println("key is empty:", err)
	}

	// It must be empty
	age, err := client.Get("age").Result()
	if err != nil {
		log.Println("key is empty:", err)
	}

	log.Println("name \t:", name)
	log.Println("age \t:", age)
}