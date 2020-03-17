package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/logger"
	"github.com/spf13/pflag"

	eventproto "github.com/lenniDespero/otus-golang/hw13/grpc/event"

	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/storage"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calendarpb struct {
	Events storage.Storage
}

func (c calendarpb) Edit(ctx context.Context, e *eventproto.EventEditRequest) (*eventproto.EventEditResponse, error) {
	log.Print(fmt.Sprintf("Got request Edit %v", e.String()))
	startDate, err := ptypes.Timestamp(e.Event.DateStarted)
	if err != nil {
		return nil, err
	}
	endDate, err := ptypes.Timestamp(e.Event.DateComplete)
	if err != nil {
		return nil, err
	}

	err = c.Events.Edit(e.Id, models.Event{ID: e.Event.Id, Title: e.Event.Title, DateStarted: startDate, DateComplete: endDate, Notice: e.Event.Notice})
	if err != nil {
		return nil, err
	}
	return &eventproto.EventEditResponse{}, nil
}

func (c calendarpb) Delete(ctx context.Context, e *eventproto.EventDeleteRequest) (*eventproto.EventDeleteResponse, error) {
	logger.Info(fmt.Sprintf("Got request Delete %v", e.String()))
	err := c.Events.Delete(e.Id)
	if err != nil {
		return nil, err
	}
	return &eventproto.EventDeleteResponse{}, nil
}

func (c calendarpb) Get(ctx context.Context, e *eventproto.EventGetByIdRequest) (*eventproto.EventGetByIdResponse, error) {
	logger.Info(fmt.Sprintf("Got request Get %v", e.String()))
	ev, err := c.Events.GetEventByID(e.Id)
	if err != nil {
		return nil, err
	}
	respEvents := make([]*eventproto.Event, 0, len(ev))
	for _, row := range ev {
		respEvents = append(respEvents, convertToProtoEvent(&row))
	}
	return &eventproto.EventGetByIdResponse{Events: respEvents}, nil
}

func (c calendarpb) GetAll(ctx context.Context, e *eventproto.EventGetAllRequest) (*eventproto.EventGetAllResponse, error) {
	logger.Info(fmt.Sprintf("Got request GetAll %v", e.String()))
	ev, err := c.Events.GetEvents()
	if err != nil {
		return nil, err
	}
	respEvents := make([]*eventproto.Event, 0, len(ev))
	for _, row := range ev {
		respEvents = append(respEvents, convertToProtoEvent(&row))
	}
	return &eventproto.EventGetAllResponse{Events: respEvents}, nil
}

func (c calendarpb) Add(ctx context.Context, e *eventproto.EventAddRequest) (*eventproto.EventAddResponse, error) {
	startDate, err := ptypes.Timestamp(e.DateStarted)
	if err != nil {
		return nil, err
	}
	endDate, err := ptypes.Timestamp(e.DateComplete)
	if err != nil {
		return nil, err
	}
	logger.Info(fmt.Sprintf("Got request Add %v", models.Event{Title: e.Title, DateStarted: startDate, DateComplete: endDate}))
	id, err := c.Events.Add(models.Event{Title: e.Title, DateStarted: startDate, DateComplete: endDate, Notice: e.Notice})
	return &eventproto.EventAddResponse{Id: id}, err
}

func convertToProtoEvent(event *models.Event) *eventproto.Event {
	dateStart, err := ptypes.TimestampProto(event.DateStarted)
	if err != nil {
		logger.Fatal("Cant't convert %v to timestamp proto", event.DateStarted)
	}
	dateComplete, err := ptypes.TimestampProto(event.DateComplete)
	if err != nil {
		logger.Fatal("Cant't convert %v to timestamp proto", event.DateStarted)
	}
	return &eventproto.Event{
		Id:           event.ID,
		Title:        event.Title,
		DateStarted:  dateStart,
		DateComplete: dateComplete,
		Notice:       event.Notice,
	}
}

func StartGrpcServer(storage storage.Storage, address string, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%s", address, port))
	if err != nil {
		logger.Fatal("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	eventproto.RegisterEventServiceServer(grpcServer, &calendarpb{Events: storage})
	grpcServer.Serve(lis)
}

func main() {
	var configPath = flag.String("config", "../../config/application.yml", "path to configuration flag")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
	conf := config.GetConfigFromFile(*configPath)
	logger.Init(conf.Log.LogLevel, conf.Log.LogFile)
	inMemoryStorage := storage.New()
	logger.Info("GRPC server start")
	StartGrpcServer(*inMemoryStorage, conf.GrpcServer.Host, conf.GrpcServer.Port)
}
