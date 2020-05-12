package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/storage/sql"
	"net"
	"os/user"
	"strconv"
	"time"

	pkg "github.com/lenniDespero/otus-golang/hw13/internal/pkg"

	"github.com/lenniDespero/otus-golang/hw13/internal/calendar"

	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/logger"
	"github.com/spf13/pflag"

	"github.com/golang/protobuf/ptypes"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calendarpb struct {
	сalendar calendar.Calendar
}

func (c calendarpb) Edit(ctx context.Context, e *pkg.EventEditRequest) (*pkg.EventEditResponse, error) {
	logger.Info(fmt.Sprintf("Got request Edit %v", e.String()))
	startDate, err := ptypes.Timestamp(e.Event.DateStarted)
	if err != nil {
		return nil, err
	}
	endDate, err := ptypes.Timestamp(e.Event.DateComplete)
	if err != nil {
		return nil, err
	}
	userId, err := getUserId()
	if err != nil {
		return nil, err
	}
	err = c.сalendar.Edit(e.Id, models.Event{ID: e.Event.Id, Title: e.Event.Title, DateStarted: startDate, DateComplete: endDate, Notice: e.Event.Notice}, userId)
	if err != nil {
		return nil, err
	}
	return &pkg.EventEditResponse{}, nil
}

func (c calendarpb) Delete(ctx context.Context, e *pkg.EventDeleteRequest) (*pkg.EventDeleteResponse, error) {
	logger.Info(fmt.Sprintf("Got request Delete %v", e.String()))
	err := c.сalendar.Delete(e.Id)
	if err != nil {
		return nil, err
	}
	return &pkg.EventDeleteResponse{}, nil
}

func (c calendarpb) Get(ctx context.Context, e *pkg.EventGetByIdRequest) (*pkg.EventGetByIdResponse, error) {
	logger.Info(fmt.Sprintf("Got request Get %v", e.String()))
	ev, err := c.сalendar.GetEventByID(e.Id)
	if err != nil {
		return nil, err
	}
	respEvents := make([]*pkg.Event, 0, len(ev))
	for _, row := range ev {
		respEvents = append(respEvents, convertToProtoEvent(&row))
	}
	return &pkg.EventGetByIdResponse{Events: respEvents}, nil
}

func (c calendarpb) GetAll(ctx context.Context, e *pkg.EventGetAllRequest) (*pkg.EventGetAllResponse, error) {
	logger.Info(fmt.Sprintf("Got request GetAll %v", e.String()))
	ev, err := c.сalendar.GetEvents()
	if err != nil {
		return nil, err
	}
	respEvents := make([]*pkg.Event, 0, len(ev))
	for _, row := range ev {
		respEvents = append(respEvents, convertToProtoEvent(&row))
	}
	return &pkg.EventGetAllResponse{Events: respEvents}, nil
}

func (c calendarpb) Add(ctx context.Context, e *pkg.EventAddRequest) (*pkg.EventAddResponse, error) {
	startDate, err := ptypes.Timestamp(e.DateStarted)
	if err != nil {
		return nil, err
	}
	endDate, err := ptypes.Timestamp(e.DateComplete)
	if err != nil {
		return nil, err
	}
	userId, err := getUserId()
	if err != nil {
		return nil, err
	}
	logger.Info(fmt.Sprintf("Got request Add %v", models.Event{Title: e.Title, DateStarted: startDate, DateComplete: endDate}))
	id, err := c.сalendar.Add(e.Title, startDate.Local(), endDate.Local(), e.Notice, userId)
	return &pkg.EventAddResponse{Id: id}, err
}

func (c calendarpb) GetDay(ctx context.Context, e *pkg.EventsGetDayRequest) (*pkg.EventsGetDayResponse, error) {
	logger.Info(fmt.Sprintf("Got request GetDay events %v", e.String()))
	now := time.Now()
	endDay := time.Until(now.Add(time.Duration(24) * time.Hour))
	timeLength := strconv.FormatInt(int64(endDay.Round(time.Minute).Minutes()), 10)
	ev, err := c.сalendar.GetEventsByStartPeriod("0", timeLength)
	if err != nil {
		return nil, err
	}
	respEvents := make([]*pkg.Event, 0, len(ev))
	for _, row := range ev {
		respEvents = append(respEvents, convertToProtoEvent(&row))
	}
	return &pkg.EventsGetDayResponse{Events: respEvents}, nil
}

func (c calendarpb) GetWeek(ctx context.Context, e *pkg.EventsGetWeekRequest) (*pkg.EventsGetWeekResponse, error) {
	logger.Info(fmt.Sprintf("Got request GetWeek events %v", e.String()))
	now := time.Now()
	endDay := time.Until(now.Add(time.Duration(24) * time.Hour * 7))
	timeLength := strconv.FormatInt(int64(endDay.Round(time.Minute).Minutes()), 10)
	ev, err := c.сalendar.GetEventsByStartPeriod("0", timeLength)
	if err != nil {
		return nil, err
	}
	respEvents := make([]*pkg.Event, 0, len(ev))
	for _, row := range ev {
		respEvents = append(respEvents, convertToProtoEvent(&row))
	}
	return &pkg.EventsGetWeekResponse{Events: respEvents}, nil
}

func (c calendarpb) GetMonth(ctx context.Context, e *pkg.EventsGetMonthRequest) (*pkg.EventsGetMonthResponse, error) {
	logger.Info(fmt.Sprintf("Got request GetMonth events %v", e.String()))
	now := time.Now()
	endDay := time.Until(now.Add(time.Duration(24) * time.Hour * 30))
	timeLength := strconv.FormatInt(int64(endDay.Round(time.Minute).Minutes()), 10)
	ev, err := c.сalendar.GetEventsByStartPeriod("0", timeLength)
	if err != nil {
		return nil, err
	}
	respEvents := make([]*pkg.Event, 0, len(ev))
	for _, row := range ev {
		respEvents = append(respEvents, convertToProtoEvent(&row))
	}
	return &pkg.EventsGetMonthResponse{Events: respEvents}, nil
}

func convertToProtoEvent(event *models.Event) *pkg.Event {
	dateStart, err := ptypes.TimestampProto(event.DateStarted)
	if err != nil {
		logger.Fatal("Cant't convert %v to timestamp proto", event.DateStarted)
	}
	dateComplete, err := ptypes.TimestampProto(event.DateComplete)
	if err != nil {
		logger.Fatal("Cant't convert %v to timestamp proto", event.DateStarted)
	}
	return &pkg.Event{
		Id:           event.ID,
		Title:        event.Title,
		DateStarted:  dateStart,
		DateComplete: dateComplete,
		Notice:       event.Notice,
	}
}

func StartGrpcServer(calendar calendar.Calendar, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		logger.Fatal("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pkg.RegisterEventServiceServer(grpcServer, &calendarpb{сalendar: calendar})
	grpcServer.Serve(lis)
}

func getUserId() (int64, error) {
	user, err := user.Current()
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(user.Uid, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func main() {
	var configPath = flag.String("config", "../../config/application.yml", "path to configuration flag")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
	conf := config.GetConfigFromFile(*configPath)
	logger.Init(conf.Log.LogLevel, conf.Log.LogFile)
	storage, err := sql.New(&conf.DBConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	calendar := calendar.New(storage)
	logger.Info("GRPC server start")
	StartGrpcServer(*calendar, conf.GrpcServer.Port)
}
