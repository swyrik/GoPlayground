package main

import (
	"os"
	"text/template"
)

func texttemplatestest() {

	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"GO", "Rust", "C++", "C#"})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name : {{.Name}}\n") // we can get the field name from the provided data
	t2.Execute(os.Stdout, struct {
		Name string
	}{"golang"})

	t2.Execute(os.Stdout, map[string]string{"Name": "java"})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, " not empty ")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
}
