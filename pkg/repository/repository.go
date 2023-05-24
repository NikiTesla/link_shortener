package repository

import (
	"errors"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo interface {
	SaveLink(originalLink, shortenedLink string) (string, error)
	GetLink(shortenedLink string) (string, error)
	IsDuplicate(shortenedLink, originalLink string) (shortDuplicate bool, origDuplicateShort string, err error)
}

// PostgresDB is an implementation of Repo interface
type PostgresDB struct {
	DB *pgxpool.Pool
}

// InMemoryDB is an implementation of Repo interface
type InMemoryDB struct {
	sync.RWMutex
	DB map[string]string
}

var ErrLinkNotFound = errors.New("there is no such shortened link in database")
var ErrShortLinkAlreadyExists = errors.New("shortened link already exists")
