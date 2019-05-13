package main

import (
	"GoCookbook/content"
	"fmt"
)

func main() {


	sourcePath := "source"
	book := content.NewBook(sourcePath)
	fmt.Println(sourcePath, book)
}
