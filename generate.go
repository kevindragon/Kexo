package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func withDir(path string, f func()) {
	currentDir, _ := os.Getwd()
	os.Chdir(POSTS_PATH)
	f()
	os.Chdir(currentDir)
}

func getFileList() []string {
	fileList := []string{}
	withDir(POSTS_PATH, func() {
		err := filepath.Walk(
			".",
			func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() && strings.HasSuffix(path, ".md") {
					fileList = append(fileList, path)
				}
				return nil
			},
		)
		if err != nil {
			panic(err)
		}
	})
	return fileList
}

func Generate() {
	fileList := getFileList()

	articles := []*Article{}
	for _, file := range fileList {
		articles = append(articles, parseFile(file))
	}

	fmt.Println(articles)
}

type Metadata struct {
	Title string   `yaml:"title"`
	Date  string   `yaml:"date"`
	Tags  []string `yaml:"tags"`
}

func (m Metadata) GetDate() time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", m.Date)
	return t
}

type Article struct {
	Header string
	Body   string
	Metadata
}

func parseHeader(header string) Metadata {
	meta := Metadata{}
	yaml.Unmarshal([]byte(header), &meta)

	return meta
}

func parseBody(body string) string {
	return string(blackfriday.MarkdownCommon([]byte(body)))
}

func readArticle(filename string) *Article {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content := string(bts)

	regex := regexp.MustCompile("(?m)^---+\\s*$")
	parts := regex.Split(content, -1)
	if len(parts) < 3 {
		return nil
	}
	header := parts[1]
	body := parts[2]
	metadata := parseHeader(header)

	return &Article{header, parseBody(body), metadata}
}

func parseFile(filename string) *Article {
	fullPath := path.Join(POSTS_PATH, filename)
	article := readArticle(fullPath)
	return article
}
