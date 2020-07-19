package parser

import (
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href, Text string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

//Parse will take in an HTML documment and will return
//a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	checkError(err)

	var ret []Link
	nodes := linkNodes(doc)
	for _, node := range nodes {
		ret = append(ret, buildLink(node))
	}
	return ret, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

//dfs is a function that does
// depth first search traversal of html tree
/*func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}*/
