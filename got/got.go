package got

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

func init() {
	funcMap := template.FuncMap{
		"JsonTagStyle": func(s string) string {
			return strings.Title(s)
		},
	}
	var err error
	tmpl, err = template.New("").Funcs(funcMap).ParseGlob("template/*.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
}

func GenerateRepository(data interface{}) error {
	return tmpl.ExecuteTemplate(os.Stdout, "repository", data)
}
