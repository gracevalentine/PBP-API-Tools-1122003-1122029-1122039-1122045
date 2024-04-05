package controller

import (
	"encoding/json"
	"net/http"

	m "EksplorasiToolsAPI/Model"
)

func HandleReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var res m.Reservation
		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Simpan reservasi ke cache
		SaveReservation(res)

		// Jadwalkan pengiriman email
		scheduleJob(res)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Reservation created"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
