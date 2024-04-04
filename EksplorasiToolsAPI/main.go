package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/go-gomail/gomail"
	"github.com/go-redis/redis"
	// "github.com/gorilla/mux"
	"github.com/robfig/cron"

	m "EksplorasiToolsAPI/Model"
	// c "EksplorasiToolsAPI/controller"
)

func main() {
	// Inisialisasi klien Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	http.HandleFunc("/reservation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var res m.Reservation
			err := json.NewDecoder(r.Body).Decode(&res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Simpan reservasi ke cache
			err = client.Set("latest_reservation", res, 0).Err()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Inisialisasi GoCRON
			c := cron.New()
			c.AddFunc("@every 1h", func() {
				// Kirim email konfirmasi dengan GoMail
				m := gomail.NewMessage()
				m.SetHeader("From", "restaurant@example.com")
				m.SetHeader("To", res.Email)
				m.SetHeader("Subject", "Reservation Confirmation")
				m.SetBody("text/html", "Thank you for your reservation, "+res.Name+"! Your table will be ready at "+res.Time+".")

				d := gomail.NewDialer("smtp.example.com", 587, "user", "password")

				// Kirim email dalam goroutine
				go func() {
					if err := d.DialAndSend(m); err != nil {
						log.Println("Failed to send email:", err)
					} else {
						log.Println("Email sent successfully")
					}
				}()
			})
			c.Start()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Reservation created"))
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	// router := mux.NewRouter()

	// router.HandleFunc("/v2/transactions/user/{id}", c.DeleteSingleProductGorm).Methods("DELETE")
	// router.HandleFunc("/v1/transactions", c.InsertNewProductsandTrans).Methods("POST")

	// http.Handle("/", router)
	// fmt.Println("Connected to port 8888")
	// log.Println("Connected to port 8888")
	// log.Fatal(http.ListenAndServe(":8888", router))
	http.ListenAndServe(":8080", nil)
}
