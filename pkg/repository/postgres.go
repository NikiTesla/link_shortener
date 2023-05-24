package repository

import (
	"context"
	"database/sql"
	"fmt"
)

// SaveLink insert new record in database with original link and unique shortened link.
// Uniqueness is checked with IsDuplicate method. If shortened link already exists, returns ErrLinkAlreadyExists.
// Transactions used
func (p *PostgresDB) SaveLink(originalLink, shortenedLink string) (string, error) {
	tx, err := p.DB.Begin(context.Background())
	if err != nil {
		return "", fmt.Errorf("cannot begin transaction: %s", err)
	}
	defer tx.Rollback(context.Background())

	shortDupl, origDuplShort, err := p.IsDuplicate(shortenedLink, originalLink)
	if err != nil {
		return "", fmt.Errorf("cannot check if duplicate: %s", err)
	}
	if shortDupl {
		return "", ErrLinkAlreadyExists
	}
	if origDuplShort != "" {
		return origDuplShort, nil
	}

	query := "INSERT INTO links(original, short) VALUES ($1, $2)"
	if _, err = tx.Exec(context.Background(), query, originalLink, shortenedLink); err != nil {
		return "", fmt.Errorf("cannot insert new link: %s", err)
	}

	if err = tx.Commit(context.Background()); err != nil {
		return "", fmt.Errorf("cannot commit transaction: %s", err)
	}

	return shortenedLink, nil
}

// GetLink searches record in database according to shortened link.
// Returns original link
func (p *PostgresDB) GetLink(shortenedLink string) (string, error) {
	tx, err := p.DB.Begin(context.Background())
	if err != nil {
		return "", fmt.Errorf("cannot begin transaction: %s", err)
	}
	defer tx.Rollback(context.Background())

	var originalLink string
	err = tx.QueryRow(context.Background(), "SELECT original from links where short = $1", shortenedLink).Scan(&originalLink)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrLinkNotFound
		} else {
			return "", fmt.Errorf("cannot get original link: %s", err)
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return originalLink, fmt.Errorf("cannot commit transaction: %s", err)
	}

	return originalLink, nil
}

// IsDuplicate checks if shortened and original links exist in database
// Gets both shortened and original link and returns bool duplicate existance of short link.
// If original links was already saved, returns it's short version as second parameter, else ""
func (p *PostgresDB) IsDuplicate(shortenedLink, originalLink string) (shortDuplicate bool, OrigDuplicateShort string, err error) {
	err = p.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT id FROM links WHERE short = $1)",
		originalLink).Scan(&shortDuplicate)
	if err != nil {
		return false, "", err
	}

	err = p.DB.QueryRow(context.Background(), "SELECT short FROM links WHERE original = $1)").Scan(&OrigDuplicateShort)
	if err != nil {
		return shortDuplicate, OrigDuplicateShort, err
	}

	return shortDuplicate, OrigDuplicateShort, err
}
