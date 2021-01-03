package links

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

//link to article,link name - "div","h2"
var thisOnly = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}

//link to author, link name - "h4"
//var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}
//link to site, link name - "div", "div", "h4"
//var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}
//link to image - "div", "div", "img"
//var signature = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "a"}

// Extractor for implementing methods of getting html body
type Extractor interface {
	Extract() ([]Link, error)
}

// Link contains found links
type Link struct {
	Name string
	Url  string
}

// sign contains the number of found links with the same tag-paths
type sign struct {
	id []string
	n  int
}

//var signs []sign

// extract searches for urls in html body
func extract(doc *html.Node, su Url) ([]Link, error) {
	log.Println("enter extract")
	defer log.Println("exit extract")

	var links []Link
	var isIn = make(map[string]bool)

	var history []string
	var hl, sl int

	var signature []string

	in := func(n *html.Node) {
		//fmt.Println(strings.Join(signature, ":"))
		if n.Type == html.ElementNode {

			history = append(history, n.Data) // save path of crawling
			l := len(history) - 1
			// if current node tag eq signature tag at node depth
			if l < sl && n.Data == signature[l] {
				hl++
			}
			// if path of crawling eq signature
			if sl == hl {
				// if tag is anchor
				if n.Data == "a" {
					//fmt.Println("history a:", strings.Join(history, " "))
					// looking for href=url in anchors attributes
					for _, a := range n.Attr {
						if a.Key != "href" {
							continue
						}
						// ignore bad url
						psu, err := url.Parse(string(su))
						if err != nil {
							continue
						}
						cu, _ := psu.Parse(a.Val)
						// if its a new url
						if !isIn[a.Val] {
							text := ""
							in := func(n *html.Node) {
								if n.Type == html.TextNode && len(n.Data) > 0 && len(text) == 0 {
									text += n.Data
								}
							}
							url := cu.Scheme + "://" + cu.Host + cu.Path

							forEachNode(n, in, nil) // find url name

							isIn[a.Val] = true // register new url

							links = append(links, Link{text, url}) // add url to link list

						}
					}
				}
			}
		} // if == ElementNode
	} //in

	out := func(n *html.Node) {
		l := len(history) - 1
		if n.Type == html.ElementNode {
			if l < sl && n.Data == signature[l] {
				hl--
			}
			history = history[:l]
		}
	}

	// scaning all link signatures
	for _, s := range countLinkSignatures(doc) {
		//if strings.Join(s.id, ":") == strings.Join(thisOnly, ":") {
		history = nil
		hl = 0

		signature = s.id
		sl = len(signature)
		forEachNode(doc, in, out)
		//}
	}

	return links, nil
}

// countLinkSignatures counts links this the same path
func countLinkSignatures(n *html.Node) []sign {
	log.Println("enter countLinkSignatures")
	defer log.Println("exit countLinkSignatures")

	var signCount = make(map[string]int)
	var s []sign

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
		s = append(s, sign{strings.Split(k, ":"), v})
	}
	return s
}

// printTree prints html node tree
func printTree(n *html.Node) {
	log.Println("enter printTree")
	defer log.Println("exit printTree")

	depth := 0

	in := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	out := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	forEachNode(n, in, out)
}

// CleanHTMLbody cleans off html-body from script tags
func CleanHTMLbody(body io.Reader) (io.Reader, error) {
	log.Println("enter CleanHTMLbody")
	defer log.Println("exit CleanHTMLbody")

	doc, err := html.Parse(body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - Body for a clean not parsed => %v", err))
	}
	i := 0
	out := func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				if c.Data == "script" {
					i++
					n.RemoveChild(c)
					if n.FirstChild != nil {
						c = n.FirstChild
					}
				}
			}
		}
	}
	forEachNode(doc, nil, out)

	log.Println("Removed scripts: ", i)

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - Cleaned body not rendered => %v", err))
	}
	rd := strings.NewReader(buf.String())

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - Cleaned body not copied => %v", err))
	}
	return rd, nil
}

// forEachNode crawls all html nodes
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

// Extract opens html-file contains urls
func (f FileName) Extract() ([]Link, error) {
	log.Println("enter FileName.Extract")
	defer log.Println("exit FileName.Extract")

	fd, err := os.Open(string(f))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - File not opened: %s => %v", f, err))
	}
	defer fd.Close()

	doc, err := html.Parse(fd)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - Html not parsed: %s => %v", f, err))
	}

	return extract(doc, Url(""))
}

type Url string

// Extract sends http-request to web-page contains urls
func (u Url) Extract() ([]Link, error) {
	log.Println("enter Url.Extract")
	defer log.Println("exit Url.Extract")

	req, err := http.NewRequest("GET", string(u), nil)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - NewRequest: %s => %v", string(u), err))
	}

	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	now := time.Now()
	for ; err != nil; time.Sleep(10 * time.Millisecond) {
		log.Println("Attempt to connect:" + string(u))
		if time.Since(now).Milliseconds() > 1000 {
			break
		}
		resp, err = http.DefaultClient.Do(req)
	}
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - Request timeout: %s => %v", string(u), err))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Error - Request StatusCode: %s => %v", string(u), err))
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error - Response body parsing: %s => %v", string(u), err))
	}
	return extract(doc, u)
}
