package content

type book struct {
	chapters map[int]chapter
}

func NewBook(sourcePath string) *book {
	var book = book{}
	return &book
}

func (p *book) GenerateHTML() {
}
