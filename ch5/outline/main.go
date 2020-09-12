// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(strings.NewReader(hh)) //os.Stdin
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	siblings(nil, doc) //передаем корень html-дерева
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}

	c := n.FirstChild
	if c != nil {
		outline(stack, c)

		if c.NextSibling != nil {
			siblings(stack, c.NextSibling)
		}
	}
}

//п.1. Регистрируем входящий узел. 
//п.1.1. Рекурсивно опускаемся по первым детям до последнего узла ветки. 
//п.1.2. Если узел последний (без детей), то возвращаемся на предыдущий уровень рекурсии.  
//п.1.3. На предыдущем уровне, проверяем текущий узел на наличие сиблингов первого ребенка 
//(от которого вернулись), если есть, то с первым сиблингом переходим в siblings.

//п.2 С первым сиблингом заходим в outline и далее по п.1, возвращаемся смотрим следующего...
//Оutline прогоняет по уровням, а siblings прокручивает по всем сиблингам узла. 

func siblings(stack []string, n *html.Node) {
	outline(stack, n)

	if s := n.NextSibling; s != nil { //
		siblings(stack, s)
	}
}

var hh string = `
<html>
	<head>
		<title>Title</title>
		<meta charset="utf-8">
		<meta name="GENERATOR" content="Microsoft FrontPage 4.0">
		<meta name="ProgId" content="FrontPage.Editor.Document">
	</head>
	<body>
		<h1>2 issues</h1>
		<table>
			<tr style='text-align: left'>
				<th>#</th>
				<th>State</th>
				<th>User</th>
				<th>Title</th>
			</tr>
			<tr>
				<td><a href='https://github.com/golang/go/issues/10535'>10535</a></td>
				<td>open</td>
				<td><a href='https://github.com/dvyukov'>dvyukov</a></td>
				<td><a href='https://github.com/golang/go/issues/10535'>x/net/html: void element &lt;link&gt; has child nodes</a></td>
			</tr>
			<tr>
				<td><a href='https://github.com/golang/go/issues/3133'>3133</a></td>
				<td>closed</td>
				<td><a href='https://github.com/ukai'>ukai</a></td>
				<td><a href='https://github.com/golang/go/issues/3133'>html/template: escape xmldesc as &amp;lt;?xml</a></td>
			</tr>
		</table>
	</body>
</html>`
