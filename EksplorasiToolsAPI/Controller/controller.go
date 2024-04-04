package controller

import (
	"encoding/json"
	"net/http"

	m "EksplorasiToolsAPI/Model"
)

func HandleReservation(ctx context.Context, client *redis.Client, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var res Reservation
		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Simpan reservasi ke cache
		err = saveReservation(ctx, client, "latest_reservation", &res)
		if err != nil {
			http.Error(w, "Failed to save reservation", http.StatusInternalServerError)
			return
		}

		// Jadwalkan pengiriman email
		// Anda perlu memperbarui fungsi scheduleJob Anda untuk menerima argumen yang sama
		err = scheduleJob(ctx, client, &res)
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

