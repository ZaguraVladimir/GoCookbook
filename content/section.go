package content

type section struct {
	num         int
	subsections map[int]string
}

func NewSection(num int, sectionPath string) *section {
	var section = section{}
	return &section
}
