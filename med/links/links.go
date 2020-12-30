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

//link to article,link name - "div","h2"
//var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}
//link to author, link name - "h4"
//var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}
//link to site, link name - "div", "div", "h4"
//var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}
//link to image - "div", "div", "img"
//var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}

type Extractor interface {
	Extract() ([]link, error)
}

type link struct {
	name string
	url  string
}

type signature struct {
	sign []string
	n    int
}

var signatures []signature

func extract(doc *html.Node) ([]link, error) {
	var links []link
	var isIn = make(map[string]bool)

	signature := countLinkSignatures(doc)[0].sign
	sl := len(signature)
	var history = make([]string, 0)
	hl := 0

	depth := 0

	in := func(n *html.Node) {
		if n.Type == html.ElementNode {
			history = append(history, n.Data)
			l := len(history) - 1
			if l < sl && n.Data == signature[l] {
				hl++
			}

			if sl == hl {
				if n.Data == "a" {
					//fmt.Println("history a:", strings.Join(history, " "))
					for _, a := range n.Attr {
						if a.Key != "href" {
							continue
						}
						if _, err := url.Parse(a.Val); err != nil {
							continue // ignore bad URLs
						}
						if !isIn[a.Val] {
							text := ""
							in := func(n *html.Node) {
								if n.Type == html.TextNode && len(n.Data) > 0 { //&& len(text) == 0 {
									text += n.Data
								}
							}
							forEachNode(n, in, nil)

							isIn[a.Val] = true

							links = append(links, link{text, a.Val})
						}
					}
				}
			}
		}
	}
	out := func(n *html.Node) {
		l := len(history) - 1
		if n.Type == html.ElementNode {
			if l < sl && n.Data == signature[l] {
				hl--
			}
			history = history[:l]
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

	return links, nil
}

// count signatures of links
func countLinkSignatures(n *html.Node) []signature {
	var signCount = make(map[string]int)
	var s []signature

	history := make([]string, 0)
	in := func(n *html.Node) {
		if n.Type == html.ElementNode {
			history = append(history, n.Data)
			if n.Data == "a" {
				signCount[strings.Join(history, ":")]++
			}
		}
	}
	out := func(n *html.Node) {
		if n.Type == html.ElementNode {
			history = history[:len(history)-1]
		}
	}
	forEachNode(n, in, out)

	for k, v := range signCount {
		s = append(s, signature{strings.Split(k, ":"), v})
	}
	return s
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

func (f FileName) Extract() ([]link, error) {
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

func (u Url) Extract() ([]link, error) {
	log.Println("Url.Extract")
	req, err := http.NewRequest("GET", string(u), nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)

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
