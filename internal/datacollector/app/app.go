package app

import (
	"github.com/ethcero/connected-pv/internal/datacollector"
	"github.com/ethcero/connected-pv/internal/datacollector/collector"
	"github.com/ethcero/connected-pv/internal/datacollector/publisher"
	"github.com/ethcero/connected-pv/pkg/scheduler"
	"log"
)

type App struct {
	config    datacollector.Config
	scheduler *scheduler.Scheduler
}

func NewApp() *App {
	return &App{
		config:    datacollector.NewConfig(),
		scheduler: scheduler.NewScheduler(5),
	}
}

func (app *App) Start() {
	log.Println("Initializing data collector")
	bus := make(chan datacollector.BusMessage)
	app.initPublisher(bus)
	app.startCollector(bus)
}

func (app *App) Stop() {
	log.Println("Stopping data collector")
	app.scheduler.Stop()
}

func (app *App) startCollector(bus chan datacollector.BusMessage) {
	// Start the scheduler
	c := collector.NewCollector(app.config.CollectorConfig)
	app.scheduler.Start(func() {
		collector.CollectAndDispatch(c, bus)
	})
}

func (app *App) initPublisher(bus chan datacollector.BusMessage) {
	p := publisher.NewPublisher(app.config.Context, app.config.PublisherConfig, app.config.IotConfig)
	p.Start()
	publisher.HandlePublish(p, bus)
}
