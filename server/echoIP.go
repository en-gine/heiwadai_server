package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// IPアドレスを取得する
func EchoMyIP() {
	url := "http://checkip.dyndns.com/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	htmlData := string(body)
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	text := extractText(doc)
	fmt.Println(text)
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type == html.ElementNode && n.Data == "script" || n.Data == "style" {
		return ""
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}

	return strings.Join(strings.Fields(text), " ")
}
