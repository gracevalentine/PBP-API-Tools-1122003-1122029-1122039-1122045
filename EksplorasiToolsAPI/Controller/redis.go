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

func saveReservation(reservation m.Reservation) {
	// Simpan reservasi ke cache
	err := client.Set("latest_reservation", reservation, 0).Err()
	if err != nil {
		log.Println("Failed to save reservation:", err)
	} else {
		log.Println("Reservation saved successfully")
	}
}
