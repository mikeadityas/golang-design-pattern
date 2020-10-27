package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type RenewTokenEvent struct {
	token string
}

type RenewTokenObserver interface {
	NotifyNewToken(RenewTokenEvent)
}

type RenewTokenSubject interface {
	AddListener(RenewTokenObserver)
	RemoveListener(RenewTokenObserver)
	Notify(RenewTokenEvent)
}

type consumerService struct {
	id              int
	token           string
	lastRetrievedAt time.Time
}

func (c *consumerService) NotifyNewToken(event RenewTokenEvent) {
	log.Printf(
		"Observer id: %d Received new token: %s after %v\n",
		c.id,
		event.token,
		time.Since(c.lastRetrievedAt),
	)
	c.token = event.token
	c.lastRetrievedAt = time.Now()
}

type secretTokenSubject struct {
	observers *sync.Map
}

func (s *secretTokenSubject) AddListener(observer RenewTokenObserver) {
	s.observers.Store(observer, struct{}{})
}

func (s *secretTokenSubject) RemoveListener(observer RenewTokenObserver) {
	s.observers.Delete(observer)
}

func (s *secretTokenSubject) Notify(event RenewTokenEvent) {
	s.observers.Range(func(key, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		key.(RenewTokenObserver).NotifyNewToken(event)
		return true
	})
}

func renewTokenWorker(ctx context.Context, periodSecond int, newTokenCh chan string) {
	ticker := time.NewTicker(time.Duration(periodSecond) * time.Second)

	newTokenCh <- renewToken()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Stopping worker...")
				ticker.Stop()
				log.Println("Worker stopped...")
				return
			case <-ticker.C:
				newTokenCh <- renewToken()
			}
		}
	}()
}

func renewToken() string {
	return fmt.Sprintf("newToken-%d", time.Now().UTC().Unix())
}

func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newTokenCh := make(chan string)

	tokenRenewPeriodSecond := 1

	go renewTokenWorker(ctx, tokenRenewPeriodSecond, newTokenCh)

	secretTokenSubj := &secretTokenSubject{
		observers: &sync.Map{},
	}

	start := time.Now()

	consumerService1 := &consumerService{id: 1, lastRetrievedAt: start}
	consumerService2 := &consumerService{id: 2, lastRetrievedAt: start}

	secretTokenSubj.AddListener(consumerService1)
	secretTokenSubj.AddListener(consumerService2)

	go func() {
		for token := range newTokenCh {
			secretTokenSubj.Notify(RenewTokenEvent{token: token})
		}
	}()

	go func() {
		// Test removing consumerService1 after 2 seconds
		select {
		case <-time.After(2 * time.Second):
			log.Println("Removing consumerService1 from observers")
			secretTokenSubj.RemoveListener(consumerService1)
			log.Println("consumerService1 removed")
		}
	}()

	waitForSignal()

	cancel()

	gracefulShutdownCtx, gracefulShutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer gracefulShutdownCancel()

	select {
	case <-gracefulShutdownCtx.Done():
		log.Println("5 seconds grace period has passed")
	}

	log.Println("Clean up finished")
}
