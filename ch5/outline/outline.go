// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package outline

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

//!+
func Run(r io.Reader) {
	doc, err := html.Parse(r)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		if err != nil {
			log.Println(err)
		}
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-
