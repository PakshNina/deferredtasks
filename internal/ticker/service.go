package ticker

import (
	"context"
	"fmt"
	"time"
)

const (
	period = time.Minute
)

// service struct
type tickerService struct {
}

// NewTickerService returns new service
func NewTickerService() *tickerService {
	return &tickerService{}
}

// Run starts service
func (ds *tickerService) Run(ctx context.Context) {
	ticker := time.NewTicker(period)
	for {
		select {
		case <-ticker.C:
			go func() {
				ds.task()
			}()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// periodic task
func (ds *tickerService) task() {
	fmt.Printf("Ticker task, %v\n", time.Now())
}
