package content

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type book struct {
	countAll int
	texts    map[int]string
	chapters map[int]*chapter
}

func NewBook(sourcePath string) *book {

	var book = book{}

	bookDir, err := os.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	chapterDirs, err := bookDir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	book.countAll = len(chapterDirs)

	// Посчитаем количество файлов с текстом, и каталогов с частями книги
	countFiles, countDirs := 0, 0
	for _, chapterDir := range chapterDirs {
		if chapterDir.IsDir() {
			countDirs++
		} else {
			countFiles++
		}
	}

	fmt.Println(countFiles, countDirs)

	book.texts = make(map[int]string, countFiles)
	book.chapters = make(map[int]*chapter, countDirs)
	for _, chapterDir := range chapterDirs {
		name := chapterDir.Name()
		fullName := filepath.Join(sourcePath, name)
		num, _ := strconv.Atoi(name[:3])
		if chapterDir.IsDir() {
			book.chapters[num] = NewChapter(num, fullName)
		} else {
			if data, err := ioutil.ReadFile(fullName); err == nil {
				book.texts[num] = string(data)
			}
		}
	}

	return &book
}

func (p *book) GenerateHTML() {
}
