package service

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/NikiTesla/link_shortener/api"
	"github.com/NikiTesla/link_shortener/pkg/environment"
)

func TestSaveOriginal(t *testing.T) {
	mock := &MockDB{make(map[string]string)}
	mock.DB["0testlink0"] = "https://www.ozon.ru/"

	server := NewShortenerServer(&environment.Environment{DB: mock})

	data := []*pb.SaveOriginalRequest{
		{OriginalLink: ""},
		{OriginalLink: "https://www.ozon.ru/"},
	}

	expected_error := []error{
		fmt.Errorf("recieved empty link"),
		nil,
	}

	for i := 0; i < len(data); i++ {
		resp, err := server.SaveOriginal(context.Background(), data[i])
		if err != nil {
			if err.Error() != expected_error[i].Error() {
				t.Errorf("error/expected error mismatch: %s / %s", err, expected_error[i])
			}
		} else {
			if len(resp.GetShortedLink()) != linkLength {
				t.Errorf("length is not equal %d", linkLength)
			}
			if expected_error[i] != nil {
				t.Errorf("expected not nil error: %s", expected_error[i])
			}
		}
	}

}

func TestGetOriginal(t *testing.T) {
	mock := &MockDB{make(map[string]string)}
	mock.DB["0testlink0"] = "https://www.ozon.ru/"

	server := NewShortenerServer(&environment.Environment{DB: mock})

	data := []*pb.GetOriginalRequest{
		{ShortedLink: ""},
		{ShortedLink: "0testlink0"},
	}
	expected_reponse := []*pb.GetOriginalResponse{
		{},
		{OriginalLink: "https://www.ozon.ru/"},
	}
	expected_error := []error{
		fmt.Errorf("recieved empty link"),
		nil,
	}

	for i := 0; i < len(data); i++ {
		resp, err := server.GetOriginal(context.Background(), data[i])
		if resp.GetOriginalLink() != expected_reponse[i].GetOriginalLink() {
			t.Errorf("response/expected response mismatch: %s / %s", resp, expected_reponse[i])
		}
		if err != nil {
			if err.Error() != expected_error[i].Error() {
				t.Errorf("error/expected error mismatch: %s / %s", err, expected_error[i])
			}
		} else {
			if expected_error[i] != nil {
				t.Errorf("expected not nil error: %s", expected_error[i])
			}
		}
	}
}

func TestGenerateShortenedLink(t *testing.T) {
	mock := &MockDB{make(map[string]string)}
	mock.DB["0testlink0"] = "https://www.ozon.ru/"

	server := NewShortenerServer(&environment.Environment{DB: mock})

	originalLink, err := server.generateShortenedLink("http://localhost:8080/test")
	if err != nil {
		t.Errorf("error is not nil: %s", err)
	}

	if len(originalLink) != linkLength {
		t.Errorf("length of shortened link is %d and not %d", len(originalLink), linkLength)
	}
}
