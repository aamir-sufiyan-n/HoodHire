package utils

import (
	"fmt"
	"hoodhire/internal/repositories"
	"hoodhire/structures/models"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func StartSubscriptionCron(db *gorm.DB, subRepo *repositories.SubscriptionRepo) {
	c := cron.New(cron.WithLocation(time.Local))

	c.AddFunc("0 0 * * *", func() {
		fmt.Println("running subscription cron job:", time.Now())

		// EXPIRING SOON

		expiringSoon, err := subRepo.GetExpiringSubscriptions(7)
		if err != nil {
			fmt.Println("cron error (expiring):", err)
			return
		} else {
			for _, sub := range expiringSoon {

				if sub.Hirer.User.Email != "" {
					go SendSubscriptionExpiryReminderMail(
						sub.Hirer.User.Email,
						sub.Hirer.FullName,
						sub.Plan.Name,
						sub.EndDate,
					)
				}
			}
		}

		// EXPIRED

		expired, err := subRepo.GetExpiredSubscriptions()
		if err != nil {
			fmt.Println("cron error (expiring):", err)
			return
		} else {
			for _, sub := range expired {

				sub.Status = "expired"
				subRepo.DB.Model(&models.Hirer{}).
					Where("id = ?", sub.HirerID).
					Update("is_pro", false)
				if err := subRepo.UpdateSubscription(&sub); err != nil {
					continue
				}

				if sub.Hirer.User.Email != "" {
					go SendSubscriptionExpiredMail(
						sub.Hirer.User.Email,
						sub.Hirer.FullName,
						sub.Plan.Name,
					)
				}
			}
		}

	})

	c.Start()
	fmt.Println("subscription cron job started")
}
