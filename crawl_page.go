package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(currentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	// fmt.Printf("Current pages: %v\nMax pages: %v\n", len(cfg.pages), cfg.maxPages)
	if !cfg.canCrawl() {
		fmt.Println("Max pages reached")
		return
	}

	parsedCurrentURL, err := url.Parse(currentURL)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(currentURL)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	first := cfg.addPageVisit(normalizedURL)

	if !first {
		return
	}

	html, err := getHTML(currentURL)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	links, err := GetURLSFromHTML(html, cfg.baseURL)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Printf("\tGetting ready to crawl %v pages\n", len(links))

	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}

}

func (cfg *config) canCrawl() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if len(cfg.pages) >= cfg.maxPages {
		return false
	}
	return true
}
