package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lenniDespero/otus-golang/hw13/internal/calendar"
	amqpClient "github.com/lenniDespero/otus-golang/hw13/internal/pkg/ampq"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/logger"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/storage/sql"
	"github.com/spf13/pflag"
	"log"
	"strconv"
	"time"
)

type scheduler struct {
	calendar calendar.CalendarInterface
	ampq     *amqpClient.Ampq
}

func newScheduler(calendar calendar.CalendarInterface, ampq *amqpClient.Ampq) *scheduler {
	return &scheduler{calendar, ampq}
}

func (s *scheduler) Start(conf config.Scheduler) {
	period, err := strconv.Atoi(conf.Period)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Debug(fmt.Sprintf("Set ticker on minutes : %d", period))

	for ; true; <-time.Tick(time.Duration(period) * time.Minute) {
		logger.Debug("Get current events")
		events, err := s.calendar.GetEventsByStartPeriod(conf.BeforeTime, conf.EventTime)
		if err != nil {
			logger.Fatal(err.Error())
		}
		for _, event := range events {
			msg, err := json.Marshal(event)
			if err != nil {
				logger.Fatal(fmt.Sprintf("Failed to encode event: %s", err.Error()))
				continue
			}
			if err := s.ampq.Publish(msg); err != nil {
				logger.Fatal(fmt.Sprintf("Failed to publish event: %s", err.Error()))
				continue
			}
			logger.Debug(fmt.Sprintf("message %s published", msg))
		}
	}
}

func main() {
	var configPath = flag.String("config", "../config/application.yml", "path to configuration flag")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
	conf := config.GetConfigFromFile(*configPath)
	logger.Init(conf.Log.LogLevel, conf.Log.LogFile)
	storage, err := sql.New(&conf.DBConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	calendar := calendar.New(storage)
	amqpBus, err := amqpClient.NewAmpq(&conf.Ampq)
	if err != nil {
		logger.Fatal(err.Error())
	}
	scheduler := newScheduler(calendar, amqpBus)
	if err != nil {
		log.Fatalf("Scheduler init error: %s", err)
	}
	scheduler.Start(conf.Scheduler)
}
