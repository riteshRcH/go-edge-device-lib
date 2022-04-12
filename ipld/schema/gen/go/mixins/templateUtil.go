package mixins

import (
	"io"
	"strings"
	"text/template"

	"github.com/riteshRcH/go-edge-device-lib/ipld/testutil"
)

func doTemplate(tmplstr string, w io.Writer, data interface{}) {
	tmpl := template.Must(template.New("").
		Funcs(template.FuncMap{
			"title": func(s string) string { return strings.Title(s) },
		}).
		Parse(testutil.Dedent(tmplstr)))
	if err := tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}
