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

	booker := content.NewBooker()
	if err := booker.AddBook("Go Cookbook", "D:\\PROG\\GoProjects\\src\\GoCookbook\\GoCookbook"); err != nil {
		log.Fatal(err)
	}

	booker.WriteFile("Go Cookbook", "D:\\PROG\\GoProjects\\src\\GoCookbook\\result\\Go Cookbook.md")
	booker.WriteFile("Go Cookbook", "D:\\PROG\\GoProjects\\src\\GoCookbook\\result\\Go Cookbook.html")
}

func testMD() {

	data, _ := ioutil.ReadFile("in.md")
	page := struct {
		Title string
		Body  string
	}{"Test MarkDown", string(data)}

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
	if err := tmpl.ExecuteTemplate(w, tmplFileName, page); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Complete: ", time.Now().Format("15:04:05"))
}
