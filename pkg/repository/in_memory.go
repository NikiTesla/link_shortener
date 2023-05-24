package repository

import (
	"fmt"
)

// SaveLink gets original shortened links and save them in database
// Uses lock for writing
func (im *InMemoryDB) SaveLink(originalLink, shortenedLink string) (string, error) {
	im.Lock()
	defer im.Unlock()

	shortDupl, origDuplShort, err := im.IsDuplicate(shortenedLink, originalLink)
	if err != nil {
		return "", fmt.Errorf("cannot check if duplicate")
	}
	if shortDupl {
		return "", ErrLinkAlreadyExists
	}
	if origDuplShort != "" {
		return origDuplShort, nil
	}

	im.DB[shortenedLink] = originalLink

	return shortenedLink, nil
}

// GetLink gets shortened link and search in database for the original one
// If there is no such saved link, returns ErrLinkNotFound error
// Uses read lock for reading
func (im *InMemoryDB) GetLink(shortenedLink string) (string, error) {
	im.RLock()
	originalLink, ok := im.DB[shortenedLink]
	im.RUnlock()

	if !ok {
		return "", ErrLinkNotFound
	}

	return originalLink, nil
}

// IsDuplicate checks if such shortened and original links already exist in storage
// Gets both shortened and original link and returns bool duplicate existance of short link.
// If original links was already saved, returns it's short version as second parameter, else ""
func (im *InMemoryDB) IsDuplicate(shortenedLink, originalLink string) (shortDuplicate bool, origDuplicateShort string, err error) {
	_, ok := im.DB[shortenedLink]
	if ok {
		shortDuplicate = true
	}

	for k, v := range im.DB {
		if v == originalLink {
			origDuplicateShort = k
		}
	}

	return shortDuplicate, origDuplicateShort, nil
}
