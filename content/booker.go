package content

import (
	"io"
	"os"
)

type Booker struct {
	books map[string]*book
}

func NewBooker() *Booker {
	return &Booker{make(map[string]*book)}
}

func (p *Booker) AddBook(name, path string) error {
	p.books[name] = NewBook(name, path)
	return nil
}

func (p *Booker) Write(name string, writer io.Writer) (int, error) {
	data := []byte(p.books[name].String())
	return writer.Write(data)
}

func (p *Booker) WriteFile(name, path string) (int, error) {

	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return p.Write(name, file)
}
