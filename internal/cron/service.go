package cron

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// cronRunner interface for cron
type cronRunner interface {
	AddFunc(spec string, cmd func()) (cron.EntryID, error)
	Start()
	Stop() context.Context
}

// cronService cron service struct
type cronService struct {
	cron cronRunner
}

// NewCronService returns new cron service
func NewCronService() *cronService {
	return &cronService{
		cron: cron.New(),
	}
}

const (
	timePatterns = "* * * * *"
)

// Run starts service
func (ds *cronService) Run(ctx context.Context) {
	if _, err := ds.cron.AddFunc(timePatterns, ds.task); err != nil {
		log.Fatal(err)
	}
	ds.cron.Start()
	<-ctx.Done()
	cronTask := ds.cron.Stop()
	<-cronTask.Done()
}

// periodic task
func (ds *cronService) task() {
	fmt.Printf("Cron task, %v\n", time.Now())
}
