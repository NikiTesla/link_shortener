package service

import (
	"fmt"

	"github.com/NikiTesla/link_shortener/pkg/repository"
)

type MockDB struct {
	DB map[string]string
}

// SaveLink gets original shortened links and save them in database
// Uses lock for writing
func (m *MockDB) SaveLink(originalLink, shortenedLink string) (string, error) {
	shortDupl, origDuplShort, err := m.IsDuplicate(shortenedLink, originalLink)
	if err != nil {
		return "", fmt.Errorf("cannot check if duplicate")
	}
	if shortDupl {
		return "", repository.ErrShortLinkAlreadyExists
	}
	if origDuplShort != "" {
		return origDuplShort, nil
	}

	m.DB[shortenedLink] = originalLink

	return shortenedLink, nil
}

// GetLink gets shortened link and search in database for the original one
// If there is no such saved link, returns ErrLinkNotFound error
// Uses read lock for reading
func (m *MockDB) GetLink(shortenedLink string) (string, error) {
	originalLink, ok := m.DB[shortenedLink]

	if !ok {
		return "", repository.ErrLinkNotFound
	}

	return originalLink, nil
}

// IsDuplicate checks if such shortened link already exists in storage
func (m *MockDB) IsDuplicate(shortenedLink, originalLink string) (shortDuplicate bool, origDuplicateShort string, err error) {
	_, ok := m.DB[shortenedLink]
	if ok {
		shortDuplicate = true
	}

	for k, v := range m.DB {
		if v == originalLink {
			origDuplicateShort = k
		}
	}

	return shortDuplicate, origDuplicateShort, nil
}
