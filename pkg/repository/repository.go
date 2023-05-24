package repository

import (
	"errors"
	"sync"

	"github.com/jackc/pgx"
)

type Repo interface {
	SaveLink(originalLink, shortenedLink string) error
	GetLink(shortenedLink string) (string, error)
}

type PostgresDB struct {
	DB *pgx.Conn
}

type InMemoryDB struct {
	sync.RWMutex
	DB map[string]string
}

var ErrLinkNotFound = errors.New("there is no such shortened link in database")
