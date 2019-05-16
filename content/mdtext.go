package content

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type mdText string

func newMDText(path string, captionLevel int) *mdText {

	result := mdText("")

	if data, err := ioutil.ReadFile(path); err == nil {

		_, name := filepath.Split(path)
		text := strings.Repeat("#", captionLevel) + name[4:len(name)-3]
		if len(data) != 0 {
			text = fmt.Sprintf("%s\r\n%s", text, string(data))
		}
		result = mdText(text)
	}
	return &result
}

func (t mdText) String() string {
	return string(t)
}
