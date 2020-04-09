package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/storage/sql"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/types"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lenniDespero/otus-golang/hw13/internal/calendar"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"

	"github.com/gorilla/mux"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/logger"
	"github.com/spf13/pflag"
)

type MyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Server struct {
	calendar calendar.CalendarInterface
	ctx      context.Context
}

type eventRequest struct {
	ID           string
	Title        string
	Notice       string
	DateStarted  time.Time
	DateComplete time.Time
}

func (err *MyError) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}

func sendResponse(msg []byte, code int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", msg)
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
	cal := calendar.New(storage)
	logger.Info("Calendar was created")
	InitServer(conf.HttpListen.Ip, conf.HttpListen.Port, cal)
}

//Init http server
func InitServer(listenIp string, listenPort string, calendar calendar.CalendarInterface) {
	server := &Server{calendar: calendar, ctx: context.Background()}
	router := mux.NewRouter()
	router.HandleFunc("/hello", server.hello).Methods("GET")
	router.HandleFunc("/add", server.add).Methods("POST")
	router.HandleFunc("/edit/{id}", server.edit).Methods("POST")
	router.HandleFunc("/get", server.get).Methods("GET")
	router.HandleFunc("/get/{id}", server.getById).Methods("GET")
	router.HandleFunc("/delete/{id}", server.delete).Methods("POST")
	router.HandleFunc("/events", server.events).Queries("time_before", "{time_before}", "time_length", "{time_length}").Methods("GET")

	srv := &http.Server{
		Addr:         listenIp + ":" + listenPort,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error while starting HTTP server", "error", err)
		}
	}()
	logger.Info("HTTP server started on host: " + listenIp + ", port: " + listenPort)

	<-done
	logger.Info("HTTP server stopped")

	ctx, cancel := context.WithTimeout(server.ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown failed", "error", err)
	}
	logger.Info("HTTP server exited properly")
}

func (s Server) hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming message",
		"host", r.Host,
		"url", r.URL.Path)
	message := []byte(`{"Message":"Hello world"}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", message)
}

func (s Server) add(w http.ResponseWriter, r *http.Request) {
	logger.Info("Ger request add")
	data := &eventRequest{}
	var err error
	data.Title = r.FormValue("title")
	data.Notice = r.FormValue("notice")
	data.DateStarted, err = time.Parse("2006-01-02 15:04:05", r.FormValue("dateStarted"))
	if err != nil {
		msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	data.DateComplete, err = time.Parse("2006-01-02 15:04:05", r.FormValue("dateComplete"))
	if err != nil {
		msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	if data.DateStarted.After(data.DateComplete) {
		msg, _ := json.Marshal(MyError{http.StatusBadRequest, "Check dates date_complete before date_started"})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	ip := r.RemoteAddr
	eventId, err := s.calendar.Add(data.Title, data.DateStarted, data.DateComplete, data.Notice, iptoInt(ip))
	if err != nil {
		if err == types.ErrDateBusy {
			msg, _ := json.Marshal(MyError{http.StatusBadRequest, err.Error()})
			sendResponse(msg, http.StatusBadRequest, w)
			return
		} else {
			msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
			sendResponse(msg, http.StatusInternalServerError, w)
			return
		}
	}
	message := fmt.Sprintf(`{"id":"%s"}`, eventId)
	msg := []byte(message)
	sendResponse(msg, http.StatusOK, w)
}

func (s Server) edit(w http.ResponseWriter, r *http.Request) {
	logger.Info("Ger request edit")
	data := &eventRequest{}
	vars := mux.Vars(r)
	var err error
	data.ID = r.FormValue("id")
	data.Title = r.FormValue("title")
	data.Notice = r.FormValue("notice")
	data.DateStarted, err = time.Parse("2006-01-02 15:04:05", r.FormValue("dateStarted"))
	if err != nil {
		msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	data.DateComplete, err = time.Parse("2006-01-02 15:04:05", r.FormValue("dateComplete"))
	if err != nil {
		msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	if data.DateStarted.After(data.DateComplete) {
		msg, _ := json.Marshal(MyError{http.StatusBadRequest, "Check dates date_complete before date_started"})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	ip := r.RemoteAddr
	err = s.calendar.Edit(vars["id"], models.Event{ID: data.ID, Title: data.Title, DateStarted: data.DateStarted, DateComplete: data.DateComplete, Notice: data.Notice}, iptoInt(ip))
	if err != nil {
		if err == types.ErrDateBusy {
			msg, _ := json.Marshal(MyError{http.StatusBadRequest, err.Error()})
			sendResponse(msg, http.StatusBadRequest, w)
			return
		} else {
			msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
			sendResponse(msg, http.StatusInternalServerError, w)
			return
		}
	}
	message := fmt.Sprintf(`{"Message":"Event with id %s was changed"}`, vars["id"])
	msg := []byte(message)
	sendResponse(msg, http.StatusOK, w)
}

func (s Server) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logger.Info("Incoming message get",
		"host", r.Host,
		"url", r.URL.Path)
	events, err := s.calendar.GetEvents()
	if err != nil {
		if err == types.ErrNotFound {
			msg, _ := json.Marshal(MyError{http.StatusNotFound, err.Error()})
			sendResponse(msg, http.StatusNotFound, w)
			return
		} else {
			msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
			sendResponse(msg, http.StatusInternalServerError, w)
			return
		}
	}
	ev, err := json.Marshal(events)
	if err != nil {
		msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	sendResponse(ev, http.StatusOK, w)
}

func (s Server) getById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	logger.Info("Incoming message get by id",
		"host", r.Host,
		"url", r.URL.Path)

	event, err := s.calendar.GetEventByID(vars["id"])
	if err != nil {
		if err == types.ErrNotFound {
			msg, _ := json.Marshal(MyError{http.StatusNotFound, err.Error()})
			sendResponse(msg, http.StatusNotFound, w)
			return
		} else {
			msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
			sendResponse(msg, http.StatusInternalServerError, w)
			return
		}
	}
	ev, err := json.Marshal(event)
	if err != nil {
		msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	sendResponse(ev, http.StatusOK, w)
}

func (s Server) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logger.Info("Incoming message delete id",
		"host", r.Host,
		"url", r.URL.Path)
	err := s.calendar.Delete(vars["id"])
	if err != nil {
		if err == types.ErrNotFound {
			msg, _ := json.Marshal(MyError{http.StatusNotFound, err.Error()})
			sendResponse(msg, http.StatusNotFound, w)
			return
		} else {
			msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
			sendResponse(msg, http.StatusInternalServerError, w)
			return
		}
	}
	message := fmt.Sprintf(`{"Message":"event with id %s was deleted"}`, vars["id"])
	msg := []byte(message)
	sendResponse(msg, http.StatusOK, w)
}

func (s Server) events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logger.Info("Incoming message events",
		"host", r.Host,
		"url", r.URL.Path)
	v := r.URL.Query()
	timeBefore := v.Get("time_before")
	timeLength := v.Get("time_length")
	if timeBefore == "" || timeLength == "" {
		msg, _ := json.Marshal(MyError{http.StatusBadRequest, "Query parameters time_before and time_length required"})
		sendResponse(msg, http.StatusBadRequest, w)
		return
	}
	events, err := s.calendar.GetEventsByStartPeriod(timeBefore, timeLength)
	if err != nil {
		if err == types.ErrNotFound {
			msg, _ := json.Marshal(MyError{http.StatusNotFound, err.Error()})
			sendResponse(msg, http.StatusNotFound, w)
			return
		} else {
			msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
			sendResponse(msg, http.StatusInternalServerError, w)
			return
		}
	}
	ev, err := json.Marshal(events)
	if err != nil {
		msg, _ := json.Marshal(MyError{http.StatusInternalServerError, err.Error()})
		sendResponse(msg, http.StatusInternalServerError, w)
		return
	}
	sendResponse(ev, http.StatusOK, w)
}

func iptoInt(ip string) int64 {
	bits := strings.Split(ip, ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])
	var sum int64
	// left shifting 24,16,8,0 and bitwise OR
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
