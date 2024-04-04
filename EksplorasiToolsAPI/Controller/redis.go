package controller

import (
	"log"

	"github.com/go-redis/redis"

	m "EksplorasiToolsAPI/Model"
)

var client *redis.Client

func init() {
	// Inisialisasi klien Redis
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func (r *Reservation) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Reservation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, r)
}

func SaveReservation(ctx context.Context, client *redis.Client, key string, res *Reservation) error {
	data, err := res.MarshalBinary()
	if err != nil {
		return err
	}
	return client.Set(ctx, key, data, 0).Err()
}

func GetReservation(ctx context.Context, client *redis.Client, key string) (*Reservation, error) {
	data, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var res Reservation
	if err := res.UnmarshalBinary([]byte(data)); err != nil {
		return nil, err
	}
	return &res, nil
}

func saveReservation(ctx context.Context, client *redis.Client, key string, res *Model.Reservation) error {
	data, err := res.MarshalBinary()
	if err != nil {
		return err
	}
	return client.Set(ctx, key, data, 0).Err()
}

func getReservation(ctx context.Context, client *redis.Client, key string) (*Model.Reservation, error) {
	data, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var res Model.Reservation
	if err := res.UnmarshalBinary([]byte(data)); err != nil {
		return nil, err
	}
	return &res, nil
}

