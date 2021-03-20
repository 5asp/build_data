package main

import (
	"fmt"
	"github.com/t0of/qbang/content"
	"github.com/t0of/qbang/db"
)

const Path = "/Users/xxuechin/Downloads/17eQbang.003.txt"

func main() {
	err := db.ConnDb()
	if err != nil {
		fmt.Println(err)
	}
	exec := content.NewContent(Path)
	exec.Do()
}
