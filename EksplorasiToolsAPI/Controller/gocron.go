package controller

import (
	"github.com/robfig/cron"

	m "EksplorasiToolsAPI/Model"
)

func scheduleJob(reservation m.Reservation) {
	c := cron.New()
	c.AddFunc("@every 15s", func() {
		sendEmail(reservation)
	})
	c.Start()
}
