package main

import (
	"GoCookbook/content"
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	//testMD()

	sourcePath := "source"
	book := content.NewBook(sourcePath)

	fmt.Println(sourcePath, book)

}

func testMD() {

	data, _ := ioutil.ReadFile("in.md")
	p := Page{"Test MarkDown", string(data)}

	w, err := os.Create("out.html")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	markDowner := func(args ...interface{}) template.HTML {

		s := blackfriday.Run([]byte(fmt.Sprintf("%s", args...)), blackfriday.WithExtensions(blackfriday.CommonExtensions))
		return template.HTML(s)
	}

	tmplFileName := "tmpl.html"
	tmpl := template.New(tmplFileName)
	tmpl.Funcs(template.FuncMap{"markDown": markDowner})
	template.Must(tmpl.ParseFiles(tmplFileName))
	if err := tmpl.ExecuteTemplate(w, tmplFileName, p); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Complete: ", time.Now().Format("15:04:05"))
}

type Page struct {
	Title string
	Body  string
}
