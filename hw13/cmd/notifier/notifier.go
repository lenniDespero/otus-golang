package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg"
	amqpClient "github.com/lenniDespero/otus-golang/hw13/internal/pkg/ampq"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/logger"
	"github.com/spf13/pflag"
	"github.com/streadway/amqp"
	"log"
)

type notifier struct {
	ampq *amqpClient.Ampq
}

func newNotifier(ampq *amqpClient.Ampq) *notifier {
	return &notifier{ampq}
}

func (n *notifier) Start() error {
	err := n.ampq.Subscribe("notifier", func(delivery amqp.Delivery) {
		event := &pkg.Event{}
		if err := json.Unmarshal(delivery.Body, event); err != nil {
			logger.Error(fmt.Sprintf("Failed to parse message: %s", err.Error()))
			return
		}
		n.notify(event)
	})
	if err != nil {
		return err
	}

	return nil
}

func (n *notifier) notify(event *pkg.Event) {
	fmt.Printf("Get event: %v", event)
}

func main() {
	var configPath = flag.String("config", "../config/application.yml", "path to configuration flag")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
	conf := config.GetConfigFromFile(*configPath)
	logger.Init(conf.Log.LogLevel, conf.Log.LogFile)
	amqpBus, err := amqpClient.NewAmpq(&conf.Ampq)
	if err != nil {
		logger.Fatal(err.Error())
	}
	notifierAgent := newNotifier(amqpBus)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Scheduler init error: %s", err.Error()))
	}
	notifierAgent.Start()
}
