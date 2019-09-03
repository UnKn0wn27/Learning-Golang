package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Note struct {
	Title       string
	Description string
}

const tmpl = `Note are:
{{range .}}
	Title: {{.Title}}. Description: {{.Description}}
	{{$note := "Sample Note/Declared Variable"}}
	{{$note}}
{{end}}
`

func main() {
	//Create slice of Note struct
	note := []Note{
		{"text/templates", "Template generates textual ouput"},
		{"html/templates", "Template generates html output"},
	}

	//create a new template name
	t := template.New("note")

	//parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//Applies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute: ", err)
		return
	}

	fmt.Println("\nNamed Templates\n")
	t, err = template.New("test").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "World")
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
