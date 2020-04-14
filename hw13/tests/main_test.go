package main

import (
	"github.com/cucumber/godog"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/storage/sql"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		godog.SuiteContext(s)
		FeatureContextHttpEvent(s)
		FeatureContextListEvents(s)
		FeatureContextNotice(s)
	}, godog.Options{
		Format:    "pretty",
		Paths:     []string{"feature"},
		Randomize: 0,
	})

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func ClearDatabase() {
	configPath := "../config/application.yml"
	conf := config.GetConfigFromFile(configPath)
	storage, err := sql.New(&conf.DBConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = storage.ClearDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	storage.ConnPool.Close()
}
