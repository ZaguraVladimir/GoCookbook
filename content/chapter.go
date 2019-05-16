package content

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

// ГЛАВА
type chapter struct {
	name  string
	parts map[int]fmt.Stringer
}

func newChapter(path string) *chapter {

	items := getItems(path)

	var chapter = chapter{}
	_, chapter.name = parseName(path)
	chapter.parts = make(map[int]fmt.Stringer, len(items))

	for _, fi := range items {
		name := fi.Name()
		fullName := filepath.Join(path, name)
		num, _ := strconv.Atoi(name[:3])
		chapter.parts[num] = newSection(fullName)
	}

	return &chapter
}

func (p *chapter) String() string {

	var b strings.Builder

	b.Grow(len(p.name) + 6)
	b.WriteString("##")
	b.WriteString(p.name)

	for i := 1; i < len(p.parts); i++ {
		b.WriteString("\r\n")
		b.WriteString(p.parts[i].String())
	}

	return b.String()
}
