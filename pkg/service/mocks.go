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
func (m *MockDB) SaveLink(originalLink, shortenedLink string) error {
	duplicate, err := m.IsDuplicate(shortenedLink)
	if err != nil {
		return fmt.Errorf("cannot check if duplicate")
	}
	if duplicate {
		return repository.ErrLinkAlreadyExists
	}

	m.DB[shortenedLink] = originalLink

	return nil
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
func (m *MockDB) IsDuplicate(shortenedLink string) (bool, error) {
	_, ok := m.DB[shortenedLink]
	if ok {
		return true, nil
	}

	return false, nil
}
