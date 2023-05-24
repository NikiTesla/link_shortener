package repository

func (p *PostgresDB) SaveLink(originalLink, shortenedLink string) error {
	return nil
}

func (p *PostgresDB) GetLink(shortenedLink string) (string, error) {
	return "", nil
}
