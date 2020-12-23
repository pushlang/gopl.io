package links

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

var signature = []string{"h4", "div", "h4", "div", "div", "a"}

type Extractor interface {
	Extract() ([]string, error)
}

func extract(doc *html.Node) ([]string, error) {
	var isIn = make(map[string]bool)
	var links []string
	var history = make([]string, 0)

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode {
			history = append(history, "\""+n.Data+"\"")
		}

		if n.Type == html.ElementNode && n.Data == "a" {
			if signature == history {
				for _, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					_, err := url.Parse(a.Val)
					if err != nil {
						continue // ignore bad URLs
					}
					log.Println(strings.Join(history, ","))
					history = nil
					if !isIn[a.Val] {
						isIn[a.Val] = true
						links = append(links, a.Val)
					}
				}
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
		return nil, fmt.Errorf("Parsing as HTML: %v", err)
	}

	return extract(doc)
}

type Url string

func (u Url) Extract() ([]string, error) {
	log.Println("Url.Extract")
	req, err := http.NewRequest("GET", string(u), nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	//fmt.Println(resp.Header)

	for err != nil {
		time.Sleep(250 * time.Millisecond)
		log.Printf("Error:%s\n", err)
		resp, err = http.DefaultClient.Do(req)
	}
	fmt.Printf("Status code:%d\n", resp.StatusCode)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Getting: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing HTML: %v", err)
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
