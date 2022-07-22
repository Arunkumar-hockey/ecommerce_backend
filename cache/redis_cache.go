package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis-13756.c274.us-east-1-3.ec2.cloud.redislabs.com:13756",
		Password: "106N4MspuvRWpVBfCGOxV6NRUA6clhVG",
		DB:       0,
	})
}

func SetData(key string, value interface{}) {
	client := GetClient()

	result, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error::::", err)
	}

	err = client.HSet(ctx, "Products", key, result).Err()
	if err != nil {
		fmt.Println("Error::::", err)
	}
	fmt.Println("Key:", key)
	fmt.Println("Value:::", value)
}

func GetData(key string) (interface{}, error) {
	client := GetClient()

	value, err := client.HGet(ctx, "Products", key).Result()
	if err == redis.Nil {
		fmt.Printf("Key %s does not exist")
		return nil, err
	}
	fmt.Println("DATA::::", value)
	var data interface{}
	err = json.Unmarshal([]byte(value), &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

func GetAllData() (interface{}, error) {
	client := GetClient()
	result, err := client.HGetAll(ctx, "Products").Result()
	if err != nil {
		fmt.Println("false.....")
		return nil, err
	}

	//m := make(map[string]string)
	redisData := make([]string, 0)
	//var list []products.Product
	for _, value := range result {
		//var product []products.Product
		//productJson, _ := json.Marshal(value)
		//unmarshalErr := json.Unmarshal(productJson,product)
		//if unmarshalErr != nil {
		//	fmt.Println("ERROR:::", unmarshalErr)
		//	return nil, err
		//}
		redisData = append(redisData, value)
	}

	//unmarshalErr := json.Unmarshal([]byte(redisData), &m)
	//if unmarshalErr != nil {
	//	fmt.Println("ERROR:::", unmarshalErr)
	//	return nil, err
	//}

	//if marshallErr := bson.Unmarshal(raw, &product); marshallErr == nil {
	//	lists = append(lists, *product)
	//}

	fmt.Println("Success.....")
	return redisData, nil
}

func DeleteData(key string) (string, error) {
	client := GetClient()
	_, err := client.HDel(ctx, "Products", key).Result()
	if err == redis.Nil {

		msg := fmt.Sprintf("Key %s does not exist")
		return msg, err
	}
	msg := fmt.Sprintf("Key Deleted Succesfull")
	return msg, nil
}
