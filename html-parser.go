package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLSFromHTML(body string, baseURL *url.URL) ([]string, error) {
	urls := []string{}
	reader := strings.NewReader(body)

	parsedHTML, err := html.Parse(reader)

	nodes := GetAllHtmlNodesOfTag(parsedHTML, "a")

	for _, node := range nodes {
		attrs := node.Attr

		for _, attr := range attrs {
			if attr.Key == "href" {
				href, err := url.Parse(attr.Val)

				if err != nil {
					fmt.Printf("could not parse href '%v': %v\n", attr.Val, err)
					continue
				}

				resolvedURL := baseURL.ResolveReference(href)

				urls = append(urls, resolvedURL.String())
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return urls, nil
}

func GetAllHtmlNodesOfTag(rootNode *html.Node, tag string) []html.Node {
	found := []html.Node{}

	if rootNode.Type == html.ElementNode && strings.ToLower(rootNode.DataAtom.String()) == strings.ToLower(tag) {
		found = append(found, *rootNode)
	}
	childrenIterator := rootNode.ChildNodes()

	for child := range childrenIterator {
		found = append(found, GetAllHtmlNodesOfTag(child, tag)...)
	}

	return found
}
