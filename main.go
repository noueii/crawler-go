package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	var baseURL string
	var maxConcurrency int
	var maxPages int

	if len(args) < 3 {
		fmt.Println("invalid arguments")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL = args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Invalid argument 'maxConcurrency'. Expected integer.\nError: %s", err.Error())
		return
	}
	maxPages, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Invalid argument 'maxPages'. Expected integer.\nError: %s", err.Error())
		return
	}

	fmt.Println("------------------------")
	fmt.Printf("Starting crawl of: %s\n", baseURL)

	cfg, err := NewConfig(baseURL, maxPages, maxConcurrency)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cfg.wg.Add(1)
	cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)
}
