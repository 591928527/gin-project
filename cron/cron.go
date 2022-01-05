package cron

import (
	"blogs/models"
	"log"
	"time"

	"github.com/robfig/cron"
)

func Crontab() {
	log.Println("Starting...")
	c := cron.New()
	c.AddFunc("*/10 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("*/10 * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
