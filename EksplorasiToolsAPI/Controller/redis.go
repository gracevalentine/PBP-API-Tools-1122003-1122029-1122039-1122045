package controller

import (
	"encoding/json"
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

// func SaveReservation(ctx context.Context, client *redis.Client, key string, res *m.Reservation) error {
// 	return client.Set(ctx, key, res, 0).Err()
// }

func SaveReservation(reservation m.Reservation) {
	// Simpan reservasi ke cache
	converted, err := json.Marshal(reservation)

	err = client.Set("latest_reservation", converted, 0).Err()
	if err != nil {
		log.Println("Failed to save reservation:", err)
	} else {
		log.Println("Reservation saved successfully")
	}
}
