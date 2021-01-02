package myshows

import (
	"sync"
	"net/http"
)

type myShows struct {
	t  string
	c  *http.Client
	mu sync.Mutex
}

func (m *myShows) SetToken(token string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.t = token
}

func (m *myShows) Auth(id, secret, user, password string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.auth(id, secret, user, password)
}

func (m *myShows) GetUnwatchedEpisodes() ([]EpisodeDesc, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.getUnwatchedEpisodes()
}

func New() *myShows {
	m := new(myShows)
	m.c = httpClient
	return m
}
