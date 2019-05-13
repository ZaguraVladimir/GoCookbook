package main

import (
	"GoCookbook/content"
	"fmt"
)

func main() {

	sourcePath := "source"
	book, err := content.NewBook(sourcePath)
	if err != nil {
		fmt.Println("ERROR!!!")
	}
	fmt.Println(sourcePath, book)
}
