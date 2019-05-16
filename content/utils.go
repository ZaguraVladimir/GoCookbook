package content

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseName(path string) (num int, name string) {

	_, name = filepath.Split(path)
	num, _ = strconv.Atoi(name[:3])

	name = strings.ReplaceAll(name, "_", "/")
	name = strings.TrimSuffix(name, ".md")
	name = name[4:]

	return num, name
}

func getItems(path string) []os.FileInfo {

	dir, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	items, err := dir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	return items
}
