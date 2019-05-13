package content

import (
	"os"
	"path/filepath"
	"strconv"
)

type book struct {
	head     map[int]string
	chapters map[int]*chapter
}

func NewBook(sourcePath string) *book {

	bookDir, err := os.Open(sourcePath)
	if err != nil {
		return nil
	}

	chapterDirs, err := bookDir.Readdir(0)
	if err != nil {
		return nil
	}

	book := book{}
	book.chapters = make(map[int]*chapter, len(chapterDirs))
	for _, chapterDir := range chapterDirs {
		chapterName := chapterDir.Name()
		i, _ := strconv.Atoi(chapterName[:3])
		book.chapters[i] = NewChapter(filepath.Join(sourcePath, chapterName))
	}

	return &book
}

func (p *book) GenerateHTML() {
}
