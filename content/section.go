package content

type section struct {
	subsections map[int]string
}

func NewSection(sectionPath string) *section {
	var section = section{}
	return &section
}
