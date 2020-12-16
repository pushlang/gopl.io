package links

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

type Extractor interface {
	Extract() ([]string, error)
}

func extract(doc *html.Node) ([]string, error) {
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := url.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

type FileName string

func (f FileName) Extract() ([]string, error) {
	fd, err := os.Open(string(f))
	if err != nil {
		log.Println(err)
	}
	doc, err := html.Parse(fd)

	if err != nil {
		return nil, fmt.Errorf("parsing as HTML: %v", err)
	}

	return extract(doc)
}

type Url string

func (u Url) Extract() ([]string, error) {
	resp, err := http.Get(string(u))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing as HTML: %v", err)
	}
	return extract(doc)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
