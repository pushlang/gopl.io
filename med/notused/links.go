package links2

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

var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}

type Extractor interface {
	Extract() ([]string, error)
}

func extract(doc *html.Node) ([]string, error) {
	var signCount = make(map[string]int)
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

	depth := 0

	in := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if hl < sl && n.Data == signature[hl] {
				fmt.Print("+", n.Data, "=", signature[hl])
				hl++
			}

			if n.Data == "a" {
				if sl == hl {
					for _, a := range n.Attr {
						if a.Key != "href" {
							continue
						}
						if _, err := url.Parse(a.Val); err != nil {
							continue // ignore bad URLs
						}
						if !isIn[a.Val] {
							fc := n.FirstChild
							if fc != nil && fc.Data == "div" { //&& fc.Type == html.TextNode { //&& len(fc.Data) > 5 {
								if fc := fc.FirstChild; fc != nil && fc.Data == "h4" {
									if fc := fc.FirstChild; fc != nil { //&& fc.Data == "h4" {
										fmt.Println("text:", fc.Data)
									}
								}
							}

							isIn[a.Val] = true
							fmt.Println("link:", a.Val)
							links = append(links, a.Val)
						}
					}
				}
			}
		}
	}
	out := func(n *html.Node) {
		if n.Type == html.ElementNode {
			//hl--
		}
	}

	_ = func(n *html.Node) { // count signatures in
		hl = len(history)

		if n.Type == html.ElementNode {
			depth++
			history = append(history, "\""+n.Data+"\"")
			if n.Data == "a" {
				signCount[strings.Join(history, ", ")]++
			}
		}
	}
	_ = func(n *html.Node) { // count signatures out
		if n.Type == html.ElementNode {
			depth--
			history = history[:len(history)-1]
		}
	}
	_ = func(n *html.Node) { // prints html tree in
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	_ = func(n *html.Node) { // prints html tree out
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(doc, in, out)

	for k, v := range signCount {
		fmt.Printf("%s:%d\n", k, v)
	}

	return links, nil
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
