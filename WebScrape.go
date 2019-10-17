package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go Web Scraper v1.0")
	fmt.Println("Created by: Chad Humphries (GrunclePug)")
	fmt.Println("---------------------------------------------------------------------------------------------------------")
	fmt.Println("Input a website URL")

	for {
		fmt.Print("$> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, newLineCharacter(), "", -1)

		if strings.Compare(input, "quit") == 0 {
			os.Exit(1)
		} else if strings.Contains(input, " ") {
			fmt.Println("Result:")
			doc, _ := html.Parse(strings.NewReader(string(httpGet(strings.TrimPrefix(input, input[0:strings.Index(input, " ")+1])))))
			node, err := getHTMLElement(input[0:strings.Index(input, " ")], doc)
			if err != nil {
				fmt.Println(err)
				main()
			}
			fmt.Println(renderNode(node))
			fmt.Println("New location:")
		} else {
			fmt.Println("Results:")
			doc, _ := html.Parse(strings.NewReader(string(httpGet(input))))
			node, err := getHTMLElement("html", doc)
			if err != nil {
				fmt.Println(err)
				main()
			}
			fmt.Println(renderNode(node))
			fmt.Println("New location:")
		}
	}
}

//httpGet sends HTTP Get request and returns the page as a slice of bytes
func httpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

//getHTMLElement grabs a certain html element from the webpage
func getHTMLElement(element string, doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == element {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <" + element + "> in the node tree.")
}

//renderNode renders html node
func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	html.Render(io.Writer(&buf), n)

	return buf.String()
}

//newLineCharacter determines the new line character for your OS
func newLineCharacter() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}
