package content

import (
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"strings"
)

type booker struct {
	books map[string]*book
}

func NewBooker() *booker {
	return &booker{make(map[string]*book)}
}

func (p *booker) AddBook(name, path string) error {
	p.books[name] = newBook(name, path)
	return nil
}

func (p *booker) Write(name string, typeResult typeResult, writer io.Writer) {

	var data []byte
	switch typeResult {
	case md:
		data = []byte(p.books[name].String())
	case html:
		data = []byte(p.books[name].String())
	}

	if _, err := writer.Write(data); err != nil {
		log.Error(err)
	}
}

func (p *booker) WriteFile(name, path string) {

	file, err := os.Create(path)
	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	var typeResult typeResult
	if strings.HasSuffix(path, ".md") {
		typeResult = md
	} else if strings.HasSuffix(path, ".html") {
		typeResult = html
	}
	p.Write(name, typeResult, file)
}
