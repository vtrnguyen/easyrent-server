package workers

import (
	"easyrent-server/internal/services"
	email "easyrent-server/internal/shared/job"
	"log"
)

func StartEmailWorker(
	emailService *services.EmailService,
) {
	go func() {

		for job := range email.Queue {

			log.Println(
				"Sending email:",
				job.To,
			)

			err := emailService.Send(job)

			if err != nil {
				log.Println(
					"Email failed:",
					err,
				)

				continue
			}

			log.Println(
				"Email sent:",
				job.To,
			)
		}

	}()
}