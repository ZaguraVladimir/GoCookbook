package content

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type MDText string

func NewMDText(path string, captionLevel int) *MDText {

	result := MDText("")

	if data, err := ioutil.ReadFile(path); err == nil {

		_, name := filepath.Split(path)
		text := strings.Repeat("#", captionLevel) + name[4:len(name)-3]
		if len(data) != 0 {
			text = fmt.Sprintf("%s\r\n%s", text, string(data))
		}
		result = MDText(text)
	}
	return &result
}

func (t MDText) String() string {
	return string(t)
}
