package content

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

// РАЗДЕЛ
type section struct {
	name  string
	parts map[int]fmt.Stringer
}

func newSection(path string) *section {

	items := getItems(path)

	var section = section{}
	_, section.name = parseName(path)
	section.parts = make(map[int]fmt.Stringer, len(items))

	for _, fi := range items {
		name := fi.Name()
		fullName := filepath.Join(path, name)
		num, _ := strconv.Atoi(name[:3])
		section.parts[num] = newMDText(fullName, 4)
	}

	return &section
}

func (p *section) String() string {

	var b strings.Builder

	b.Grow(len(p.name) + 6)
	b.WriteString("###")
	b.WriteString(p.name)

	for i := 1; i < len(p.parts); i++ {
		b.WriteString("\r\n")
		b.WriteString(p.parts[i].String())
	}

	return b.String()
}
