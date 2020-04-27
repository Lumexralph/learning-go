package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/Lumexralph/go_learning/github"
)

const tmpl = `{{.TotalCount}} issues:
{{range .Items}}------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// Producing output with a template is a two-step process.
// First we must parse the template into a suitable internal representation,
// and then execute it on specific inputs.
var issuelist, err = template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(tmpl)
var reports = template.Must(issuelist, err)

// create the report
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("issue.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := reports.Execute(f, result); err != nil {
		log.Fatal(err)
	}
}
