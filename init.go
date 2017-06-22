package main

import (
	"io/ioutil"
	"os"
)

func initDir() {
	if _, err := os.Stat("source/_posts"); os.IsNotExist(err) {
		os.MkdirAll("source/_posts", 0755)
	}
}

func initConfig() {
	fn := "_config.yml"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		config := []byte(
			`# This is config file`,
		)
		err := ioutil.WriteFile(fn, config, 0644)
		if err != nil {
			panic("Can not create config file _config.yml")
		}
	}
}

func InitBlog() {
	initDir()
	initConfig()
}
