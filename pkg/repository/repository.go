package repository

import (
	"errors"
	"sync"

	"github.com/jackc/pgx/v5"
)

type Repo interface {
	SaveLink(originalLink, shortenedLink string) error
	GetLink(shortenedLink string) (string, error)
}

type Postgres struct {
	db *pgx.Conn
}

type InMemory struct {
	sync.RWMutex
	m map[string]string
}

var ErrLinkNotFound = errors.New("there is no such shortened link in database")
