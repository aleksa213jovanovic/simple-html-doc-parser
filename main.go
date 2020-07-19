package main

import (
	"fmt"
	"htmlParser/parser"
	"log"
	"os"
)

var stringHTML = `
<html>
	<body>
		<h1>Hello!</h1>
		<a href="/dog">
			A link to another page
			<span>Something in a span</span>
		</a>
		<b>Bold text!</b>
		<a href="/page-two">A link to a second page</a>
		</a>
	</body>
</html>
`

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	fileHTML, err := os.Open("/home/aleksa/Desktop/response.html")
	checkError(err)
	links, err := parser.Parse(fileHTML)
	checkError(err)
	for _, link := range links {
		fmt.Printf("link = %s    text = %s\n", link.Href, link.Text)
	}
}
