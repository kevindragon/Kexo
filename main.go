package main

import (
	//"fmt"
	//"github.com/russross/blackfriday"
	//"io/ioutil"
	"os"
)

func main() {
	InitBlog()

	if len(os.Args) <= 1 {
		Help()
	} else {
		RunCommand(os.Args[1:])
	}

	/*
		bts, _ := ioutil.ReadFile("source_test/_posts/我来学Kotlin-基础之包.md")
		output := blackfriday.MarkdownCommon(bts)
		fmt.Println(string(output))
	*/
}
