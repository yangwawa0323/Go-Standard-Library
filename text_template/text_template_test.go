package texttemplate

import (
	"fmt"
	"io"
	"os"
	"testing"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func Test_Parse(t *testing.T) {
	sweaters := Inventory{"Wool", 17}
	tmpl, err := template.New("test").
		Parse("{{.Count}} items are mode of {{.Material}}")
	if err != nil {
		t.Fatal("Parse template error: ", err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		t.Fatal("Execute template error: ", err)
	}
}

func Test_Template_Trim_Space(t *testing.T) {
	sweaters := []Inventory{
		{"Wool", 17},
		{"Ken", 34},
		{"Yangwawa", 47},
	}
	tmpl_str := `
	{{/* "." the dot is the variable of the secondary parameter of tmpl.Execute() */}}
	{{ range $index, $item := .}}
	   {{$index}}. {{$item.Count}} items are mode of {{$item.Material}}
	{{ end }}
	`

	tmpl, err := template.New("test").
		Parse(tmpl_str)
	if err != nil {
		t.Fatal("Parse template error: ", err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		t.Fatal("Execute template error: ", err)
	}
}

func Test_URLQueryEscaper(t *testing.T) {
	var keyword string = `It is a demo of escape url string`
	var name string = "yang@wawa"
	// static method
	keyword = template.URLQueryEscaper(keyword)
	name = template.URLQueryEscaper(name)
	uri := fmt.Sprintf("http://localhost/search?keyword=%s&name=%s", keyword, name)
	io.WriteString(os.Stdout, uri)
}

type User struct {
	Name  string
	Email string
}

func Test_ParseFiles(t *testing.T) {
	tmpl := template.Must(template.ParseFiles("./template_01.txt", "./template_02.txt"))
	tmpl.Execute(os.Stdout, &User{"Yangwawa", "12238747@qq.com"})
}
