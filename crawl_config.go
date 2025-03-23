package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.pages[normalizedURL]

	if ok {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func NewConfig(baseURL string, maxPages, maxConcurrentRequests int) (*config, error) {
	parsedURL, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	return &config{
		pages:              make(map[string]int),
		maxPages:           maxPages,
		baseURL:            parsedURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrentRequests),
		wg:                 &sync.WaitGroup{},
	}, nil

}
