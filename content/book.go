package content

import (
	"fmt"
	"path/filepath"
	"strings"
)

// КНИГА
type book struct {
	name  string
	parts map[int]fmt.Stringer
}

func NewBook(name, path string) *book {

	items := getItems(path)

	var book = book{}
	book.name = name
	book.parts = make(map[int]fmt.Stringer, len(items))

	for _, fi := range items {

		num, _ := parseName(fi.Name())

		fullName := filepath.Join(path, fi.Name())
		if fi.IsDir() {
			book.parts[num] = NewChapter(fullName)
		} else {
			book.parts[num] = NewMDText(fullName, 2)
		}
	}

	return &book
}

func (p *book) GenerateHTML() {
}

func (p *book) String() string {

	var b strings.Builder

	b.Grow(len(p.name) + 1)
	b.WriteString("#")
	b.WriteString(p.name)

	for i := 1; i < len(p.parts); i++ {
		b.WriteString("\r\n")
		b.WriteString(p.parts[i].String())
	}

	return b.String()
}
