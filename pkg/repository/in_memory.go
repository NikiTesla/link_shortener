package repository

// SaveLink gets original shortened links and save them in database
// Uses lock for writing
func (im *InMemory) SaveLink(originalLink, shortenedLink string) error {
	im.Lock()
	im.m[shortenedLink] = originalLink
	im.Unlock()

	return nil
}

// GetLink gets shortened link and search in database for the original one
// If there is no such saved link, returns ErrLinkNotFound error
// Uses read lock for reading
func (im *InMemory) GetLink(shortenedLink string) (string, error) {
	im.RLock()
	originalLink, ok := im.m[shortenedLink]
	im.RUnlock()

	if !ok {
		return "", ErrLinkNotFound
	}

	return originalLink, nil
}
