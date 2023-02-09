package crons

import "github.com/robfig/cron"

func RunCrons() {
	cron := cron.New()
	cron.AddFunc("* * * * *", func() {
		ExpireShortenUrlCronJob()
	})
	cron.Start()
	select {}
}
