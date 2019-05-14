package content

import (
	"os"
	"path/filepath"
	"strconv"
)

type chapter struct {
	num      int
	head     map[int]string
	sections map[int]*section
}

func NewChapter(num int, chapterPath string) *chapter {

	var chapter = chapter{num: num}

	chapterDir, err := os.Open(chapterPath)
	if err != nil {
		return nil
	}

	chapterDirs, err := chapterDir.Readdir(0)
	if err != nil {
		return nil
	}

	// HEAD
	//chapter.sections = make(map[int]*section, len(chapterDirs))
	//for _, chapterDir := range chapterDirs {
	//	sectionName := chapterDir.Name()
	//	i, _ := strconv.Atoi(sectionName[:3])
	//	chapter.sections[i] = NewSection(filepath.Join(chapterPath, "sections", sectionName))
	//}

	// SECTIONS
	chapter.sections = make(map[int]*section, len(chapterDirs))
	for _, chapterDir := range chapterDirs {
		sectionName := chapterDir.Name()
		num, _ := strconv.Atoi(sectionName[:3])
		chapter.sections[num] = NewSection(num, filepath.Join(chapterPath, "sections", sectionName))
	}

	return &chapter
}
