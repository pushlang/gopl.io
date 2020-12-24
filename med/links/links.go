package links

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/html"
)

var signature = []string{"h4", "div", "h4", "div", "div", "a"}

type Extractor interface {
	Extract() ([]string, error)
}

func extract(doc *html.Node) ([]string, error) {
	var isIn = make(map[string]bool)
	/*type link struct {
		href string
		text string
	}
	var links []link*/
	var links []string

	sl := len(signature)
	var history = make([]string, 0, sl)
	hl := 0

	visitNode := func(n *html.Node) {
		hl = len(history)

		fc := n.FirstChild

		if n.Type == html.ElementNode && fc != nil && fc.Type == html.TextNode && len(fc.Data) > 5 {
			if hl == 0 {
				if n.Data == signature[0] {
					fmt.Println("text:", fc.Data)
					history = append(history, n.Data)
				}
			} else {
				if hl <= sl && history[hl-1] == signature[hl-1] {
					fmt.Println("text:", fc.Data)
					history = append(history, n.Data)
				} else {
					history = nil
				}
			}
		}

		if n.Type == html.ElementNode && n.Data == "a" {
			if len(signature) == len(history) {
				for _, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					_, err := url.Parse(a.Val)
					if err != nil {
						continue // ignore bad URLs
					}
					//log.Println(strings.Join(history, ","))
					history = nil
					if !isIn[a.Val] {
						isIn[a.Val] = true

						fmt.Println("link:", a.Val)
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
