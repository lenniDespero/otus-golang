package sql

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
)

// Storage struct
type Storage struct {
	ConnPool *pgxpool.Pool
}

//New returns new storage
func New(dbconf *config.DBConfig) (*Storage, error) {
	storage := &Storage{}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbconf.User, dbconf.Password, dbconf.Host, dbconf.Port, dbconf.Database)
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse connection configs")
	}
	cfg.MaxConns = 8
	cfg.ConnConfig.TLSConfig = nil
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: 5 * time.Minute,
		Timeout:   1 * time.Second,
	}).DialContext

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to postgres")
	}
	storage.ConnPool = pool
	return storage, nil
}

func (s Storage) Add(event models.Event) (string, error) {
	panic("implement me")
}

func (s Storage) Edit(id string, event models.Event) error {
	panic("implement me")
}

func (s Storage) GetEvents() ([]models.Event, error) {
	panic("implement me")
}

func (s Storage) GetEventByID(id string) ([]models.Event, error) {
	panic("implement me")
}

func (s Storage) Delete(id string) error {
	panic("implement me")
}
