package repository

import "fmt"

// SaveLink insert new record in database with original link and unique shortened link.
// Uniqueness is checked with IsDuplicate method. If shortened link already exists, returns ErrLinkAlreadyExists.
// Transactions used
func (p *PostgresDB) SaveLink(originalLink, shortenedLink string) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return fmt.Errorf("cannot begin transaction: %s", err)
	}
	defer tx.Rollback()

	duplicate, err := p.IsDuplicate(shortenedLink)
	if err != nil {
		return fmt.Errorf("cannot check if duplicate: %s", err)
	}
	if duplicate {
		return ErrLinkAlreadyExists
	}

	query := "INSERT INTO links(original, short) VALUES ($1, $2)"
	if _, err = tx.Exec(query, originalLink, shortenedLink); err != nil {
		return fmt.Errorf("cannot insert new link: %s", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("cannot commit transaction: %s", err)
	}

	return nil
}

// GetLink searches record in database according to shortened link.
// Returns original link
func (p *PostgresDB) GetLink(shortenedLink string) (string, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return "", fmt.Errorf("cannot begin transaction: %s", err)
	}
	defer tx.Rollback()

	var originalLink string
	err = tx.QueryRow("SELECT original from links where short = $1", shortenedLink).Scan(&originalLink)
	if err != nil {
		return "", fmt.Errorf("cannot get original link: %s", err)
	}

	if err = tx.Commit(); err != nil {
		return originalLink, fmt.Errorf("cannot commit transaction: %s", err)
	}

	return originalLink, nil
}

// IsDuplicate checks if shortened link exists in database
func (p *PostgresDB) IsDuplicate(shortenedLink string) (bool, error) {
	var duplicate bool
	err := p.DB.QueryRow("SELECT EXISTS(SELECT id FROM links WHERE short = $1)",
		shortenedLink).Scan(&duplicate)

	return duplicate, err
}
