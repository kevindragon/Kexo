package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func main() {
	bts, _ := ioutil.ReadFile("source/_posts/我来学Kotlin-基础之包.md")

	output := blackfriday.MarkdownCommon(bts)

	fmt.Println(string(output))
}
