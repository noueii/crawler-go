package main

import "fmt"

func printReport(pages map[string]int, url string) {
	fmt.Println("=========================")
	fmt.Printf("\t REPORT for %s\n", url)
	fmt.Println("=========================")

	for key, val := range pages {
		fmt.Printf("Found %v internal links to %s\n", val, key)
	}
}
