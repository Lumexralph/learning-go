package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/Lumexralph/go_learning/github"
)

const tmpl = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
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

	f, err := os.Create("issue.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := reports.Execute(f, result); err != nil {
		log.Fatal(err)
	}
}
