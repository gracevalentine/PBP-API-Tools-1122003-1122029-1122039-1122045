package controller

import (
	"context"
	"encoding/json"
	"net/http"

	m "EksplorasiToolsAPI/Model"

	"github.com/go-redis/redis"
)

func HandleReservation(ctx context.Context, client *redis.Client, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var res m.Reservation
		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// simpan reservasi ke cache
		err = saveReservation(client, "latest_reservation", &res)
		if err != nil {
			http.Error(w, "Failed to save reservation", http.StatusInternalServerError)
			return
		}

		err = scheduleJob(res)
		if err != nil {
			http.Error(w, "Failed to schedule email", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Reservation created"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
