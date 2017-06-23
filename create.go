package main

import (
	"io/ioutil"
	"path"
	"strings"
	"time"
)

func Create(args []string) {
	if len(args) <= 0 {
		HelpCmd(NewCmd)
		return
	}

	if len(args) == 1 {
		createPost(args[0])
		return
	}
}

var postTpl string = `---
title: {{title}}
date: {{date}}
tags:
---

`

func createPost(identity string) {
	filename := path.Join(POSTS_PATH, identity+".md")
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")
	content := strings.Replace(postTpl, "{{title}}", identity, 1)
	content = strings.Replace(content, "{{date}}", date, 1)
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
