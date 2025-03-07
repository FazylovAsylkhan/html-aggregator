package scheduler

import (
	"context"

	"github.com/FazylovAsylkhan/html-aggregator/internal/database"
	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
	baspana "github.com/FazylovAsylkhan/html-aggregator/internal/usecase/baspana_market"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Scheduler struct{
	cron *cron.Cron
	log *logrus.Logger
	baspana *baspana.Baspana
}

func Init(db *database.Queries) *Scheduler {
	scheduler := Scheduler{
		cron: cron.New(),
		log: logger.New(),
		baspana: baspana.Init(db),
	}
	scheduler.log.SetFormatter(&logger.GeneralFormatter{})
	

	return &scheduler
}

func (s Scheduler) Start() {
	_, err := s.cron.AddFunc("*/5 * * * *", s.sendRequest)
	if err != nil {
		s.log.Infof("scheduler failed: %v", err)
	}

	s.cron.Start()
	select {}
}

func (s Scheduler) sendRequest() {
	s.log.Info("cron: start sending request")
	_, err := s.baspana.LoadPosts(context.Background())
	if err != nil {
		s.log.Infof("request failed: %v", err)
	}
	s.log.Info("cron: finished sending request successfully")
}