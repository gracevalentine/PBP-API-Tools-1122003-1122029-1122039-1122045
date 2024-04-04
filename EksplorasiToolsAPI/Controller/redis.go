package controller

import (
	"github.com/go-redis/redis"

	"EksplorasiToolsAPI/Model"
	m "EksplorasiToolsAPI/Model"
)

var client *redis.Client

func init() {
	// inisialisasi klien Redis
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
	})
}

func SaveReservation(client *redis.Client, key string, res *m.Reservation) error {
	data, err := res.MarshalBinary()
	if err != nil {
		return err
	}
	return client.Set(key, data, 0).Err()
}

func GetReservation(client *redis.Client, key string) (*m.Reservation, error) {
	data, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	var res m.Reservation
	if err := res.UnmarshalBinary([]byte(data)); err != nil {
		return nil, err
	}
	return &res, nil
}

func saveReservation(client *redis.Client, key string, res *m.Reservation) error {
	data, err := res.MarshalBinary()
	if err != nil {
		return err
	}
	return client.Set(key, data, 0).Err()
}

func getReservation(client *redis.Client, key string) (*m.Reservation, error) {
	data, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	var res Model.Reservation
	if err := res.UnmarshalBinary([]byte(data)); err != nil {
		return nil, err
	}
	return &res, nil
}
